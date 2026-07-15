# Bun CLI Command Reference

Complete reference for all `bun` CLI commands, subcommands, and their flags.

> **Source of truth**: <https://bun.sh/docs/cli>. For package manager internals, see [package-manager.md](package-manager.md). For configuration, see [config.md](config.md).

## Command Overview

| Command | Description |
|---------|-------------|
| `bun <file>` | Execute a source file (shorthand for `bun run <file>`) |
| `bun run <file\|script>` | Execute a file or package.json script |
| `bun install` | Install all dependencies from package.json |
| `bun add <pkg>` | Add a dependency to the project |
| `bun remove <pkg>` | Remove a dependency from the project |
| `bun build <entry>` | Bundle JavaScript/TypeScript files |
| `bun test` | Run the Jest-compatible test runner |
| `bun x` / `bunx` | Auto-install and run a package |
| `bun create` | Scaffold a project from a template |
| `bun init` | Initialize a new Bun project interactively |
| `bun pm` | Package manager utilities |
| `bun upgrade` | Upgrade Bun to the latest version |
| `bun link` | Register a local package as linkable |
| `bun unlink` | Unregister a linked package |
| `bun completions` | Install shell completions |
| `bun dev` | Run the dev server (or a `dev` package.json script) |

## Global Flags

These flags apply to `bun` when running files or scripts:

| Flag | Type | Description |
|------|------|-------------|
| `--watch` | boolean | Auto-restart process on file change (hard restart) |
| `--hot` | boolean | Hot reload code without restarting process (preserves global state) |
| `--silent` | boolean | Don't print the script command |
| `--eval` / `-e` | string | Evaluate argument as a script |
| `--print` / `-p` | string | Evaluate and print the result |
| `--preload` / `-r` | string | Import a module before other modules |
| `--require` | string | Alias of `--preload` |
| `--import` | string | Alias of `--preload` |
| `--no-install` | boolean | Disable auto-install |
| `--install` | string | Configure auto-install: `auto`, `fallback`, `force` |
| `-i` | boolean | Auto-install dependencies (equivalent to `--install=fallback`) |
| `--prefer-offline` | boolean | Skip staleness checks for packages |
| `--prefer-latest` | boolean | Always check npm for latest matching versions |
| `--smol` | boolean | Use less memory, run GC more often |
| `--expose-gc` | boolean | Expose `gc()` on the global object |
| `--no-deprecation` | boolean | Suppress deprecation warnings |
| `--throw-deprecation` | boolean | Make deprecation warnings into errors |
| `--title` | string | Set the process title |
| `--zero-fill-buffers` | boolean | Force `Buffer.allocUnsafe()` to zero-fill |
| `--no-addons` | boolean | Throw error if `process.dlopen` is called |
| `--unhandled-rejections` | string | `strict`, `throw`, `warn`, `none`, `warn-with-error-code` |
| `--console-depth` | number | Default depth for console.log inspection (default: 2) |
| `--inspect` | string | Activate Bun's debugger |
| `--inspect-wait` | string | Activate debugger, wait for connection before executing |
| `--inspect-brk` | string | Activate debugger, set breakpoint on first line and wait |
| `--no-clear-screen` | boolean | Don't clear terminal on reload with `--hot`/`--watch` |
| `--no-macros` | boolean | Disable macros from being executed |
| `--define` / `-d` | string | Substitute K:V while parsing |
| `--drop` | string | Remove function calls (e.g., `--drop=console`) |
| `--loader` / `-l` | string | Parse files with `.ext:loader` |
| `--conditions` | string | Pass custom resolution conditions |
| `--main-fields` | string | Main fields to lookup in package.json |
| `--preserve-symlinks` | boolean | Preserve symlinks when resolving files |
| `--preserve-symlinks-main` | boolean | Preserve symlinks for the main entry point |
| `--extension-order` | string | File extension resolution order (default: `.tsx,.ts,.jsx,.js,.json`) |
| `--tsconfig-override` | string | Specify custom tsconfig.json |
| `--jsx-factory` | string | Function called for JSX elements (classic runtime) |
| `--jsx-fragment` | string | Function called for JSX fragments |
| `--jsx-import-source` | string | Module specifier for jsx/jsxs factory (default: `react`) |
| `--jsx-runtime` | string | `automatic` (default) or `classic` |
| `--jsx-side-effects` | boolean | Treat JSX elements as having side effects |
| `--ignore-dce-annotations` | boolean | Ignore tree-shaking annotations |
| `--port` | number | Default port for `Bun.serve` |
| `--fetch-preconnect` | string | Preconnect to a URL while code loads |
| `--max-http-header-size` | number | Max HTTP header size in bytes (default: 16384) |
| `--dns-result-order` | string | DNS lookup order: `verbatim`, `ipv4first`, `ipv6first` |
| `--use-system-ca` | boolean | Use system's trusted certificate authorities |
| `--use-openssl-ca` | boolean | Use OpenSSL's default CA store |
| `--use-bundled-ca` | boolean | Use bundled CA store |
| `--redis-preconnect` | boolean | Preconnect to `$REDIS_URL` at startup |
| `--sql-preconnect` | boolean | Preconnect to PostgreSQL at startup |
| `--user-agent` | string | Default User-Agent header for HTTP requests |
| `--env-file` | string | Load environment variables from file(s) |
| `--cwd` | string | Absolute path to resolve files and entrypoints from |
| `--config` / `-c` | string | Path to Bun config file (default: `$cwd/bunfig.toml`) |
| `--bun` / `-b` | boolean | Force script to use Bun's runtime instead of Node.js |

## bun run

Executes a source file (JS/TS/JSX/TSX) or a `package.json` script. Can also pipe code from stdin.

```bash
bun run index.js
bun run index.tsx
bun index.tsx                    # shorthand
bun run dev                      # runs package.json script
bun run                          # lists available scripts
echo "console.log('Hello')" | bun run -
```

### Resolution Order for `bun run`

1. `package.json` scripts
2. Source files (with extension resolution)
3. Binaries from project packages
4. System commands

### `bun run` Specific Flags

| Flag | Type | Description |
|------|------|-------------|
| `--silent` | boolean | Don't print the script command |
| `--if-present` | boolean | Exit without error if entrypoint doesn't exist |
| `--filter` / `-F` | string | Run a script in all workspace packages matching the pattern |
| `--workspaces` | boolean | Run a script in all workspace packages |
| `--parallel` | boolean | Run multiple scripts concurrently with prefixed output |
| `--sequential` | boolean | Run multiple scripts sequentially with prefixed output |
| `--no-exit-on-error` | boolean | Continue running other scripts when one fails |
| `--elide-lines` | number | Lines of script output shown with `--filter` (default: 10) |
| `--bun` / `-b` | boolean | Force the script to use Bun instead of Node.js |
| `--shell` | string | Control the shell used: `bun` or `system` |

## bun install

Installs all `dependencies`, `devDependencies`, `optionalDependencies`, and `peerDependencies` from `package.json`. Runs lifecycle scripts (`{pre|post}install`, `{pre|post}prepare`). Writes a `bun.lock` lockfile.

```bash
bun install
bun install react
bun install react@19.1.1
bun install --production
bun install --frozen-lockfile
bun install --global cowsay
bun ci                              # alias for --frozen-lockfile
```

### Key Flags

| Flag | Type | Description |
|------|------|-------------|
| `--production` | boolean | Don't install devDependencies |
| `--no-save` | boolean | Don't update package.json or save a lockfile |
| `--omit` | string | Exclude `dev`, `optional`, or `peer` dependencies |
| `--dev` | boolean | Add to devDependencies |
| `--optional` | boolean | Add to optionalDependencies |
| `--peer` | boolean | Add to peerDependencies |
| `--exact` | boolean | Add exact version instead of `^` range |
| `--yarn` | boolean | Write a yarn.lock file (yarn v1) |
| `--frozen-lockfile` | boolean | Disallow changes to lockfile |
| `--save-text-lockfile` | boolean | Save a text-based lockfile |
| `--lockfile-only` | boolean | Generate lockfile without installing |
| `--global` / `-g` | boolean | Install globally |
| `--force` | boolean | Always request latest versions, reinstall all |
| `--dry-run` | boolean | Don't install anything |
| `--no-cache` | boolean | Ignore manifest cache entirely |
| `--no-verify` | boolean | Skip integrity verification of new packages |
| `--ignore-scripts` | boolean | Skip lifecycle scripts |
| `--filter` | string | Install for matching workspaces |
| `--backend` | string | `clonefile`, `hardlink`, `symlink`, `copyfile` |
| `--linker` | string | `isolated` or `hoisted` |
| `--registry` | string | Use a specific registry |
| `--ca` / `--cafile` | string | Certificate Authority |
| `--cache-dir` | string | Cache directory path |
| `--concurrent-scripts` | number | Max concurrent lifecycle script jobs |
| `--network-concurrency` | number | Max concurrent network requests (default: 48) |
| `--minimum-release-age` | number | Filter packages published more recently than N seconds ago |
| `--cpu` | string | Override CPU architecture (e.g., `x64`, `arm64`) |
| `--os` | string | Override OS (e.g., `linux`, `darwin`) |
| `--only-missing` | boolean | Only add deps not already present |
| `--trust` | boolean | Add to trustedDependencies and install |
| `--verbose` / `--silent` / `--no-progress` / `--no-summary` | boolean | Output control |

## bun add

Adds a specific package as a dependency.

```bash
bun add preact
bun add zod@3.20.0
bun add --dev @types/react
bun add -d @types/react
bun add --optional lodash
bun add --peer @types/bun
bun add react --exact
bun add --global cowsay
bun add git@github.com:moment/moment.git
```

### Key Flags (subset of install flags)

| Flag | Short | Description |
|------|-------|-------------|
| `--dev` | `-d` / `-D` | Add to devDependencies |
| `--optional` | | Add to optionalDependencies |
| `--peer` | | Add to peerDependencies |
| `--exact` | `-E` | Use exact version, not `^` range |
| `--global` | `-g` | Install globally |
| `--force` | `-f` | Always request latest, reinstall all |
| `--production` | `-p` | Don't install devDependencies |
| `--config` | `-c` | Path to config file |
| `--analyze` | `-a` | Recursively analyze with Bun's bundler |
| `--yarn` | `-y` | Write yarn.lock (yarn v1) |

Also inherits all networking, cache, and lockfile flags from `bun install`.

## bun remove

Removes a package from the project's dependencies.

```bash
bun remove ts-node
bun remove react react-dom
```

Shares the same flag set as `bun install` (lockfile, networking, cache, etc.).

## bun x / bunx

Auto-installs and runs packages from npm. Equivalent of `npx` or `yarn dlx`.

```bash
bunx cowsay "Hello world!"
bunx prisma migrate
bunx prettier foo.js
bunx uglify-js@3.14.0 app.js
bunx -p @angular/cli ng new my-app
bunx --bun vite dev
```

### Flags

| Flag | Type | Description |
|------|------|-------------|
| `--bun` | boolean | Force the command to run with Bun instead of Node.js |
| `-p` / `--package` | string | Specify package when binary name differs from package name |
| `--no-install` | boolean | Skip installation if package is not already installed |
| `--verbose` | boolean | Enable verbose output during installation |
| `--silent` | boolean | Suppress output during installation |

## bun create

Scaffolds a new Bun project from a template. Sources include: React components, `create-<template>` npm packages, GitHub repos, or local templates.

```bash
bun create ./MyComponent.jsx       # from a React component
bun create remix                   # from npm (downloads create-remix)
bun create <user>/<repo>           # from GitHub
bun create <user>/<repo> mydir     # from GitHub with custom destination
```

### Flags

| Flag | Description |
|------|-------------|
| `--force` | Overwrite existing files |
| `--no-install` | Skip installing node_modules and tasks |
| `--no-git` | Don't initialize a git repository |
| `--open` | Start and open in-browser after finish |

### Environment Variables

- `GITHUB_API_DOMAIN` — GitHub domain for downloads (GitHub Enterprise/proxy)
- `GITHUB_TOKEN` / `GITHUB_ACCESS_TOKEN` — Access private repos, avoid rate limits

## bun init

Scaffolds a new Bun project interactively. Creates `package.json`, `tsconfig.json`, entry point, and `README.md`.

```bash
bun init                           # in current directory
bun init my-app
bun init path/to/directory
bun init -y                        # accept all defaults
bun init --react                    # React project
bun init --react=tailwind           # React + Tailwind CSS
bun init --react=shadcn             # React + @shadcn/ui
bun init -m                         # minimal (type defs only)
```

### Flags

| Flag | Short | Description |
|------|-------|-------------|
| `--yes` | `-y` | Accept all default prompts without asking |
| `--minimal` | `-m` | Only initialize type definitions (skip app scaffolding) |
| `--react` | | Scaffold a React project. Accepts: `tailwind`, `shadcn` presets |

## bun build

Bundles JavaScript/TypeScript files using Bun's native bundler. See [bundler.md](bundler.md) for the full bundler reference.

```bash
bun build ./index.tsx --outdir ./out
bun build ./cli.tsx --outfile mycli --compile
bun build ./index.tsx --outdir ./out --watch
bun build ./index.tsx --outdir ./out --minify
```

## bun test

Runs tests using Bun's built-in Jest-compatible test runner. See [test-runner.md](test-runner.md) for the full test runner reference.

```bash
bun test
bun test ./test/specific-file.test.ts
bun test --test-name-pattern addition
bun test --watch
bun test --coverage
```

## bun pm

Package manager utilities. See [package-manager.md](package-manager.md) for the full `pm` subcommand reference.

```bash
bun pm pack
bun pm ls
bun pm cache
bun pm whoami
bun pm trust --all
```

## bun upgrade

Upgrades Bun itself to the latest stable or canary version.

```bash
bun upgrade                        # upgrade to latest stable
bun upgrade --canary                # upgrade to latest canary build
bun upgrade --stable                # switch back to stable from canary
```

If installed via a package manager (Homebrew, Scoop, npm), use that manager instead:

```bash
brew upgrade bun
scoop update bun
npm install -g bun
```

Verify with:

```bash
bun --version
bun --revision
```

## bun link

Registers a local package as "linkable" so it can be symlinked into other projects.

```bash
cd /path/to/cool-pkg
bun link                            # registers the package

cd /path/to/my-app
bun link cool-pkg                   # symlinks into node_modules
```

Shares the same flag set as `bun install` (lockfile, networking, cache, etc.) plus `--global` / `-g`.

## bun unlink

Unregisters a previously linked package.

```bash
cd /path/to/cool-pkg
bun unlink
```

Shares the same flag set as `bun link`.

## bun completions

Installs shell completions for bash, zsh, and fish. The completion scripts are auto-generated from the same parameter definitions used for argument parsing and help text.

```bash
bun completions
```

## bun dev

There is no dedicated `bun dev` CLI command. The `bun dev` shorthand works because `bun run` can be invoked without the `run` keyword. If a built-in `bun` command has the same name, the built-in command takes precedence, but `dev` is not a built-in, so it resolves to the `dev` script in `package.json`:

```json
{
  "scripts": {
    "dev": "bun server.ts"
  }
}
```

```bash
bun dev                             # shorthand for `bun run dev`
```

## Watch Mode vs Hot Mode

| Feature | `--watch` | `--hot` |
|---------|-----------|---------|
| Behavior | Hard restarts process | Soft reloads code |
| Global state | Lost | Preserved (`globalThis` persists) |
| HTTP servers | Restarted | Not restarted, handler updated |
| Speed | Fast (native FS watchers: kqueue/inotify) | Faster (no process restart) |

## References

- <https://bun.sh/docs/cli>
- <https://bun.sh/docs/runtime>
