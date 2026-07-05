# The Proxy System

How rustup's proxy/shim mechanism dispatches to real toolchain binaries, plus the `run`, `which`, `doc`, and `man` commands.

## Concept

The executables in `~/.cargo/bin` (`cargo`, `rustc`, `rustdoc`, `rustfmt`, `rust-analyzer`, etc.) are **not** the real tools. They are **shims** that delegate to the real binary in the active toolchain. This is what makes `cargo +nightly build` and per-directory toolchain switching work transparently.

## Multi-Call Binary

The `rustup` binary is a **multi-call binary** — its behavior changes based on the name used to invoke it (the `arg0`):

| Invoked as | Mode entered |
|------------|--------------|
| `rustup` | `rustup_mode` — the CLI for managing toolchains |
| `rustup-init` / `rustup-setup` | `setup_mode` — the installer |
| `cargo`, `rustc`, `rustdoc`, `rustfmt`, `rust-analyzer`, `rust-lldb`, `rust-gdb`, `rust-gdbgui`, `clippy-driver`, `cargo-clippy`, `cargo-miri`, `rls` | `proxy_mode` — dispatch to the real tool |

### How the shims are created

- **Unix:** proxies are **symlinks** (or hardlinks, controlled by `RUSTUP_HARDLINK_PROXIES`) to the `rustup` binary.
- **Windows:** proxies are **copies** of the `rustup` executable.

During install or update, rustup ensures that for every tool in the Rust distribution a corresponding executable named after that tool exists in `~/.cargo/bin`.

## Static Proxy List

| Proxy | Component | Description |
|-------|-----------|-------------|
| `rustc` | `rustc` | The Rust compiler |
| `rustdoc` | `rustc` | Documentation generator |
| `cargo` | `cargo` | Package manager and build tool |
| `rust-lldb` | `rustc` | LLDB debugger wrapper with Rust pretty-printing |
| `rust-gdb` | `rustc` | GDB debugger wrapper with Rust pretty-printing |
| `rust-gdbgui` | `rustc` | GDB GUI debugger wrapper |
| `rust-analyzer` | `rust-analyzer` | Language server for IDE support |
| `rustfmt` | `rustfmt` | Code formatter |
| `cargo-fmt` | `rustfmt` | Code formatter invoked via cargo |
| `clippy-driver` | `clippy` | Linting tool driver |
| `cargo-clippy` | `clippy` | Linting tool invoked via cargo |
| `cargo-miri` | `miri` | Experimental MIR interpreter for UB checking |
| `rls` | `rls` | **(Deprecated)** Legacy IDE tool, replaced by rust-analyzer |

`rust-analyzer`, `rustfmt`, and `cargo-fmt` are "dup tools" — tools commonly installed by Cargo as well. rustup takes care not to overwrite a user's own installation of these.

See [component.md](component.md) for component details.

## Proxy Dispatch Flow

```
~/.cargo/bin/cargo  (symlink/copy → rustup binary)
  → rustup detects arg0 = "cargo", enters proxy_mode
  → Toolchain resolution (see override.md):
      1. Check first arg for "+toolchain" prefix
      2. Check RUSTUP_TOOLCHAIN env var
      3. Walk up dir tree for rustup override (settings.toml)
      4. Walk up dir tree for rust-toolchain.toml / rust-toolchain
      5. Use default toolchain
  → Install-on-run (if RUSTUP_AUTO_INSTALL=1 and toolchain missing)
  → Resolve real binary: ~/.rustup/toolchains/<toolchain>/bin/cargo
  → Prepare environment (see below)
  → exec the real tool binary (execve on Unix)
```

## Environment Set by Proxies

Before executing the real tool, the proxy sets these environment variables:

| Variable | Description |
|----------|-------------|
| `RUSTUP_TOOLCHAIN` | The name of the selected toolchain (e.g. `stable-x86_64-unknown-linux-gnu`) |
| `RUSTUP_HOME` | Path to rustup metadata directory |
| `CARGO_HOME` | Path to cargo's home directory (synchronized for consistency) |
| `RUST_RECURSION_COUNT` | Counter incremented to detect/prevent infinite proxy recursion |
| `RUSTUP_TOOLCHAIN_SOURCE` | (Unstable) indicates how `RUSTUP_TOOLCHAIN` was determined |
| `LD_LIBRARY_PATH` (Linux) | Toolchain's `lib` directory added to the dynamic linker path |
| `DYLD_FALLBACK_LIBRARY_PATH` (macOS) | Toolchain's `lib` directory added on macOS |

See [environment.md](environment.md) for the full environment variable reference.

## `+toolchain` Shorthand

Any proxied command accepts a `+<toolchain>` prefix on its first argument. The proxy strips it and uses it as the toolchain override (priority #1 in resolution — see [override.md](override.md)).

```
cargo +nightly build          # == rustup run nightly cargo build
rustc +stable --version
```

If the `cargo` in `$PATH` is **not** the rustup proxy (e.g. a system cargo), the `+toolchain` arg reaches the real tool and fails:

```
error: no such command: `+stable`
        Cargo does not handle `+toolchain` directives.
        Did you mean to invoke `cargo` through rustup` instead?
```

Fix: ensure `~/.cargo/bin` (the rustup proxies) is first in `$PATH`.

## `rustup run`

Run a command with a specific toolchain, overriding all other resolution mechanisms.

```
rustup run <toolchain> <command> [args...]
```

```
rustup run nightly rustc --version
rustup run nightly bash         # run a shell configured for nightly
```

| Flag | Description |
|------|-------------|
| `--install` | Install the toolchain if it is missing before running |

Without `--install`, if the toolchain is not installed the command fails — unless `RUSTUP_AUTO_INSTALL=1`, in which case it auto-installs. See [environment.md](environment.md).

`rustup run` is equivalent to the `+toolchain` shorthand on a proxied command:

```
rustup run nightly cargo build   ==   cargo +nightly build
```

## `rustup which`

Show the full path to the binary that a proxy resolves to for a given command.

```
rustup which cargo
# /home/user/.rustup/toolchains/stable-x86_64-unknown-linux-gnu/bin/cargo
```

| Flag | Description |
|------|-------------|
| `--toolchain <name>` | Show the binary for a specific toolchain |

## `rustup doc`

Open the offline Rust documentation (from the `rust-docs` component) in a web browser.

| Flag | Description |
|------|-------------|
| `--book` | Open *The Rust Programming Language* book |
| `--std` | Open the Standard Library API reference |
| `--toolchain <name>` | Use documentation from a specific toolchain |
| `--path` | Print the path to the documentation without opening it |

```
rustup doc --book
rustup doc --std
rustup doc --path            # print path only
```

## `rustup man`

*(Unix only)* View the man page for a given Rust tool command.

```
rustup man cargo
rustup man rustc
```

## Edge Cases

### Recursion protection

`RUST_RECURSION_COUNT` is incremented on each proxy invocation to prevent infinite loops. This could occur if a custom toolchain's `cargo` binary somehow invokes the `cargo` proxy again.

### Cargo fallback for custom toolchains

If a custom toolchain (created with `rustup toolchain link`) does not include `cargo`, the proxy falls back to `cargo` from a release channel, preferring **nightly**, then **beta**, then **stable**.

### Auto-install on first use

When a toolchain is selected but not installed (and `RUSTUP_AUTO_INSTALL=1`), rustup auto-installs it before dispatching. This applies to proxy invocations, `rustup run`, and `rust-toolchain.toml` discovery.

## Related

- [override.md](override.md) — full toolchain resolution precedence
- [environment.md](environment.md) — environment variables set by proxies
- [toolchain.md](toolchain.md) — custom toolchains and the cargo fallback
