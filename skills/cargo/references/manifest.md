# The Manifest Format (Cargo.toml)

The `Cargo.toml` file is the **manifest** — a TOML file that declares a package's identity, dependencies, targets, features, profiles, and metadata. Cargo reads it to build the package. For the conceptual model, see [concepts.md](concepts.md).

> **Source of truth**: <https://doc.rust-lang.org/cargo/reference/manifest.html>.

## Top-Level Sections

| Section | Description |
|---------|-------------|
| `cargo-features` | Unstable, nightly-only features (e.g. `cargo-features = ["codegen-backend"]`) |
| `[package]` | Package identity and metadata |
| `[lib]` | Library target (singular — one library per package) |
| `[[bin]]` | Binary target (array — multiple binaries allowed) |
| `[[example]]` | Example target (array) |
| `[[test]]` | Integration test target (array) |
| `[[bench]]` | Benchmark target (array) |
| `[dependencies]` | Library dependencies — see [dependencies.md](dependencies.md) |
| `[dev-dependencies]` | Dependencies for examples, tests, benchmarks |
| `[build-dependencies]` | Dependencies for the build script — see [build-scripts.md](build-scripts.md) |
| `[target]` | Platform-specific dependencies (e.g. `[target.'cfg(unix)'.dependencies]`) |
| `[features]` | Conditional compilation features — see [features.md](features.md) |
| `[lints]` | Lint level configuration (respected as of Rust 1.74) |
| `[hints]` | Compilation hints (respected as of Rust 1.90; no stable hints yet) |
| `[badges]` | Registry status badges (deprecated for crates.io display) |
| `[patch]` | Override dependencies — see [registries.md](registries.md) |
| `[replace]` | Override dependencies (**deprecated** — use `[patch]`) |
| `[profile]` | Build profile settings — see [profiles.md](profiles.md) |
| `[workspace]` | Workspace definition — see [workspaces.md](workspaces.md) |

## The `[package]` Section

The only field Cargo strictly requires is `name`. Publishing to a registry requires more.

### Core Identity

| Field | Type | Default | Notes |
|-------|------|---------|-------|
| `name` | string | — | **Required.** Alphanumeric, `-`, `_`. On crates.io: ASCII only, max 64 chars, no reserved names (no Windows special names like `con`, `aux`). |
| `version` | string | `"0.0.0"` | SemVer: `MAJOR.MINOR.PATCH` with optional pre-release (`-alpha.1`) and build metadata (`+21AF26D3`). Required for publishing. |
| `edition` | string | `"2015"` (if absent) | `"2015"`, `"2018"`, `"2021"`, `"2024"`. Affects all crates in the package. See [concepts.md](concepts.md). |
| `rust-version` | string | — | Minimum supported Rust toolchain (MSRV), e.g. `"1.70.0"`. Enforced at build (errors or warns if toolchain is older). |
| `authors` | string[] | — | **Deprecated.** Optional email in angled brackets: `"Name <email>"`. |

### Metadata (required by crates.io)

| Field | Type | Notes |
|-------|------|-------|
| `description` | string | Short blurb (plain text, not Markdown). Required by crates.io. |
| `license` | string | SPDX 2.3 expression. Supports `AND`, `OR`, `WITH`. Must be from SPDX list 3.20. Either `license` or `license-file` required for crates.io. |
| `license-file` | string (path) | Path to a file containing the license text. |
| `readme` | string/bool | Path to README. Auto-detects `README.md`, `README.txt`, `README`. `false` disables auto-detect; `true` assumes `README.md`. |
| `homepage` | string (URL) | Package homepage. Should differ from `documentation`/`repository`. |
| `repository` | string (URL) | Source repository URL. |
| `documentation` | string (URL) | Docs URL. If absent, crates.io links to docs.rs. |
| `keywords` | string[] | Search keywords. crates.io: max 5, max 20 chars each, ASCII, start alphanumeric, only `_`, `-`, `+`. |
| `categories` | string[] | crates.io categories. max 5, must match the official category slug list. |
| `metadata` | table | Arbitrary data for external tools; ignored by Cargo. |

### Build and Linking

| Field | Type | Default | Notes |
|-------|------|---------|-------|
| `build` | string/bool | `"build.rs"` | Path to the build script. `build = false` disables auto-detection. See [build-scripts.md](build-scripts.md). |
| `links` | string | — | Name of the native library this package links (e.g. `"git2"`). Only one package per `links` value. Requires a build script. Enables `DEP_<LINKS>_*` propagation. |

### Publishing and File Selection

| Field | Type | Notes |
|-------|------|-------|
| `publish` | bool/string[] | `false` disables publishing. Array restricts to named registries. |
| `exclude` | string[] | gitignore-style patterns to exclude when packaging. |
| `include` | string[] | gitignore-style patterns to include. Mutually exclusive with `exclude`; use `!` for exclusions within includes. |

**Always excluded**: sub-packages (subdirs with their own `Cargo.toml`), the `target` directory.
**Always included**: `Cargo.toml`, minimized `Cargo.lock`, and `license-file` if specified.

### include/exclude Pattern Syntax (gitignore-style)

| Pattern | Matches |
|---------|--------|
| `foo` | `foo` anywhere (≡ `**/foo`) |
| `/foo` | `foo` only at package root |
| `foo/` | a directory named `foo` anywhere |
| `*` | zero+ chars except `/` |
| `?` | one char except `/` |
| `[ab]` / `[a-z]` | character ranges |
| `**/foo/bar` | `bar` under any `foo` |
| `foo/**` | everything inside `foo` |
| `a/**/b` | zero+ dirs between `a` and `b` |
| `!pattern` | negation |

### Workspace and Misc

| Field | Type | Notes |
|-------|------|-------|
| `workspace` | string (path) | Path to workspace root `Cargo.toml`. Cannot coexist with `[workspace]`. See [workspaces.md](workspaces.md). |
| `default-run` | string | Default binary for `cargo run` when multiple exist. |
| `resolver` | string | Resolver version: `"1"`, `"2"`, `"3"`. See [features.md](features.md). |
| `autolib` | bool | Disable library auto-discovery. |
| `autobins` | bool | Disable binary auto-discovery. |
| `autoexamples` | bool | Disable example auto-discovery. |
| `autotests` | bool | Disable test auto-discovery. |
| `autobenches` | bool | Disable benchmark auto-discovery. |

## Target Tables

Configure specific build targets. See [concepts.md](concepts.md) for the target model.

### `[lib]` — Library (singular)

```toml
[lib]
name = "foo"
path = "src/lib.rs"
crate-type = ["rlib", "cdylib"]
```

### `[[bin]]` — Binary (array)

```toml
[[bin]]
name = "my-bin"
path = "src/bin/my-bin.rs"
required-features = ["webp"]
```

### Common target fields

| Field | Type | Default | Notes |
|-------|------|---------|-------|
| `name` | string | from filename / package name | Target name |
| `path` | string | convention-based | Source file path |
| `test` | bool | `true` | Whether `cargo test` includes this target |
| `doctest` | bool | `true` (lib only) | Whether to run doctests |
| `bench` | bool | `true` | Whether `cargo bench` includes this target |
| `doc` | bool | `true` | Whether `cargo doc` includes this target |
| `plugin` | bool | `false` | **Deprecated and unused.** Superseded by `proc-macro`. |
| `proc-macro` | bool | `false` | Procedural macro crate — see [concepts.md](concepts.md) |
| `harness` | bool | `true` | Whether to use the libtest harness. `false` → you define `main()` and use a custom test runner. |
| `edition` | string | package edition | Per-target edition override |
| `crate-type` | string[] | `["bin"]` or `["lib"]` | rustc `--crate-type` — see [concepts.md](concepts.md) |
| `required-features` | string[] | — | Features that must be enabled to build this target — see [features.md](features.md) |

## `[lints]` (Rust 1.74+)

Override lint levels from the manifest:

```toml
[lints.rust]
unsafe_code = "forbid"
missing_docs = "warn"

[lints.clippy]
enum_glob_use = "deny"
```

Full form with priority:

```toml
[lints.rust]
unsafe_code = { level = "forbid", priority = 0 }
```

- **Levels**: `forbid`, `deny`, `warn`, `allow`
- **Priority**: signed int; higher numbers override lower. Default `0`.
- **Tool grouping**: the part before `::` in a lint name selects the group (`rust`, `clippy`, `rustdoc`...). Falls back to `rust` if no `::`.
- Can be inherited from the workspace: `lints.workspace = true`.

## `[badges]` (deprecated for crates.io)

```toml
[badges]
maintenance = { status = "actively-developed" }
```

Status values: `actively-developed`, `passively-maintained`, `as-is`, `experimental`, `looking-for-maintainer`, `deprecated`, `none`. crates.io no longer displays these; put badges in your README instead.

## Virtual Manifest

A manifest with `[workspace]` but no `[package]` is a **virtual manifest** — it exists only to define a workspace root. See [workspaces.md](workspaces.md).

## Minimal Example

```toml
[package]
name = "my-crate"
version = "0.1.0"
edition = "2021"
rust-version = "1.74"
license = "MIT OR Apache-2.0"
description = "A short description"
repository = "https://github.com/user/my-crate"

[dependencies]
serde = { version = "1.0", features = ["derive"] }

[dev-dependencies]
proptest = "1.0"

[features]
default = ["std"]
std = []

[[bin]]
name = "my-bin"
required-features = ["std"]
```

## References

- <https://doc.rust-lang.org/cargo/reference/manifest.html>
- <https://doc.rust-lang.org/cargo/reference/cargo-targets.html>
- <https://doc.rust-lang.org/cargo/reference/lints.html>

## Related

- [dependencies.md](dependencies.md) — dependency specification
- [features.md](features.md) — the `[features]` table
- [workspaces.md](workspaces.md) — the `[workspace]` table
- [profiles.md](profiles.md) — the `[profile.*]` tables
- [build-scripts.md](build-scripts.md) — the `build`/`links` fields
