# Bun Runtime API

Reference for the `Bun` global object, built-in globals, module resolution, and the runtime's key subsystems. Bun uses the JavaScriptCore (JSC) engine (not V8) and is written in Zig and C++.

> **Source of truth**: <https://bun.sh/docs/runtime>. For the Bun shell, see [shell.md](shell.md). For the bundler, see [bundler.md](bundler.md).

## Architecture

- **Engine**: JavaScriptCore (JSC, Apple's JS engine from Safari), not V8
- **Implementation**: Zig and C++, with Rust components in the shell
- **Startup**: ~4x faster than Node.js on Linux (5.2ms vs 25.1ms for Hello World)
- **Transpiler**: Built-in, handles TS/JSX natively

## Globals

Bun implements globals from three sources:

### Web API Globals

`AbortController`, `AbortSignal`, `Blob`, `ByteLengthQueuingStrategy`, `CountQueuingStrategy`, `Crypto`, `crypto`, `CryptoKey`, `CustomEvent`, `Event`, `ErrorEvent`, `CloseEvent`, `MessageEvent`, `EventTarget`, `fetch`, `FormData`, `globalThis`, `Headers`, `JSON`, `performance`, `queueMicrotask()`, `ReadableStream`, `ReadableByteStreamController`, `ReadableStreamDefaultController`, `ReadableStreamDefaultReader`, `reportError`, `Response`, `Request`, `setImmediate()`, `setInterval()`, `setTimeout()`, `clearImmediate()`, `clearInterval()`, `clearTimeout()`, `ShadowRealm`, `SubtleCrypto`, `DOMException`, `TextDecoder`, `TextEncoder`, `TransformStream`, `TransformStreamDefaultController`, `URL`, `URLSearchParams`, `WebAssembly`, `WritableStream`, `WritableStreamDefaultController`, `WritableStreamDefaultWriter`, `atob()`, `btoa()`, `alert`, `confirm`, `prompt`

### Node.js Globals

`Buffer`, `__dirname`, `__filename`, `exports`, `global`, `module`, `process`, `require()`

### Bun-Specific Globals

`Bun` (the primary global object), `BuildMessage`, `ResolveMessage`, `HTMLRewriter` (from Cloudflare)

## `Bun` Global Object

### Properties

| Property | Type | Description |
|----------|------|-------------|
| `Bun.version` | string | Current Bun version string |
| `Bun.revision` | string | Git commit hash of current build |
| `Bun.env` | object | Alias for `process.env` |
| `Bun.main` | string | Absolute path to entrypoint |
| `Bun.cwd` | string | Current working directory |

### HTTP Server (`Bun.serve`)

```ts
const server = Bun.serve({
  fetch(req) { return new Response("Hello"); },
  port: 3000,
});
```

| Option | Description |
|--------|-------------|
| `fetch` | Request handler function |
| `routes` | Static and dynamic routes (`"/users/:id"`, wildcards, per-method) |
| `port` | Port number (default: 3000) |
| `hostname` | Hostname to bind |
| `tls` | TLS/SSL options |
| `websocket` | WebSocket handlers |
| `error` | Error handler |
| `maxRequestBodySize` | Max request body size |
| `reusePort` | SO_REUSEPORT for load balancing |

#### Routes API (v1.2.3+)

Supports static routes (`"/api/status"`), dynamic routes (`"/users/:id"`), per-method handlers (`{ GET, POST, ... }`), wildcards (`"/api/*"`), redirects, and serving `Bun.file()` directly.

#### `export default` Syntax

Instead of calling `Bun.serve()`, export a default object:

```ts
export default {
  fetch(req) { return new Response("Hello"); },
  port: 3000,
};
```

#### Server Methods

| Method | Description |
|--------|-------------|
| `server.stop(force?)` | Stop accepting connections |
| `server.reload(options)` | Update handlers without restarting |
| `server.url` | Server URL |
| `server.port` / `server.hostname` | Server address info |
| `server.pendingRequests` / `server.pendingWebSockets` | Active connection counters |
| `server.requestIP(request)` | Get client IP address |
| `server.timeout(request, seconds)` | Override idle timeout per-request |
| `server.ref()` / `server.unref()` | Keep/allow process to exit |
| `server.upgrade(request, options?)` | Upgrade HTTP to WebSocket |
| `server.publish(topic, data)` | Publish to WebSocket subscribers |
| `server.subscriberCount(topic)` | Count WebSocket subscribers |

### File I/O

#### `Bun.file(path)` — Create a BunFile reference

```ts
const file = Bun.file("./data.txt");
const text = await file.text();
const json = await file.json();
const stream = file.stream();
const buffer = await file.arrayBuffer();
const bytes = await file.bytes();
```

| Method | Description |
|--------|-------------|
| `file.text()` | Read as string |
| `file.json()` | Read as JSON |
| `file.stream()` | Read as `ReadableStream` |
| `file.arrayBuffer()` | Read as `ArrayBuffer` |
| `file.bytes()` | Read as `Uint8Array` |
| `file.writer(options?)` | Get `FileSink` for incremental writing |
| `file.exists()` | Check if file exists |
| `file.delete()` | Delete the file |
| `file.size` | File size in bytes |

`Bun.file()` also accepts file descriptors (`Bun.file(fd)`) and `file://` URLs.

#### `Bun.write(destination, data)`

Multi-tool write function. Accepts strings, Blobs, ArrayBuffers, Responses, TypedArrays.

```ts
await Bun.write("./output.txt", "Hello World");
await Bun.write(Bun.stdout, "Hello stdout");
await Bun.write(Bun.file("./data.bin"), new Uint8Array([1, 2, 3]));
```

#### Standard Streams

`Bun.stdin`, `Bun.stdout`, `Bun.stderr` are `BunFile` instances.

#### FileSink

| Method | Description |
|--------|-------------|
| `.write(chunk)` | Write a chunk |
| `.flush()` | Flush buffered data |
| `.end()` | End writing |
| `.ref()` / `.unref()` | Process lifecycle control |
| `highWaterMark` | Configurable buffer threshold |

### Child Processes

#### `Bun.spawn` / `Bun.spawnSync`

```ts
const proc = Bun.spawn(["echo", "hello"], {
  cwd: "/tmp",
  env: { ...process.env, FOO: "bar" },
  stdout: "pipe",
  stderr: "pipe",
  stdin: "pipe",
});

const output = await new Response(proc.stdout).text();
const exitCode = await proc.exited;
```

| Property/Method | Description |
|-----------------|-------------|
| `proc.stdout` / `proc.stderr` | Output as `ReadableStream` (or `Buffer` for sync) |
| `proc.stdin` | `FileSink` for writing input |
| `proc.pid` | Process ID |
| `proc.exited` | Promise resolving on exit |
| `proc.exitCode` / `proc.signalCode` / `proc.killed` | Exit status |
| `proc.kill(signal?)` | Kill process |
| `proc.ref()` / `proc.unref()` | Detach from parent lifecycle |
| `proc.send(message)` | IPC — send message to child |
| `proc.disconnect()` | Disconnect IPC channel |
| `proc.resourceUsage()` | Get CPU time, max RSS, etc. |

**Options**: `cwd`, `env`, `stdin`/`stdout`/`stderr` (pipe, inherit, ignore, BunFile, fd), `onExit` callback, `ipc` handler, `serialization` ("json" or "advanced"), `signal` (AbortSignal), `timeout`, `killSignal`, `maxBuffer`, `terminal` (PTY support).

#### PTY/Terminal

```ts
const proc = Bun.spawn({
  cmd: ["bash"],
  terminal: { cols: 80, rows: 24, data, exit, drain },
});
proc.terminal.write("ls\n");
proc.terminal.resize(120, 40);
proc.terminal.setRawMode(true);
proc.terminal.close();
```

### Hashing and Crypto

#### Password Hashing

```ts
const hash = await Bun.password.hash("password", "argon2id"); // default
const hash = await Bun.password.hash("password", "bcrypt");
const isValid = await Bun.password.verify("password", hash);
```

Sync variants: `Bun.password.hashSync()`, `Bun.password.verifySync()`.

#### Non-Crypto Hashing

```ts
Bun.hash(data, seed?);        // 64-bit Wyhash (default)
Bun.hash.wyhash(data);
Bun.hash.crc32(data);
Bun.hash.adler32(data);
Bun.hash.cityHash32(data);
Bun.hash.cityHash64(data);
Bun.hash.xxHash32(data);
Bun.hash.xxHash64(data);
Bun.hash.xxHash3(data);
Bun.hash.murmur32v3(data);
Bun.hash.murmur32v2(data);
Bun.hash.murmur64v2(data);
Bun.hash.rapidhash(data);
```

#### CryptoHasher (Incremental)

```ts
const hasher = new Bun.CryptoHasher("sha256");
hasher.update("hello ");
hasher.update("world");
const digest = hasher.digest("hex"); // or "base64", or Uint8Array
```

Algorithms: blake2b256/512, blake2s256, md4, md5, ripemd160, sha1/224/256/384/512/512-224/512-256, sha3-224/256/384/512, shake128/256. HMAC supported with key parameter.

### Transpiler

```ts
const transpiler = new Bun.Transpiler({
  loader: "tsx",
  define: { DEBUG: "false" },
  target: "browser",
  tsconfig: "./tsconfig.json",
});

const code = transpiler.transformSync("const x: number = 1");
const asyncCode = await transpiler.transform("const x: number = 1");
const { exports, imports } = transpiler.scan("export const x = 1");
const justImports = transpiler.scanImports("import { x } from './y'");
```

**TranspilerOptions**: `define`, `loader` ("js"|"jsx"|"ts"|"tsx"), `target` ("browser"|"bun"|"node"), `tsconfig`, `macro`, `exports` (eliminate/replace), `trimUnusedImports`, `minifyWhitespace`, `inline`.

### Networking

| API | Description |
|-----|-------------|
| `Bun.listen(options)` | Start TCP server |
| `Bun.connect(options)` | Connect via TCP |
| `Bun.udp_socket(options)` | UDP socket |
| `Bun.dns.lookup(hostname, options)` | DNS resolution |
| `Bun.dns.prefetch(hostname)` | Prefetch DNS |
| `Bun.dns.getCacheStats()` | DNS cache stats |
| `new WebSocket(url)` | WebSocket client |

### Database and Storage

| API | Description |
|-----|-------------|
| `new Database("file.db")` from `bun:sqlite` | SQLite3 driver |
| `Bun.SQL` / `Bun.sql` | SQL client (PostgreSQL, MySQL, SQLite) |
| `Bun.RedisClient` / `Bun.redis` | Redis (Valkey) client |
| `Bun.S3` | S3-compatible object storage |

### Compression

| API | Description |
|-----|-------------|
| `Bun.gzipSync(data, options?)` | GZIP compress |
| `Bun.gunzipSync(data)` | GZIP decompress |
| `Bun.deflateSync(data, options?)` | DEFLATE compress |
| `Bun.inflateSync(data)` | DEFLATE decompress |
| `Bun.zstdCompressSync(data, options?)` | Zstandard compress (sync) |
| `Bun.zstdCompress(data)` | Zstandard compress (async) |
| `Bun.zstdDecompressSync(data)` | Zstandard decompress (sync) |
| `Bun.zstdDecompress(data)` | Zstandard decompress (async) |

### Stream Utilities

| API | Description |
|-----|-------------|
| `Bun.readableStreamToText(stream)` | Stream to string |
| `Bun.readableStreamToJSON(stream)` | Stream to JSON object |
| `Bun.readableStreamToArrayBuffer(stream)` | Stream to ArrayBuffer |
| `Bun.readableStreamToBytes(stream)` | Stream to Uint8Array |
| `Bun.readableStreamToBlob(stream)` | Stream to Blob |
| `Bun.readableStreamToArray(stream)` | Stream to array of chunks |
| `Bun.readableStreamToFormData(stream, boundary?)` | Stream to FormData |

### Utility Functions

| API | Description |
|-----|-------------|
| `Bun.sleep(ms)` | Async sleep (returns Promise) |
| `Bun.sleepSync(ms)` | Blocking sleep |
| `Bun.nanoseconds()` | Nanoseconds since process start |
| `Bun.which(bin, options?)` | Find executable path |
| `Bun.randomUUIDv7(encoding?, timestamp?)` | Generate UUID v7 (monotonic, sortable) |
| `Bun.peek(promise)` | Read promise result without await (if settled) |
| `Bun.peek.status(promise)` | Check promise status: "fulfilled", "pending", "rejected" |
| `Bun.deepEquals(a, b, strict?)` | Deep equality check |
| `Bun.inspect(value)` | Serialize object as console.log would |
| `Bun.inspect.custom` | Symbol for custom inspection |
| `Bun.inspect.table(data, properties, options)` | Format tabular data |
| `Bun.escapeHTML(value)` | Escape HTML entities |
| `Bun.stringWidth(text, options?)` | Column width (handles ANSI, emoji, wide chars) |
| `Bun.stripANSI(text)` | Strip ANSI escape codes |
| `Bun.wrapAnsi(text, columns, options?)` | Wrap text to column width |
| `Bun.fileURLToPath(url)` | Convert `file://` URL to absolute path |
| `Bun.pathToFileURL(path)` | Convert absolute path to `file://` URL |
| `Bun.openInEditor(path, options?)` | Open file in default editor |
| `Bun.resolveSync(specifier, root)` | Resolve module path |

### Other Built-in Modules

| Module/API | Description |
|-------------|-------------|
| `Bun.build()` | Bundler (see [bundler.md](bundler.md)) |
| `Bun.FileSystemRouter` | File-system based routing |
| `HTMLRewriter` | CSS-selector-based HTML transformation (from Cloudflare) |
| `Bun.WebView` | Headless browser |
| `Bun.CSRF.generate` / `Bun.CSRF.verify` | CSRF token generation/verification |
| `Bun.Cookie` / `Bun.CookieMap` | HTTP cookie APIs |
| `Bun.Glob` | File globbing |
| `Bun.plugin` | Module loaders/plugins |
| `Bun.semver` | Semantic versioning API |
| `Bun.TOML.parse` | TOML parsing |
| `Bun.markdown` | Markdown parsing/rendering |
| `Bun.color` | Color formatting (CSS, ANSI, hex) |
| `Bun.Image` | Image decoding/transformation/encoding |
| `Bun.mmap` | Memory-mapped files |
| `Bun.gc` | Trigger garbage collection |
| `Bun.generateHeapSnapshot` | Generate heap snapshot |
| `Bun.ArrayBufferSink` | Incremental buffer construction |
| `Bun.allocUnsafe` | Allocate uninitialized buffer |
| `Bun.concatArrayBuffers` | Concatenate array buffers |
| `bun:ffi` | Foreign Function Interface |
| `bun:test` | Jest-compatible test runner (see [test-runner.md](test-runner.md)) |
| `bun:jsc` | Low-level JSC internals (`serialize`, `deserialize`, `estimateShallowMemoryUsageOf`) |
| `new Worker()` | Web Workers (multi-threading) |
| `Node-API` | Native add-on support |

## `import.meta` Properties

| Property | Description |
|----------|-------------|
| `import.meta.dir` | Directory of current file |
| `import.meta.dirname` | Alias for `import.meta.dir` |
| `import.meta.file` | Current filename |
| `import.meta.path` | Absolute path to current file |
| `import.meta.filename` | Alias for `import.meta.path` |
| `import.meta.url` | `file://` URL of current file |
| `import.meta.main` | `true` if entrypoint |
| `import.meta.env` | Alias for `process.env` |
| `import.meta.resolve(specifier)` | Resolve module specifier to `file://` URL |

## Module Resolution

Bun supports both **ES modules** (`import`/`export`) and **CommonJS** (`require()`/`module.exports`).

### Extension Resolution Order

For extensionless imports, Bun tries in order:

`.tsx` → `.jsx` → `.mts` → `.ts` → `.mjs` → `.js` → `.cts` → `.cjs` → `.json` → `index.*` variants

### TypeScript File Extension Substitution

Importing `"./hello.js"` resolves to `"./hello.ts"` — Bun substitutes `.js` with the corresponding TypeScript extension.

### Package Resolution

Uses the Node.js algorithm with `node_modules` scanning. Supports:

- `exports` field with conditional exports
- Condition priority: `"bun"` → `"node-addons"` → `"node"` → `"require"` → `"import"` → `"default"`
- `NODE_PATH`
- `tsconfig.json` `paths`
- `package.json` `imports` subpath maps

## Watch Mode vs Hot Mode

| Feature | `--watch` | `--hot` |
|---------|-----------|---------|
| Behavior | Hard restarts process | Soft reloads code |
| Global state | Lost | Preserved (`globalThis` persists) |
| HTTP servers | Restarted | Not restarted, handler updated |
| Speed | Fast (native FS watchers: kqueue/inotify) | Faster (no process restart) |

Use `--watch` for development where a clean restart is desired. Use `--hot` for faster iteration where preserving in-memory state matters (e.g., development servers with active connections).

## Auto-Install

Bun can automatically install missing dependencies when running files. Configured via `--install` flag or `[install]` in `bunfig.toml`:

| Mode | Behavior |
|------|----------|
| `auto` (default) | Resolve from local `node_modules`; auto-install if not found |
| `force` | Always auto-install, even if `node_modules` exists |
| `disable` | Never auto-install |
| `fallback` | Check local first, then auto-install missing (CLI: `bun -i`) |

## References

- <https://bun.sh/docs/runtime>
- <https://bun.sh/docs/runtime/globals>
- <https://bun.sh/docs/runtime/bun-apis>
- <https://bun.sh/docs/runtime/http/server>
- <https://bun.sh/docs/runtime/file-io>
- <https://bun.sh/docs/runtime/child-process>
- <https://bun.sh/docs/runtime/hashing>
- <https://bun.sh/docs/runtime/utils>
- <https://bun.sh/docs/runtime/module-resolution>
- <https://bun.sh/docs/runtime/watch-mode>
