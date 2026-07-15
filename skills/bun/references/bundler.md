# Bun Bundler

Reference for Bun's native bundler: `bun build` CLI and `Bun.build()` JavaScript API. Covers targets, formats, loaders, minification, sourcemaps, code splitting, standalone executables, and all options.

> **Source of truth**: <https://bun.sh/docs/bundler>. For CLI command reference, see [cli.md](cli.md).

## Overview

Bun's native bundler is accessible via the `bun build` CLI command or the `Bun.build()` JavaScript API.

```bash
bun build ./index.tsx --outdir ./out
```

```ts
const result = await Bun.build({
  entrypoints: ['./index.tsx'],
  outdir: './out',
});
```

## Targets

The `target` option specifies the intended execution environment for the bundle.

| Target | Description |
|--------|-------------|
| `"browser"` | **Default.** Prioritizes the `"browser"` export condition. Built-in `node:*` modules are importable but `fs.readFile` etc. won't work at runtime. |
| `"bun"` | For bundles that run in the Bun runtime. Generated bundles are marked with `// @bun` pragma. If entrypoint has a Bun shebang, defaults to `"bun"`. |
| `"node"` | Prioritizes the `"node"` export condition. Does not polyfill `Bun` global or `bun:*` modules. |

If a `format: "cjs"` is used with `target: "bun"`, the `// @bun @bun-cjs` pragma is added.

## Formats

| Format | Description |
|--------|-------------|
| `"esm"` | **Default.** ES Module format. Supports top-level await and `import.meta`. |
| `"cjs"` | **Experimental.** CommonJS format. When chosen, default target changes from `"browser"` to `"node"`. |
| `"iife"` | **Experimental.** Immediately Invoked Function Expression. |

## Loaders (Content Types)

| Extension(s) | Loader | Behavior |
|-------------|--------|----------|
| `.js` `.jsx` `.cjs` `.mjs` `.mts` `.cts` `.ts` `.tsx` | Built-in transpiler | Transpiles TS/JSX to JS. Dead code elimination and tree shaking. Does not down-convert syntax. |
| `.json` | JSON | Parsed and inlined as JS object |
| `.jsonc` | JSONC | JSON with comments, parsed and inlined |
| `.toml` | TOML | Parsed and inlined as JS object |
| `.yaml` `.yml` | YAML | Parsed and inlined as JS object |
| `.txt` | Text | Contents read and inlined as string |
| `.html` | HTML | Referenced assets (scripts, stylesheets, images) are bundled |
| `.css` | CSS | Bundled into a single `.css` file in outdir |
| `.node` `.wasm` | Asset | Treated as assets during bundling |

Unrecognized extensions are treated as external files — copied as-is into `outdir`, import resolved as a path.

### Custom Loaders

```ts
await Bun.build({
  loader: {
    ".png": "dataurl",
    ".txt": "file",
  },
});
```

Available loader types: `"js"`, `"jsx"`, `"ts"`, `"tsx"`, `"css"`, `"json"`, `"jsonc"`, `"toml"`, `"yaml"`, `"text"`, `"file"`, `"napi"`, `"wasm"`, `"html"`.

## Bun.build() API Options

### `entrypoints` (Required)

Array of paths to entrypoint files. Bun generates one bundle per entrypoint.

### `files`

A map of file paths to their contents for in-memory bundling. Supports virtual files that don't exist on disk, overriding files on disk, or mixing disk and virtual files. File contents can be `string`, `Blob`, `TypedArray`, or `ArrayBuffer`.

```ts
const result = await Bun.build({
  entrypoints: ["/app/index.ts"],
  files: {
    "/app/index.ts": `import { greet } from "./greet.ts"; console.log(greet("World"));`,
    "/app/greet.ts": `export function greet(name: string) { return "Hello, " + name + "!"; }`,
  },
});
```

### `outdir`

Output directory. If not provided, bundled code is not written to disk; it's returned as `BuildArtifact` objects.

### `outfile`

Write output to a specific file (single entrypoint only).

### `target`

See Targets table. Default: `"browser"`.

### `format`

See Formats table. Default: `"esm"`.

### `splitting`

Boolean. Default `false`. Enables code splitting — shared modules are split into separate chunk files.

### `plugins`

Array of `BunPlugin` objects. Shared with the runtime plugin system.

### `env`

Controls how environment variables are handled during bundling:

| Value | Description |
|-------|-------------|
| `"inline"` | Injects all `process.env.FOO` references as string literals |
| `"disable"` | Disables env var injection (default) |
| `"FOO_*"` (prefix) | Only inlines env vars matching the prefix |

### `sourcemap`

| Value | Description |
|-------|-------------|
| `"none"` | No sourcemap (default) |
| `"linked"` | Separate `*.js.map` file with `//# sourceMappingURL` comment |
| `"external"` | Separate `*.js.map` file without the comment (includes `//# debugId`) |
| `"inline"` | Sourcemap appended as base64 payload in the bundle |

### `minify`

- `true`: enables all minification
- Object for granular control: `{ whitespace: true, syntax: true, identifiers: true }`

### `external`

Array of import paths to exclude from the bundle. The import statement is left as-is for runtime resolution. Wildcard `*` marks all imports as external.

### `packages`

| Value | Description |
|-------|-------------|
| `"bundle"` | Package dependencies are included in the bundle (default) |
| `"external"` | Package dependencies are excluded |

### `naming`

Customizes generated file names. Default: `"[dir]/[name].[ext]"`.

Tokens: `[name]`, `[ext]`, `[hash]`, `[dir]`.

Can be a string or an object:

```ts
naming: {
  entry: '[dir]/[name].[ext]',
  chunk: '[name]-[hash].[ext]',
  asset: '[name]-[hash].[ext]',
}
```

### `root`

Root directory of the project. If unspecified, computed as the first common ancestor of all entrypoint files.

### `publicPath`

A prefix added to import paths in bundled code (e.g., `"https://cdn.example.com/"`).

### `define`

A map of global identifiers to replace at build time. Keys are identifier names, values are JSON strings.

### `banner` / `footer`

Strings added to the top/bottom of the final bundle (e.g., `'"use client";'`).

### `drop`

Array of function calls to remove (e.g., `["console", "debugger"]`).

### `features`

Enables compile-time feature flags for dead code elimination using `import { feature } from "bun:bundle"`. The `feature("FLAG")` function is replaced with `true`/`false` at bundle time.

```ts
await Bun.build({
  features: ["PREMIUM"],
});
```

```ts
import { feature } from "bun:bundle";
if (feature("PREMIUM")) { initPremiumFeatures(); }
```

### `optimizeImports`

Array of package names to optimize barrel file imports. Skips parsing unused submodules of re-export index files. Works automatically for packages with `"sideEffects": false`.

### `metafile`

Generates build metadata:

- `true`: returns in result
- string path: writes JSON
- `{ json: "path", markdown: "path" }`: writes both formats
- `--metafile-md` for LLM-friendly markdown output

### `bytecode`

Boolean. Generates bytecode for JS/TS entrypoints to improve startup times. Requires `target: "bun"`. CommonJS generates `.jsc` files alongside entrypoints. ESM requires `compile: true`.

### `conditions`

Array of strings for package.json `exports` conditions used when resolving imports.

### `throw`

Boolean. Default `true`. When `true`, rejects with `AggregateError` on failure. When `false`, returns `{ success: false }`.

### `tsconfig`

Custom tsconfig.json file path for path resolution.

### `ignoreDCEAnnotations` / `emitDCEAnnotations`

Booleans. Control dead code elimination annotations (`@__PURE__`).

## Build Output

```ts
interface BuildOutput {
  outputs: BuildArtifact[];
  success: boolean;
  logs: Array<BuildMessage | ResolveMessage>;
  metafile?: BuildMetafile;
}

interface BuildArtifact extends Blob {
  kind: "entry-point" | "chunk" | "asset" | "sourcemap" | "bytecode";
  path: string;
  loader: Loader;
  hash: string | null;
  sourcemap: BuildArtifact | null;
}
```

`BuildArtifact` objects extend `Blob` and can be used with `await output.text()`, `await output.bytes()`, `await output.arrayBuffer()`, `new Response(output)`, or `Bun.write()`.

## CLI Flags

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--target` | string | `"browser"` | `browser`, `bun`, or `node` |
| `--format` | string | `"esm"` | `esm`, `cjs`, or `iife` |
| `--outdir` | string | `"dist"` | Output directory |
| `--outfile` | string | | Write output to a specific file |
| `--production` | boolean | | Set `NODE_ENV=production` and enable minification |
| `--bytecode` | boolean | | Use bytecode cache when compiling |
| `--compile` | boolean | | Generate standalone Bun executable |
| `--minify` | boolean | | Enable all minification |
| `--minify-syntax` | boolean | | Minify syntax and inline constants |
| `--minify-whitespace` | boolean | | Minify whitespace |
| `--minify-identifiers` | boolean | | Minify variable/function identifiers |
| `--keep-names` | boolean | | Preserve original function/class names when minifying |
| `--splitting` | boolean | | Enable code splitting |
| `--no-bundle` | boolean | | Transpile only, do not bundle |
| `--sourcemap` | string | `"none"` | `linked`, `inline`, `external`, or `none` |
| `--env` | string | `"disable"` | Inline env vars (`"inline"`, prefix like `FOO_PUBLIC_*`) |
| `--external` / `-e` | string | | Exclude modules (supports wildcards) |
| `--packages` | string | `"bundle"` | `external` or `bundle` |
| `--root` | string | | Root directory |
| `--public-path` | string | | Prefix for import paths |
| `--banner` | string | | Add banner to output |
| `--footer` | string | | Add footer to output |
| `--define` | string | | Replace global identifiers at build time |
| `--drop` | string | | Drop function calls (e.g., `console`, `debugger`) |
| `--loader` / `-l` | string | | Custom loader (`.ext:loader`) |
| `--entry-naming` | string | `"[dir]/[name].[ext]"` | Entry point filenames |
| `--chunk-naming` | string | `"[name]-[hash].[ext]"` | Chunk filenames |
| `--asset-naming` | string | `"[name]-[hash].[ext]"` | Asset filenames |
| `--conditions` | string | | Custom resolution conditions |
| `--css-chunking` | boolean | | Chunk CSS files together |
| `--watch` | boolean | | Rebuild on file changes |
| `--no-clear-screen` | boolean | | Don't clear terminal on rebuild |
| `--react-fast-refresh` | boolean | | Enable React Fast Refresh |
| `--react-compiler` | boolean | | Run React Compiler (experimental) |
| `--feature` | string | | Enable compile-time feature flags |
| `--metafile` | string | | Write build metadata JSON |
| `--metafile-md` | string | | Write build metadata as markdown |
| `--optimize-imports` | boolean | | Optimize barrel file imports |
| `--emit-dce-annotations` | boolean | | Re-emit DCE annotations |
| `--compile-exec-argv` | string | | Prepend arguments to standalone executable's execArgv |
| `--windows-hide-console` | boolean | | Prevent console window on Windows |
| `--windows-icon` | string | | Set icon for Windows executable |
| `--windows-title` | string | | Set Windows executable product name |
| `--windows-publisher` | string | | Set Windows executable company name |
| `--windows-version` | string | | Set Windows executable version |
| `--windows-description` | string | | Set Windows executable description |
| `--windows-copyright` | string | | Set Windows executable copyright |
| `--app` | boolean | | Build a web app using Bun Bake (experimental) |
| `--server-components` | boolean | | Enable React Server Components (experimental) |

## Standalone Executables

Generate a standalone Bun executable containing the bundle:

```bash
bun build ./cli.tsx --outfile mycli --compile
```

The resulting binary is a self-contained executable with Bun's runtime embedded. Customize Windows executable metadata with `--windows-*` flags.

## Examples

### Code Splitting

```bash
bun build ./entry-a.ts ./entry-b.ts --outdir ./out --splitting
```

### Environment Variables Inline

```bash
bun build ./index.tsx --outdir ./out --env inline
```

### External Modules

```bash
bun build ./index.tsx --outdir ./out --external lodash --external react
```

### Define Constants

```bash
bun build ./index.tsx --outdir ./out --define STRING='"value"' --define 'nested.boolean=true'
```

### Granular Minification

```bash
bun build ./index.tsx --outdir ./out --minify-whitespace --minify-identifiers --minify-syntax
```

### Drop Console/Debugger

```bash
bun build ./index.tsx --outdir ./out --drop console --drop debugger
```

### React Fast Refresh

```bash
bun build ./index.tsx --outdir ./out --react-fast-refresh
```

## References

- <https://bun.sh/docs/bundler>
