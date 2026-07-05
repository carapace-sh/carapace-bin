# Build Profiles

Profiles control compiler settings (optimization, debug info, LTO, etc.) for each build. Cargo selects a profile based on the command and `--profile` flag, and applies it to every compilation unit in the build.

> **Source of truth**: <https://doc.rust-lang.org/cargo/reference/profiles.html>.

## Standard Profiles

| Profile | Default command(s) | Inherits from |
|---------|-------------------|---------------|
| `dev` | `cargo build`, `cargo run`, `cargo check`, `cargo rustc` | — (built-in) |
| `release` | `cargo build --release`, `cargo install` | — (built-in) |
| `test` | `cargo test` | `dev` |
| `bench` | `cargo bench` | `release` |
| `doc` | `cargo doc` | — (built-in, minimal) |

Select explicitly with `--profile <name>` or `--release` (equivalent to `--profile release`).

## Default Values

### `[profile.dev]`

| Setting | Default |
|---------|---------|
| `opt-level` | `0` |
| `debug` | `true` (≡ `2` / `"full"`) |
| `split-debuginfo` | platform-specific (`unpacked` on macOS) |
| `strip` | `"none"` |
| `debug-assertions` | `true` |
| `overflow-checks` | `true` |
| `lto` | `false` |
| `panic` | `"unwind"` |
| `incremental` | `true` |
| `codegen-units` | `256` |
| `rpath` | `false` |

### `[profile.release]`

| Setting | Default |
|---------|---------|
| `opt-level` | `3` |
| `debug` | `false` |
| `split-debuginfo` | platform-specific |
| `strip` | `"none"` |
| `debug-assertions` | `false` |
| `overflow-checks` | `false` |
| `lto` | `false` |
| `panic` | `"unwind"` |
| `incremental` | `false` |
| `codegen-units` | `16` |
| `rpath` | `false` |

### `[profile.test]`

Inherits all defaults from `dev`.

### `[profile.bench]`

Inherits all defaults from `release`.

## Profile Settings Reference

### `opt-level`

Controls `-C opt-level`. Values:

| Value | Effect |
|-------|--------|
| `0` | No optimizations |
| `1` | Basic optimizations |
| `2` | Some optimizations |
| `3` | All optimizations |
| `"s"` | Optimize for binary size |
| `"z"` | Optimize for size + disable loop vectorization |

> Level `3` can sometimes be slower than `2`. `"s"`/`"z"` aren't guaranteed smaller — benchmark.

### `debug`

Controls `-C debuginfo`:

| Value | Effect |
|-------|--------|
| `0`, `false`, `"none"` | No debug info (release default) |
| `"line-directives-only"` | Line directives only (profiling) |
| `"line-tables-only"` | Minimal: filename/line for backtraces |
| `1`, `"limited"` | Without type/variable info |
| `2`, `true`, `"full"` | Full debug info (dev default) |

### `split-debuginfo`

Controls `-C split-debuginfo`. Determines whether debug info is embedded in the executable or in a separate adjacent file. Default is `unpacked` on macOS (when debug enabled); otherwise platform-specific per rustc. Some options are nightly-only.

### `strip`

Controls `-C strip`:

| Value | Effect |
|-------|--------|
| `"none"` (default) | No stripping |
| `"debuginfo"` | Strip debug info |
| `"symbols"` | Strip symbols |
| `true` | ≡ `"symbols"` |
| `false` | ≡ `"none"` |

### `debug-assertions`

Controls `-C debug-assertions`. Enables `cfg(debug_assertions)` and the `debug_assert!` macro. Default `true` for `dev`/`test`, `false` for `release`/`bench`.

### `overflow-checks`

Controls `-C overflow-checks`. Causes a panic on integer overflow when enabled. Default `true` for `dev`/`test`, `false` for `release`/`bench`.

### `lto`

Controls `-C lto`, `-C linker-plugin-lto`, `-C embed-bitcode`:

| Value | Effect |
|-------|--------|
| `false` | "thin local LTO" across local codegen units only (no cross-crate LTO; off if `codegen-units=1` or `opt-level=0`) |
| `true` / `"fat"` | Full cross-crate LTO across the entire dependency graph |
| `"thin"` | Scalable incremental LTO (similar gains to fat, faster) |
| `"off"` | Disables LTO entirely |

### `panic`

Controls `-C panic`:

| Value | Effect |
|-------|--------|
| `"unwind"` (default) | Unwind stack on panic |
| `"abort"` | Terminate process on panic |

> **Tests, benchmarks, build scripts, and proc-macros ignore `panic`** — the test harness requires `unwind`. When using `abort` and building tests, all dependencies are forced to build with `unwind`.

### `incremental`

Controls `-C incremental`. Saves extra info to disk for faster recompiles. Only used for workspace members and path dependencies. Overridable via `CARGO_INCREMENTAL` env var or `build.incremental` config.

### `codegen-units`

Controls `-C codegen-units`. Splits a crate into N units for parallel codegen. Default `256` for incremental builds, `16` for non-incremental. Must be > 0. Lower values produce better-optimized code (at the cost of compile time).

### `rpath`

Controls `-C rpath`. Embeds rpath info in the binary for dynamic library lookup. Default `false`.

## Custom Profiles

Define a custom profile with `inherits`:

```toml
[profile.release-lto]
inherits = "release"
lto = true

[profile.dev-no-debug]
inherits = "dev"
debug = false
```

Use with `cargo build --profile release-lto`. The output goes to `target/release-lto/`.

## Profile Inheritance

- `test` inherits from `dev`
- `bench` inherits from `release`
- Custom profiles inherit from the profile named in `inherits`
- Unspecified settings fall back to the inherited profile, ultimately to built-in defaults

## Profile Overrides

### Package-Specific Overrides

Override settings for a specific dependency:

```toml
[profile.dev.package.foo]
opt-level = 3
```

Target a specific version:
```toml
[profile.dev.package."foo:2.1.0"]
opt-level = 3
```

Override all non-workspace dependencies:
```toml
[profile.dev.package."*"]
opt-level = 2
```

### Build Override

Override settings for build scripts, proc-macros, and their dependencies:

```toml
[profile.dev.build-override]
opt-level = 3
codegen-units = 256
debug = false
```

> If a dependency is used both as a normal and a build dependency, and `build-override` is set, Cargo may build it **twice** (once per role), increasing initial build time.

### Override Precedence (first match wins)

1. `[profile.dev.package.<name>]` — named package override
2. `[profile.dev.package."*"]` — any non-workspace member
3. `[profile.dev.build-override]` — build scripts, proc-macros, and their deps
4. `[profile.dev]` — settings in `Cargo.toml`
5. Built-in Cargo defaults

> `panic`, `lto`, and `rpath` **cannot** be overridden in package or build-override sections.

## Profile Settings Location

Profile settings in `Cargo.toml` are read **only from the workspace root manifest**. Dependency manifests' `[profile]` sections are ignored.

Profiles can also be overridden from config files and environment variables (see [config.md](config.md)), which take precedence over `Cargo.toml`.

## Profiles and the Target Directory

Each profile gets its own subdirectory:

| Profile | Directory |
|---------|-----------|
| `dev` | `target/debug/` |
| `release` | `target/release/` |
| `test` | `target/debug/` (inherits dev) |
| `bench` | `target/release/` (inherits release) |
| custom (e.g. `release-lto`) | `target/release-lto/` |

Incremental compilation data is stored per-profile.

## Generics and Optimization

Where generic code is **instantiated** determines which profile's optimization applies. Raising `opt-level` for a dependency that defines generic functions may not help if those generics are instantiated in your crate (where your crate's profile applies).

- **opt-level 2 or 3**: a crate will not share monomorphized generics from other crates
- **opt-level 1**: some optimization while still allowing monomorphized items to be shared

## The Unit Graph

Each `(package, target, profile)` triple is a **unit** in Cargo's build graph. Profile overrides can cause a single package to appear as multiple units (e.g., built once normally and once with build-override settings). See [concepts.md](concepts.md).

## References

- <https://doc.rust-lang.org/cargo/reference/profiles.html>

## Related

- [concepts.md](concepts.md) — the unit graph and build pipeline
- [config.md](config.md) — profile overrides via config and environment variables
- [commands.md](commands.md) — `--profile`, `--release` flags
