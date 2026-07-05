---
name: rustup
description: >
  Use when working with the rustup CLI — the Rust toolchain multiplexer. Covers
  installing/updating/uninstalling toolchains, channel syntax (stable/beta/nightly,
  dated, versioned), cross-compilation targets, components and profiles, directory
  and file overrides (rust-toolchain.toml), the proxy/shim dispatch system
  (cargo/rustc/rustdoc/rustfmt/rust-analyzer), `rustup run`, `+toolchain` shorthand,
  self-management, shell completions, and environment variables.
  Triggers on: "rustup", "rustup install", "rustup update", "rustup toolchain",
  "rustup target", "rustup component", "rustup override", "rustup run", "rustup which",
  "rustup doc", "rustup man", "rustup self", "rustup show", "rustup set", "rustup home",
  "rustup completions", "rustup check", "rustup default", "+toolchain",
  "rust-toolchain.toml", "RUSTUP_HOME", "RUSTUP_TOOLCHAIN", "RUSTUP_AUTO_INSTALL",
  "rustup profile", "rustup proxy", "rust-analyzer proxy", "cargo-miri".
user-invocable: true
---

# rustup In-Depth Reference

Reference for [rustup](https://rust-lang.github.io/rustup/) — the Rust toolchain installer and multiplexer. Covers the full CLI surface, the proxy dispatch system, toolchain resolution, cross-compilation, and configuration.

## Data Flow

```
Command line
  → rustup binary (multi-call: arg0 selects mode)
    → rustup_mode    : rustup <subcommand> ...   (management)
    → proxy_mode     : cargo / rustc / ...        (tool dispatch)
    → setup_mode     : rustup-init                (installer)
      → toolchain resolution
        1. +toolchain shorthand
        2. RUSTUP_TOOLCHAIN env var
        3. rustup override (dir override, closest wins)
        4. rust-toolchain.toml / rust-toolchain file (closest wins)
        5. default toolchain
      → install-on-run (if RUSTUP_AUTO_INSTALL=1)
      → exec real toolchain binary with configured env
```

## Sub-Resources

Load the reference that matches your task. When in doubt, load multiple references.

| Keywords | Reference |
|----------|----------|
| install, uninstall, update, default, check, show, set, home, completions, help, global flags, --verbose, --quiet, --version, --no-self-update, top-level commands | [references/cli.md](references/cli.md) |
| toolchain, install, uninstall, list, link, custom toolchain, channel, stable, beta, nightly, versioned, dated, target triple, host, spec syntax, toolchain name | [references/toolchain.md](references/toolchain.md) |
| target, cross-compilation, target add, target list, target remove, rust-std, std library, linker, --target, --toolchain, --installed | [references/target.md](references/target.md) |
| component, add, remove, list, rustc, cargo, rustfmt, rust-docs, rust-analyzer, clippy, miri, rust-src, llvm-tools, rls, --target, --toolchain, profile, minimal, default, complete | [references/component.md](references/component.md) |
| override, override set, override unset, override list, rust-toolchain.toml, rust-toolchain file, toolchain resolution, precedence, RUSTUP_TOOLCHAIN, +toolchain, settings.toml, directory override | [references/override.md](references/override.md) |
| proxy, shim, cargo, rustc, rustdoc, rustfmt, rust-analyzer, clippy-driver, cargo-clippy, cargo-miri, rls, rust-gdb, rust-lldb, multi-call binary, rustup run, +toolchain shorthand, which, doc, man, install-on-run, recursion protection | [references/proxy.md](references/proxy.md) |
| environment variables, RUSTUP_HOME, CARGO_HOME, RUSTUP_TOOLCHAIN, RUSTUP_DIST_SERVER, RUSTUP_UPDATE_ROOT, RUSTUP_AUTO_INSTALL, RUSTUP_LOG, RUSTUP_TERM_COLOR, RUSTUP_IO_THREADS, settings.toml, configuration | [references/environment.md](references/environment.md) |

## Quick Guide

- **How do I install a toolchain?** → [references/toolchain.md](references/toolchain.md)
- **How do I switch the default toolchain?** → [references/cli.md](references/cli.md)
- **How do I cross-compile to another target?** → [references/target.md](references/target.md)
- **How do I add clippy/rustfmt/rust-analyzer?** → [references/component.md](references/component.md)
- **How do I pin a toolchain per project?** → [references/override.md](references/override.md)
- **How does `cargo +nightly build` work?** → [references/proxy.md](references/proxy.md)
- **How does rustup decide which toolchain to use?** → [references/override.md](references/override.md)
- **How do I run a command with a specific toolchain?** → [references/proxy.md](references/proxy.md)
- **What components and profiles are available?** → [references/component.md](references/component.md)
- **What environment variables does rustup use?** → [references/environment.md](references/environment.md)
- **How do I generate shell completions for rustup?** → [references/cli.md](references/cli.md)
- **How do I update or uninstall rustup itself?** → [references/cli.md](references/cli.md)

## Cross-Project References

- For **cargo** (the package manager/build tool proxied by rustup), see cargo's own documentation. This skill covers only rustup's proxy dispatch and environment setup for cargo, not cargo's command surface.
- For **rust-analyzer** as an LSP server (capabilities, configuration, extensions), see the rust-analyzer documentation. This skill covers only the rustup component/proxy aspect.
