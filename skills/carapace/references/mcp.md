# Carapace MCP Server Guide

The carapace MCP server exposes shell completion capabilities to AI assistants and other MCP clients via the [Model Context Protocol](https://modelcontextprotocol.io/).

## Starting the Server

The MCP server runs over stdio (JSON-RPC 2.0) and is started with:

```bash
carapace --mcp
```

The server reads JSON-RPC requests from stdin and writes responses to stdout. It supports both single requests and batch requests (JSON arrays). Notifications (requests without an `id`) are silently skipped.

## Client Configuration

### Crush

Add to `crush.json`:

```json
{
  "mcp": {
    "carapace": {
      "type": "stdio",
      "command": "carapace",
      "args": ["--mcp"]
    }
  }
}
```

### Other MCP Clients

Any MCP client that supports stdio transport can connect. The server name is `carapace` and it reports protocol version `2024-11-05`.

## Protocol

The server implements JSON-RPC 2.0 with these methods:

|| Method | Description |
|--------|-------------|
| `initialize` | Returns server info and capabilities |
| `tools/list` | Returns the four tool definitions |
| `tools/call` | Executes a tool by name |

Capabilities: `{tools: {}}` — only tools are supported (no resources, no prompts).

## Tools

### `complete_command` — Shell Command Completion

Returns context‑aware, dynamic completions for shell commands.

**Input Schema:**

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `args` | `string[]` | Yes | Command line arguments to complete |
| `executable` | `string` | No | Path to the executable providing the completion (requires `bridge`) |
| `bridge` | `string` | No | Bridge providing the completion (e.g. `carapace-bin`, `cobra`, `zsh`, `fish`, `bash`, `argcomplete`, `click`) |

**How it works:**

The first element in `args` is the command name (not a flag). The remaining elements are the arguments to complete, including the partial word being completed.

The behavior depends on which optional parameters are provided:

| `executable` | `bridge` | Behavior | Confirmation |
|-------------|----------|----------|--------------|
| not set | not set | Default completion using `carapace` from `PATH` | Not required |
| not set | set | Explicit bridge using `carapace <command>/<bridge> export <args...>` | Not required |
| set | set | Custom executable with specified bridge | Required |

**Mode details:**

1. **Default** (no `executable`, no `bridge`): Invokes `carapace <command> export <args...>` using the carapace binary from `PATH`. This is the standard completion flow.

2. **Bridge only** (no `executable`, `bridge` set): Uses the explicit bridge variant. Invokes `carapace <command>/<bridge> export <args...>`. For example, with `bridge: "zsh"`, it completes `tail` using zsh's completion system. The `carapace-bin/<bridge>` syntax is also supported to use an explicit bridge within carapace-bin (e.g. `bridge: "carapace-bin/cobra"`).

3. **Executable + bridge** (`executable` and `bridge` both set): Uses the given executable for completion. **This requires user confirmation** since it executes an arbitrary binary.
   - For `carapace-bin` bridge: Invokes `<executable> <command> export <args...>` directly (default choice).
   - For `carapace-bin/<bridge>`: Invokes `<executable> <command>/<bridge> export <args...>` (explicit bridge with custom executable).
   - For other bridges (cobra, argcomplete, etc.): Invokes `carapace <command>/<bridge> export <args...>` with the executable's directory prepended to `PATH`, so the bridge action resolves the custom executable.

**Constraints:**

- `args` must contain at least 2 elements (command + argument to complete)
- Command must not start with `-` (flags are not completer names)
- Arguments must not contain NUL bytes (`\0`)
- `bridge` is required when `executable` is set

**Response:** Plain text containing the completion output (one completion candidate per line). On error, the response has `isError: true` with the error message as text.

**Examples:**

```json
// Complete git branches after "git checkout " (default mode)
{"args": ["git", "checkout", ""]}

// Complete docker subcommands (default mode)
{"args": ["docker", ""]}

// Complete tail using zsh bridge (bridge mode)
{"args": ["tail", ""], "bridge": "zsh"}

// Complete git using cobra bridge (bridge mode)
{"args": ["git", ""], "bridge": "cobra"}

// Complete using a custom carapace-bin executable (executable mode - needs confirmation)
{"args": ["git", "checkout", ""], "executable": "/path/to/carapace", "bridge": "carapace-bin"}

// Complete using a custom executable with cobra bridge (executable mode - needs confirmation)
{"args": ["myapp", ""], "executable": "/path/to/myapp", "bridge": "cobra"}

// Complete using a custom carapace-bin executable with explicit bridge variant
{"args": ["git", ""], "executable": "/path/to/carapace", "bridge": "carapace-bin/cobra"}
```

**CLI equivalents:**

| Mode | CLI equivalent |
|------|---------------|
| Default | `carapace <cmd> export <args...>` |
| Bridge only | `carapace <cmd>/<bridge> export <args...>` |
| Executable + carapace-bin | `<executable> <cmd> export <args...>` |
| Executable + carapace-bin/bridge | `<executable> <cmd>/<bridge> export <args...>` |
| Executable + other bridge | `PATH=<dir>:$PATH carapace <cmd>/<bridge> export <args...>` |

### `complete_macro` — Macro Completion

Returns context‑aware, dynamic completions for a carapace macro.

**Input Schema:**

|| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `macro` | `string` | Yes | Macro name (e.g. `tools.git.Refs`) |
| `args` | `string[]` | Yes | Arguments to complete |
| `executable` | `string` | No | Path to the carapace executable providing the macro |

**How it works:**

1. **Default** (no `executable`): Invokes the current carapace binary with `_carapace macro <macro> <args...>`.

2. **Custom executable** (`executable` set): Resolves the executable path and invokes it with the same arguments. **This requires user confirmation** since it executes an arbitrary binary.

**Constraints:**

- `macro` is required
- `args` is required
- Arguments must not contain NUL bytes (`\0`)

**Response:** Plain text containing the completion output (JSON with `values` array). On error, the response has `isError: true` with the error message as text.

**Examples:**

```json
// Complete git refs (default mode)
{"macro": "tools.git.Refs", "args": [""]}

// Complete git refs with opts (default mode)
{"macro": "tools.git.Refs", "args": ["{localbranches: true, remotebranches: true, tags: true}", ""]}

// Complete using a custom carapace executable (needs confirmation)
{"macro": "tools.git.Refs", "args": [""], "executable": "/path/to/carapace"}
```

**CLI equivalent:** `carapace _carapace macro <macro> <args...>`

### `list_macros` — Available Macros

Lists all available carapace macros with their names, signatures, descriptions, and Go source references.

**Input Schema:** No parameters (empty object).

**Response:** JSON array of macro objects as text:

```json
[
  {
    "name": "carapace.tools.git.Refs",
    "signature": "{localbranches: false, remotebranches: false, tags: false}",
    "description": "complete refs",
    "reference": "github.com/carapace-sh/carapace-bin/pkg/actions/tools/git#ActionRefs"
  },
  {
    "name": "carapace.tools.git.Tags",
    "signature": "—",
    "description": "complete tags",
    "reference": "github.com/carapace-sh/carapace-bin/pkg/actions/tools/git#ActionTags"
  }
]
```

**Fields:**

| Field | Description |
|-------|-------------|
| `name` | Fully qualified macro name with `carapace.` prefix (e.g. `carapace.tools.git.Refs`) |
| `signature` | Argument signature — see carapace-macro skill for format details. `—` (em dash) means no arguments (MacroN) |
| `description` | Short description of what the macro completes |
| `reference` | Go import path with `#FunctionName` for looking up source code |

**Use cases:**

- Discover which macros exist before using them in YAML specs or Go code
- Look up argument signatures to format macro calls correctly
- Find the Go source reference to understand parameters in detail

**CLI equivalent:** `carapace --macro`

### `codegen` — Generate Go Code from YAML Spec

Generates Go completer source code from a carapace YAML spec file. This is useful for converting a user spec into a native Go completer that can be built into carapace-bin.

**Input Schema:**

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `path` | `string` | Yes | Path to the YAML spec file (relative or absolute) |

**How it works:**

1. Resolves `path` to an absolute path (relative paths are resolved from the server's working directory)
2. Invokes `carapace --codegen <path>` internally
3. Scans the output for `.go` filenames that were generated
4. Returns a sorted JSON array of generated file paths

**Response:** JSON array of generated `.go` file paths as text:

```json
["cmd/root.go", "cmd/cache.go"]
```

On error (missing path, invalid spec, codegen failure), the response has `isError: true` with the error message as text.

**CLI equivalent:** `carapace --codegen <path>`

## MCP Tool to CLI Flag Mapping

| MCP Tool | CLI Command | Notes |
|----------|-------------|-------|
| `complete_command` | `carapace <cmd> export <args...>` | Default mode; bridge/executable modes use different invocations |
| `list_macros` | `carapace --macro` | Iterates `actions.Macros` map |
| `complete_macro` | `carapace _carapace macro <macro> <args...>` | Invokes macro completion on current or custom executable |
| `codegen` | `carapace --codegen <path>` | Shells out to the same executable |

## Relationship to Other Skills

| Skill | When to use instead of this skill |
|-------|----------------------------------|
| `carapace-macro` | Formatting macro arguments in specs/YAML, understanding MacroN/MacroI/MacroV types |
| `carapace-spec` | Writing YAML spec files (input for codegen) |
| `carapace-action` | Creating/modifying shared Go actions that become macros |
| `carapace-scrape` | Generating specs from CLI source code |
| `carapace-integrate` | Integrating carapace library into cobra-based CLIs |

## Error Handling

The MCP server returns errors in two ways:

1. **JSON-RPC errors** — For protocol-level failures (unknown method, invalid tool call parameters). Uses standard JSON-RPC error codes:
   - `-32601` — Method not found
   - `-32602` — Invalid params

2. **Tool errors** — For tool execution failures (command not found, invalid arguments, codegen failures). Returned as `isError: true` with the error message in the text content.
