# Cross-Compilation Targets

How rustup manages cross-compilation targets — the `rustup target` subcommands and the cross-compilation workflow.

## Concept

A **target** is a platform you can compile for. Each toolchain ships the Rust standard library (`rust-std`) for the **host** platform by default. To cross-compile to another platform, you install that platform's `rust-std` with `rustup target add`.

The host platform's standard library is installed automatically when a toolchain is first installed — no `rustup target add` needed.

## `rustup target` Subcommands

| Subcommand | Description |
|------------|-------------|
| `rustup target list` | List all available targets for the active toolchain |
| `rustup target add <target>...` | Install the standard library for one or more targets |
| `rustup target remove <target>...` | Remove a previously installed target |

### `target list` flags

| Flag | Description |
|------|-------------|
| `--installed` | List only installed targets |
| `--toolchain <name>` | Operate on a specific toolchain instead of the active one |

### `target add` / `target remove` flags

| Flag | Description |
|------|-------------|
| `--toolchain <name>` | Operate on a specific toolchain instead of the active one |

## Cross-Compilation Workflow

1. Install the target's standard library:

   ```
   rustup target add arm-linux-androideabi
   ```

2. Build for that target with cargo:

   ```
   cargo build --target=arm-linux-androideabi
   ```

### Important limitation

`rustup target add` installs only the Rust **standard library** for a target. A **linker** and other platform tools must still be installed manually. For example:

- Android cross-compilation requires the **Android NDK**.
- Linux cross-compilation from macOS requires a cross-linker (e.g. `cross`, `zig cc`, or a `*-linux-gnu-gcc` toolchain).

Configure the linker in cargo's config (`~/.cargo/config.toml` or `.cargo/config.toml`):

```toml
[target.arm-linux-androideabi]
linker = "arm-linux-androideabi-gcc"
```

## Multiple targets per toolchain

A single toolchain supports multiple targets simultaneously. For example, all four x86 Windows targets can be added to one toolchain:

```
rustup target add x86_64-pc-windows-msvc
rustup target add x86_64-pc-windows-gnu
rustup target add i686-pc-windows-msvc
rustup target add i686-pc-windows-gnu
```

## Operating on a non-default toolchain

By default `rustup target` operates on the **active** toolchain (see [override.md](override.md)). Use `--toolchain` to target a different one:

```
rustup target add wasm32-unknown-unknown --toolchain nightly
rustup target list --installed --toolchain nightly
```

## Common targets

| Target | Platform |
|--------|----------|
| `wasm32-unknown-unknown` | WebAssembly (no OS) |
| `wasm32-wasip1` | WebAssembly with WASI Preview 1 |
| `x86_64-pc-windows-msvc` | Windows x86_64 MSVC ABI |
| `x86_64-pc-windows-gnu` | Windows x86_64 GNU ABI |
| `x86_64-apple-darwin` | macOS Intel |
| `aarch64-apple-darwin` | macOS Apple Silicon |
| `x86_64-unknown-linux-gnu` | Linux x86_64 GNU |
| `aarch64-unknown-linux-gnu` | Linux ARM64 GNU |
| `arm-linux-androideabi` | Android ARM |
| `aarch64-linux-android` | Android ARM64 |
| `thumbv7em-none-eabi` | Embedded ARMv7E-M (no OS) |

Run `rustup target list` to see all available targets for the active toolchain.

## Related

- [toolchain.md](toolchain.md) — installing and managing toolchains
- [component.md](component.md) — components (targets are added per-toolchain via `rust-std`)
