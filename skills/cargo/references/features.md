# Features and the Resolver

Cargo's **features** mechanism provides conditional compilation and optional dependencies. The **resolver** selects which versions and features are activated across the dependency graph. This document covers the `[features]` table, feature unification, the resolver versions (v1/v2/v3), and command-line feature flags.

> **Source of truth**: <https://doc.rust-lang.org/cargo/reference/features.html> and <https://doc.rust-lang.org/cargo/reference/resolver.html>.

## What Features Are

A **feature** is a named flag defined in `[features]`. Cargo passes enabled features to rustc via `--cfg feature="<name>"`, and code tests for them:

```rust
#[cfg(feature = "webp")]
mod webp;
```

Features are **additive** — enabling a feature should only add functionality, never remove it. This is critical because of feature unification (see below). Any combination of features must be safe to enable together.

## The `[features]` Table

```toml
[features]
default = ["ico", "webp"]
webp = []
bmp = []
png = []
ico = ["bmp", "png"]        # enabling "ico" also enables "bmp" and "png"
parallel = ["jpeg-decoder/rayon"]   # enables a feature on a dependency
avif = ["dep:ravif", "dep:rgb"]     # enables optional dependencies
```

Each feature maps to an **array of other features or optional dependency references** that it activates. A feature with `[]` enables nothing else — it's just a flag for `#[cfg]`.

### Feature Name Rules

- Characters: Unicode XID (most letters), may start with `_` or digits `0`-`9`
- After the first character: may also contain `-`, `+`, `.`
- On crates.io: ASCII alphanumeric, `_`, `-`, `+` only; max 300 features per crate

## The `default` Feature

By default all features are disabled, **except** those listed in `default`:

```toml
[features]
default = ["std"]
std = []
```

`default` is enabled unless the user opts out:
- CLI: `--no-default-features`
- Dependency declaration: `default-features = false`

> **Warning**: Removing a feature from the `default` list is a SemVer-incompatible change for downstream consumers.

## Optional Dependencies as Features

An `optional = true` dependency implicitly creates a feature of the same name:

```toml
[dependencies]
gif = { version = "0.11", optional = true }
# implicitly: features.gif = ["dep:gif"]
```

### The `dep:` Syntax (Rust 1.60+)

The `dep:` prefix explicitly refers to an optional dependency **without** creating an implicit feature. This lets you group optional deps under a custom feature name:

```toml
[dependencies]
ravif = { version = "0.6", optional = true }
rgb = { version = "0.8", optional = true }

[features]
avif = ["dep:ravif", "dep:rgb"]   # no implicit "ravif"/"rgb" features created
```

Without `dep:`, the optional dep key in `[features]` would also enable the dependency. With `dep:`, you control exactly when the dependency is enabled.

### The `?` Weak Dependency Syntax (Rust 1.60+)

`"package-name/feature-name"` enables a feature on a dependency **and** enables the dependency if it's optional. Adding `?` after the package name means "only if something else already enabled it":

```toml
[dependencies]
serde = { version = "1.0", optional = true }
rgb = { version = "0.8", optional = true }

[features]
serde = ["dep:serde", "rgb?/serde"]
# "rgb?/serde" enables serde on rgb ONLY if rgb is already enabled by something else
```

Without `?`, `"rgb/serde"` would force-enable `rgb` whenever `serde` feature is on, which is often undesirable.

## Enabling Features on Dependencies

In a dependency declaration:

```toml
[dependencies]
serde = { version = "1.0", features = ["derive"] }
flate2 = { version = "1.0", default-features = false, features = ["zlib-rs"] }
```

Or via the `[features]` table with `package/feature` syntax:

```toml
[dependencies]
jpeg-decoder = { version = "0.1", default-features = false }

[features]
parallel = ["jpeg-decoder/rayon"]
```

## Feature Unification

When multiple packages depend on the same crate, Cargo builds that crate **once** with the **union** of all features requested across all dependents.

Example: if `foo` enables `winapi`'s `fileapi` and `bar` enables `winapi`'s `std`, then `winapi` is built with `fileapi + std`. There is only one copy of `winapi` in the build.

**Why features must be additive**: because unification means any feature could be enabled alongside any other. A feature that disables functionality would break other dependents that rely on it.

### Mutually Exclusive Features

Avoid them. If unavoidable, use a compile error:

```rust
#[cfg(all(feature = "foo", feature = "bar"))]
compile_error!("features \"foo\" and \"bar\" cannot be enabled at the same time");
```

Consider using a single feature with an enum-like value pattern instead (though Cargo features are boolean, not valued).

### Inspecting Resolved Features

```sh
cargo tree -e features              # show features in the dep graph
cargo tree -f "{p} {f}"             # compact: package + comma-separated features
cargo tree -e features -i foo       # inverted: what enables features on foo
cargo tree --duplicates             # find crates built with multiple feature sets
```

## The Resolver

The resolver selects one version per package and one feature set, given all the version requirements and feature requests in the graph. It minimizes the number of crates built (preferring fewer versions).

### Resolver Versions

| Version | Default for | Behavior |
|---------|-------------|----------|
| **v1** | edition < 2021 | Unifies features across all dependency kinds; `--features` only works for the current directory's package |
| **v2** | edition >= 2021 | Avoids feature unification in specific cases (below); `--features` works for any `-p`/`--workspace` selected package |
| **v3** | opt-in (`resolver = "3"`) | Like v2 plus MSRV-aware selection (uses `rust-version` to prefer compatible versions) |

Set the resolver in `[package]` or `[workspace]`:

```toml
[package]
resolver = "2"
```

### Resolver v2 Behavior

v2 avoids unifying features in three situations:

1. **Platform-specific dependencies** — features on deps for targets not being built are ignored
2. **Build-dependencies and proc-macros** — do not share features with the same crate used as a normal dependency
3. **Dev-dependencies** — do not activate features unless building a test/example/bench target

This prevents a build-dependency from accidentally pulling in `std` on a crate that is otherwise used in a `no_std` context.

**Tradeoff**: v2 can cause the same crate to be built multiple times with different feature sets, increasing build times. Use `cargo tree --duplicates` to detect this.

### Resolver v2 Command-Line Changes

| Flag | v1 | v2 |
|------|----|----|
| `--features` | current dir's package only | any package selected with `-p` or `--workspace` |
| `--no-default-features` | current dir's package only | all workspace members |

```sh
# v2: enable features on multiple packages without package/ prefix
cargo build -p foo -p bar --features feat1,feat2
```

### MSRV-Aware Resolution (v3 / `resolver.incompatible-rust-versions`)

```toml
# Cargo.toml
[workspace]
resolver = "3"
```

Or via config:
```toml
[resolver]
incompatible-rust-versions = "fallback"   # "allow" (default) or "fallback"
```

When `fallback`, the resolver prefers versions whose `rust-version` is compatible with the current toolchain, falling back to newer versions only if no compatible one exists. Supported as of Rust 1.84.

## Command-Line Feature Flags

| Flag | Description |
|------|-------------|
| `--features <FEATURES>` | Enable features (comma or space separated; `package/feature` syntax allowed). Quote if spaces: `--features "foo bar"`. |
| `--all-features` | Enable all features of all selected packages. |
| `--no-default-features` | Do not enable the `default` feature. |
| `-F <FEATURES>` | Short form of `--features`. |

### `--features` Syntax

```sh
cargo build --features foo
cargo build --features "foo bar"          # space-separated
cargo build --features foo,bar            # comma-separated
cargo build --features foo/bar            # feature on a dependency
cargo build -p foo --features feat1       # resolver v2: feature on selected package
cargo build --all-features
cargo build --no-default-features --features std
```

## `required-features` on Targets

A build target (binary, example, test, bench) can declare features that must be enabled for it to be built:

```toml
[[bin]]
name = "my-binary"
required-features = ["webp"]
```

When `webp` is not enabled, `cargo build` skips this target.

## Build Script Feature Detection

Build scripts receive enabled features as environment variables: `CARGO_FEATURE_<NAME>` (uppercased, `-` → `_`). See [build-scripts.md](build-scripts.md).

## SemVer Compatibility for Features

| Safe in a minor release | Unsafe (breaking) |
|------------------------|-------------------|
| Add a new feature or optional dependency | Remove a feature or optional dependency |
| Add a feature to a dependency's feature list | Move existing public code behind a feature |
| — | Remove a feature from a feature list (e.g. from `default`) |

## References

- <https://doc.rust-lang.org/cargo/reference/features.html>
- <https://doc.rust-lang.org/cargo/reference/resolver.html>
- <https://doc.rust-lang.org/cargo/reference/semver.html>

## Related

- [dependencies.md](dependencies.md) — optional dependencies and version requirements
- [workspaces.md](workspaces.md) — workspace-level resolver setting
- [commands.md](commands.md) — `--features`, `--all-features`, `--no-default-features` flags
- [build-scripts.md](build-scripts.md) — `CARGO_FEATURE_*` env vars
