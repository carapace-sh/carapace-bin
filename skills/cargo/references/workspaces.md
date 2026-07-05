# Workspaces

A **workspace** is a collection of related packages that share a `Cargo.lock`, output directory, and various settings. Workspaces let you manage a multi-crate project as a single unit.

> **Source of truth**: <https://doc.rust-lang.org/cargo/reference/workspaces.html>.

## Workspace Definition

A workspace is defined by a `[workspace]` section in a `Cargo.toml`. The manifest containing `[workspace]` is the **workspace root**. There are two forms:

**Root package** (has both `[package]` and `[workspace]`):
```toml
[package]
name = "my-project"
version = "0.1.0"

[workspace]
members = ["crate-a", "crate-b"]
```

**Virtual manifest** (only `[workspace]`, no `[package]`):
```toml
[workspace]
members = ["crate-a", "crate-b"]
resolver = "2"
```

A virtual manifest exists solely to define the workspace. It cannot be built directly.

## Members and Exclusion

```toml
[workspace]
members = ["crate-a", "crate-b", "libs/*"]
exclude = ["libs/old-crate"]
```

- `members` supports glob patterns (`*`, `**`)
- `exclude` removes paths that would otherwise be matched by `members` globs
- A subdirectory with its own `Cargo.toml` that isn't listed is **not** a workspace member (it's a standalone package, unless it sets `workspace = ".."` to opt into the parent workspace)

### `default-members`

```toml
[workspace]
members = ["crate-a", "crate-b", "tools/*"]
default-members = ["crate-a"]
```

When `cargo build` is run from the workspace root without `-p`, only `default-members` are built. If unset, all members are built.

## Shared Resources

All workspace members share:

| Resource | Location |
|----------|----------|
| `Cargo.lock` | One lockfile at the workspace root |
| `target/` directory | One output directory at the workspace root |
| `Cargo.lock` resolution | A single resolved dependency graph across all members |

This means dependencies are resolved **across all members together** — feature unification (see [features.md](features.md)) considers all members' feature requests.

## The `resolver` Field

The resolver version is set at the workspace level:

```toml
[workspace]
resolver = "2"
```

If a workspace root package also sets `resolver` in `[package]`, the workspace-level setting takes precedence. See [features.md](features.md) for resolver behavior.

## Workspace Inheritance

### `[workspace.package]` — Shared Package Fields

Declare shared `[package]` fields at the workspace root:

```toml
# Workspace root
[workspace.package]
version = "1.0.0"
authors = ["Team <team@example.com>"]
edition = "2021"
license = "MIT"
repository = "https://github.com/org/repo"
```

Members inherit with `.workspace = true`:

```toml
# Member package
[package]
name = "my-crate"
version.workspace = true
edition.workspace = true
license.workspace = true
```

**Inheritable fields**: `version`, `authors`, `edition`, `rust-version`, `description`, `homepage`, `repository`, `license`, `license-file`, `keywords`, `categories`, `publish`, `exclude`, `include`, `documentation`, `readme`, `metadata`.

### `[workspace.dependencies]` — Shared Dependencies

Declare dependencies once at the workspace root:

```toml
# Workspace root
[workspace.dependencies]
serde = "1.0"
rand = { version = "0.8", features = ["small_rng"] }
my-crate = { path = "crates/my-crate" }
tokio = { version = "1", features = ["full"] }
```

Members inherit:

```toml
# Member package
[dependencies]
serde.workspace = true
tokio.workspace = true
```

A member can override individual fields while inheriting the rest:

```toml
[dependencies]
serde = { workspace = true, features = ["rc"] }
```

### `[workspace.lints]` — Shared Lint Configuration

```toml
# Workspace root
[workspace.lints.rust]
unsafe_code = "forbid"

# Member
[lints]
workspace = true
```

### `[workspace.metadata]` — Arbitrary Shared Data

```toml
[workspace.metadata]
something = "value"
```

Ignored by Cargo; available to external tools. Accessible via `cargo metadata` output.

## Path Dependencies Within a Workspace

Path dependencies pointing to other workspace members are the norm:

```toml
# crates/app/Cargo.toml
[dependencies]
my-lib = { path = "../my-lib" }
```

Cargo recognizes that `../my-lib` is a workspace member and uses it directly (not via a registry). When `app` is published, Cargo rewrites the path dependency to a registry version requirement if `my-lib` has a `version` field; if not, the publish fails (path deps without versions can't be published).

## The Root Package Pattern

A common layout:

```
my-project/
├── Cargo.toml          # [package] + [workspace] members
├── Cargo.lock
├── target/
├── crates/
│   ├── lib/
│   │   └── Cargo.toml  # [package] name="lib", workspace=".."
│   └── bin/
│       └── Cargo.toml  # [package] name="bin", depends on lib
```

The root `Cargo.toml` has both `[package]` (the root crate) and `[workspace]` (listing the sub-crates as members). Alternatively, use a virtual manifest at the root and put all crates in subdirectories.

## `cargo` Commands in a Workspace

| Command | Behavior |
|---------|----------|
| `cargo build` | Builds `default-members` (or all members if unset) |
| `cargo build -p <name>` | Builds a specific member |
| `cargo build --workspace` | Builds all members |
| `cargo test --workspace` | Tests all members |
| `cargo run -p <name>` | Runs a specific member's binary |

The `-p` / `--package` flag accepts a **Package ID Spec** — see [lockfile.md](lockfile.md) and [commands.md](commands.md).

## References

- <https://doc.rust-lang.org/cargo/reference/workspaces.html>

## Related

- [manifest.md](manifest.md) — the `[workspace]` section and `workspace` field
- [features.md](features.md) — workspace-level resolver and feature unification
- [dependencies.md](dependencies.md) — workspace dependency inheritance
- [lockfile.md](lockfile.md) — shared Cargo.lock
