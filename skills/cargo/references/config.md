# Configuration

Cargo configuration via `.cargo/config.toml` — the hierarchy, precedence, `--config` flag, environment variable mapping, and all configuration sections.

> **Source of truth**: <https://doc.rust-lang.org/cargo/reference/config.html>.

## File Locations and Hierarchy

Cargo probes `.cargo/config.toml` in the current directory and **all parent directories**, merging values. Given an invocation in `/projects/foo/bar/baz`, the order (highest to lowest precedence):

1. `/projects/foo/bar/baz/.cargo/config.toml`
2. `/projects/foo/bar/.cargo/config.toml`
3. `/projects/foo/.cargo/config.toml`
4. `/projects/.cargo/config.toml`
5. `/.cargo/config.toml`
6. `$CARGO_HOME/config.toml` (lowest) — Unix: `~/.cargo/config.toml`, Windows: `%USERPROFILE%\.cargo\config.toml`

When invoked from a workspace root, Cargo does **not** read config files from crates within the workspace.

> Cargo also reads `.cargo/config` (no extension) for backward compatibility. If both exist, the file **without** the extension wins. The `.toml` form (added in 1.39) is preferred.

## Merge Rules

- **Scalars** (numbers, strings, booleans): deeper config directory values win; `$CARGO_HOME` has lowest priority
- **Arrays**: joined together, higher-precedence items placed **later**

## Overall Precedence (highest to lowest)

1. `--config` command-line option values
2. Environment variables (`CARGO_*`)
3. Local (project) `.cargo/config.toml` files (hierarchical)
4. `$CARGO_HOME/config.toml` (user/global)

## The `--config` Flag

Accepts TOML `key=value` syntax or a path to an extra config file:

```sh
# KEY=VALUE
cargo --config net.git-fetch-with-cli=true fetch

# Path to extra config file
cargo --config ./path/to/extra-config.toml fetch

# Multiple uses (merged left-to-right)
cargo --config 'build.rustdocflags = ["--html-in-header", "header.html"]' …

# Complex keys with cfg expressions
cargo --config "target.'cfg(all(target_arch = \"arm\", target_os = \"none\"))'.runner = 'my-runner'" …

# Profile overrides
cargo --config profile.dev.package.image.opt-level=3 …
```

`--config` values take precedence over environment variables, which take precedence over config files.

## Environment Variable Mapping

For each config key `foo.bar`, the env var `CARGO_FOO_BAR` can be used. Conversion:
- Uppercase
- Dots and dashes → underscores

Example: `target.x86_64-unknown-linux-gnu.runner` → `CARGO_TARGET_X86_64_UNKNOWN_LINUX_GNU_RUNNER`

Env vars take precedence over TOML config files. Only integer, boolean, string, and some array values are supported via env vars. See [environment.md](environment.md) for the full variable list.

## Config-Relative Paths

| Source | Path relative to |
|--------|------------------|
| Environment variables | current working directory |
| `--config KEY=VALUE` | current working directory |
| Config files | parent directory of the `.cargo` directory containing the config file |

For executables without path separators, Cargo searches `PATH`.

## Including Extra Config Files

```toml
include = ["frodo.toml", "samwise.toml"]
```

Or with inline tables (optional flag):

```toml
include = [
    { path = "required.toml" },
    { path = "optional.toml", optional = true },
]
```

- Only `.toml` paths accepted
- Loaded left-to-right; later files take precedence
- Recursive (included files can have `include`)
- The including file's own values take highest precedence

## Configuration Sections

### `[alias]` — Command Aliases

```toml
[alias]
b = "build"
t = ["test", "--color", "always"]
```

Built-in aliases: `b`=build, `c`=check, `d`=doc, `t`=test, `r`=run, `rm`=remove. Aliases are recursive. Cannot redefine existing built-in commands.

### `[build]` — Build Settings

| Key | Type | Default | Env Var | Description |
|-----|------|---------|---------|-------------|
| `build.jobs` | int/string | # CPUs | `CARGO_BUILD_JOBS` | Max parallel compiler processes. Negative = CPUs + value. `"default"` resets. |
| `build.rustc` | path | `"rustc"` | `CARGO_BUILD_RUSTC`, `RUSTC` | Rust compiler executable |
| `build.rustc-wrapper` | path | — | `CARGO_BUILD_RUSTC_WRAPPER`, `RUSTC_WRAPPER` | Wrapper for rustc |
| `build.rustc-workspace-wrapper` | path | — | `CARGO_BUILD_RUSTC_WORKSPACE_WRAPPER`, `RUSTC_WORKSPACE_WRAPPER` | Wrapper for workspace members only |
| `build.rustdoc` | path | `"rustdoc"` | `CARGO_BUILD_RUSTDOC`, `RUSTDOC` | Doc generator |
| `build.target` | string/array | host triple | `CARGO_BUILD_TARGET` | Target triple(s). Supports `"host-tuple"` substitution. Can be a custom target spec path. |
| `build.target-dir` | path | `"target"` | `CARGO_BUILD_TARGET_DIR`, `CARGO_TARGET_DIR` | Compiler output directory |
| `build.build-dir` | path | `build.target-dir` | `CARGO_BUILD_BUILD_DIR` | Intermediate artifacts. Supports `{workspace-root}`, `{cargo-cache-home}`, `{workspace-path-hash}`. |
| `build.rustflags` | string/array | — | `CARGO_BUILD_RUSTFLAGS`, `CARGO_ENCODED_RUSTFLAGS`, `RUSTFLAGS` | Extra rustc flags |
| `build.rustdocflags` | string/array | — | `CARGO_BUILD_RUSTDOCFLAGS`, `CARGO_ENCODED_RUSTDOCFLAGS`, `RUSTDOCFLAGS` | Extra rustdoc flags |
| `build.incremental` | bool | from profile | `CARGO_BUILD_INCREMENTAL`, `CARGO_INCREMENTAL` | Incremental compilation override |
| `build.dep-info-basedir` | path | — | `CARGO_BUILD_DEP_INFO_BASEDIR` | Strips prefix from dep-info file paths |

#### `build.rustflags` Source Priority (first wins)

1. `CARGO_ENCODED_RUSTFLAGS` (flags separated by `\x1f`)
2. `RUSTFLAGS`
3. All matching `target.<triple>.rustflags` + `target.<cfg>.rustflags` joined
4. `build.rustflags` config value

### `[env]` — Environment Variables for Build Scripts/Compilation

```toml
[env]
OPENSSL_DIR = "/opt/openssl"
TMPDIR = { value = "/home/tmp", force = true }       # override existing
# OPENSSL_DIR = { value = "vendor/openssl", relative = true }  # config-relative path
```

- Does **not** override existing env vars by default
- `force = true`: override existing
- `relative = true`: value is relative to the config file's `.cargo` parent; env var gets absolute path

### `[target.<triple>]` and `[target.<cfg>]`

| Key | Type | Env Var | Description |
|-----|------|---------|-------------|
| `target.<triple>.linker` | path | `CARGO_TARGET_<triple>_LINKER` | Linker for this target |
| `target.<cfg>.linker` | path | — | Linker for cfg match (triple takes precedence) |
| `target.<triple>.runner` | string/array | `CARGO_TARGET_<triple>_RUNNER` | Runner for executables |
| `target.<cfg>.runner` | string/array | — | Runner for cfg match |
| `target.<triple>.rustflags` | string/array | `CARGO_TARGET_<triple>_RUSTFLAGS` | Target-specific rustc flags |
| `target.<cfg>.rustflags` | string/array | — | cfg-specific flags (multiple matches joined) |
| `target.<triple>.rustdocflags` | string/array | `CARGO_TARGET_<triple>_RUSTDOCFLAGS` | Target-specific rustdoc flags |

#### `[target.<triple>.<links>]` — Override Build Scripts

```toml
[target.x86_64-unknown-linux-gnu.foo]
rustc-link-lib = ["foo"]
rustc-link-search = ["/path/to/foo"]
rustc-flags = "-L /some/path"
rustc-cfg = ['key="value"']
rustc-env = { key = "value" }
rustc-cdylib-link-arg = ["…"]
metadata_key1 = "value"
```

See [build-scripts.md](build-scripts.md).

### `[net]` — Network Settings

| Key | Type | Default | Env Var | Description |
|-----|------|---------|---------|-------------|
| `net.retry` | int | `3` | `CARGO_NET_RETRY` | Network retry count |
| `net.git-fetch-with-cli` | bool | `false` | `CARGO_NET_GIT_FETCH_WITH_CLI` | Use `git` CLI instead of built-in libgit2 |
| `net.offline` | bool | `false` | `CARGO_NET_OFFLINE` | Avoid network access |
| `net.ssh.known-hosts` | string[] | — | not supported | SSH host keys (known_hosts format) |

### `[http]` — HTTP Settings

| Key | Type | Default | Env Var(s) | Description |
|-----|------|---------|------------|-------------|
| `http.debug` | bool | `false` | `CARGO_HTTP_DEBUG` | HTTP request debugging |
| `http.proxy` | string | — | `CARGO_HTTP_PROXY`, `HTTPS_PROXY`, `https_proxy`, `http_proxy` | HTTP proxy (libcurl format) |
| `http.timeout` | int | `30` | `CARGO_HTTP_TIMEOUT`, `HTTP_TIMEOUT` | Per-request timeout (seconds) |
| `http.low-speed-limit` | int | `10` | `CARGO_HTTP_LOW_SPEED_LIMIT` | Slow connection threshold (bytes/sec) |
| `http.cainfo` | path | — | `CARGO_HTTP_CAINFO` | CA bundle path |
| `http.proxy-cainfo` | path | falls back to `cainfo` | `CARGO_HTTP_PROXY_CAINFO` | Proxy CA bundle |
| `http.check-revoke` | bool | `true` (Windows) / `false` | `CARGO_HTTP_CHECK_REVOKE` | TLS revocation checks (Windows) |
| `http.ssl-version` | string/{min,max} | — | `CARGO_HTTP_SSL_VERSION` | `"default"`, `"tlsv1"`, `"tlsv1.0"`-`"tlsv1.3"` |
| `http.multiplexing` | bool | `true` | `CARGO_HTTP_MULTIPLEXING` | HTTP/2 multiplexing |
| `http.user-agent` | string | Cargo's version | `CARGO_HTTP_USER_AGENT` | Custom user-agent |

### `[term]` — Terminal Output

| Key | Type | Default | Env Var | Description |
|-----|------|---------|---------|-------------|
| `term.quiet` | bool | `false` | `CARGO_TERM_QUIET` | Suppress log messages |
| `term.verbose` | bool | `false` | `CARGO_TERM_VERBOSE` | Detailed messages |
| `term.color` | string | `"auto"` | `CARGO_TERM_COLOR` | `"auto"`, `"always"`, `"never"` |
| `term.hyperlinks` | bool | auto | `CARGO_TERM_HYPERLINKS` | Hyperlinks in terminal |
| `term.unicode` | bool | auto | `CARGO_TERM_UNICODE` | Non-ASCII unicode output |
| `term.progress.when` | string | `"auto"` | `CARGO_TERM_PROGRESS_WHEN` | `"auto"`, `"always"`, `"never"` |
| `term.progress.width` | int | — | `CARGO_TERM_PROGRESS_WIDTH` | Progress bar width |
| `term.progress.term-integration` | bool | auto | `CARGO_TERM_PROGRESS_TERM_INTEGRATION` | Report to terminal emulator |

### `[cargo-new]`

| Key | Type | Default | Env Var | Description |
|-----|------|---------|---------|-------------|
| `cargo-new.vcs` | string | `"git"` or `"none"` | `CARGO_CARGO_NEW_VCS` | `git`, `hg`, `pijul`, `fossil`, `none` |

### `[install]`

| Key | Type | Default | Env Var | Description |
|-----|------|---------|---------|-------------|
| `install.root` | path | `$CARGO_HOME` | `CARGO_INSTALL_ROOT` | Root for `cargo install` executables |

### `[registry]` and `[registries.<name>]`

See [registries.md](registries.md) for the full registry configuration reference.

### `[source.<name>]` — Source Replacement

```toml
[source.crates-io]
replace-with = "my-vendor-source"

[source.my-vendor-source]
directory = "vendor"   # or: registry, local-registry, git, branch, tag, rev
```

See [registries.md](registries.md).

### `[patch.<registry>]`

Same as `[patch]` in `Cargo.toml` but from config. Config patches take precedence over manifest patches. Relative `path` deps resolve relative to the config file. See [registries.md](registries.md).

### `[profile.<name>]`

Override profile settings globally. Each key has env var `CARGO_PROFILE_<name>_<KEY>`. See [profiles.md](profiles.md).

### `[credential-alias]`

Define aliases for credential providers:

```toml
[credential-alias.my-provider]
command = ["/path/to/credential-helper", "--arg"]
```

### `[future-incompat-report]`

| Key | Type | Default | Env Var |
|-----|------|---------|---------|
| `frequency` | string | `"always"` | `CARGO_FUTURE_INCOMPAT_REPORT_FREQUENCY` | `"always"` or `"never"` |

### `[cache]`

| Key | Type | Default | Env Var | Description |
|-----|------|---------|---------|-------------|
| `cache.auto-clean-frequency` | string | `"1 day"` | `CARGO_CACHE_AUTO_CLEAN_FREQUENCY` | `"never"`, `"always"`, or `<int><unit>` |

Global cache auto-cleaning: network files unused 3 months deleted; non-network unused 1 month. Disabled if offline.

### `[doc]`

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| `doc.browser` | string/array | `BROWSER` env or system default | Browser for `cargo doc --open` |

## Credentials File

`$CARGO_HOME/credentials.toml` (auto-managed by `cargo login`/`cargo logout`). Also reads `.cargo/credentials` (no extension) for backward compatibility.

```toml
[registry]
token = "…"

[registries.<name>]
token = "…"
```

### Credential Provider Precedence (per registry)

1. `registries.<name>.credential-provider`
2. `registry.credential-provider` (crates.io)
3. `registry.global-credential-providers` (fallback list, default `["cargo:token"]`)

## References

- <https://doc.rust-lang.org/cargo/reference/config.html>

## Related

- [environment.md](environment.md) — `CARGO_*` environment variable reference
- [profiles.md](profiles.md) — profile overrides via config
- [registries.md](registries.md) — registry and source config
- [build-scripts.md](build-scripts.md) — build script override via `[target.*.<links>]`
