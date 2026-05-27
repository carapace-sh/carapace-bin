# MCP

The [Model Context Protocol][MCP] server provides access to [Carapace] features.

![](./mcp/mcp.cast)

Simply add it to your [config].

```json

{
  "$schema": "https://charm.land/crush.json",
  "mcp": {
    "carapace": {
      "type": "stdio",
      "command": "carapace",
      "args": [
        "--mcp"
      ]
    }
  }
}
```

## Tools

- **complete** Return context‑aware, dynamic completions for shell commands.
- **list_macros** List available macros and their signatures.

[Carapace]:https://carapace.sh
[config]:https://github.com/charmbracelet/crush#mcps
[MCP]:https://modelcontextprotocol.io/docs/getting-started/intro
