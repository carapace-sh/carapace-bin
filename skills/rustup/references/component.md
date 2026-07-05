# Components and Profiles

How rustup manages components — the installable pieces of a Rust toolchain — and profiles that group components.

## Concept

A **component** is an installable piece of a Rust toolchain: the compiler, the package manager, documentation, the formatter, the linter, the language server, etc. Components are added and removed per toolchain with `rustup component`.

A **profile** is a named bundle of components installed by default with a new toolchain.

## `rustup component` Subcommands

| Subcommand | Description |
|------------|-------------|
| `rustup component list` | List available and installed components |
| `rustup component add <component>...` | Add component(s) to the active toolchain |
| `rustup component remove <component>...` | Remove component(s) |

### `component list` flags

| Flag | Description |
|------|-------------|
| `--installed` | List only installed components |
| `--toolchain <name>` | Operate on a specific toolchain |

### `component add` / `component remove` flags

| Flag | Description |
|------|-------------|
| `--toolchain <name>` | Operate on a specific toolchain instead of the active one |
| `--target <target>` | Add/remove the component for a specific target (e.g. `rust-std` per target) |

## Components

| Component | Description |
|-----------|-------------|
| `rustc` | The Rust compiler and `rustdoc` |
| `cargo` | Package manager and build tool |
| `rustfmt` | Automatic code formatter |
| `rust-std` | Standard library (one per target, e.g. `rust-std-x86_64-pc-windows-msvc`) |
| `rust-docs` | Local copy of Rust documentation |
| `rust-analyzer` | Language server for IDE/editor support |
| `clippy` | Lint tool for common mistakes |
| `miri` | Experimental Rust interpreter for undefined-behavior checking |
| `rust-src` | Local copy of the Rust standard library source code |
| `rust-mingw` | GCC linker and platform libraries for the `x86_64-pc-windows-gnu` target |
| `llvm-tools` | Collection of LLVM tools (unstable) |
| `rustc-dev` | Compiler development libraries (for building tools that link against rustc) |
| `rls` | **(Deprecated)** Legacy IDE tool, replaced by `rust-analyzer` |

### Per-target components

Some components are per-target — most notably `rust-std`, which has a separate install per target triple (e.g. `rust-std-x86_64-unknown-linux-gnu`, `rust-std-wasm32-unknown-unknown`). Use `rustup target add` for these rather than `rustup component add`.

## Profiles

A profile determines which components are installed by default with a new toolchain. Set the default profile with `rustup set profile <name>` or override per-install with `--profile`.

| Profile | Components | Use case |
|---------|------------|----------|
| `minimal` | `rustc`, `cargo`, `rust-std` | CI/CD, minimal installations |
| `default` | `minimal` + `rust-docs`, `rustfmt`, `clippy` | General development |
| `complete` | All available components | Full dev environment (rarely succeeds) |

`complete` includes everything, but some components are not available on all platforms, so a `complete` install often fails. Prefer `default` and add components explicitly.

### Setting the profile

```
rustup set profile minimal          # set default for future installs
rustup install nightly --profile default   # override per-install
rustup show profile                 # show current default profile
```

The profile is stored in `${RUSTUP_HOME}/settings.toml` — see [environment.md](environment.md).

## Operating on a non-default toolchain

By default `rustup component` operates on the **active** toolchain (see [override.md](override.md)). Use `--toolchain` to target a different one:

```
rustup component add rust-src --toolchain nightly
rustup component list --installed --toolchain stable
```

## Related

- [toolchain.md](toolchain.md) — `--profile` and `--component` flags on `toolchain install`
- [target.md](target.md) — `rust-std` per target is managed via `rustup target`
- [proxy.md](proxy.md) — which proxy binary maps to which component
