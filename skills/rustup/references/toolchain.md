# Toolchain Management

How rustup manages toolchains — the `rustup toolchain` subcommands, the toolchain specification syntax, and custom toolchains.

## `rustup toolchain` Subcommands

| Subcommand | Description |
|------------|-------------|
| `rustup toolchain install <toolchain>` | Install a toolchain |
| `rustup toolchain uninstall <toolchain>...` | Uninstall one or more toolchains |
| `rustup toolchain remove <toolchain>...` | Alias for `uninstall` |
| `rustup toolchain list` | List installed toolchains |
| `rustup toolchain link <name> <path>` | Create a custom toolchain by linking to a local Rust build |

`rustup install` and `rustup uninstall` are shorthands for `rustup toolchain install` / `rustup toolchain uninstall` — they share the same flags.

### `toolchain install` flags

| Flag | Description |
|------|-------------|
| `--profile <name>` | Installation profile: `minimal`/`default`/`complete` — see [component.md](component.md) |
| `--component <name>` | Additional component to install (repeatable) |
| `--target <target>` | Additional target to install (repeatable) |
| `--force` | Force installation even if components are missing |
| `--allow-downgrade` | Allow rustup to downgrade to find required components |
| `--force-non-host` | Allow installing a non-host toolchain |
| `--no-self-update` | Suppress rustup self-update during install |

### `toolchain list` flags

| Flag | Description |
|------|-------------|
| `--verbose` / `-v` | Show rustc version info for each toolchain |
| `--quiet` / `-q` | Single-column output (useful for scripting) |

## Toolchain Specification Syntax

A toolchain name encodes a release channel, optional date, and optional host triple:

```
<channel>[-<date>][-<host>]
```

| Component | Form | Examples |
|-----------|------|----------|
| `<channel>` | `stable` \| `beta` \| `nightly` \| `<version>` | `stable`, `1.70.0`, `1.70` |
| `<version>` | `<major.minor>` \| `<major.minor.patch>` | `1.70`, `1.70.0` |
| `<date>` | `YYYY-MM-DD` | `2024-01-15` |
| `<host>` | target triple (or abbreviated on Windows) | `x86_64-pc-windows-msvc`, `msvc`, `gnu` |

### Examples

| Spec | Meaning |
|------|---------|
| `stable` | Latest stable release |
| `beta` | Latest beta release |
| `nightly` | Latest nightly build |
| `1.70.0` | Exact version 1.70.0 |
| `1.70` | Latest patch in the 1.70.x series |
| `nightly-2024-01-15` | Nightly from a specific date |
| `stable-x86_64-pc-windows-msvc` | Stable with full host triple |
| `stable-msvc` | Stable with abbreviated host (Windows MSVC ABI) |
| `stable-gnu` | Stable with GNU ABI on Windows |
| `nightly-2024-01-15-x86_64-pc-windows-msvc` | Full date + host |

### Versioned channels

- `1.70.0` pins to an exact patch release.
- `1.70` resolves to the latest patch release in the 1.70.x series.

### Dated toolchains

Dated toolchains are primarily used with the **nightly** channel to pin to a specific nightly build. This is common for projects that depend on unstable features available only on a particular nightly.

```
rustup install nightly-2024-01-15
rustup default nightly-2024-01-15
```

### Host triples

The host triple identifies the platform ABI. On Windows, abbreviated forms are accepted:

- `msvc` — expands to the host arch with the MSVC ABI (e.g. `x86_64-pc-windows-msvc`)
- `gnu` — expands to the host arch with the GNU ABI (e.g. `x86_64-pc-windows-gnu`)

On other platforms the host triple is typically omitted (rustup infers the host).

## Custom Toolchains

A custom toolchain is a locally-built Rust sysroot that rustup tracks by name. Create one with `rustup toolchain link`:

```
rustup toolchain link my-toolchain /path/to/rust/build/sysroot
```

- The `<path>` must point to a directory laid out like a standard Rust sysroot (containing `bin/`, `lib/`, etc.).
- rustup creates a directory symlink to `<path>` in `${RUSTUP_HOME}/toolchains/<name>`.
- Custom toolchains appear in `rustup toolchain list` and can be selected with `+my-toolchain`, `rustup default my-toolchain`, or `rustup run my-toolchain <cmd>`.
- Custom toolchains do not auto-update.

### Cargo fallback for custom toolchains

If a custom toolchain does not include `cargo`, rustup falls back to `cargo` from a release channel, preferring **nightly**, then **beta**, then **stable**.

## Where Toolchains Live

Installed toolchains are stored under `${RUSTUP_HOME}/toolchains/<toolchain-name>/`:

```
~/.rustup/toolchains/
  stable-x86_64-unknown-linux-gnu/
    bin/    # rustc, cargo, rustdoc, ...
    lib/    # std, libraries for the dynamic linker path
    share/  # documentation
  nightly-2024-01-15-x86_64-unknown-linux-gnu/
  my-toolchain -> /path/to/custom/sysroot
```

See [environment.md](environment.md) for `RUSTUP_HOME` configuration.

## Related

- [component.md](component.md) — components and profiles installed with a toolchain
- [target.md](target.md) — cross-compilation targets added to a toolchain
- [override.md](override.md) — how the active toolchain is selected
