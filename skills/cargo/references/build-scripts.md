# Build Scripts

Build scripts (`build.rs`) are custom Rust programs Cargo compiles and runs **just before** building a package. They handle linking to native libraries, generating code, and platform-specific configuration.

> **Source of truth**: <https://doc.rust-lang.org/cargo/reference/build-scripts.html>.

## When to Use a Build Script

- Compile or link to C/C++/other non-Rust libraries
- Find system libraries via pkg-config or probing
- Generate Rust source code before compilation (parsers, protobufs, bindings)
- Set platform-specific `cfg` flags or environment variables for the crate

Cargo does not aim to replace specialized build tools (make, cmake) but integrates with them.

## Life Cycle

1. **Compile**: Cargo compiles `build.rs` (using the host toolchain, with `[build-dependencies]`).
2. **Run**: Cargo executes the compiled build script. The script communicates by printing `cargo::<directive>` lines to stdout.
3. **Rebuild triggers**: If no `rerun-if` directives are emitted, Cargo re-runs the script whenever **any file in the package** changes. Use `rerun-if` to narrow this.
4. **Completion**: Exit code `0` → proceed to build the package. Non-zero → halt with error.

The build script's current directory is the package root (where `Cargo.toml` lives).

## Declaring a Build Script

By default, Cargo looks for `build.rs` at the package root:

```toml
[package]
build = "build.rs"   # default; can be a custom path
```

Disable with `build = false`.

## Build Dependencies

Build scripts can depend on other crates via `[build-dependencies]`:

```toml
[build-dependencies]
cc = "1.0"
pkg-config = "0.3"
```

- Build scripts **cannot** access `[dependencies]` or `[dev-dependencies]` (not built yet)
- `[build-dependencies]` are **not** available to the package itself unless also listed in `[dependencies]`
- Build dependencies are compiled for the **host** platform (even when cross-compiling)

## Build Script Inputs (Environment Variables)

| Variable | Description |
|----------|-------------|
| `CARGO_MANIFEST_DIR` | Directory containing `Cargo.toml` |
| `CARGO_PKG_NAME` | Package name |
| `CARGO_PKG_VERSION` | Full version (e.g. `"1.2.3"`) |
| `CARGO_PKG_VERSION_MAJOR` / `_MINOR` / `_PATCH` / `_PRE` | Version components |
| `CARGO_PKG_AUTHORS` | Colon-separated authors |
| `CARGO_PKG_DESCRIPTION` / `_HOMEPAGE` / `_REPOSITORY` / `_LICENSE` | Metadata |
| `CARGO_CFG_TARGET_OS` / `_ARCH` / `_ENV` / `_FAMILY` / `_VENDOR` / `_POINTER_WIDTH` / `_ENDIAN` | Target platform cfg values |
| `CARGO_CFG_UNIX` / `CARGO_CFG_WINDOWS` | Set on unix-like / windows-like targets |
| `CARGO_FEATURE_<NAME>` | Set for each enabled feature (uppercased, `-`→`_`) |
| `OUT_DIR` | Where to write generated artifacts |
| `HOST` | Host platform triple |
| `TARGET` | Target platform triple |
| `NUM_JOBS` | Max parallel jobs |
| `OPT_LEVEL` | Optimization level (`0`-`3`, `s`, `z`) |
| `DEBUG` | `"true"` or `"false"` |
| `PROFILE` | Profile name (`debug`, `release`) |
| `CARGO_ENCODED_RUSTFLAGS` | rustc flags separated by `\x1f` |

> **Critical**: Do **not** use `cfg!()` or `#[cfg]` in build scripts — those check the **host** platform, not the target. Read `CARGO_CFG_*` env vars instead.

### `OUT_DIR`

Cargo does **not** clean `OUT_DIR` between builds. Contents may persist across rebuilds for incremental support. Don't assume it's empty.

## Build Script Outputs (Directives)

The script communicates with Cargo by printing lines starting with `cargo::` to stdout. Other output is ignored.

> **MSRV note**: The `cargo::` (double-colon) syntax requires Rust 1.77. For older versions, use `cargo:KEY=VALUE` (single colon).

Output is hidden during normal builds. Use `-vv` to see it. All stdout is saved to `target/debug/build/<pkg>/output`.

### Linking

| Directive | Effect |
|-----------|--------|
| `cargo::rustc-link-lib=[KIND[:MODIFIERS]=]NAME[:RENAME]` | Pass `-l` to rustc. `KIND`: `dylib`, `static`, `framework`. Only passed to the library target (or all targets if no library). |
| `cargo::rustc-link-search=[KIND=]PATH` | Pass `-L` to rustc. `KIND`: `dependency`, `crate`, `native`, `framework`, `all`. |
| `cargo::rustc-flags=FLAGS` | Pass space-separated `-l`/`-L` flags (equivalent to the above). |
| `cargo::rustc-link-arg=FLAG` | Pass `-C link-arg=FLAG` to supported targets (bins, cdylib, examples, tests, benches). |
| `cargo::rustc-link-arg-bin=BIN=FLAG` | Pass link-arg only when building binary `BIN`. |
| `cargo::rustc-link-arg-bins=FLAG` | Pass link-arg to all binaries. |
| `cargo::rustc-link-arg-tests=FLAG` | Pass link-arg to tests. |
| `cargo::rustc-link-arg-examples=FLAG` | Pass link-arg to examples. |
| `cargo::rustc-link-arg-benches=FLAG` | Pass link-arg to benchmarks. |
| `cargo::rustc-cdylib-link-arg=FLAG` | Pass link-arg only for `cdylib` targets. |

### Compilation

| Directive | Effect |
|-----------|--------|
| `cargo::rustc-cfg=KEY[="VALUE"]` | Pass `--cfg` to rustc. Bare (`cargo::rustc-cfg=abc` → `#[cfg(abc)]`) or key/value (`key="value"`). Must be registered with `rustc-check-cfg` or `unexpected_cfgs` lint allowed. |
| `cargo::rustc-check-cfg=CHECK_CFG` | Register expected custom cfgs for the `unexpected_cfgs` lint (MSRV 1.80). Syntax mirrors rustc's `--check-cfg`. |
| `cargo::rustc-env=VAR=VALUE` | Set an env var at compile time, retrievable via `env!()` in the crate. |

### Rebuild Triggers

| Directive | Effect |
|-----------|--------|
| `cargo::rerun-if-changed=PATH` | Re-run the script only if `PATH` (file or directory) changes (mtime-based). Without any `rerun-if`, Cargo scans the entire package. |
| `cargo::rerun-if-env-changed=NAME` | Re-run if the named **global** env var changes (e.g. `CC`). Not for Cargo-set vars like `TARGET`. As of Rust 1.46, `env!()`/`option_env!()` in source auto-detects — this is only needed for vars not referenced by those macros. |

### Communication

| Directive | Effect |
|-----------|--------|
| `cargo::warning=MESSAGE` | Display a warning after the script finishes. Only shown for path dependencies by default; use `-vv` for all. |
| `cargo::error=MESSAGE` | Display an error and **fail the build** (MSRV 1.84). |
| `cargo::metadata=KEY=VALUE` | Set metadata for **dependent packages'** build scripts (used with `links`). |

## The `links` Field

```toml
[package]
links = "foo"
```

Declares the package links to native library `libfoo`.

- Only **one** package per `links` value (prevents duplicate symbols)
- A package with `links` **must** have a build script
- The build script should use `cargo::rustc-link-lib` to link the library

### Metadata Propagation

If package `bar` has `links = "baz"` and emits `cargo::metadata=key=value`, then a **direct dependent** package `foo`'s build script gets `DEP_BAZ_KEY=value`. Metadata is only passed to immediate dependents, not transitive.

## The `*-sys` Convention

Packages named `foo-sys` follow a convention:
1. **Link** to native `libfoo` (probe system, then build from source)
2. **Declare** FFI types/functions — no high-level abstractions
3. A companion crate (`foo`, not `-sys`) provides safe wrappers

Benefits: centralizes linking logic, satisfies the one-package-per-`links` rule, and lets other `-sys` crates access `DEP_*` metadata.

## Code Generation

Build scripts can generate Rust source into `OUT_DIR` and include it:

```rust
// In build.rs: write generated code to OUT_DIR
let out_dir = std::env::var("OUT_DIR").unwrap();
let dest = Path::new(&out_dir).join("generated.rs");
// ... write generated code ...

// In lib.rs:
include!(concat!(env!("OUT_DIR"), "/generated.rs"));
```

## Overriding Build Scripts via Config

If a manifest has `links`, the build script can be overridden in `.cargo/config.toml`:

```toml
[target.x86_64-unknown-linux-gnu.foo]
rustc-link-lib = ["foo"]
rustc-link-search = ["/path/to/foo"]
rustc-cfg = ['key="value"']
rustc-env = { key = "value" }
rustc-cdylib-link-arg = ["…"]
metadata_key1 = "value"
```

This skips compiling/running the build script entirely. `warning`, `rerun-if-changed`, and `rerun-if-env-changed` are ignored. See [config.md](config.md).

## Jobserver Protocol

Cargo and rustc use the GNU make **jobserver** protocol to coordinate concurrency. Each build script inherits **one job slot** and should use only one CPU. For parallelism, use the [`jobserver` crate](https://crates.io/crates/jobserver) (e.g., `cc` crate's `parallel` feature).

## Example: Linking a C Library

```rust
// build.rs
fn main() {
    cc::Build::new()
        .file("src/foo.c")
        .compile("foo");

    println!("cargo::rerun-if-changed=src/foo.c");
}
```

```rust
// lib.rs
extern "C" { fn foo() -> i32; }

pub fn call_foo() -> i32 { unsafe { foo() } }
```

## References

- <https://doc.rust-lang.org/cargo/reference/build-scripts.html>

## Related

- [manifest.md](manifest.md) — the `build` and `links` fields
- [dependencies.md](dependencies.md) — `[build-dependencies]`
- [config.md](config.md) — overriding build scripts via `[target.*.<links>]`
- [environment.md](environment.md) — full environment variable list
- [profiles.md](profiles.md) — `build-override` profile settings
