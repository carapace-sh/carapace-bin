# Environment Variables and Configuration

The environment variables rustup reads and sets, and the `settings.toml` configuration file.

## Environment Variables

| Variable | Default | Description |
|----------|---------|-------------|
| `RUSTUP_HOME` | `~/.rustup` | Root directory for toolchains, settings, and metadata |
| `CARGO_HOME` | `~/.cargo` | Root directory for cargo cache and proxy binaries |
| `RUSTUP_TOOLCHAIN` | (none) | Override the active toolchain (priority #2 in resolution) |
| `RUSTUP_DIST_SERVER` | `https://static.rust-lang.org` | URL for downloading Rust resources |
| `RUSTUP_UPDATE_ROOT` | `https://static.rust-lang.org/rustup` | URL for rustup self-update downloads |
| `RUSTUP_AUTO_INSTALL` | `1` | Auto-install missing toolchains when referenced (`1`/`0`) |
| `RUSTUP_LOG` | (none) | Enable custom logging (tracing_subscriber directive syntax) |
| `RUSTUP_TERM_COLOR` | `auto` | Color output: `auto`, `always`, `never` |
| `RUSTUP_IO_THREADS` | auto (max 8) | (Unstable) Number of IO threads for unpacking |
| `RUSTUP_HARDLINK_PROXIES` | (unset) | (Unstable) Use hardlinks instead of symlinks for proxies |
| `RUSTUP_PERMIT_COPY_RENAME` | (unset) | (Unstable, Linux only) Allow copy fallback when `rename()` fails across devices |
| `RUSTUP_OVERRIDE_UNIX_FALLBACK_SETTINGS` | (unset) | (Unstable) Path to system-wide fallback settings file |
| `RUSTUP_TOOLCHAIN_SOURCE` | (set by rustup) | (Unstable) Indicates how `RUSTUP_TOOLCHAIN` was determined |

### `RUSTUP_HOME`

The root directory where rustup stores toolchains, settings, and metadata. Structure:

```
~/.rustup/
  settings.toml           # configuration (see below)
  toolchains/             # installed toolchains
    stable-x86_64-unknown-linux-gnu/
    nightly-2024-01-15-x86_64-unknown-linux-gnu/
  download/               # download cache
  tmp/                    # temporary files during install
```

Override with `RUSTUP_HOME` to relocate (e.g. for a portable install or a different disk).

### `CARGO_HOME`

The root directory for cargo's cache and the proxy binaries. The proxies live in `${CARGO_HOME}/bin/`. Override to relocate cargo's data.

### `RUSTUP_TOOLCHAIN`

Forces the active toolchain for the current process and any proxy invocations. This is priority #2 in toolchain resolution (see [override.md](override.md)) — beaten only by the `+toolchain` command-line shorthand.

```
RUSTUP_TOOLCHAIN=nightly cargo build   # == cargo +nightly build
```

### `RUSTUP_AUTO_INSTALL`

When `1` (the default), rustup automatically installs a toolchain that is selected but not installed — whether via `+toolchain`, `RUSTUP_TOOLCHAIN`, a directory override, or a `rust-toolchain.toml` file. Set to `0` to disable and get an error instead.

Controlled persistently with `rustup set auto-install <enable|disable|check-only>`.

### `RUSTUP_DIST_SERVER` / `RUSTUP_UPDATE_ROOT`

Override the download URLs. Useful for mirrors or offline/air-gapped environments with a local mirror:

```
RUSTUP_DIST_SERVER=https://mirrors.example.com/rust-static
RUSTUP_UPDATE_ROOT=https://mirrors.example.com/rustup
```

### `RUSTUP_LOG`

Custom logging using `tracing_subscriber` directive syntax. When unset, `--verbose` sets `DEBUG` and `--quiet` sets `WARN`. Example:

```
RUSTUP_LOG=rustup=debug rustup update
RUSTUP_LOG=trace rustup install nightly
```

### `RUSTUP_TERM_COLOR`

Controls color output regardless of terminal detection:

| Value | Behavior |
|-------|----------|
| `auto` | Color if stdout is a tty (default) |
| `always` | Always emit color codes |
| `never` | Never emit color codes |

### `RUSTUP_IO_THREADS`

Number of IO threads used for unpacking archives during install/update. Defaults to auto-detected CPU count, capped at 8.

## `settings.toml`

The configuration file at `${RUSTUP_HOME}/settings.toml`. Managed by the rustup CLI — do **not** edit by hand (the schema is not part of the public interface).

```toml
version = "12"
default_host_tuple = "x86_64-unknown-linux-gnu"
default_toolchain = "stable"
profile = "default"
auto_self_update = "enable"
auto_install = true

[overrides]
"/home/user/projects/my-project" = "nightly-2024-01-15"
```

| Field | Set by | Description |
|-------|--------|-------------|
| `version` | (internal) | Metadata schema version (current: `"12"`) |
| `default_host_tuple` | `rustup set default-host` | Default host target triple for new toolchains (`default_host_triple` accepted as a legacy alias) |
| `default_toolchain` | `rustup default` | Global default toolchain |
| `profile` | `rustup set profile` | Default installation profile — see [component.md](component.md) |
| `auto_self_update` | `rustup set auto-self-update` | `enable`/`disable`/`check-only` |
| `auto_install` | `rustup set auto-install` | `true`/`false` |
| `[overrides]` | `rustup override set/unset` | Directory → toolchain map — see [override.md](override.md) |

### System-wide fallback (Unix)

On Unix, if the user has no `settings.toml` (or no `default_toolchain` set), rustup consults `/etc/rustup/settings.toml` for a system-wide `default_toolchain`. The path can be overridden with the unstable `RUSTUP_OVERRIDE_UNIX_FALLBACK_SETTINGS` env var. `NotFound` and `PermissionDenied` errors on this file are handled silently.

## `rustup set` commands

The CLI commands that mutate `settings.toml`:

| Command | Effect |
|---------|--------|
| `rustup set default-host <triple>` | Set `default_host_triple` |
| `rustup set profile <name>` | Set `profile` |
| `rustup set auto-self-update <mode>` | Set `auto_self_update` (`enable`/`disable`/`check-only`) |
| `rustup set auto-install <mode>` | Set `auto_install` (`enable`/`disable`) |

## Related

- [override.md](override.md) — the `[overrides]` section and toolchain resolution
- [component.md](component.md) — the `profile` setting
- [proxy.md](proxy.md) — environment variables set by proxies at dispatch time
