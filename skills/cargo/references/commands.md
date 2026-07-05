# Command Reference

The cargo CLI command and subcommand reference, organized by group. Covers flags, common options, and usage. For configuration of flags via env/config, see [config.md](config.md).

## Top-Level `cargo`

```
cargo [OPTIONS] [COMMAND]
```

| Flag | Description |
|------|-------------|
| `--list` | List installed commands (including subcommands on `$PATH` like `cargo-watch`) |
| `--explain=<CODE>` | Detailed explanation of a rustc error message |
| `-C <DIR>` | Change to directory before doing anything (nightly-only) |
| `-V`, `--version` | Print version info and exit |
| `-h`, `--help` | Print help |

### Persistent Flags (apply to all subcommands)

| Flag | Description |
|------|-------------|
| `--color <auto\|always\|never>` | Coloring |
| `--config <KEY=VALUE>` | Override a config value (TOML dot notation; repeatable) |
| `--frozen` | Equivalent to `--locked` + `--offline` |
| `--locked` | Assert `Cargo.lock` is unchanged; fail if an update is needed |
| `--offline` | Run without network access (deps must be cached) |
| `-Z <FLAG>` | Unstable (nightly-only) flag; `cargo -Z help` for details |
| `-q`, `--quiet` | Suppress cargo log messages |
| `-v`, `--verbose` | Verbose output (`-vv` for very verbose / build script output) |

### Subcommand Discovery

`cargo` looks for `cargo-<name>` executables on `$PATH`. This is how `cargo-watch` (`cargo watch`), `cargo-edit` (`cargo add` before it was built-in), and `cargo-expand` work. `cargo --list` shows all discovered subcommands.

## Build Commands

### `cargo build`

Compile a local package and all dependencies.

| Flag | Description |
|------|-------------|
| `-p, --package <SPEC>` | Package(s) to build (repeatable) |
| `--workspace` | Build all packages in the workspace |
| `--exclude <SPEC>` | Exclude packages (repeatable) |
| `-r, --release` | Build in release mode (`--profile release`) |
| `--profile <NAME>` | Build with the specified profile |
| `-F, --features <FEAT>` | Features to activate (repeatable; comma/space separated) |
| `--all-features` | Activate all features |
| `--no-default-features` | Do not activate the `default` feature |
| `--lib` | Build only this package's library |
| `--bin <NAME>` | Build only the specified binary (repeatable) |
| `--bins` | Build all binaries |
| `--example <NAME>` | Build only the specified example (repeatable) |
| `--examples` | Build all examples |
| `--test <NAME>` | Build only the specified test target (repeatable) |
| `--tests` | Build all test targets |
| `--bench <NAME>` | Build only the specified bench target (repeatable) |
| `--benches` | Build all bench targets |
| `--all-targets` | Build all targets |
| `--target <TRIPLE>` | Build for the target triple (repeatable) |
| `--target-dir <DIR>` | Directory for generated artifacts |
| `--manifest-path <PATH>` | Path to `Cargo.toml` |
| `--message-format <FMT>` | Error format (repeatable): `human`, `json`, `short`, `json-diagnostic-rendered-ansi`, `json-render-diagnostics` |
| `--timings` | Output a build timing report |
| `--unit-graph` | Output build graph as JSON (unstable) |
| `--keep-going` | Do not abort on first error |
| `--ignore-rust-version` | Ignore `rust-version` specification |
| `--future-incompat-report` | Output a future incompatibility report |
| `-j, --jobs <N>` | Number of parallel jobs (default: # CPUs) |
| `-h, --help` | Print help |

### `cargo check`

Like `cargo build` but only checks for errors without producing artifacts. Faster. Same flags as `build` (minus profile-specific ones where irrelevant). Uses `dev` profile by default, `--release` for release.

### `cargo test`

Execute all unit and integration tests and build examples.

| Flag (beyond build flags) | Description |
|------|-------------|
| `--doc` | Run only documentation tests |
| `--no-run` | Compile but don't run tests |
| `--no-fail-fast` | Run all tests regardless of failure |
| `-q, --quiet` | One character per test instead of one line |

Test binaries go to `target/debug/deps/`. Tests use the `test` profile (inherits `dev`). Args after `--` are passed to the test binary.

### `cargo bench`

Execute all benchmarks. Uses the `bench` profile (inherits `release`). Same target-selection flags as build, plus `--no-run`, `--no-fail-fast`.

### `cargo run`

Build and run a binary or example.

```sh
cargo run                      # run the default binary
cargo run --bin my-bin         # run a specific binary
cargo run --example my-example # run an example
cargo run -- arg1 arg2         # args after -- passed to the program
```

| Flag | Description |
|------|-------------|
| `--bin <NAME>` | Name of the binary to run |
| `--example <NAME>` | Name of the example to run |
| `-p, --package <SPEC>` | Package with the target to run |
| `-r, --release` | Run in release mode |
| `--profile <NAME>` | Run with the specified profile |
| `-F, --features`, `--all-features`, `--no-default-features` | Feature selection |

### `cargo doc`

Build documentation into `target/doc`.

| Flag | Description |
|------|-------------|
| `--open` | Open docs in a browser after building |
| `--no-deps` | Don't build docs for dependencies |
| `--document-private-items` | Include private items |
| `-r, --release` | Build in release mode |
| `--profile <NAME>` | Build with specified profile |

### `cargo clean`

Remove artifacts from the `target` directory.

| Flag | Description |
|------|-------------|
| `-p, --package <SPEC>` | Clean only this package's artifacts |
| `--target <TRIPLE>` | Clean for a specific target |
| `--profile <NAME>` | Clean artifacts of a specific profile |
| `-r, --release` | Clean release artifacts |
| `--doc` | Clean only the documentation directory |
| `--target-dir <DIR>` | Artifact directory |
| `--manifest-path <PATH>` | Path to `Cargo.toml` |
| `-n, --dry-run` | Display what would be deleted without deleting |

### `cargo fix`

Automatically apply lint warning fixes from rustc.

| Flag (beyond build flags) | Description |
|------|-------------|
| `--edition` | Fix in preparation for the next edition |
| `--edition-idioms` | Fix warnings to migrate to edition idioms |
| `--broken-code` | Fix even if there are existing compiler errors |
| `--allow-dirty` | Fix even with a dirty working directory |
| `--allow-staged` | Fix even with staged changes |
| `--allow-no-vcs` | Fix even without a VCS |

### `cargo rustc`

Compile a package and pass extra options to rustc. Args after `--` are passed to rustc.

| Flag (beyond build flags) | Description |
|------|-------------|
| `--print <INFO>` | Output compiler info without compiling (e.g. `cfg`, `crate-name`) |
| `--crate-type <TYPE>` | Comma-separated crate types to emit |

### `cargo rustdoc`

Build docs passing custom flags to rustdoc. Args after `--` are forwarded to rustdoc. Same flags as `cargo doc` plus `--output-format` (unstable: `html`, `json`).

## Manifest Commands

### `cargo add`

Add a dependency to `Cargo.toml`.

```sh
cargo add serde
cargo add serde --features derive
cargo add tokio --no-default-features
cargo add foo --git https://github.com/u/foo --branch dev
cargo add bar --path ../bar
cargo add baz --dev
cargo add qux --build
cargo add quux --target 'cfg(windows)'
```

| Flag | Description |
|------|-------------|
| `--git <URL>` | Git repository location |
| `--branch <BRANCH>` / `--tag <TAG>` / `--rev <REV>` | Git ref |
| `--path <PATH>` | Filesystem path to local crate |
| `--registry <NAME>` | Registry to use |
| `--dev` | Add as dev-dependency |
| `--build` | Add as build-dependency |
| `--target <TRIPLE>` | Add as target-specific dependency |
| `--optional` | Mark as optional |
| `--no-optional` | Mark as required |
| `--default-features` | Re-enable default features |
| `--no-default-features` | Disable default features |
| `-F, --features <FEAT>` | Features to activate (repeatable) |
| `--rename <NAME>` | Rename the dependency |
| `--public` / `--no-public` | Mark as public/private (unstable) |
| `--base <PATH>` | Path base for local crate (unstable) |
| `--manifest-path <PATH>` | Path to `Cargo.toml` |
| `-p, --package <SPEC>` | Package to modify |
| `-n, --dry-run` | Don't write the manifest |

Without arguments, interactive mode prompts for the dependency.

### `cargo remove`

Remove a dependency from `Cargo.toml`.

```sh
cargo remove serde
cargo remove serde --dev
cargo remove serde --build
cargo remove serde --target 'cfg(windows)'
```

| Flag | Description |
|------|-------------|
| `--dev` | Remove from dev-dependencies |
| `--build` | Remove from build-dependencies |
| `--target <TRIPLE>` | Remove from target-dependencies |
| `--manifest-path <PATH>` | Path to `Cargo.toml` |
| `-p, --package <SPEC>` | Package to remove from |
| `-n, --dry-run` | Don't write the manifest |

### `cargo update`

Update dependencies in `Cargo.lock`.

```sh
cargo update                       # update all
cargo update -p serde              # update just serde
cargo update -p serde --precise 1.0.85
cargo update --dry-run
cargo update -w                    # only workspace packages
```

| Flag | Description |
|------|-------------|
| `-p, --package <SPEC>` | Package to update (repeatable) |
| `--precise <VERSION>` | Update to exactly this version |
| `--recursive` / `--aggressive` | Force update all deps of the spec too |
| `-b, --breaking` | Update to latest SemVer-breaking version (unstable) |
| `-w, --workspace` | Only update workspace packages |
| `-n, --dry-run` | Don't write the lockfile |
| `--manifest-path <PATH>` | Path to `Cargo.toml` |
| `--ignore-rust-version` | Ignore `rust-version` |

### `cargo tree`

Display a tree visualization of the dependency graph.

```sh
cargo tree
cargo tree -d                           # show duplicates
cargo tree -i serde                     # inverted: what depends on serde
cargo tree -e features                  # show features
cargo tree -f "{p} {f}"                 # custom format
cargo tree --depth 2
cargo tree --target all
```

| Flag | Description |
|------|-------------|
| `-p, --package <SPEC>` | Root of the tree (repeatable) |
| `-i, --invert <SPEC>` | Invert the tree, focus on a package |
| `-d, --duplicates` | Show only multi-version deps (implies `-i`) |
| `-e, --edges <KIND>` | Edge kinds: `normal`, `build`, `dev`, `features`, `all` |
| `-f, --format <FMT>` | Format string |
| `--depth <N>` | Max display depth |
| `--target <TRIPLE>` | Filter by target (`all` for all) |
| `--exclude <SPEC>` | Exclude workspace members |
| `--prune <SPEC>` | Hide packages from the tree |
| `--prefix <INDENT>` | Indentation style |
| `--charset <SET>` | Character set |
| `--no-dedupe` | Don't de-duplicate shared deps |
| `-F, --features`, `--all-features`, `--no-default-features` | Feature selection |

### `cargo metadata`

Output the resolved dependency graph as JSON.

```sh
cargo metadata
cargo metadata --no-deps
cargo metadata --format-version 1
cargo metadata --filter-platform x86_64-unknown-linux-gnu
```

| Flag | Description |
|------|-------------|
| `--no-deps` | Only workspace member info, don't fetch deps |
| `--format-version <N>` | Output format version (currently `1`) |
| `--filter-platform <TRIPLE>` | Only include deps matching this target |
| `-F, --features`, `--all-features`, `--no-default-features` | Feature selection |
| `--manifest-path <PATH>` | Path to `Cargo.toml` |

### `cargo generate-lockfile`

Generate a new `Cargo.lock` (or update the existing one).

| Flag | Description |
|------|-------------|
| `--manifest-path <PATH>` | Path to `Cargo.toml` |
| `--ignore-rust-version` | Ignore `rust-version` |
| `--publish-time <TIME>` | Latest publish time (unstable) |

### `cargo locate-project`

Print a JSON representation of a `Cargo.toml` location.

| Flag | Description |
|------|-------------|
| `--workspace` | Locate the workspace root `Cargo.toml` |
| `--manifest-path <PATH>` | Path to `Cargo.toml` |
| `--message-format <FMT>` | Output representation |

### `cargo fetch`

Fetch dependencies into the local cache without building.

| Flag | Description |
|------|-------------|
| `--target <TRIPLE>` | Fetch for this target (repeatable) |
| `--manifest-path <PATH>` | Path to `Cargo.toml` |

### `cargo vendor`

Vendor all dependencies locally.

```sh
cargo vendor
cargo vendor third-party/vendor
cargo vendor > .cargo/config.toml
```

| Flag | Description |
|------|-------------|
| `-s, --sync <TOML>` | Additional `Cargo.toml` to sync (repeatable) |
| `--no-delete` | Don't delete older crates in vendor dir |
| `--versioned-dirs` | Include version in subdir names |
| `--respect-source-config` | Respect `[source]` config |
| `--manifest-path <PATH>` | Path to `Cargo.toml` |

### `cargo pkgid`

Print a fully qualified package ID specification.

```sh
cargo pkgid
cargo pkgid -p serde
```

See [lockfile.md](lockfile.md) for the Package ID Spec format.

## Package Commands

### `cargo new`

Create a new cargo package at `<path>` (creates a new directory).

```sh
cargo new my-project
cargo new my-lib --lib
cargo new my-project --edition 2021
cargo new my-project --vcs none
cargo new my-project --name custom-name
```

| Flag | Description |
|------|-------------|
| `--bin` | Binary template (default) |
| `--lib` | Library template |
| `--edition <YEAR>` | Edition to set |
| `--name <NAME>` | Package name (defaults to directory name) |
| `--vcs <VCS>` | `git`, `hg`, `pijul`, `fossil`, `none` |
| `--registry <NAME>` | Registry to use |

### `cargo init`

Create a new cargo package in the **current** directory (no new directory).

Same flags as `cargo new`.

### `cargo install`

Install a Rust binary.

```sh
cargo install ripgrep              # from crates.io
cargo install --git https://github.com/u/foo
cargo install --path ./my-crate
cargo install --list               # list installed
cargo install foo --version 1.0.0
```

| Flag | Description |
|------|-------------|
| `--git <URL>` | Install from a git repo |
| `--branch` / `--tag` / `--rev` | Git ref |
| `--path <PATH>` | Install from a local path |
| `--registry <NAME>` / `--index <URL>` | Registry source |
| `--vers` / `--version <VER>` | Version to install |
| `--bin <NAME>` | Install only this binary (repeatable) |
| `--bins` | Install all binaries |
| `--example <NAME>` / `--examples` | Install examples instead |
| `--root <DIR>` | Installation directory |
| `--list` | List all installed packages and versions |
| `-f, --force` | Force overwriting existing installs |
| `--no-track` | Don't save tracking info |
| `--debug` | Build in debug mode (`dev` profile) |
| `--profile <NAME>` | Install with this profile |
| `-n, --dry-run` | Checks without installing (unstable) |
| `-F, --features`, `--all-features`, `--no-default-features` | Feature selection |
| `-j, --jobs <N>` | Parallel jobs |

### `cargo uninstall`

Remove an installed binary.

```sh
cargo uninstall ripgrep
cargo uninstall --bin rg
cargo uninstall -p ripgrep
```

| Flag | Description |
|------|-------------|
| `-p, --package <SPEC>` | Package to uninstall |
| `--bin <NAME>` | Only uninstall this binary |
| `--root <DIR>` | Directory to uninstall from |

### `cargo search`

Search packages in the registry.

```sh
cargo search serde
cargo search --limit 20
```

| Flag | Description |
|------|-------------|
| `--limit <N>` | Max results (default 10, max 100) |
| `--index <URL>` | Registry index URL |
| `--registry <NAME>` | Registry to search |

### `cargo info`

Display detailed info about a package.

```sh
cargo info serde
cargo info --registry my-registry
```

## Publishing Commands

### `cargo publish`

Upload a package to the registry. See [registries.md](registries.md).

| Flag | Description |
|------|-------------|
| `-n, --dry-run` | Validate without uploading |
| `--no-verify` | Skip build verification |
| `--allow-dirty` | Allow uncommitted VCS changes |
| `--registry <NAME>` / `--index <URL>` | Target registry |
| `--token <TOKEN>` | Token override |
| `-p, --package <SPEC>` | Package(s) to publish |
| `--workspace` | Publish all packages |
| `--exclude <SPEC>` | Exclude packages |
| `--target <TRIPLE>` | Build for target |

### `cargo package`

Assemble a package into a distributable `.crate` tarball.

| Flag | Description |
|------|-------------|
| `-l, --list` | Print files included without creating the archive |
| `--no-verify` | Don't verify by building |
| `--no-metadata` | Ignore missing-metadata warnings |
| `--allow-dirty` | Allow dirty working directories |
| `--exclude-lockfile` | Don't include the lock file |
| `-p, --package <SPEC>` | Package(s) to assemble |
| `--workspace` | Assemble all packages |
| `--exclude <SPEC>` | Exclude packages |

### `cargo login` / `cargo logout` / `cargo owner` / `cargo yank`

See [registries.md](registries.md) for details.

## Inspection Commands

### `cargo report`

Generate and display reports (unstable).

```sh
cargo report timings
cargo report sessions
cargo report rebuilds
cargo report future-incompatibilities   # alias: future-incompat
```

| Subcommand | Flags |
|------------|-------|
| `timings` | `--id <ID>`, `--open`, `--manifest-path` |
| `sessions` | `--limit <N>`, `--manifest-path` |
| `rebuilds` | `--id <ID>`, `--manifest-path` |
| `future-incompatibilities` | `--id <ID>`, `-p <SPEC>` |

### `cargo config`

Inspect configuration values (unstable).

```sh
cargo config get <key>
cargo config get <key> --show-origin
cargo config get <key> --format <FMT>
cargo config get <key> --merged <bool>
```

### `cargo version`

Print version information.

## Package ID Spec

Many commands accept a **Package ID Spec** (`-p` flag) to select a package. Formats:

| Format | Example | Matches |
|--------|---------|--------|
| name | `serde` | any version of `serde` |
| name@version | `serde@1.0` | version requirement |
| URL spec | `https://github.com/rust-lang/crates.io-index#serde@1.0.189` | exact |
| local path | `./my-crate` | the package at that path |

See [lockfile.md](lockfile.md) for the full spec and how it appears in `Cargo.lock` and `cargo metadata`.

## Common Flag Patterns

These flags appear on most build/test/bench commands:

| Flag | Meaning |
|------|---------|
| `--manifest-path <PATH>` | Path to `Cargo.toml` (default: search upward from CWD) |
| `--target <TRIPLE>` | Cross-compile for this target |
| `--target-dir <DIR>` | Output directory (overrides `build.target-dir`) |
| `-j, --jobs <N>` | Parallel jobs |
| `-F, --features`, `--all-features`, `--no-default-features` | Feature selection |
| `--keep-going` | Continue after errors |
| `--message-format <FMT>` | Compiler message format |
| `--timings` | Build timing report |
| `--ignore-rust-version` | Ignore MSRV |

## References

- <https://doc.rust-lang.org/cargo/commands/cargo.html>
- <https://doc.rust-lang.org/cargo/reference/pkgid-spec.html>

## Related

- [config.md](config.md) — flag equivalents in config and env vars
- [features.md](features.md) — `--features` semantics
- [profiles.md](profiles.md) — `--profile`, `--release`
- [lockfile.md](lockfile.md) — Package ID Spec
- [registries.md](registries.md) — publishing commands in depth
