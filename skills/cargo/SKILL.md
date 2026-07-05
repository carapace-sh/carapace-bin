---
name: cargo
description: >
  Use when working with the cargo CLI — Rust's package manager. Covers the
  manifest format (Cargo.toml), dependencies, features, workspaces, build
  profiles, build scripts, configuration (.cargo/config.toml), registries and
  publishing, environment variables, the lockfile and cache, and the command
  reference. Triggers on: "cargo", "cargo build", "cargo check", "cargo test",
  "cargo run", "cargo add", "cargo update", "cargo publish", "cargo new",
  "cargo init", "cargo install", "cargo doc", "cargo bench", "cargo tree",
  "cargo metadata", "Cargo.toml", "Cargo.lock", "crate", "workspace",
  "feature", "profile", "build script", "build.rs", "registry", "crates.io",
  "CARGO_HOME", "rustc", "rustdoc", "rust-version", "edition".
user-invocable: true
---

# Cargo — Rust's Package Manager In-Depth Reference

Comprehensive reference for the cargo CLI (<https://doc.rust-lang.org/cargo/>), covering the manifest format, dependency specification, the features system and resolver, workspaces, build profiles, build scripts, configuration, registries and publishing, environment variables, the lockfile and cache layout, and the full command reference.

## Data Flow

```
Cargo.toml (manifest)
  → dependency resolver (registry index + Cargo.lock)
    → fetch (registry cache / git / path / vendor)
      → build scripts (build.rs, OUT_DIR)
        → rustc / rustdoc (per profile, per target)
          → target/<profile>/  (artifacts)
            → run / test / bench / install
```

## Sub-Resources

Load the reference that matches your task. When in doubt, load multiple references.

| Keywords | Reference |
|----------|----------|
| concept, package, crate, target, binary, library, example, test, bench, edition, registry, Cargo.toml, Cargo.lock, cargo home, build pipeline, crate types, cdylib, rlib, dylib, staticlib, proc-macro, target triple, host vs target, cross-compilation, manifest vs lockfile | [references/concepts.md](references/concepts.md) |
| manifest, Cargo.toml, `[package]`, name, version, edition, rust-version, authors, description, license, license-file, keywords, categories, readme, homepage, repository, documentation, metadata, publish, include, exclude, build, links, default-run, resolver, autolib, autobins, autoexamples, autotests, autobenches, `[lib]`, `[[bin]]`, `[[example]]`, `[[test]]`, `[[bench]]`, crate-type, required-features, doctest, `[lints]`, `[badges]`, cargo-features | [references/manifest.md](references/manifest.md) |
| dependency, [dependencies], [dev-dependencies], [build-dependencies], [target.*.dependencies], version requirement, SemVer, caret requirement, tilde requirement, wildcard, git dependency, branch, tag, rev, path dependency, registry dependency, optional dependency, default-features, features, rename, package, workspace dependency, dependency inheritance, platform-specific, cfg expression | [references/dependencies.md](references/dependencies.md) |
| feature, [features], default, dep: syntax, optional dependency as feature, package/feature, weak dependency, ?/ syntax, feature unification, additive, mutually exclusive, resolver v1, resolver v2, resolver v3, --features, --all-features, --no-default-features, required-features, CARGO_FEATURE, feature inspection, cargo tree -e features, SemVer compatibility | [references/features.md](references/features.md) |
| workspace, [workspace], members, exclude, virtual manifest, default-members, resolver, workspace.package, workspace.dependencies, workspace inheritance, workspace metadata, root package, shared target directory, path dependencies in workspace | [references/workspaces.md](references/workspaces.md) |
| profile, [profile.dev], [profile.release], [profile.test], [profile.bench], opt-level, debug, debug-assertions, overflow-checks, lto, panic, incremental, codegen-units, rpath, strip, split-debuginfo, inherits, custom profile, --profile, --release, package override, build-override, unit graph, target directory layout, monomorphization, generics | [references/profiles.md](references/profiles.md) |
| build script, build.rs, build dependency, [build-dependencies], links, *-sys, OUT_DIR, cargo::rustc-link-lib, cargo::rustc-link-search, cargo::rustc-flags, cargo::rustc-cfg, cargo::rustc-check-cfg, cargo::rustc-env, cargo::rerun-if-changed, cargo::rerun-if-env-changed, cargo::warning, cargo::error, cargo::metadata, DEP_, jobserver, code generation, override build script, build script env vars | [references/build-scripts.md](references/build-scripts.md) |
| config, configuration, .cargo/config.toml, config hierarchy, precedence, --config, environment variable mapping, CARGO_ prefix, [build], [target], [registry], [source], [net], [http], [term], [env], [alias], [cargo-new], [profile], [credential-alias], [future-incompat-report], [cache], credentials.toml, credential provider, include, rustflags, rustdocflags, rustc-wrapper, target-dir, jobs, proxy, timeout | [references/config.md](references/config.md) |
| registry, crates.io, index, git protocol, sparse protocol, sparse+, config.json, authentication, token, credential provider, cargo:token, cargo:wincred, cargo:macos-keychain, cargo:libsecret, cargo login, cargo logout, cargo owner, cargo yank, cargo publish, cargo search, cargo info, custom registry, alternate registry, source replacement, replace-with, [patch], [replace], local registry, directory source, vendor, cargo vendor, publish restriction | [references/registries.md](references/registries.md) |
| command, subcommand, flag, cargo build, cargo check, cargo test, cargo bench, cargo run, cargo doc, cargo clean, cargo fix, cargo rustc, cargo rustdoc, cargo add, cargo remove, cargo update, cargo tree, cargo metadata, cargo vendor, cargo fetch, cargo new, cargo init, cargo install, cargo uninstall, cargo search, cargo info, cargo login, cargo logout, cargo owner, cargo yank, cargo publish, cargo package, cargo pkgid, cargo locate-project, cargo generate-lockfile, cargo report, cargo config, cargo version, common flags, --manifest-path, --locked, --frozen, --offline, -Z, pkgid spec, message-format, timings | [references/commands.md](references/commands.md) |
| environment variable, CARGO_HOME, CARGO_BIN_NAME, CARGO_BUILD_*, CARGO_CFG_TARGET_*, CARGO_CRATE_NAME, CARGO_ENCODED_RUSTFLAGS, CARGO_ENCODED_RUSTDOCFLAGS, CARGO_INCREMENTAL, CARGO_MANIFEST_DIR, CARGO_MANIFEST_LINKS, CARGO_NET_*, CARGO_PKG_*, CARGO_REGISTRY_*, CARGO_TARGET_DIR, CARGO_TARGET_TMPDIR, CARGO_TERM_*, CARGO_LOG, CARGO_MAKEFLAGS, RUSTFLAGS, RUSTDOCFLAGS, RUSTC, RUSTDOC, RUSTC_WRAPPER, [env] config, force, relative | [references/environment.md](references/environment.md) |
| lockfile, Cargo.lock, lockfile format, package ID, source id, checksum, lockfile version, v3, v4, commit, generate-lockfile, update, precise, dry-run, breaking, cache, registry cache, git cache, ~/.cargo, target directory, target/debug, target/release, fingerprint, incremental, auto-clean | [references/lockfile.md](references/lockfile.md) |

## Quick Guide

- **What are cargo's core concepts?** → [references/concepts.md](references/concepts.md)
- **How do I write Cargo.toml?** → [references/manifest.md](references/manifest.md)
- **How do I specify dependencies?** → [references/dependencies.md](references/dependencies.md)
- **How do features and the resolver work?** → [references/features.md](references/features.md)
- **How do workspaces work?** → [references/workspaces.md](references/workspaces.md)
- **How do build profiles work?** → [references/profiles.md](references/profiles.md)
- **How do I write a build script?** → [references/build-scripts.md](references/build-scripts.md)
- **How does cargo configuration work?** → [references/config.md](references/config.md)
- **How do registries and publishing work?** → [references/registries.md](references/registries.md)
- **What commands does cargo have?** → [references/commands.md](references/commands.md)
- **What environment variables does cargo use?** → [references/environment.md](references/environment.md)
- **How does Cargo.lock work?** → [references/lockfile.md](references/lockfile.md)
- **How do I add a dependency?** → [references/dependencies.md](references/dependencies.md) and [references/commands.md](references/commands.md)
- **How do I enable features?** → [references/features.md](references/features.md) and [references/commands.md](references/commands.md)
- **How do I configure rustc flags?** → [references/config.md](references/config.md) and [references/environment.md](references/environment.md)
- **How do I publish to crates.io?** → [references/registries.md](references/registries.md)
- **How do I vendor dependencies?** → [references/registries.md](references/registries.md) and [references/commands.md](references/commands.md)
- **How do I override a dependency?** → [references/registries.md](references/registries.md) (`[patch]`, source replacement)
- **How do I inspect the dependency graph?** → [references/commands.md](references/commands.md) (`cargo tree`, `cargo metadata`)
- **How do I set up a custom registry?** → [references/config.md](references/config.md) and [references/registries.md](references/registries.md)
- **How do I do reproducible builds?** → [references/lockfile.md](references/lockfile.md) and [references/commands.md](references/commands.md) (`--locked`, `--frozen`)

## Cross-Project References

- For the **rustc** and **rustdoc** compilers themselves (flag reference, codegen options, lint catalog, target specs), see the official [rustc book](https://doc.rust-lang.org/rustc/) and [rustdoc book](https://doc.rust-lang.org/rustdoc/) — this skill covers only how cargo *invokes* them.
- For the **Rust language** (editions, attributes, `cfg` expressions, macros, the reference), see the [Rust Reference](https://doc.rust-lang.org/reference/).
- For carapace-specific cargo completion (the `cargo` completer, macros like `tools.cargo.*`), see the **carapace** skill (in this repo) → `references/spec.md` within that skill.
