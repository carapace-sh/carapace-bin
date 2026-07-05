# Registries and Publishing

How Cargo interacts with registries (crates.io and custom), the index formats, authentication, publishing, source replacement, `[patch]`, local registries, and vendoring.

> **Source of truth**: <https://doc.rust-lang.org/cargo/reference/registries.html>.

## What a Registry Is

A **registry** is the source from which Cargo installs crates and fetches dependencies. It has:
- An **index** — searchable metadata for every crate/version
- An optional **web API** — for publishing

The default registry is [crates.io](https://crates.io/).

## The Registry Index

The index is a collection of metadata files. Each package has one file, named after the package (lowercase). Each version is one JSON line.

### Index Directory Layout

| Name length | Path |
|-------------|------|
| 1 char | `1/<name>` |
| 2 chars | `2/<name>` |
| 3 chars | `3/<first-char>/<name>` |
| 4+ chars | `<first-two>/<second-two>/<name>` |

Example: `cargo` → `ca/rg/cargo`.

### `config.json` (index root)

| Key | Description |
|-----|-------------|
| `dl` | URL for downloading `.crate` files. Supports `{crate}`, `{version}`, `{prefix}`, `{lowerprefix}`, `{sha256-checksum}` markers. |
| `api` | Base URL for the web API (optional; required for `cargo publish`). |
| `auth-required` | bool — whether all operations (including downloads) require auth. |

### Per-Version JSON Entry

| Field | Description |
|-------|-------------|
| `name` | Package name |
| `vers` | SemVer version |
| `deps` | Array of dependency objects (`name`, `req`, `features`, `optional`, `default_features`, `target`, `kind`, `registry`, `package`) |
| `cksum` | SHA-256 of the `.crate` file |
| `features` | Map of feature → enabled features/deps |
| `yanked` | bool — whether this version is yanked |
| `links` | The `links` string (or null) |
| `v` | Schema version (1 or 2) |
| `features2` | Extended features (`dep:` / `pkg?/feat`); requires `v: 2` |
| `rust_version` | Optional MSRV |
| `pubtime` | Optional ISO8601 publish timestamp |

> The `yanked` field is the only one that may change after publication.

## Protocols: Git vs Sparse

| Protocol | URL format | Mechanism | Performance |
|----------|------------|-----------|-------------|
| **Git** | (no prefix) `https://github.com/rust-lang/crates.io-index` | Clones the entire git repo | Slow initial clone; efficient incremental fetch |
| **Sparse** | `sparse+https://index.crates.io/` | Individual HTTP requests per crate metadata file | Fast; only downloads relevant metadata |

Sparse uses `ETag`/`Last-Modified` headers with `If-None-Match`/`If-Modified-Since` for 304 caching.

### Protocol Selection for crates.io

```toml
[registries.crates-io]
protocol = "sparse"   # default in modern Cargo; also "git"
```

Env var: `CARGO_REGISTRIES_CRATES_IO_PROTOCOL`. The sparse protocol is the default for crates.io in modern Cargo.

> Don't offer both protocols for the same custom registry — `Cargo.lock` stores the registry URL, causing conflicts. crates.io is a special case handled internally.

## Authentication

### Credential Providers

| Provider | Description |
|----------|-------------|
| `cargo:token` | Default. Stores tokens in `$CARGO_HOME/credentials.toml`. Also checks `CARGO_REGISTRIES_<NAME>_TOKEN` env vars. |
| `cargo:wincred` | Windows Credential Manager |
| `cargo:macos-keychain` | macOS Keychain |
| `cargo:libsecret` | GNOME/libsecret (GNOME Keyring, KDE Wallet, KeePassXC) |
| `cargo:token-from-stdout <cmd>` | Launches a subprocess that returns a token on stdout |
| Custom aliases | Via `[credential-alias]` in config |

### Provider Configuration

```toml
[registry]
global-credential-providers = ["cargo:token", "cargo:libsecret", "cargo:macos-keychain", "cargo:wincred"]
```

Later entries have higher precedence. Per-registry: `registries.<name>.credential-provider`.

### Token Storage

```toml
# $CARGO_HOME/credentials.toml
[registries.my-registry]
token = "854DvwSlUwEHtIo3kWy6x7UCPKHfzCmy"
```

Also via: `--token` CLI flag, `CARGO_REGISTRY_TOKEN` (crates.io), `CARGO_REGISTRIES_<NAME>_TOKEN` env vars.

## Custom Registries

### Defining

```toml
# .cargo/config.toml
[registries]
my-registry = { index = "https://my-intranet:8080/git/index" }
```

Or sparse: `index = "sparse+https://my-intranet/index/"`.

Env var: `CARGO_REGISTRIES_MY_REGISTRY_INDEX=https://...`

### Using as a Dependency Source

```toml
[dependencies]
other-crate = { version = "1.0", registry = "my-registry" }
```

> crates.io does not accept packages that depend on crates from other registries.

### Publishing to a Custom Registry

```sh
cargo login --registry=my-registry
cargo publish --registry=my-registry
```

Or set a default:
```toml
[registry]
default = "my-registry"
```

## Registry Configuration Reference

| Key | Env Var | Description |
|-----|---------|-------------|
| `registries.<name>.index` | `CARGO_REGISTRIES_<name>_INDEX` | Registry index URL |
| `registries.<name>.token` | `CARGO_REGISTRIES_<name>_TOKEN` | Auth token (credentials file) |
| `registries.<name>.credential-provider` | `CARGO_REGISTRIES_<name>_CREDENTIAL_PROVIDER` | Per-registry credential provider |
| `registries.crates-io.protocol` | `CARGO_REGISTRIES_CRATES_IO_PROTOCOL` | `"git"` or `"sparse"` |
| `registry.default` | `CARGO_REGISTRY_DEFAULT` | Default registry name |
| `registry.token` | `CARGO_REGISTRY_TOKEN` | crates.io token |
| `registry.credential-provider` | `CARGO_REGISTRY_CREDENTIAL_PROVIDER` | crates.io credential provider |
| `registry.global-credential-providers` | `CARGO_REGISTRY_GLOBAL_CREDENTIAL_PROVIDERS` | Ordered provider list |

## Publishing Commands

### `cargo publish`

1. Checks `package.publish` restrictions
2. Packages a `.crate` file (compressed source)
3. Verifies by extracting and building
4. Uploads to the registry
5. Polls until the version appears in the index

| Flag | Description |
|------|-------------|
| `--dry-run` | Checks without uploading |
| `--no-verify` | Skip build verification |
| `--allow-dirty` | Allow uncommitted VCS changes |
| `--registry <name>` | Target registry |
| `--token <token>` | Token override |
| `--index <url>` | Index URL override |

### `cargo login`

Saves an auth token. Reads from stdin.
```sh
cargo login                         # default registry
cargo login --registry my-registry
```

### `cargo logout`

Removes the token locally (does not revoke on the server).
```sh
cargo logout
cargo logout --registry my-registry
```

### `cargo owner`

```sh
cargo owner --list foo
cargo owner --add username foo
cargo owner --remove username foo
cargo owner --add github:org:team foo    # team owner
```

- **Named owners**: full rights (publish, yank, add/remove)
- **Team owners**: restricted (publish, yank; no owner management)

### `cargo yank`

Removes a version from the index. **Does not delete data** — the crate remains downloadable.

```sh
cargo yank foo@1.0.7
cargo yank --undo foo@1.0.7
```

Semantics: no new deps can target a yanked version, but existing `Cargo.lock` files continue to work.

### `cargo search`

```sh
cargo search serde
cargo search --limit 20
cargo search --registry my-registry
```

Results in TOML format. Default limit 10, max 100.

### `cargo info`

```sh
cargo info serde
cargo info --registry my-registry
```

Displays detailed metadata about a package.

## Publish Restriction

```toml
[package]
publish = ["my-registry"]   # only this registry
publish = false              # no publishing
```

## Source Replacement

Redirect registry/git sources to another data source (mirror, local copy):

```toml
[source.crates-io]
replace-with = "my-vendor-source"

[source.my-vendor-source]
directory = "vendor"        # or: registry, local-registry, git
```

| Source type | Config key |
|-------------|------------|
| Registry | `registry = "<url>"` |
| Local registry | `local-registry = "<path>"` |
| Directory (unpacked) | `directory = "<path>"` |
| Git | `git = "<url>"` (plus `branch`/`tag`/`rev`) |

Rules:
- Source code must be **exactly the same** from both sources
- A replacement **cannot** contain crates not in the original
- Not for patching (use `[patch]`) or private registries

## `[patch]` — Overriding Dependencies

```toml
[patch.crates-io]
uuid = { path = "../path/to/uuid" }

[patch."https://github.com/your/repository"]
my-library = { path = "../my-library/path" }
```

- Applies **transitively** to the entire dependency graph
- Can only be defined at the **workspace root**
- Can patch in versions that don't exist on the registry yet
- Multiple versions: use `package` to rename

```toml
[patch.crates-io]
serde = { git = 'https://github.com/serde-rs/serde.git' }
serde2 = { git = '...', package = 'serde', branch = 'v2' }
```

Can also be in `.cargo/config.toml` (config patches take precedence).

## `[replace]` (Deprecated)

```toml
[replace]
"foo:0.1.0" = { git = 'https://github.com/example/foo.git' }
```

Use `[patch]` instead.

## Local Registries

A subset of a registry on the local filesystem: `*.crate` files + an `index` directory. Managed by `cargo-local-registry`:

```toml
[source.crates-io]
replace-with = "my-local"

[source.my-local]
local-registry = "path/to/registry"
```

## Directory Sources and Vendoring

Unpacked crate source on disk. Managed by `cargo vendor`:

```sh
cargo vendor                    # into "vendor/"
cargo vendor third-party/vendor
cargo vendor > .cargo/config.toml   # emit config
```

```toml
[source.crates-io]
replace-with = "vendored-sources"

[source.vendored-sources]
directory = "vendor"
```

| Flag | Description |
|------|-------------|
| `--versioned-dirs` | Include version in subdir names |
| `--respect-source-config` | Honor existing `[source]` config |
| `--no-delete` | Don't clean vendor dir first |
| `-s, --sync` | Additional `Cargo.toml` to sync |

Vendored sources are **read-only**. To modify, use `[patch]` or a `path` dependency.

## References

- <https://doc.rust-lang.org/cargo/reference/registries.html>
- <https://doc.rust-lang.org/cargo/reference/registry-index.html>
- <https://doc.rust-lang.org/cargo/reference/registry-authentication.html>
- <https://doc.rust-lang.org/cargo/reference/source-replacement.html>
- <https://doc.rust-lang.org/cargo/reference/overriding-dependencies.html>
- <https://doc.rust-lang.org/cargo/reference/publishing.html>

## Related

- [config.md](config.md) — `[registry]`, `[registries]`, `[source]` config sections
- [dependencies.md](dependencies.md) — `registry` dependency field
- [lockfile.md](lockfile.md) — how `Cargo.lock` stores registry sources
- [commands.md](commands.md) — `cargo publish`, `cargo login`, `cargo yank`, `cargo vendor`, etc.
