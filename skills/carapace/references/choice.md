# Carapace Choices, Variants & Bridges

How carapace resolves which completer to use when multiple sources exist for the same command.

## Completer Sources

A single command can have completers from multiple sources:

| Source | Location | Group |
|--------|----------|-------|
| Native Go completers | `completers/` (built-in) | `common`, `unix`, `linux`, `darwin`, `bsd`, `windows`, `android` |
| User YAML specs | `${UserConfigDir}/carapace/specs/` | `user` |
| System YAML specs | `${UserConfigDir}/carapace/specs/` | `system` |
| Known bridges | hardcoded in `bridges.go` (~80 commands) | `bridge` |
| Shell bridges | enabled via `CARAPACE_BRIDGES` | `bridge` |
| Overlays | `${UserConfigDir}/carapace/overlays/` | modifies existing variant |

## Group Priority

When multiple groups provide a completer for the same command, they resolve in this order (highest wins):

1. **`user`** — user YAML specs
2. **`system`** — system YAML specs
3. **OS-specific** (`linux`, `darwin`, `windows`, `android`, `bsd`) — matches current OS
4. **`unix`** — Unix-like
5. **`common`** — cross-platform
6. **`bridge`** — lowest priority

A **choice** (set via `--choice`) overrides all of the above — see [Choices](#choices) below.

Within the same group, variants are ordered alphabetically by name.

## Listing Variants

```sh
carapace --list {command}       # list all variants for a command
carapace --list @{group}       # list completers in a group
carapace --list gh@bridge      # list bridge variants for 'gh'
```

## Choices

When multiple variants exist for a command, use `--choice` to select the preferred one:

```sh
carapace --choice {name}[/{variant}][@{group}]
```

Examples:

```sh
# Prefer bsd sed over the default
carapace --choice sed@bsd

# Prefer a specific tldr variant
carapace --choice tldr/tldr-python-client

# Use the cobra bridge for gh instead of the native completer
carapace --choice gh/cobra@bridge
```

Choices are persisted as simple text files in `${UserConfigDir}/carapace/choices/`:

```
carapace/
└── choices/
    ├── sed              # content: sed@bsd
    └── tldr             # content: tldr/tldr-python-client
```

To remove a choice, delete the file: `rm ~/.config/carapace/choices/gh`

A choice gives that variant the **highest priority**, overriding all group ordering.

## Bridges

Bridges delegate completion to another engine when carapace lacks a native completer.

### Implicit Bridges (Shell-Based)

Enabled via `CARAPACE_BRIDGES`. These discover any command that has completion registered in the source shell:

```sh
export CARAPACE_BRIDGES='zsh,fish,bash,inshellisense'
```

| Bridge | What it uses | Notes |
|--------|-------------|-------|
| `zsh` | Zsh's completion system | Requires zsh installed |
| `fish` | Fish's completion system | Requires fish installed |
| `bash` | Bash's completion system | Requires bash and completions loaded |
| `inshellisense` | [inshellisense](https://github.com/microsoft/inshellisense) | Requires inshellisense installed |

Default order: `zsh,fish,bash,inshellisense`. The completer list is **cached** — clear with `carapace --clear-cache` after installing new shells or commands.

> Shell bridges should be a fallback — framework bridges produce better completions with less overhead.

### Explicit Bridges (Framework-Based)

These must be set explicitly via `--choice` or a user spec because carapace **doesn't auto-detect** which framework a tool was built with by default. For example, carapace doesn't know that `gh` (GitHub CLI) uses cobra — you must tell it. (There is an experimental `bridge.Detect()` function in carapace-bridge that probes a command to auto-detect its framework, but it is not used at runtime.)

**Via choice:**
```sh
carapace --choice gh/cobra@bridge
```

**Via user spec** (in `${UserConfigDir}/carapace/specs/gh.yaml`):
```yaml
# yaml-language-server: $schema=https://carapace.sh/schemas/command.json
name: gh
description: GitHub CLI
parsing: disabled
completion:
  positionalany: ["$carapace.bridge.Cobra([gh])"]
```

### Known Bridges

carapace-bin includes a hardcoded map of ~80 commands with their known framework (in `cmd/carapace/cmd/completers/bridges.go`). These are low-priority `bridge` group defaults — a native completer or user choice always overrides them.

Full list of bridge types:

| Bridge | Framework/Source | Macro |
|--------|-----------------|-------|
| `argcomplete` | [kislyuk/argcomplete](https://github.com/kislyuk/argcomplete) | `$carapace.bridge.Argcomplete([cmd])` |
| `argcomplete_v1` | argcomplete (legacy) | `$carapace.bridge.ArgcompleteV1([cmd])` |
| `aws` | [aws/aws-cli](https://github.com/aws/aws-cli) | `$carapace.bridge.Aws([cmd])` |
| `bash` | [bash](https://www.gnu.org/software/bash/) | `$carapace.bridge.Bash([cmd])` |
| `carapace` | [carapace-sh/carapace](https://github.com/carapace-sh/carapace) | `$carapace.bridge.Carapace([cmd])` |
| `carapace-bin` | carapace-bin registry | `$carapace.bridge.CarapaceBin([cmd])` |
| `clap` | [clap-rs/clap](https://github.com/clap-rs/clap) (experimental) | `$carapace.bridge.Clap([cmd])` |
| `click` | [pallets/click](https://github.com/pallets/click) | `$carapace.bridge.Click([cmd])` |
| `cobra` | [spf13/cobra](https://github.com/spf13/cobra) | `$carapace.bridge.Cobra([cmd])` |
| `complete` | [posener/complete](https://github.com/posener/complete) | `$carapace.bridge.Complete([cmd])` |
| `fish` | [fish](https://fishshell.com/) | `$carapace.bridge.Fish([cmd])` |
| `gcloud` | [Google Cloud SDK](https://docs.cloud.google.com/sdk/gcloud) | `$carapace.bridge.Gcloud([cmd])` |
| `inshellisense` | [inshellisense](https://github.com/microsoft/inshellisense) | `$carapace.bridge.Inshellisense([cmd])` |
| `jj` | [jj-vcs](https://www.jj-vcs.dev) | `$carapace.bridge.Jj([cmd])` |
| `kingpin` | [alecthomas/kingpin](https://github.com/alecthomas/kingpin) | `$carapace.bridge.Kingpin([cmd])` |
| `kitten` | [kitty](https://sw.kovidgoyal.net/kitty/) | `$carapace.bridge.Kitten([cmd])` |
| `powershell` | [powershell](https://microsoft.com/powershell) | `$carapace.bridge.Powershell([cmd])` |
| `urfavecli` | [urfave/cli](https://github.com/urfave/cli) | `$carapace.bridge.Urfavecli([cmd])` |
| `urfavecli_v1` | urfave/cli (legacy) | `$carapace.bridge.UrfavecliV1([cmd])` |
| `yargs` | [yargs/yargs](https://github.com/yargs/yargs) | `$carapace.bridge.Yargs([cmd])` |
| `zsh` | [zsh](https://www.zsh.org/) | `$carapace.bridge.Zsh([cmd])` |

Custom shell bridge configurations can be placed in `${UserConfigDir}/carapace/bridge/`.

### How Bridge Resolution Works

The `ActionBridge` function (from carapace-bridge) resolves completions at runtime:

1. **Check choices first** — looks up `~/.config/carapace/choices/<name>`; if a choice exists with `group == "bridge"` (or empty), uses its variant as the bridge type
2. **Fall back to `CARAPACE_BRIDGES`** — iterates the env var in order, checking if the command is known to each shell bridge (bash, fish, zsh, inshellisense)
3. **Return empty** — if neither yields a match, returns no completions

### Determining the Right Bridge

An agent can look up a tool's framework by:
1. Checking the tool's dependencies (e.g. `go list -m` for Go tools, `pip show` for Python)
2. Searching the tool's source code for framework imports (e.g. `import "github.com/spf13/cobra"`)
3. Checking the tool's documentation or `--help` output for framework clues
4. Checking the hardcoded map in `cmd/carapace/cmd/completers/bridges.go` for known mappings
5. Using the experimental `bridge.Detect()` function from carapace-bridge, which probes a command to auto-detect its completion framework by testing each bridge type

Common framework signatures:
- **Cobra** (Go): `import "github.com/spf13/cobra"`, completion command `cmd completion <shell>`
- **Click** (Python): `import click`, `@click.command()`
- **Argcomplete** (Python): `import argcomplete`, `argcomplete.autocomplete(parser)`
- **urfave/cli** (Go): `import "github.com/urfave/cli/v2"`, `cli.App{}`
- **Yargs** (Node.js): `require('yargs')`, `yargs.command(...)`
- **Clap** (Rust): `use clap::Parser`, `#[derive(Parser)]`
- **Gcloud** (Python): uses argcomplete internally — check for `gcloud` binary and `CLOUDSDK_*` env vars

## Troubleshooting

- **Completions not appearing after install**: Run `carapace --clear-cache` to refresh the completer list.
- **Wrong completer used**: Check `carapace --list <cmd>` to see variants, then use `--choice` to override.
- **Bridge completions empty**: Ensure the target shell/framework is installed and completions are registered in it.
- **Changes not taking effect**: Restart your shell or re-source the init snippet.