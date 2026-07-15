# Bun Configuration (bunfig.toml)

Reference for Bun's configuration file `bunfig.toml`, covering all sections, fields, defaults, and environment variable mappings.

> **Source of truth**: <https://bun.sh/docs/runtime/bunfig>. For CLI flags, see [cli.md](cli.md). For environment variables, see [environment.md](environment.md).

## Overview

`bunfig.toml` is Bun's configuration file. It's optional — Bun works without it. It only contains Bun-specific settings; Bun relies on `package.json` and `tsconfig.json` for other config.

### Global vs Local

- **Local**: Put `bunfig.toml` in your project root alongside `package.json`.
- **Global**: Create a `.bunfig.toml` file at `$HOME/.bunfig.toml` or `$XDG_CONFIG_HOME/.bunfig.toml`.
- **Merging**: If both exist, they are shallow-merged, with local overriding global.
- **CLI flags** override `bunfig` settings where applicable.

## Top-Level (Runtime) Fields

### `preload`

- **Type**: Array of strings
- **Description**: Scripts/plugins to execute before running a file or script.

```toml
preload = ["./preload.ts"]
```

### `jsx`

Configure how Bun handles JSX. Also settable in `tsconfig.json` `compilerOptions`.

```toml
jsx = "react"
jsxFactory = "h"
jsxFragment = "Fragment"
jsxImportSource = "react"
```

### `smol`

- **Type**: Boolean
- **Description**: Enable `smol` mode — reduces memory usage at the cost of performance.

```toml
smol = true
```

### `logLevel`

- **Type**: `"debug"` | `"warn"` | `"error"`

```toml
logLevel = "debug"
```

### `define`

Replaces global identifiers with constant expressions wherever they appear. Values are parsed as JSON (single-quoted strings supported; `'undefined'` becomes `undefined` in JS).

```toml
[define]
"process.env.bagel" = "'lox'"
```

### `loader`

Maps file extensions to loaders for file types Bun doesn't support natively.

Supported loaders: `jsx`, `js`, `ts`, `tsx`, `css`, `file`, `json`, `toml`, `wasm`, `napi`, `base64`, `dataurl`, `text`.

```toml
[loader]
".bagel" = "tsx"
```

### `telemetry`

- **Type**: Boolean
- **Default**: `true` (enabled)
- **Description**: Enables/disables anonymous analytics/crash reports. Equivalent to `DO_NOT_TRACK` env var.

```toml
telemetry = false
```

### `env`

- **Type**: Boolean or object
- **Default**: `true` (loads `.env` files automatically)
- **Description**: Configure automatic `.env` file loading. Set to `false` to disable (useful in production/CI). Files passed explicitly with `--env-file` still load.

```toml
env = false
```

```toml
[env]
file = false
```

### `console`

```toml
[console]
depth = 3                           # default depth for console.log() inspection (default: 2)
```

CLI flag: `--console-depth`.

## `[serve]`

Configures `Bun.serve` and `bun run` when serving HTTP.

### `serve.port`

- **Type**: Number
- **Default**: `3000`
- **Environment variables**: `BUN_PORT` or `PORT`
- **CLI flag**: `--port`

```toml
[serve]
port = 3000
```

## `[test]`

Configures the `bun test` command. See [test-runner.md](test-runner.md) for test runner details.

| Field | Type | Default | Description |
|-------|------|---------|-------------|
| `root` | string | `.` | Root directory to run tests from |
| `preload` | array | | Scripts to load before tests |
| `pathIgnorePatterns` | array | | Glob patterns to exclude from test discovery |
| `smol` | boolean | `false` | Reduce memory usage during tests |
| `coverage` | boolean | `false` | Generate coverage profile |
| `coverageThreshold` | number\|object | | Coverage threshold (exits non-zero if not met) |
| `coverageSkipTestFiles` | boolean | `false` | Skip test files when computing coverage |
| `coverageIgnoreSourcemaps` | boolean | `false` | Report against transpiled output |
| `coveragePathIgnorePatterns` | string\|array | | Exclude files from coverage reports |
| `coverageReporter` | array | `["text"]` | Coverage reporters (`text`, `lcov`) |
| `coverageDir` | string | `"coverage"` | Directory for coverage reports |
| `randomize` | boolean | `false` | Run tests in random order |
| `seed` | number | | Random seed for reproducible randomization |
| `rerunEach` | number | `0` | Re-run each test file N times |
| `retry` | number | `0` | Default retry count for failed tests |
| `concurrentTestGlob` | string | | Glob for files that run tests concurrently |
| `onlyFailures` | boolean | `false` | Only display failed tests in output |

### `[test.reporter]`

```toml
[test.reporter]
dots = true                         # one dot per test
junit = "test-results.xml"          # JUnit XML output path
```

### Coverage Threshold Examples

```toml
# Single threshold (90% line and function level)
[test]
coverageThreshold = 0.9

# Separate thresholds per metric
[test]
coverageThreshold = { lines = 0.7, functions = 0.8, statements = 0.9 }
```

### Path Ignore Patterns

```toml
[test]
pathIgnorePatterns = ["vendor/**", "submodules/**", "fixtures/**"]
```

## `[install]`

Configures `bun install` behavior. See [package-manager.md](package-manager.md) for package manager details.

| Field | Type | Default | Description |
|-------|------|---------|-------------|
| `optional` | boolean | `true` | Install optional dependencies |
| `dev` | boolean | `true` | Install dev dependencies |
| `peer` | boolean | `true` | Install peer dependencies |
| `production` | boolean | `false` | Don't install devDependencies |
| `exact` | boolean | `false` | Use exact version, not `^` range |
| `ignoreScripts` | boolean | `false` | Skip lifecycle scripts |
| `concurrentScripts` | number | 2x CPU cores | Max concurrent lifecycle scripts |
| `saveTextLockfile` | boolean | `true` | Save text-based `bun.lock` (vs binary `bun.lockb`) |
| `auto` | string | `"auto"` | Auto-install behavior: `auto`, `force`, `disable`, `fallback` |
| `prefer` | string | `"online"` | Registry resolution: `online`, `offline`, `latest` |
| `frozenLockfile` | boolean | `false` | Disallow lockfile changes |
| `dryRun` | boolean | `false` | Don't install anything |
| `globalDir` | string | `~/.bun/install/global` | Global packages directory |
| `globalBinDir` | string | `~/.bun/bin` | Global binaries directory |
| `linkWorkspacePackages` | boolean | `true` | Link workspace packages in monorepo |
| `linker` | string | varies | `isolated` or `hoisted` |
| `globalStore` | boolean | `false` | Share package installations across projects |
| `minimumReleaseAge` | number | `null` | Only install packages published N+ seconds ago |

### `install.auto` Values

| Value | Description |
|-------|-------------|
| `"auto"` | Resolve from local `node_modules` if exists; otherwise auto-install |
| `"force"` | Always auto-install, even if `node_modules` exists |
| `"disable"` | Never auto-install |
| `"fallback"` | Check local first, then auto-install missing packages (CLI: `bun -i`) |

### `install.prefer` Values

| Value | Description | CLI equivalent |
|-------|-------------|----------------|
| `"online"` | Check registry for stale packages as needed | |
| `"offline"` | Skip staleness checks, resolve from cache | `--prefer-offline` |
| `"latest"` | Always check npm for latest versions | `--prefer-latest` |

### `install.linker` Values

| Value | Description |
|-------|-------------|
| `"hoisted"` | Link dependencies in a shared `node_modules` directory |
| `"isolated"` | Link dependencies inside each package installation (pnpm-style) |

Default: `"isolated"` for new workspaces, `"hoisted"` for new single-package projects and existing projects (pre-v1.3.2).

### `install.registry`

```toml
# Simple string
[install]
registry = "https://registry.npmjs.org"

# With token
registry = { url = "https://registry.npmjs.org", token = "123456" }

# With username/password
registry = "https://username:password@registry.npmjs.org"
```

### `install.scopes`

Configure a registry for a particular npm scope (e.g., `@myorg/<package>`). Supports `$variable` notation for environment variables.

```toml
[install.scopes]
# Registry as string
myorg = "https://username:password@registry.myorg.com/"

# With username/password (env vars supported)
myorg = { username = "myusername", password = "$npm_password", url = "https://registry.myorg.com/" }

# With token (env vars supported)
myorg = { token = "$npm_token", url = "https://registry.myorg.com/" }
```

### `install.ca` / `install.cafile`

```toml
[install]
ca = "-----BEGIN CERTIFICATE-----\n...\n-----END CERTIFICATE-----"
cafile = "path/to/cafile"
```

### `install.cache`

```toml
[install.cache]
dir = "~/.bun/install/cache"        # default cache directory
disable = false                     # don't load from global cache
disableManifest = false             # always resolve latest from registry
```

### `install.lockfile`

```toml
[install.lockfile]
save = true                         # generate lockfile on install (default)
print = "yarn"                      # generate non-Bun lockfile alongside bun.lock
```

### `install.security.scanner`

```toml
[install.security]
scanner = "@oven/bun-security-scanner"
```

When configured: auto-install is disabled, packages are scanned before installation, installation is cancelled if fatal issues are found, and security warnings are displayed.

### `install.publicHoistPattern` / `install.hoistPattern`

When using `"isolated"` linker:

```toml
[install]
publicHoistPattern = ["*eslint*", "*prettier*"]  # hoist to root node_modules
hoistPattern = ["*"]                                # hoist to virtual store root
```

### `install.minimumReleaseAge`

```toml
[install]
minimumReleaseAge = 259200          # 3 days
minimumReleaseAgeExcludes = ["@types/bun", "typescript"]
```

## `[run]`

Configures `bun run`. Only loads local project's `bunfig.toml` (no global `.bunfig.toml` check).

| Field | Type | Default | Description |
|-------|------|---------|-------------|
| `shell` | string | `"bun"` on Windows, `"system"` elsewhere | Shell used for package.json scripts |
| `bun` | boolean | `true` if `node` not in `$PATH` | Auto-alias `node` to `bun` |
| `silent` | boolean | `false` | Suppress printing the command being run |
| `elide-lines` | number | `10` | Lines of output per script with `--filter` |
| `noOrphans` | boolean | | Don't leave orphan processes (kills descendants on exit) |

### `run.bun`

Auto-alias `node` to `bun`. Prepends `$PATH` with a `node` symlink pointing to the `bun` binary. Works recursively for all scripts/executables invoked by `bun run` or `bun`. Equivalent to prefixing commands with `--bun`.

### `run.shell`

Controls whether `package.json` scripts are executed with Bun's built-in shell or the system shell (`/bin/sh` on Unix, `cmd.exe` on Windows). See [shell.md](shell.md) for the Bun shell.

### `run.noOrphans`

Don't leave orphan processes behind. Bun watches the parent process and exits when it goes away (even on `SIGKILL`). On exit, recursively `SIGKILL`s all descendant processes. Linux/macOS only.

Environment variable: `BUN_FEATURE_FLAG_NO_ORPHANS=1`.

## Configuration Precedence

Configuration is resolved in this order (highest to lowest):

1. **CLI flags** (e.g., `--port 4000`)
2. **Local `bunfig.toml`** (in project root)
3. **Global `.bunfig.toml`** (`$HOME/.bunfig.toml` or `$XDG_CONFIG_HOME/.bunfig.toml`)
4. **Built-in defaults**

For `[run]` section specifically, only local `bunfig.toml` is checked (no global).

## References

- <https://bun.sh/docs/runtime/bunfig>
