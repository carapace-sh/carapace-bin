# Environment Variables

Cargo's environment variables: configuration overrides (the `CARGO_*` mapping), variables set for build scripts and compilation, and runtime variables. For the config-file equivalents, see [config.md](config.md).

> **Source of truth**: <https://doc.rust-lang.org/cargo/reference/environment-variables.html>.

## Configuration Override Variables

For each config key `foo.bar`, the env var `CARGO_FOO_BAR` overrides it (uppercase, dots/dashes → underscores). These take precedence over config files but below `--config`. Only scalars and some arrays are supported.

| Env Var | Config Key |
|---------|------------|
| `CARGO_BUILD_JOBS` | `build.jobs` |
| `CARGO_BUILD_RUSTC` | `build.rustc` |
| `CARGO_BUILD_RUSTC_WRAPPER` | `build.rustc-wrapper` |
| `CARGO_BUILD_RUSTC_WORKSPACE_WRAPPER` | `build.rustc-workspace-wrapper` |
| `CARGO_BUILD_RUSTDOC` | `build.rustdoc` |
| `CARGO_BUILD_RUSTFLAGS` | `build.rustflags` |
| `CARGO_BUILD_RUSTDOCFLAGS` | `build.rustdocflags` |
| `CARGO_BUILD_TARGET` | `build.target` |
| `CARGO_BUILD_TARGET_DIR` | `build.target-dir` |
| `CARGO_BUILD_INCREMENTAL` | `build.incremental` |
| `CARGO_BUILD_DEP_INFO_BASEDIR` | `build.dep-info-basedir` |
| `CARGO_NET_OFFLINE` | `net.offline` |
| `CARGO_NET_RETRY` | `net.retry` |
| `CARGO_NET_GIT_FETCH_WITH_CLI` | `net.git-fetch-with-cli` |
| `CARGO_HTTP_DEBUG` | `http.debug` |
| `CARGO_HTTP_PROXY` | `http.proxy` |
| `CARGO_HTTP_TIMEOUT` | `http.timeout` |
| `CARGO_HTTP_LOW_SPEED_LIMIT` | `http.low-speed-limit` |
| `CARGO_HTTP_CAINFO` | `http.cainfo` |
| `CARGO_HTTP_CHECK_REVOKE` | `http.check-revoke` |
| `CARGO_HTTP_SSL_VERSION` | `http.ssl-version` |
| `CARGO_HTTP_MULTIPLEXING` | `http.multiplexing` |
| `CARGO_HTTP_USER_AGENT` | `http.user-agent` |
| `CARGO_TERM_QUIET` | `term.quiet` |
| `CARGO_TERM_VERBOSE` | `term.verbose` |
| `CARGO_TERM_COLOR` | `term.color` |
| `CARGO_TERM_PROGRESS_WHEN` | `term.progress.when` |
| `CARGO_TERM_PROGRESS_WIDTH` | `term.progress.width` |
| `CARGO_INSTALL_ROOT` | `install.root` |
| `CARGO_REGISTRY_DEFAULT` | `registry.default` |
| `CARGO_REGISTRY_TOKEN` | `registry.token` |
| `CARGO_REGISTRIES_<NAME>_INDEX` | `registries.<name>.index` |
| `CARGO_REGISTRIES_<NAME>_TOKEN` | `registries.<name>.token` |
| `CARGO_REGISTRIES_CRATES_IO_PROTOCOL` | `registries.crates-io.protocol` |
| `CARGO_CARGO_NEW_VCS` | `cargo-new.vcs` |
| `CARGO_FUTURE_INCOMPAT_REPORT_FREQUENCY` | `future-incompat-report.frequency` |
| `CARGO_CACHE_AUTO_CLEAN_FREQUENCY` | `cache.auto-clean-frequency` |
| `CARGO_RESOLVER_INCOMPATIBLE_RUST_VERSIONS` | `resolver.incompatible-rust-versions` |
| `CARGO_PROFILE_<NAME>_<KEY>` | `profile.<name>.<key>` |

### rustflags Priority

For `build.rustflags`, four mutually exclusive sources are checked in order (first wins):

1. `CARGO_ENCODED_RUSTFLAGS` — flags separated by `\x1f` (ASCII Unit Separator)
2. `RUSTFLAGS` — space-separated flags
3. `target.<triple>.rustflags` + `target.<cfg>.rustflags` (joined)
4. `build.rustflags` config value

The same priority applies to `rustdocflags` (`CARGO_ENCODED_RUSTDOCFLAGS`, `RUSTDOCFLAGS`).

### Direct Compiler Override Variables

| Env Var | Description |
|---------|-------------|
| `RUSTC` | Override the rustc executable (also `CARGO_BUILD_RUSTC`) |
| `RUSTDOC` | Override the rustdoc executable (also `CARGO_BUILD_RUSTDOC`) |
| `RUSTFLAGS` | Extra rustc flags (see priority above) |
| `RUSTDOCFLAGS` | Extra rustdoc flags (see priority above) |
| `RUSTC_WRAPPER` | Wrapper for rustc (also `CARGO_BUILD_RUSTC_WRAPPER`) |
| `RUSTC_WORKSPACE_WRAPPER` | Wrapper for workspace members only |

## Variables Set by Cargo (Build Scripts and Compilation)

These are set in the environment when Cargo compiles a crate and when it runs build scripts. For build-script usage, see [build-scripts.md](build-scripts.md).

### Package Information

| Variable | Description |
|----------|-------------|
| `CARGO_PKG_NAME` | Package name |
| `CARGO_PKG_VERSION` | Full version (e.g. `"1.2.3"`) |
| `CARGO_PKG_VERSION_MAJOR` | Major version |
| `CARGO_PKG_VERSION_MINOR` | Minor version |
| `CARGO_PKG_VERSION_PATCH` | Patch version |
| `CARGO_PKG_VERSION_PRE` | Pre-release version |
| `CARGO_PKG_AUTHORS` | Colon-separated authors |
| `CARGO_PKG_DESCRIPTION` | Description |
| `CARGO_PKG_HOMEPAGE` | Homepage URL |
| `CARGO_PKG_REPOSITORY` | Repository URL |
| `CARGO_PKG_LICENSE` | License |
| `CARGO_PKG_LICENSE_FILE` | License file path |
| `CARGO_PKG_README` | README path |
| `CARGO_PKG_RUST_VERSION` | MSRV from manifest |
| `CARGO_MANIFEST_DIR` | Directory containing `Cargo.toml` |
| `CARGO_MANIFEST_LINKS` | The `links` manifest value |
| `CARGO_CRATE_NAME` | Name of the crate being compiled |
| `CARGO_BIN_NAME` | Name of the binary being compiled |
| `CARGO_PRIMARY_PACKAGE` | Set if the package being built is the primary one |

### Target Configuration

| Variable | Description |
|----------|-------------|
| `CARGO_CFG_TARGET_ARCH` | CPU target architecture |
| `CARGO_CFG_TARGET_ENDIAN` | `little` or `big` |
| `CARGO_CFG_TARGET_ENV` | Target environment ABI |
| `CARGO_CFG_TARGET_FAMILY` | Target family (e.g. `unix`, `windows`) |
| `CARGO_CFG_TARGET_FEATURE` | Comma-separated CPU target features |
| `CARGO_CFG_TARGET_OS` | Target operating system |
| `CARGO_CFG_TARGET_POINTER_WIDTH` | CPU pointer width (e.g. `64`) |
| `CARGO_CFG_TARGET_VENDOR` | Target vendor |
| `CARGO_CFG_UNIX` | Set on unix-like platforms |
| `CARGO_CFG_WINDOWS` | Set on windows-like platforms |

> **Build scripts**: use these `CARGO_CFG_*` vars instead of `cfg!()`/`#[cfg]`, which check the **host** platform.

### Features

| Variable | Description |
|----------|-------------|
| `CARGO_FEATURE_<NAME>` | Set for each enabled feature (uppercased, `-`→`_`) |

### Build Script Specific

| Variable | Description |
|----------|-------------|
| `OUT_DIR` | Where build scripts write generated artifacts |
| `HOST` | Host platform triple |
| `TARGET` | Target platform triple |
| `NUM_JOBS` | Max parallel jobs |
| `OPT_LEVEL` | Optimization level (`0`-`3`, `s`, `z`) |
| `DEBUG` | `"true"` or `"false"` |
| `PROFILE` | Profile name (`debug`, `release`) |
| `CARGO_ENCODED_RUSTFLAGS` | rustc flags separated by `\x1f` |

### Output Directories

| Variable | Description |
|----------|-------------|
| `CARGO_TARGET_DIR` | Location for generated artifacts (overrides `build.target-dir`) |
| `CARGO_TARGET_TMPDIR` | Only set when building integration tests/benchmarks |

## Runtime Variables

| Variable | Description |
|----------|-------------|
| `CARGO_HOME` | Cargo's data directory (default `~/.cargo`). Stores registry cache, git checkouts, config, installed binaries. |
| `CARGO_INCREMENTAL` | `1` forces incremental on; `0` forces off (overrides profile) |
| `CARGO_LOG` | Debug log level for cargo itself (`debug`, `info`, `warn`, `error`, `trace`). Uses `env_logger`. |
| `CARGO_MAKEFLAGS` | Parameters for Cargo's jobserver (parallelize subprocesses) |
| `CARGO_CACHE_RUSTC_INFO` | `0` disables caching of compiler version info |

## `[env]` Config Section

The `[env]` table in `.cargo/config.toml` sets env vars for build scripts, rustc, `cargo run`, and `cargo build`:

```toml
[env]
OPENSSL_DIR = "/opt/openssl"
TMPDIR = { value = "/home/tmp", force = true }
# OPENSSL_DIR = { value = "vendor/openssl", relative = true }
```

- Does not override existing env vars by default
- `force = true` overrides existing
- `relative = true` resolves the path relative to the config file location

See [config.md](config.md).

## References

- <https://doc.rust-lang.org/cargo/reference/environment-variables.html>

## Related

- [config.md](config.md) — config file sections and env var mapping
- [build-scripts.md](build-scripts.md) — which env vars build scripts read
- [concepts.md](concepts.md) — `CARGO_HOME` directory layout
