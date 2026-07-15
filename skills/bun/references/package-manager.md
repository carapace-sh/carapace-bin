# Bun Package Manager

Reference for Bun's package manager: `bun install`, `bun add`, `bun remove`, `bun link`, `bun unlink`, and the `bun pm` subcommands. Covers lockfile format, linkers, workspaces, trusted dependencies, and the backend system.

> **Source of truth**: <https://bun.sh/docs/cli/install>. For CLI flag reference, see [cli.md](cli.md). For bunfig.toml `[install]` section, see [config.md](config.md).

## Overview

Bun's package manager is built into the Bun runtime — no separate binary needed. It installs dependencies from `package.json`, resolves versions using npm's registry, and writes a lockfile (`bun.lock` or legacy `bun.lockb`).

## bun install

Installs all `dependencies`, `devDependencies`, `optionalDependencies`, and `peerDependencies` from `package.json`. Runs lifecycle scripts (`{pre|post}install`, `{pre|post}prepare`). Writes a `bun.lock` lockfile.

```bash
bun install                         # install all dependencies
bun install react                   # install a specific package
bun install react@19.1.1            # install a specific version
bun install --production            # skip devDependencies
bun install --frozen-lockfile       # CI mode, don't modify lockfile
bun install --global cowsay         # install globally
bun ci                              # alias for --frozen-lockfile
```

### Lifecycle Scripts

Bun runs these lifecycle scripts during install:

| Hook | When |
|------|------|
| `preinstall` | Before installation begins |
| `install` | During installation |
| `postinstall` | After installation completes |
| `preprepare` / `prepare` / `postprepare` | During `prepare` phase |

Lifecycle scripts from dependencies only run if they are listed in `trustedDependencies` in `package.json`. Use `--trust` to add a package to `trustedDependencies` and run its scripts.

### Skip Scripts

```bash
bun install --ignore-scripts         # skip all lifecycle scripts
bun pm untrusted                     # see untrusted dependencies
bun pm trust --all                   # trust all untrusted deps
bun pm trust lodash                  # trust specific package
```

### Workspaces

Bun supports monorepo workspaces via the `workspaces` field in root `package.json`:

```json
{
  "workspaces": ["packages/*", "apps/*"]
}
```

```bash
bun install --filter 'ba*'           # install in matching workspaces
bun install --filter '!pkg-c'        # exclude matching workspaces
bun install --filter './packages/pkg-a'  # specific workspace path
bun run --filter 'ba*' build          # run script in matching workspaces
bun run --workspaces build            # run script in all workspaces
```

### Backends

Bun optimizes installation using platform-specific file system operations:

| Backend | Description |
|---------|-------------|
| `clonefile` | macOS only, uses APFS clonefile syscall |
| `hardlink` | Hardlinks files (cross-platform) |
| `symlink` | Symlinks files (cross-platform) |
| `copyfile` | Copies files (cross-platform, slowest) |

```bash
bun install --backend hardlink
```

### Linkers

Controls how `bun install` lays out dependencies in `node_modules`:

| Linker | Description |
|--------|-------------|
| `hoisted` | Link dependencies in a shared `node_modules` directory (npm-style) |
| `isolated` | Link dependencies inside each package installation (pnpm-style) |

```bash
bun install --linker hoisted
bun install --linker isolated
```

Default: `"isolated"` for new workspaces, `"hoisted"` for new single-package projects and existing projects (pre-v1.3.2).

### Global Store

When using `"isolated"` linker, `globalStore = true` shares package installations across projects in a global virtual store at `<cache>/links/`. Makes warm installs after `rm -rf node_modules` much faster.

```toml
[install]
globalStore = true
```

Environment variable: `BUN_INSTALL_GLOBAL_STORE`.

### Hoist Patterns

When using `"isolated"` linker:

```toml
[install]
publicHoistPattern = ["*eslint*", "*prettier*"]  # hoist to root node_modules
hoistPattern = ["*"]                                # hoist to virtual store root (node_modules/.bun)
```

## Lockfile

Bun writes a lockfile to track exact dependency versions and integrity.

### Text Lockfile (bun.lock, default since v1.2)

Human-readable, diff-friendly, and YAML-like. Default since Bun v1.2.

```bash
bun install                          # writes bun.lock
bun install --save-text-lockfile     # force text lockfile
bun install --yarn                   # also write yarn.lock (yarn v1)
```

### Binary Lockfile (bun.lockb, legacy)

Binary format used before Bun v1.2. To use it:

```toml
[install]
saveTextLockfile = false
```

### Lockfile-Only Mode

```bash
bun install --lockfile-only          # generate lockfile without installing
```

### Frozen Lockfile

```bash
bun install --frozen-lockfile         # disallow changes to lockfile
bun ci                               # alias
```

Fails if `package.json` and `bun.lock` disagree. Useful for CI/CD.

## Trusted Dependencies

Bun does not run lifecycle scripts from dependencies by default — only from packages listed in `trustedDependencies` in `package.json`.

```json
{
  "trustedDependencies": ["eslint", "prettier"]
}
```

```bash
bun add --trust lodash               # add to trustedDependencies and install
bun pm untrusted                     # list untrusted dependencies
bun pm trust --all                   # trust all untrusted
bun pm trust lodash                  # trust specific package
bun pm default-trusted               # list default trusted deps
```

## bun pm Subcommands

### `bun pm pack`

Create a tarball of the current workspace.

```bash
bun pm pack
```

| Flag | Description |
|------|-------------|
| `--dry-run` | Show what would be included without writing to disk |
| `--destination <dir>` | Directory to save the tarball in |
| `--filename <name>` | Exact file name for the tarball |
| `--ignore-scripts` | Skip pre/postpack and prepare scripts |
| `--gzip-level <0-9>` | Gzip compression level (default: 9) |
| `--quiet` | Print only the tarball filename |

### `bun pm bin`

Print the path to the `bin` directory.

```bash
bun pm bin                           # local project bin directory
bun pm bin -g                        # global bin directory
```

### `bun pm ls` / `bun list`

Print a list of installed dependencies.

```bash
bun pm ls                            # top-level only
bun pm ls --all                      # all dependencies (nth-order)
bun pm ls --trusted                  # only trusted dependencies
```

### `bun pm whoami`

Print your npm username (requires login).

```bash
bun pm whoami
```

### `bun pm hash`

Generate and print the hash of the current lockfile.

```bash
bun pm hash                          # print lockfile hash
bun pm hash-string                   # print the string used to hash
bun pm hash-print                    # print hash stored in lockfile
```

### `bun pm cache`

Manage Bun's global module cache.

```bash
bun pm cache                        # print the cache path
bun pm cache rm                    # clear the cache
```

### `bun pm migrate`

Migrate another package manager's lockfile without installing anything.

```bash
bun pm migrate
```

### `bun pm untrusted`

Print current untrusted dependencies with scripts.

```bash
bun pm untrusted
```

### `bun pm trust`

Run scripts for untrusted dependencies and add to `trustedDependencies`.

```bash
bun pm trust <names>
bun pm trust --all                  # trust all untrusted dependencies
```

### `bun pm default-trusted`

Print the default trusted dependencies list.

```bash
bun pm default-trusted
```

### `bun pm version`

Display current package version and increment it.

```bash
bun pm version                       # display current version
bun pm version patch                 # increment patch
bun pm version minor
bun pm version major
bun pm version 1.2.3                 # set specific version
bun pm version premajor --preid=beta # prerelease
bun pm version from-git              # use latest git tag
```

| Flag | Description |
|------|-------------|
| `--no-git-tag-version` | Skip git operations |
| `--allow-same-version` | Don't throw if version is the same |
| `--message=<val>` / `-m` | Custom commit message (use `%s` for version) |
| `--preid=<val>` | Prerelease identifier (e.g., `beta`) |
| `--force` / `-f` | Bypass dirty git history check |

Supports: `patch`, `minor`, `major`, `premajor`, `preminor`, `prepatch`, `prerelease`, `from-git`, or specific versions.

### `bun pm pkg`

Manage `package.json` data with get, set, delete, and fix operations. Supports dot and bracket notation.

```bash
bun pm pkg get name
bun pm pkg get name version
bun pm pkg get                       # get all fields
bun pm pkg get scripts.build         # nested field
bun pm pkg set name="my-package"
bun pm pkg set scripts.test="jest" version=2.0.0
bun pm pkg set private=true --json
bun pm pkg delete description
bun pm pkg delete scripts.test contributors[0]
bun pm pkg fix                       # fix/normalize package.json
```

## bun add

Adds a specific package as a dependency. See [cli.md](cli.md) for the full flag reference.

```bash
bun add preact
bun add zod@3.20.0
bun add --dev @types/react
bun add --peer @types/bun
bun add --exact react
bun add --global cowsay
bun add git@github.com:moment/moment.git
bun add zod@https://registry.npmjs.org/zod/-/zod-3.21.4.tgz
```

### Dependency Types

| Flag | Section in package.json |
|------|------------------------|
| (default) | `dependencies` |
| `--dev` / `-d` | `devDependencies` |
| `--optional` | `optionalDependencies` |
| `--peer` | `peerDependencies` |

### Version Specifiers

| Syntax | Result |
|--------|--------|
| `bun add zod` | `^<latest>` (caret range) |
| `bun add zod --exact` | `<latest>` (exact version) |
| `bun add zod@3.20.0` | `3.20.0` (exact) |
| `bun add zod@^3.0.0` | `^3.0.0` |
| `bun add zod@latest` | `^<latest>` |
| `bun add git@github.com:moment/moment.git` | Git URL |
| `bun add zod@https://...tgz` | Tarball URL |

## bun remove

Removes a package from the project's dependencies.

```bash
bun remove ts-node
bun remove react react-dom
```

Shares the same flag set as `bun install`.

## bun link

Registers a local package as "linkable" so it can be symlinked into other projects.

```bash
cd /path/to/cool-pkg
bun link                              # registers the package

cd /path/to/my-app
bun link cool-pkg                     # symlinks into node_modules
```

Shares the same flag set as `bun install` plus `--global` / `-g`.

## bun unlink

Unregisters a previously linked package.

```bash
cd /path/to/cool-pkg
bun unlink
```

Shares the same flag set as `bun link`.

## Environment Variables

| Variable | Description |
|----------|-------------|
| `BUN_INSTALL_GLOBAL_DIR` | Global packages directory |
| `BUN_INSTALL_BIN` | Global binaries directory |
| `BUN_INSTALL_GLOBAL_STORE` | Enable global virtual store |
| `npm_token` | Referenced in `install.scopes` as `$npm_token` |
| `npm_password` | Referenced in `install.scopes` as `$npm_password` |

See [environment.md](environment.md) for the full list.

## References

- <https://bun.sh/docs/cli/install>
- <https://bun.sh/docs/cli/add>
- <https://bun.sh/docs/cli/remove>
- <https://bun.sh/docs/cli/pm>
- <https://bun.sh/docs/cli/link>
- <https://bun.sh/docs/cli/unlink>
- <https://bun.sh/docs/runtime/bunfig>
