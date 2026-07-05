# Core Concepts

Cargo's conceptual model: packages, crates, targets, the build pipeline, and the relationship between the manifest, lockfile, and cache.

## Package vs Crate

These terms are often conflated but have distinct meanings in Cargo:

| Term | Meaning |
|------|---------|
| **Package** | A `Cargo.toml` manifest plus the source files it references. The unit of publishing to a registry. A package has a name, version, and metadata. |
| **Crate** | A compilation unit — what rustc compiles. A package can contain multiple crates (a library crate + several binary crates). The crate root is `src/lib.rs` (library) or `src/main.rs` (binary). |

A single package may contain:
- One **library crate** (optional) — `src/lib.rs`, exposed to other packages as a dependency
- Multiple **binary crates** — `src/main.rs` (named after the package) and `src/bin/*.rs` (each file is a separate binary)
- Multiple **example** binaries — `examples/*.rs`
- Multiple **integration tests** — `tests/*.rs`
- Multiple **benchmarks** — `benches/*.rs`

## Crate Types

The `crate-type` field (or `--crate-type` flag) controls what rustc emits:

| Crate type | Output | Typical use |
|------------|--------|-------------|
| `bin` | Executable | Binary application |
| `lib` | `rlib` (Rust library) | Default for `src/lib.rs`; consumable by other Rust crates |
| `rlib` | Rust static library | The default library form |
| `dylib` | Dynamic Rust library | Dynamic linking between Rust crates (platform-specific filename) |
| `cdylib` | C-compatible dynamic library | FFI from other languages (`.so`/`.dylib`/`.dll`) |
| `staticlib` | C-compatible static library | FFI, linked into non-Rust programs (`.a`/`.lib`) |
| `proc-macro` | Procedural macro crate | Compile-time macro expansion; must be `dylib`-like |

A library package defaults to `rlib`. To produce multiple outputs, list them:

```toml
[lib]
crate-type = ["cdylib", "staticlib", "rlib"]
```

## Targets

A **Cargo target** is a single buildable unit within a package. The target tables configure them:

| Target type | Table | Default source | Auto-discovery |
|-------------|-------|----------------|----------------|
| Library | `[lib]` (singular) | `src/lib.rs` | Yes, if `src/lib.rs` exists |
| Binary | `[[bin]]` (array) | `src/main.rs` + `src/bin/*.rs` | Yes |
| Example | `[[example]]` (array) | `examples/*.rs` | Yes |
| Test | `[[test]]` (array) | `tests/*.rs` | Yes |
| Benchmark | `[[bench]]` (array) | `benches/*.rs` | Yes |

Auto-discovery can be disabled per type (`autobins = false`, `autolib = false`, `autoexamples = false`, `autotests = false`, `autobenches = false`). Disabling is required when a target has no conventional source file and you must specify `path` explicitly.

Each target supports: `name`, `path`, `test`, `doctest`, `bench`, `doc`, `proc-macro`, `harness`, `edition`, `crate-type`, `required-features`. See [manifest.md](manifest.md) for the full target table reference.

## Edition

The **Rust edition** controls language-level behavior and is set per package:

| Edition | Default | Notable changes |
|---------|---------|-----------------|
| `2015` | (if `edition` omitted) | Original Rust |
| `2018` | — | Module system changes, async/await, non-lexical lifetimes |
| `2021` | — | Edition-level prelude changes, disjoint closure captures, resolver v2 default |
| `2024` | — | `gen` keyword, unsafe outer attributes, `unsafe_op_in_unsafe_fn` warn-by-default |

The edition applies to all crates in the package. Mixing editions across dependencies is fully supported — each crate is compiled with its own edition. The default edition when the field is absent is `"2015"`; `cargo new` generates the latest stable edition.

The edition also selects the default **resolver version**: edition >= 2021 uses resolver v2 (see [features.md](features.md)).

## The Build Pipeline

```
1. Read Cargo.toml (manifest) for every workspace member
2. Resolve dependencies:
   a. Query registry index (sparse or git) for available versions
   b. Apply version requirements, features, platform filters
   c. Run the resolver to select a single version per package
   d. Read/update Cargo.lock
3. Fetch dependencies not in the local cache:
   a. Download .crate files → ~/.cargo/registry/cache/
   b. Extract → ~/.cargo/registry/src/
   c. Or checkout git deps → ~/.cargo/git/checkouts/
4. Build the unit graph:
   a. Each (package, target, profile) triple → a "unit"
   b. Apply profile overrides (package-specific, build-override)
5. For each unit, in dependency order:
   a. Run build script (build.rs) if present → emits cargo:: directives
   b. Compile with rustc (passing rustflags, features, cfg, etc.)
   c. Link → produce artifact in target/<profile>/
6. Run post-build steps (test execution, binary run, install, etc.)
```

### Fingerprinting

Cargo uses **fingerprinting** to skip unchanged work. Each unit has a fingerprint stored in `target/<profile>/.fingerprint/<pkg-hash>/`. The fingerprint incorporates:
- Source file mtimes/hashes
- Dependency fingerprints
- Feature set
- Profile settings
- rustc version
- Environment variables (via `rerun-if-env-changed`)

If the fingerprint matches, the unit is not rebuilt. This is why touching a file can sometimes not trigger a rebuild if its content (and mtime) is unchanged.

### The Unit Graph

The **unit graph** is Cargo's internal model of what to build. Each node is a `(package, target, profile)` triple. A single package can appear as multiple units when:
- Profile overrides differ (e.g., a dependency used both normally and as a build dependency under `build-override`)
- Different feature sets are required (under resolver v2; see [features.md](features.md))
- Different targets need the same crate compiled for host vs. target

The `--unit-graph` flag (unstable) dumps the graph as JSON. `cargo metadata` exposes the resolved dependency graph in a stable JSON format.

## Target Triple and Cross-Compilation

A **target triple** identifies the platform a crate is compiled for, e.g. `x86_64-unknown-linux-gnu`. The triple encodes architecture, vendor, OS, and ABI.

- **Host triple** — the platform cargo/rustc runs on
- **Target triple** — the platform the compiled crate runs on

When they differ, you are **cross-compiling**. Cargo distinguishes:
- `host` dependencies (build scripts, proc-macros) — compiled for the host
- `target` dependencies — compiled for the target triple

Select the target with `--target <triple>`, `build.target` config, or `CARGO_BUILD_TARGET`. A target must be installed via `rustup target add <triple>` (or a custom target spec JSON file).

The `CARGO_CFG_TARGET_*` env vars expose target properties to build scripts (see [build-scripts.md](build-scripts.md) and [environment.md](environment.md)).

## Manifest vs Lockfile

| Aspect | `Cargo.toml` (manifest) | `Cargo.lock` (lockfile) |
|--------|-------------------------|-------------------------|
| Author-written | Yes — you create and edit it | No — cargo generates and maintains it |
| Content | Version *requirements* (e.g. `serde = "1.0"`) | Exact *selected versions* + checksums |
| In git | Yes | For binaries: yes; for libraries: usually no (see below) |
| Purpose | Declares intent | Ensures reproducible builds |
| Shared across workspace | One per package; workspace root may be virtual | One per workspace |

### Should `Cargo.lock` be committed?

| Project type | Commit lockfile? |
|--------------|------------------|
| Binary (application) | **Yes** — ensures reproducible builds |
| Library | **No** (convention) — downstream consumers generate their own; but committing is not harmful and many libraries now do |

The `--locked` flag asserts the lockfile is up-to-date and fails if a dependency would need updating. `--frozen` combines `--locked` and `--offline`.

See [lockfile.md](lockfile.md) for the lockfile format.

## The Cargo Home (`~/.cargo`)

The `CARGO_HOME` directory (default `~/.cargo`) stores:

| Path | Contents |
|------|----------|
| `bin/` | Installed binaries (`cargo install`, `rustup` shims) |
| `registry/` | Registry index, cache, and extracted source |
| `registry/index/` | Cloned/fetched index (git or sparse) |
| `registry/cache/` | Downloaded `.crate` files (compressed) |
| `registry/src/` | Extracted crate sources |
| `git/` | Git dependency checkouts (`db/` bare repos, `checkouts/` working trees) |
| `config.toml` | User-global config (see [config.md](config.md)) |
| `credentials.toml` | Registry auth tokens |

See [lockfile.md](lockfile.md) for cache details and [config.md](config.md) for config file locations.

## Proc Macros

**Procedural macros** are special crates that run at compile time to generate code. They are a distinct crate type:

```toml
[lib]
proc-macro = true
```

Proc-macro crates:
- Compile to a host-triple dynamic library (even when cross-compiling)
- Cannot export anything other than the four macro kinds (function-like, derive, attribute)
- Run during compilation of dependent crates, not at runtime
- Are built with the **host** toolchain and host profile settings (see [profiles.md](profiles.md) `build-override`)

## References

- <https://doc.rust-lang.org/cargo/reference/manifest.html>
- <https://doc.rust-lang.org/cargo/reference/cargo-targets.html>
- <https://doc.rust-lang.org/cargo/reference/build-cache.html>
- <https://doc.rust-lang.org/edition-guide/>

## Related

- [manifest.md](manifest.md) — the Cargo.toml format
- [lockfile.md](lockfile.md) — the Cargo.lock format and cache layout
- [features.md](features.md) — features and the resolver
- [profiles.md](profiles.md) — build profiles
- [build-scripts.md](build-scripts.md) — build scripts
