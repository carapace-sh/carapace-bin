---
name: carapace-mcp
description: >
  Use when user needs to understand, configure, or use the carapace MCP server — including
  tool descriptions, input schemas, response formats, CLI equivalents, and MCP client setup.
  Covers the three tools (complete, list_macros, codegen), the JSON-RPC protocol, stdio transport,
  and how MCP tools map to carapace CLI flags. Triggers on: "carapace mcp", "MCP server",
  "MCP tools", "carapace MCP", "complete tool", "list_macros tool", "codegen tool",
  "configure MCP", "MCP setup".
user-invocable: true
---

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

| Method | Description |
|--------|-------------|
| `initialize` | Returns server info and capabilities |
| `tools/list` | Returns the three tool definitions |
| `tools/call` | Executes a tool by name |

Capabilities: `{tools: {}}` — only tools are supported (no resources, no prompts).

## Tools

### `complete` — Shell Command Completion

Returns context‑aware, dynamic completions for shell commands.

**Input Schema:**

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `args` | `string[]` | Yes | Command line arguments to complete |

**How it works:**

The first element in `args` is the command name (not a flag). The remaining elements are the arguments to complete, including the partial word being completed. The server invokes `carapace <command> export <args...>` internally and returns the raw completion output.

**Constraints:**

- `args` must contain at least 2 elements (command + argument to complete)
- Command must not start with `-` (flags are not completer names)
- Arguments must not contain NUL bytes (`\0`)

**Response:** Plain text containing the completion output (one completion candidate per line). On error, the response has `isError: true` with the error message as text.

**Examples:**

```json
// Complete git branches after "git checkout "
{"args": ["git", "checkout", ""]}

// Complete docker subcommands
{"args": ["docker", ""]}

// Complete kubectl flags for get pods
{"args": ["kubectl", "get", "pods", "-"]}
```

**CLI equivalent:** `carapace <command> export <args...>`

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
| `complete` | `carapace <cmd> export <args...>` | Shells out to the same executable |
| `list_macros` | `carapace --macro` | Iterates `actions.Macros` map |
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
