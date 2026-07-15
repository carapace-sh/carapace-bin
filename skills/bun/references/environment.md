# Bun Environment Variables

Reference for environment variables recognized by Bun, `.env` file handling, and environment-related configuration.

> **Source of truth**: <https://bun.sh/docs/runtime/bunfig> and <https://bun.sh/docs/runtime/env>. For bunfig.toml configuration, see [config.md](config.md).

## Environment Variables Recognized by Bun

| Variable | Related Setting | Description |
|----------|----------------|-------------|
| `DO_NOT_TRACK` | `telemetry` | Disables analytics/telemetry (same as `telemetry = false` in bunfig.toml) |
| `BUN_PORT` | `serve.port` | Default port for `Bun.serve` |
| `PORT` | `serve.port` | Alternative env var for `Bun.serve` port |
| `BUN_INSTALL_GLOBAL_DIR` | `install.globalDir` | Directory for globally installed packages |
| `BUN_INSTALL_BIN` | `install.globalBinDir` | Directory where globally-installed package bins are linked |
| `BUN_INSTALL_GLOBAL_STORE` | `install.globalStore` | Enable/configure global virtual store |
| `BUN_FEATURE_FLAG_NO_ORPHANS` | `run.noOrphans` | Enable no-orphans mode (set to `=1`) |
| `NODE_ENV` | | Standard Node.js env var, affects bundler `--production` flag |
| `CLAUDECODE` | test runner | When set to `1`, quiet output for Claude Code (failure details only) |
| `REPL_ID` | test runner | When set to `1`, quiet output for Replit |
| `AGENT` | test runner | When set to `1`, generic AI agent flag for quiet test output |
| `GITHUB_TOKEN` | `bun create` | Access private repos and avoid rate limits (alias: `GITHUB_ACCESS_TOKEN`) |
| `GITHUB_API_DOMAIN` | `bun create` | GitHub domain for downloads (GitHub Enterprise/proxy) |
| `npm_token` | `install.scopes` | Referenced via `$npm_token` in scope registry config |
| `npm_password` | `install.scopes` | Referenced via `$npm_password` in scope registry config |

## `.env` File Handling

### Automatic Loading

Bun loads `.env` files by default when running files, scripts, or tests. This behavior can be configured:

```toml
# bunfig.toml
env = false              # disable automatic .env loading
```

```toml
[env]
file = false             # alternative syntax
```

### Explicit `.env` File Loading

Use the `--env-file` flag to load environment variables from specific file(s):

```bash
bun --env-file .env.production run index.ts
bun --env-file .env --env-file .env.local run index.ts
```

Files passed with `--env-file` are loaded even when automatic `.env` loading is disabled.

### `.env` File Resolution Order

Bun loads `.env` files in this order (later files override earlier ones):

1. `.env`
2. `.env.local` (or `.env.<NODE_ENV>.local`)
3. `.env.development` / `.env.production` / `.env.test` (based on `NODE_ENV`)

## `Bun.env`

`Bun.env` is an alias for `process.env`. It provides access to all environment variables:

```ts
console.log(Bun.env.HOME);
console.log(Bun.env.NODE_ENV);
```

## Environment Variables in `install.scopes`

In `bunfig.toml`, the `[install.scopes]` section supports referencing environment variables with `$variable` notation:

```toml
[install.scopes]
myorg = { token = "$npm_token", url = "https://registry.myorg.com/" }
myorg = { username = "myusername", password = "$npm_password", url = "https://registry.myorg.com/" }
```

This allows keeping secrets out of the config file and in environment variables instead.

## Environment Variables in Bundling

The `Bun.build()` API and `bun build` CLI support environment variable injection:

```bash
# Inline all process.env.* references as string literals
bun build ./index.tsx --outdir ./out --env inline

# Only inline matching prefix
bun build ./index.tsx --outdir ./out --env "FOO_PUBLIC_*"
```

```ts
await Bun.build({
  entrypoints: ['./index.tsx'],
  outdir: './out',
  env: "inline",           // or "disable", or "FOO_PUBLIC_*"
});
```

## Environment Variables in the Bun Shell

The `$` template literal tag respects environment variables:

```ts
import { $ } from "bun";

// Inline env vars (like bash)
await $`FOO=foo some-command`;

// String interpolation (escaped by default)
const val = "456";
await $`FOO=${val} some-command`;

// Per-command environment
await $`some-command`.env({ ...process.env, FOO: "bar" });

// Global default environment
$.env({ FOO: "bar" });
```

## XDG Config Home

The global `.bunfig.toml` file is searched at:

1. `$HOME/.bunfig.toml`
2. `$XDG_CONFIG_HOME/.bunfig.toml`

If `XDG_CONFIG_HOME` is set, it takes precedence over `$HOME` for the global config file location.

## References

- <https://bun.sh/docs/runtime/bunfig>
- <https://bun.sh/docs/runtime/env>
- <https://bun.sh/docs/runtime/shell>
