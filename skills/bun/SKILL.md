---
name: bun
description: >
  Use when working with the bun CLI — Bun's all-in-one JavaScript runtime and
  toolkit. Covers the CLI command reference (install, add, remove, run, build,
  test, x, create, init, pm, upgrade, link, unlink), bunfig.toml configuration,
  environment variables, the package manager and lockfile, the test runner,
  the bundler, the Bun runtime API (Bun.serve, Bun.file, Bun.spawn, Bun.write,
  hashing, transpiler), the Bun shell ($ template literal), module resolution
  and loaders, and globals. Triggers on: "bun", "bunx", "bun install", "bun add",
  "bun remove", "bun run", "bun build", "bun test", "bun x", "bun create",
  "bun init", "bun pm", "bun upgrade", "bun link", "bun unlink", "bunfig.toml",
  "bun.lock", "Bun.serve", "Bun.file", "Bun.write", "Bun.spawn", "Bun.hash",
  "Bun.password", "bun:test", "Bun.build", "Bun.Transpiler", "bun shell",
  "$ template", "import.meta", "JavaScriptCore", "Bun.serve", "HTMLRewriter",
  "Bun.SQL", "Bun.RedisClient", "Bun.Glob", "Bun.plugin", "bun:ffi".
user-invocable: true
---

# Bun — JavaScript Runtime & Toolkit In-Depth Reference

Comprehensive reference for the Bun CLI (<https://bun.sh/docs>), covering the command reference, configuration, package manager, test runner, bundler, runtime API, shell, and module resolution. Bun is an all-in-one JavaScript runtime built on JavaScriptCore (not V8), with a native bundler, transpiler, test runner, and package manager written in Zig and C++.

## Data Flow

```
Source file (.js/.ts/.jsx/.tsx)
  → transpiler (JS/TS/JSX → vanilla JS, dead code elimination)
    → module resolver (node_modules, exports field, tsconfig paths)
      → runtime (JavaScriptCore engine executes)
        → globals (Web API + Node.js + Bun-specific)
          → Bun API (serve, file, spawn, write, hash, etc.)
```

## Sub-Resources

Load the reference that matches your task. When in doubt, load multiple references.

| Keywords | Reference |
|----------|----------|
| command, subcommand, flag, bun install, bun add, bun remove, bun run, bun x, bunx, bun create, bun init, bun build, bun test, bun pm, bun upgrade, bun completions, bun link, bun unlink, bun dev, bun ci, CLI reference, --watch, --hot, --silent, --eval, --print, --preload, --inspect, --define, --loader, --env-file, --port, --smol, --no-install, --install, --bun, --filter, --workspaces, --parallel, --sequential | [references/cli.md](references/cli.md) |
| bunfig.toml, configuration, config file, [install], [test], [run], [serve], preload, jsx, smol, logLevel, define, loader, telemetry, env, console, registry, scopes, ca, cafile, cache, lockfile, linker, globalStore, publicHoistPattern, hoistPattern, frozenLockfile, dryRun, auto, prefer, saveTextLockfile, ignoreScripts, concurrentScripts, minimumReleaseAge, security scanner, globalDir, globalBinDir, linkWorkspacePackages, noOrphans, elide-lines, shell, run.bun, run.silent | [references/config.md](references/config.md) |
| package manager, bun install, bun add, bun remove, bun link, bun unlink, bun pm, bun pm pack, bun pm bin, bun pm ls, bun pm whoami, bun pm cache, bun pm migrate, bun pm untrusted, bun pm trust, bun pm hash, bun pm version, bun pm pkg, bun.lock, lockfile, binary lockfile, text lockfile, trustedDependencies, backend, clonefile, hardlink, symlink, copyfile, linker, hoisted, isolated, global install, workspace, filter, omit, production, exact, yarn, frozen-lockfile, dry-run, force, no-cache, network-concurrency, concurrent-scripts, registry, scopes, ca, cafile, backend | [references/package-manager.md](references/package-manager.md) |
| test runner, bun test, bun:test, test, expect, describe, beforeAll, beforeEach, afterEach, afterAll, mock, jest.fn, snapshot, toMatchSnapshot, coverage, --coverage, --timeout, --rerun-each, --retry, --concurrent, --max-concurrency, --randomize, --seed, --bail, --test-name-pattern, --reporter, --reporter-outfile, --dots, --update-snapshots, --watch, --preload, --todo, test.concurrent, test.serial, test.failing, test.todo, JUnit, CI, CLAUDECODE, AGENT, pathIgnorePatterns, coverageThreshold | [references/test-runner.md](references/test-runner.md) |
| bundler, bun build, Bun.build, entrypoints, outdir, outfile, target, browser, bun, node, format, esm, cjs, iife, splitting, minify, minify-syntax, minify-whitespace, minify-identifiers, keep-names, sourcemap, linked, inline, external, external, define, drop, loader, banner, footer, naming, entry-naming, chunk-naming, asset-naming, root, publicPath, public-path, packages, no-bundle, compile, bytecode, --compile, standalone executable, windows-icon, windows-title, metafile, metafile-md, env, conditions, jsx, react-fast-refresh, react-compiler, css-chunking, optimizeImports, feature, features, BuildArtifact, BuildOutput, BuildMessage, ResolveMessage | [references/bundler.md](references/bundler.md) |
| runtime, Bun global, Bun.serve, Bun.file, Bun.write, Bun.spawn, Bun.spawnSync, Bun.hash, Bun.password, Bun.hash, Bun.CryptoHasher, Bun.Transpiler, Bun.peek, Bun.deepEquals, Bun.inspect, Bun.sleep, Bun.sleepSync, Bun.nanoseconds, Bun.which, Bun.randomUUIDv7, Bun.escapeHTML, Bun.stringWidth, Bun.stripANSI, Bun.wrapAnsi, Bun.fileURLToPath, Bun.pathToFileURL, Bun.resolveSync, Bun.gzipSync, Bun.gunzipSync, Bun.deflateSync, Bun.inflateSync, Bun.zstdCompressSync, Bun.zstdDecompressSync, Bun.readableStreamToText, Bun.readableStreamToJSON, Bun.readableStreamToArray, Bun.env, Bun.main, Bun.version, Bun.revision, Bun.Terminal, HTMLRewriter, Bun.Glob, Bun.plugin, Bun.semver, Bun.TOML, Bun.markdown, Bun.color, Bun.Image, Bun.mmap, Bun.gc, Bun.generateHeapSnapshot, Bun.ArrayBufferSink, Worker, import.meta, globals, Web API, Node.js globals, module resolution, load​er, extension order, tsconfig paths, exports field, conditions, watch mode, hot mode, --watch, --hot, auto-install | [references/runtime.md](references/runtime.md) |
| shell, Bun shell, $ template literal, $.quiet, $.text, $.json, $.lines, $.blob, $.nothrow, $.throws, $.env, $.cwd, $.braces, $.escape, redirection, pipe, command substitution, builtin commands, cd, ls, rm, echo, pwd, cat, touch, mkdir, which, mv, exit, true, false, yes, seq, dirname, basename, .sh file, shell injection, argument injection, security | [references/shell.md](references/shell.md) |
| environment variable, BUN_PORT, PORT, BUN_INSTALL_GLOBAL_DIR, BUN_INSTALL_BIN, BUN_INSTALL_GLOBAL_STORE, BUN_FEATURE_FLAG_NO_ORPHANS, DO_NOT_TRACK, NODE_ENV, npm_token, GITHUB_TOKEN, GITHUB_API_DOMAIN, CLAUDECODE, REPL_ID, AGENT, XDG_CONFIG_HOME, env file, --env-file, .env | [references/environment.md](references/environment.md) |

## Quick Guide

- **What CLI commands does bun have?** → [references/cli.md](references/cli.md)
- **How do I configure bun with bunfig.toml?** → [references/config.md](references/config.md)
- **How does the package manager work?** → [references/package-manager.md](references/package-manager.md)
- **How do I run tests?** → [references/test-runner.md](references/test-runner.md)
- **How do I bundle code?** → [references/bundler.md](references/bundler.md)
- **What runtime APIs does Bun provide?** → [references/runtime.md](references/runtime.md)
- **How does the Bun shell work?** → [references/shell.md](references/shell.md)
- **What environment variables does bun use?** → [references/environment.md](references/environment.md)
- **How do I install dependencies?** → [references/package-manager.md](references/package-manager.md)
- **How do I add a dependency?** → [references/cli.md](references/cli.md) and [references/package-manager.md](references/package-manager.md)
- **How do I run a file or script?** → [references/cli.md](references/cli.md)
- **How do I configure the test runner?** → [references/config.md](references/config.md) and [references/test-runner.md](references/test-runner.md)
- **How do I create a standalone executable?** → [references/bundler.md](references/bundler.md)
- **How do I start an HTTP server?** → [references/runtime.md](references/runtime.md)
- **How do I read/write files?** → [references/runtime.md](references/runtime.md)
- **How do I spawn child processes?** → [references/runtime.md](references/runtime.md)
- **How do I use the Bun shell?** → [references/shell.md](references/shell.md)
- **How do I configure auto-install?** → [references/config.md](references/config.md)

## Cross-Project References

- For **Node.js** APIs (`fs`, `path`, `http`, `stream`, `crypto`, `util`, etc.) and Node.js module system internals, see the [Node.js documentation](https://nodejs.org/docs/latest/api/) — Bun implements Node.js compatibility but this skill covers only Bun-specific behavior and differences.
- For **TypeScript** configuration (`tsconfig.json`, `compilerOptions`, JSX settings), see the [TypeScript handbook](https://www.typescriptlang.org/docs/) — Bun reads `tsconfig.json` but this skill documents only Bun-specific configuration in `bunfig.toml`.
- For carapace-specific bun completion (the `bun` completer in `completers/common/bun_completer/`), see the **carapace** skill (in this repo).
