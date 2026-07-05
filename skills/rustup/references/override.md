# Overrides and Toolchain Resolution

How rustup selects the active toolchain — the override hierarchy, directory overrides, and `rust-toolchain.toml` files.

## Resolution Precedence

When rustup (or a proxy) needs to determine the active toolchain, it checks these sources in order. The first match wins:

1. **`+toolchain` shorthand** on the command line (e.g. `cargo +nightly build`) — highest priority
2. **`RUSTUP_TOOLCHAIN`** environment variable
3. **`rustup override`** directory override (closest to CWD wins)
4. **`rust-toolchain.toml` / `rust-toolchain`** file (closest to CWD wins)
5. **`rustup default`** global default toolchain — lowest priority

Directory overrides (#3) and toolchain files (#4) are both resolved by **proximity** — whichever is closer to the current directory wins. A `rust-toolchain.toml` in the current directory beats a directory override set on a parent directory.

## `rustup override` — Directory Overrides

| Subcommand | Description |
|------------|-------------|
| `rustup override set <toolchain>` | Set a toolchain override for the current directory |
| `rustup override unset` | Remove the override for the current directory |
| `rustup override list` | List all configured directory overrides |

### `override set` flags

| Flag | Description |
|------|-------------|
| `--path <path>` | Set the override for a different directory (instead of CWD) |

### `override unset` flags

| Flag | Description |
|------|-------------|
| `--path <path>` | Remove the override for a specific path |
| `--nonexistent` | Remove overrides for all directories that no longer exist |

```
rustup override set nightly-2024-01-15
rustup override set 1.70.0 --path /home/user/work/legacy-app
rustup override unset
rustup override list
```

### How overrides are stored

Overrides are stored in **`${RUSTUP_HOME}/settings.toml`** under the `[overrides]` section. Keys are **canonicalized** (absolute, resolved) directory paths; values are toolchain names.

```toml
[overrides]
"/home/user/projects/my-project" = "nightly-2024-01-15"
"/home/user/work/legacy-app" = "1.70.0"
```

- Directory keys are canonicalized to ensure consistent matching regardless of relative paths or symlinks.
- An override applies to the directory **and all its child directories** (walked recursively).
- The closest override to the current directory wins when multiple parent directories have overrides.

> The schema of `settings.toml` is **not** part of rustup's public interface — use the CLI to query and set settings. See [environment.md](environment.md).

## `rust-toolchain.toml` / `rust-toolchain` Files

A file named `rust-toolchain.toml` or `rust-toolchain` in a project directory pins the toolchain for that project. These files are suitable for version control so that all contributors use the same toolchain.

### Discovery

rustup walks **up** the directory tree from the current directory to the filesystem root, looking for these files. The closest one wins.

### TOML format

Both filenames support the TOML format:

```toml
[toolchain]
channel = "nightly-2024-01-15"
components = [ "rustfmt", "rustc-dev" ]
targets = [ "wasm32-unknown-unknown", "thumbv7em-none-eabi" ]
profile = "minimal"
```

| Key | Type | Description |
|-----|------|-------------|
| `channel` | string | Toolchain spec (`<channel>[-<date>]`) or a custom toolchain name. See [toolchain.md](toolchain.md). |
| `path` | string | Absolute path to a custom toolchain. **Mutually exclusive with `channel`.** When `path` is used, `components`, `targets`, and `profile` have no effect. |
| `profile` | string | `minimal`/`default`/`complete`. If unset, uses the value from `rustup set profile`. See [component.md](component.md). |
| `components` | string list | Additional components to install (additive with the profile). |
| `targets` | string list | Targets to install for cross-compilation. The host target is included automatically. |

Rules:
- The `[toolchain]` section is **mandatory**.
- At least one property must be specified.
- `channel` and `path` are **mutually exclusive**.

### Legacy plain-text format

For backwards compatibility, `rust-toolchain` files (NOT `rust-toolchain.toml`) also support a single-line plain-text format containing just a bare toolchain name:

```
nightly-2024-01-15
```

- The file must be **US-ASCII** encoded (no BOM on Windows).
- The parser auto-detects the format based on content structure.
- This format is **not** available for `rust-toolchain.toml` files.

### Both files present

If both `rust-toolchain` and `rust-toolchain.toml` exist in the same directory, the `rust-toolchain` file (without `.toml`) is read for backwards compatibility.

### Path security

Path-based toolchains in `rust-toolchain.toml` must use **absolute paths**. Relative paths are rejected because they could enable arbitrary code execution through directory traversal attacks.

### Portability

Toolchain files **cannot specify architecture-specific channels** (like `nightly-x86_64-unknown-linux-gnu`) unless that toolchain is already installed. This ensures files can be shared across different host architectures. Use the channel-only form (`nightly-2024-01-15`) for portability.

### Install-on-discovery

When a `rust-toolchain.toml` file is discovered and the specified channel is not installed, rustup **automatically installs it** (if `RUSTUP_AUTO_INSTALL=1`), including any `components` and `targets` listed. See [environment.md](environment.md).

## `+toolchain` Shorthand

The `+<toolchain>` prefix on the first argument of a proxied command (or `rustup run`) overrides the active toolchain for that single invocation. This is priority #1 — it beats everything else.

```
cargo +nightly build          # equivalent to: rustup run nightly cargo build
cargo +stable --version
rustc +nightly-2024-01-15 --version
```

The `+toolchain` syntax is handled by the rustup proxy before the real tool ever sees it. See [proxy.md](proxy.md).

## Related

- [proxy.md](proxy.md) — how the proxy uses toolchain resolution before dispatching
- [environment.md](environment.md) — `RUSTUP_TOOLCHAIN`, `RUSTUP_AUTO_INSTALL`, `settings.toml`
- [toolchain.md](toolchain.md) — toolchain name syntax
