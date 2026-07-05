# Specifying Dependencies

How to declare dependencies in `Cargo.toml`: version requirements, sources (registry, git, path), platform-specific dependencies, optional dependencies, and workspace dependency inheritance. For the `[features]` interaction with optional deps, see [features.md](features.md).

> **Source of truth**: <https://doc.rust-lang.org/cargo/reference/specifying-dependencies.html>.

## Dependency Tables

| Table | Scope |
|-------|-------|
| `[dependencies]` | Library and all binary targets |
| `[dev-dependencies]` | Examples, tests, benchmarks only (not the library or binaries) |
| `[build-dependencies]` | Build script (`build.rs`) only — see [build-scripts.md](build-scripts.md) |
| `[target.'cfg(...)'.dependencies]` | Platform-specific (any of the above kinds) |

A dependency listed in multiple tables is built multiple times (once per kind) with possibly different settings. `dev-dependencies` and `build-dependencies` are **not** available to the library/binary code — only to their respective contexts.

## Version Requirements

Cargo uses **Semantic Versioning** requirements. A version requirement is a string that constrains which versions are acceptable.

| Syntax | Name | Meaning |
|--------|------|---------|
| `"1.0"` | Caret (default) | `>=1.0.0, <2.0.0` — compatible updates |
| `"1.2.3"` | Caret (default) | `>=1.2.3, <2.0.0` |
| `"0.2"` | Caret | `>=0.2.0, <0.3.0` (0.x is per-minor) |
| `"0.0.3"` | Caret | `>=0.0.3, <0.0.4` (0.0.x is exact-ish) |
| `"^1.0"` | Explicit caret | Same as `"1.0"` |
| `"~1.0"` | Tilde | `>=1.0.0, <1.1.0` — patch updates only |
| `"~1.2.3"` | Tilde | `>=1.2.3, <1.3.0` |
| `"=1.2.3"` | Equal | Exactly `1.2.3` |
| `"=1.2"` | Equal | `>=1.2.0, <1.3.0` |
| `">=1.0"` | Comparison | `>=1.0.0` (no upper bound) |
| `">=1.0, <2.0"` | Comparison | `>=1.0.0, <2.0.0` |
| `"<2.0"` | Comparison | `<2.0.0` |
| `"*"` | Wildcard | Any version |
| `"1.*"` | Wildcard | `>=1.0.0, <2.0.0` |

The default (bare `"1.0"`) is a **caret requirement**. Cargo treats versions with different major numbers as incompatible (unless the major is `0`, where minor bumps are incompatible, and `0.0.x` is exact).

### Pre-release and Build Metadata

Pre-release versions (`1.0.0-alpha`) are only matched by requirements that explicitly include a pre-release for the same `major.minor.patch`. This means `"1.0"` will **not** match `1.0.0-alpha.1`; you need `"1.0.0-alpha.1"` or similar.

Build metadata (`+build`) is ignored for matching purposes.

## Dependency Sources

### Registry (default: crates.io)

```toml
[dependencies]
serde = "1.0"
rand = { version = "0.8", features = ["small_rng"] }
```

### Git

```toml
[dependencies]
regex = { git = "https://github.com/rust-lang/regex" }
my-crate = { git = "https://github.com/user/repo", branch = "dev" }
tagged = { git = "https://github.com/user/repo", tag = "v1.0" }
pinned = { git = "https://github.com/user/repo", rev = "abc123" }
```

| Specifier | Default | Notes |
|-----------|---------|-------|
| `branch` | `master` (or `main`) | Git branch |
| `tag` | — | Git tag (e.g. `"v1.0"`) |
| `rev` | — | Specific commit SHA |

If none of `branch`/`tag`/`rev` is specified, Cargo uses the default branch.

### Path

```toml
[dependencies]
my-crate = { path = "../my-crate" }
```

Path dependencies are local filesystem crates. They are **not** published to a registry as-is; when publishing, a path dependency without a `version` must have a registry version available (Cargo rewrites it). Path dependencies with `version` will fall back to the registry version when the package is published.

### Registry (custom)

```toml
[dependencies]
other-crate = { version = "1.0", registry = "my-registry" }
```

See [registries.md](registries.md) for custom registry setup.

## Full Dependency Spec Fields

| Field | Type | Default | Notes |
|-------|------|---------|-------|
| `version` | string | — | Version requirement |
| `registry` | string | default registry | Custom registry name |
| `git` | string (url) | — | Git source URL |
| `branch` | string | default branch | Git branch |
| `tag` | string | — | Git tag |
| `rev` | string | — | Git commit |
| `path` | string (path) | — | Local path |
| `optional` | bool | `false` | Creates an implicit feature — see [features.md](features.md) |
| `default-features` | bool | `true` | Whether to enable the dependency's `default` feature |
| `features` | string[] | `[]` | Features to enable on this dependency |
| `rename` | string | — | Rename the dependency in code (`extern crate <rename>`) |
| `package` | string | — | Use a different package name than the key (for renaming) |

### Renaming Dependencies

Two mechanisms:

**`package`** (renames the package, the key is what you call it):
```toml
[dependencies]
my-types = { package = "types", version = "1.0" }
# In code: use my_types; (Rust converts - to _)
```

**`rename`** (the `cargo add --rename` form; same as `package` but inversed keying) — both produce the same effect: the dependency is imported under the key name while the real package name is in `package`.

## Platform-Specific Dependencies

Use `[target]` with cfg expressions to restrict dependencies to specific platforms:

```toml
[target.'cfg(windows)'.dependencies]
winapi = "0.3"

[target.'cfg(unix)'.dependencies]
nix = "0.27"

[target.'cfg(target_arch = "x86_64")'.dependencies]
native-simd = "0.1"

[target.x86_64-unknown-linux-gnu.dependencies]
linux-feature = "0.1"
```

A cfg expression can combine multiple predicates:
```toml
[target.'cfg(all(target_arch = "arm", target_os = "none"))'.dependencies]
cortex-m = "0.7"
```

Target triples (e.g. `x86_64-unknown-linux-gnu`) match exactly; cfg expressions (`cfg(...)`) match predicates. Triple-specific tables take precedence over cfg-specific ones.

The same table keys apply (`dependencies`, `dev-dependencies`, `build-dependencies`).

## Optional Dependencies

```toml
[dependencies]
gif = { version = "0.11", optional = true }
```

Marking a dependency `optional = true`:
- Does **not** compile it by default
- Creates an **implicit feature** with the same name (unless `dep:` is used to suppress — see [features.md](features.md))
- The dependency is only built when the corresponding feature is enabled

## Workspace Dependency Inheritance

Dependencies declared at the workspace level can be inherited by members:

```toml
# Workspace root Cargo.toml
[workspace.dependencies]
serde = "1.0"
rand = { version = "0.8", features = ["small_rng"] }
my-crate = { path = "../my-crate" }
```

```toml
# Member package Cargo.toml
[dependencies]
serde.workspace = true
rand.workspace = true
```

A member can override specific fields:
```toml
[dependencies]
serde = { workspace = true, features = ["rc"] }
```

See [workspaces.md](workspaces.md) for the full workspace inheritance model.

## Adding Dependencies via CLI

`cargo add` modifies `Cargo.toml` automatically:

```sh
cargo add serde                          # from crates.io
cargo add serde --features derive
cargo add tokio --no-default-features
cargo add my-crate --path ../my-crate
cargo add foo --git https://github.com/u/foo --branch dev
cargo add bar --dev                      # dev-dependency
cargo add baz --build                    # build-dependency
cargo add qux --target 'cfg(windows)'
cargo add serde --optional
cargo add serde --rename my-serde
cargo add serde --dry-run                # preview without writing
```

See [commands.md](commands.md) for the full `cargo add` flag reference.

## SemVer Compatibility Notes

When you publish a crate, Cargo's resolver uses version requirements to pick compatible versions. Be aware:

- A `0.x.y` version bump in a dependency is a **breaking** change (since `0.x` caret is per-minor)
- Removing a feature from `default` is a SemVer-incompatible change
- Changing which features you enable on a dependency is safe, but adding a new optional dependency is safe, removing one is not

See [features.md](features.md) for feature-level SemVer rules.

## References

- <https://doc.rust-lang.org/cargo/reference/specifying-dependencies.html>

## Related

- [features.md](features.md) — features and optional dependency interaction
- [workspaces.md](workspaces.md) — workspace dependency inheritance
- [build-scripts.md](build-scripts.md) — build-dependencies
- [registries.md](registries.md) — custom registries and source replacement
- [commands.md](commands.md) — `cargo add`, `cargo remove`, `cargo update`
