---
name: carapace
description: >
  Use when working with carapace shell completion — integrating into CLIs, writing YAML specs,
  creating custom actions, looking up macros, setting up shell completion, configuring
  choices/bridges, using the MCP server, or adding man page documentation.
user-invocable: true
---

# Carapace Usage Reference

Reference for using [carapace](https://carapace.sh) shell completion — integrating into CLIs, writing specs, creating actions, and configuring completions.

## Sub-Resources

Load the reference that matches your task. When in doubt, load multiple references.

| Keywords | Reference |
|----------|----------|
| Integrate carapace, cobra, PreRun, PreInvoke, bridge, standalone mode, flag config, non-POSIX flags | [references/integrate.md](references/integrate.md) |
| Custom actions, ActionValues, ActionCallback, ActionExecCommand, modifiers, naming, caching, batch, MultiParts, UID, tags | [references/action.md](references/action.md) |
| YAML spec, user spec, flags, subcommands, parsing modes, completion, macros, modifiers, runnable specs | [references/spec.md](references/spec.md) |
| Macro format, MacroN, MacroI, MacroV, $files, macro signatures, modifier chaining | [references/macro.md](references/macro.md) |
| Choices, variants, bridges, CARAPACE_BRIDGES, --choice, --list, group priority | [references/choice.md](references/choice.md) |
| Setup, shell integration, environment variables, overlays, extensions, install | [references/setup.md](references/setup.md) |
| Environment variable completion, env YAML, conditions, ActionNames, ActionValues | [references/env.md](references/env.md) |
| Convert spec to Go, codegen, macro-to-Go mapping, modifier translation | [references/convert.md](references/convert.md) |
| Scrape spec, generate from source, Docker, supported frameworks, patch-and-container | [references/scrape.md](references/scrape.md) |
| MCP server, complete tool, list_macros tool, codegen tool, JSON-RPC, stdio transport | [references/mcp.md](references/mcp.md) |
| Man pages, carapace-man, UID documentation, YAML man pages, inline descriptions, scheme, host, path | [references/man.md](references/man.md) |

## Quick Guide

- **How do I add carapace to my cobra CLI?** → [references/integrate.md](references/integrate.md)
- **How do I create a custom completion action?** → [references/action.md](references/action.md)
- **How do I write a YAML user spec?** → [references/spec.md](references/spec.md)
- **How do I format macro arguments?** → [references/macro.md](references/macro.md)
- **How do completers/bridges/choices work?** → [references/choice.md](references/choice.md)
- **How do I set up shell completion?** → [references/setup.md](references/setup.md)
- **How do I add environment variable completions?** → [references/env.md](references/env.md)
- **How do I convert a spec to a native Go completer?** → [references/convert.md](references/convert.md)
- **How do I generate a spec from CLI source code?** → [references/scrape.md](references/scrape.md)
- **How do I use the MCP server?** → [references/mcp.md](references/mcp.md)
- **How do I add inline documentation for completion values?** → [references/man.md](references/man.md)

## Cross-Project References

For internal library development (Action API internals, traverse engine, shell formatters, style system), use the **carapace-dev** skill (in the carapace repo).
