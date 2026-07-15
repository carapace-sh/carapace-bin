# Bun Shell

Reference for Bun's built-in shell: the `$` template literal tag, builtins, redirection, piping, environment, and security model. The Bun shell does not invoke a system shell — it's a re-implementation running in-process, written in Rust.

> **Source of truth**: <https://bun.sh/docs/runtime/shell>. For spawning child processes programmatically, see [runtime.md](runtime.md) (`Bun.spawn`).

## Overview

Bun includes a cross-platform shell implementation that runs in-process. It does not invoke `/bin/sh` or any external shell. All interpolated variables are treated as single literal strings, preventing command injection.

```ts
import { $ } from "bun";
await $`echo "Hello World!"`;
```

## The `$` Template Literal Tag

The `$` function is a tagged template literal that executes shell commands. It returns a `ShellPromise` that resolves to `{ stdout: Buffer, stderr: Buffer }` by default.

```ts
const result = await $`echo "Hello"`;
console.log(result.stdout.toString());  // "Hello\n"
```

### Methods on the Call Result

| Method | Description |
|--------|-------------|
| `.quiet()` | Suppress stdout/stderr output |
| `.text()` | Read output as string (also calls `.quiet()`) |
| `.json()` | Read output as parsed JSON |
| `.lines()` | Async iterable for reading output line-by-line |
| `.blob()` | Read output as `Blob` |
| `.nothrow()` | Disable throwing on non-zero exit codes |
| `.throws(boolean)` | Enable/disable throwing on non-zero exit codes |
| `.env(object)` | Set environment variables for a single command |
| `.cwd(path)` | Set working directory for a single command |

```ts
const text = await $`echo "Hello"`.text();
const data = await $`cat data.json`.json();

for await (const line of $`cat file.txt`.lines()) {
  console.log(line);
}

await $`false`.nothrow();  // don't throw on non-zero exit
```

### Static Methods on `$`

| Method | Description |
|--------|-------------|
| `$.nothrow()` | Global default: don't throw on non-zero exit |
| `$.throws(boolean)` | Global default for throwing behavior |
| `$.env(object)` | Set global default environment variables |
| `$.env()` | Reset global environment to default |
| `$.cwd(path)` | Set global default working directory |
| `$.braces(string)` | Brace expansion (e.g., `echo {1,2,3}`) |
| `$.escape(string)` | Escape a string for safe use in shell |

```ts
$.cwd("/tmp");
$.env({ FOO: "bar" });

await $`pwd`;           // runs in /tmp
await $`echo $FOO`;    // outputs "bar"

$.env();               // reset to default
$.cwd(process.cwd());  // reset to default
```

## Redirection Operators

| Operator | Description |
|----------|-------------|
| `<` | Redirect stdin |
| `>` or `1>` | Redirect stdout (overwrite) |
| `2>` | Redirect stderr |
| `&>` | Redirect both stdout and stderr |
| `>>` or `1>>` | Redirect stdout (append) |
| `2>>` | Redirect stderr (append) |
| `&>>` | Redirect both stdout and stderr (append) |
| `1>&2` | Redirect stdout to stderr |
| `2>&1` | Redirect stderr to stdout |

### Redirect to JavaScript Objects

Supported output targets for `>`: `Buffer`, `Uint8Array`, `Uint16Array`, `Uint32Array`, `Int8Array`, `Int16Array`, `Int32Array`, `Float32Array`, `Float64Array`, `ArrayBuffer`, `SharedArrayBuffer`, `Bun.file(path)`, `Bun.file(fd)`.

```ts
await $`echo "hello" > ${Bun.file("/tmp/out.txt")}`;
await $`echo "hello" > ${new Uint8Array(1024)}`;
```

### Redirect from JavaScript Objects

Supported input sources for `<`: All the same types as output, plus `Response` (reads from body).

```ts
const response = new Response("piped input");
await $`cat < ${response}`;

const file = Bun.file("./input.txt");
await $`grep "pattern" < ${file}`;
```

## Piping (`|`)

Standard bash-style pipe between commands is supported. You can also pipe JavaScript objects into commands.

```ts
await $`echo "hello" | wc -c`;

const response = new Response("piped data");
await $`cat < ${response} | wc -c`;
```

## Command Substitution (`$(...)`)

Use `$(...)` syntax to insert the output of one command into another.

```ts
const files = await $`ls *.ts`.text();
await $`echo ${files}`;
```

**Important**: Backtick syntax does **not** work for command substitution. Bun uses the `raw` property of template literals, so backticks are not interpreted as command substitution.

## Environment Variables

```ts
// Inline env vars (like bash)
await $`FOO=foo bun -e '...'`;

// String interpolation (escaped by default)
const val = "456";
await $`FOO=${val} some-command`;

// Per-command environment
await $`some-command`.env({ ...process.env, FOO: "bar" });

// Global default environment
$.env({ FOO: "bar" });
```

**Input is escaped by default** — interpolated values are treated as single literal strings, preventing shell injection.

## Working Directory

```ts
// Per-command
await $`pwd`.cwd("/tmp");

// Global default
$.cwd("/tmp");
```

## Builtin Commands

These are implemented natively (cross-platform, no external binary needed):

| Command | Notes |
|---------|-------|
| `cd` | Change working directory |
| `ls` | List files; supports `-l` for long format |
| `rm` | Remove files and directories |
| `echo` | Print text |
| `pwd` | Print working directory |
| `bun` | Run bun inside bun |
| `cat` | Concatenate files |
| `touch` | Create/update file timestamps |
| `mkdir` | Create directories |
| `which` | Locate a command |
| `mv` | Move files/directories (partially implemented — missing cross-device) |
| `exit` | Exit the shell |
| `true` | Returns exit code 0 |
| `false` | Returns exit code 1 |
| `yes` | Repeatedly output a string |
| `seq` | Print sequences of numbers |
| `dirname` | Strip last component from file path |
| `basename` | Strip directory and suffix from file path |

## `.sh` File Loader

Bun can execute `.sh` files directly:

```bash
bun ./script.sh
```

These scripts run cross-platform (Windows, Linux, macOS) using the Bun Shell interpreter instead of `/bin/sh`.

## Globs and Brace Expansion

- Globs (`**`, `*`, `{expansion}`) are supported natively.
- `$.braces(str)` implements brace expansion: `echo {1,2,3}` expands to `["echo 1", "echo 2", "echo 3"]`.

## Concurrency

Operations run **concurrently**, unlike bash/zsh which execute sequentially by default. Bun's shell operations are scheduled concurrently for better performance.

## Security

### Command Injection Prevention

Bun Shell **does not invoke a system shell**. All interpolated variables are treated as single literal strings, preventing **command injection**.

### Argument Injection Caveat

**Argument injection** is still possible — if the target program interprets its arguments as flags. For example:

```ts
const branch = "--upload-pack=malicious";
await $`git ls-remote origin ${branch}`;
```

User input must still be sanitized before passing to shell commands.

### Explicit Shell Spawning

If you explicitly spawn a new shell (e.g., `bash -c "echo ${userInput}"`), Bun's protections no longer apply — the external shell interprets the interpolated string.

## Escaping

By default, all interpolated strings are escaped. To skip escaping, wrap in `{ raw: 'str' }`:

```ts
const cmd = "{ raw: 'rm -rf /tmp/*' }";
await $`bash -c ${cmd}`;
```

Use `$.escape(str)` to manually escape a string for shell use:

```ts
const userInput = '; rm -rf /';
const escaped = $.escape(userInput);
await $`echo ${escaped}`;
```

## Implementation Notes

- Written in Rust with a handwritten lexer, parser, and interpreter
- Runs in-process, no external shell binary needed
- Operations run concurrently for performance
- Cross-platform: works on Windows, Linux, and macOS

## References

- <https://bun.sh/docs/runtime/shell>
