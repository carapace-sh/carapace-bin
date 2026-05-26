---
name: carapace-spec
description: >
  Use when user needs to create a carapace user spec (YAML completion file). Covers schema,
  flags, subcommands, macros, modifiers, parsing modes, runnable specs, and environment variables.
  Triggers on: "create completion", "carapace spec", "user spec", "completion yaml", or any request to write a spec file.
user-invocable: true
---

# Carapace User Spec Creation Guide

Create shell completion specs for [carapace](https://carapace.sh) using YAML files.

## Location

User specs stored in: `${UserConfigDir}/carapace/specs/`

| OS | Path |
|----|------|
| Linux | `~/.config/carapace/specs/` |
| macOS | `~/Library/Application Support/carapace/specs/` |
| Windows | `%APPDATA%\carapace\specs\` |

> Filename must match spec name (e.g., `mycmd.yaml` for `name: mycmd`)

## Schema

Add to top of spec file for IDE validation:

```yaml
# yaml-language-server: $schema=https://carapace.sh/schemas/command.json
```

## Basic Structure

```yaml
name: mycmd
description: My command description
group: Some Group
hidden: false
parsing: interspersed
flags:
  -v, --verbose: be verbose
  -f, --file=: input file
persistentflags:
  --help: show help
commands:
  - name: sub
    description: subcommand
    flags:
      -x: exclusive flag
```

## Flags

### Modifier Suffixes

| Suffix | Meaning | Example |
|--------|---------|---------|
| `=` | takes argument | `--file=` |
| `*` | repeatable | `--tag*` |
| `?` | optional argument | `--output?` |
| `&` | hidden | `--secret&` |
| `!` | required | `--config!` |

```yaml
flags:
  -b: bool flag (no suffix)
  -v=: shorthand with value
  --repeatable*: repeatable flag
  -o, --optarg?: shorthand + optional arg
  --hidden&: hidden flag
  --required!: required flag
```

### Extended Format

For `nargs` or other options, use object syntax:

```yaml
flags:
  --two-args=: {description: consumes two args, nargs: 2}
  --any-args=: {description: consumes multiple, nargs: -1}
```

### Non-POSIX Shorthands

When built with carapace-pflag:

```yaml
flags:
  -single: non-posix shorthand
  -short, -long: both shorthand
```

## Parsing Modes

| Mode | Behavior |
|------|----------|
| `interspersed` | mixed flags and positional arguments (default) |
| `non-interspersed` | first positional stops flag parsing |
| `disabled` | no flag parsing (use for bridged completions) |

```yaml
parsing: disabled  # required for bridge macros
```

## Completion

### Flag Completion

```yaml
completion:
  flag:
    file: ["$files"]
    verbose: ["true", "false"]
```

### Positional Completion

```yaml
completion:
  positional:
    - ["$directories"]           # first positional
    - ["$executables"]           # second positional
  positionalany: ["$files"]      # all remaining positions
```

### Dash Completion (for flag-like args)

```yaml
completion:
  dash:
    - ["-a", "--alpha"]
    - ["-b", "--beta"]
  dashany: ["-x", "--xxx"]
```

## Documentation

```yaml
documentation:
  command: "Main command description"
  flag:
    file: "Input file path"
  positional:
    - "First argument: source"
    - "Second argument: destination"
  positionalany: "Additional files"
  dash:
    - "Short option"
    - "Long option"
```

## Macros

### Core Macros

```yaml
completion:
  positional:
    - ["$directories"]           # complete directories
    - ["$files([.go, .mod])"]     # filter by suffix
    - ["$executables", "$executables([~/.local/bin])"]
    - ["$message(error text)"]    # show error
    - ["$spec(other.yaml)"]       # embed another spec
```

### Executing Commands

```yaml
run: "$sh(echo one; echo two)"
run: "$bash(echo -e 'one\\ntwo')"
run: "$pwsh(echo one`ntwo)"
```

### List Available Macros

Use carapace MCP tool: `list_macros` to get all available macros with signatures.

Example output includes: `choice.Choices`, `color.HexColors`, `fs.Files`, `net.hosts`, `tools.git.Refs`, `tools.docker.Containers`, and hundreds more.

## Modifiers

Change completion behavior with ` ||| ` delimiter:

```yaml
# Generic: apply to all values
["$files", "$chdir(/tmp)"]

# Specific: apply to preceding value only
["$files ||| $chdir(/tmp)"]
```

### Common Modifiers

| Modifier | Purpose | Example |
|----------|---------|---------|
| `$chdir(<dir>)` | change directory | `["$files", "$chdir($gitdir)"]` |
| `$filter([val])` | remove values | `["one two three", "$filter([two])"]` |
| `$list(,)` | join with delimiter | `["a b c", "$list(,)"]` |
| `$multiparts(/)` | split by delimiter | `["one/two", "$multiparts([/])"]` |
| `$nospace(/)` | no space after char | `["dir/", "$nospace(/)"]` |
| `$prefix(pre)` | prepend to values | `["file", "$prefix(http://)"]` |
| `$suffix(suf)` | append to values | `["file", "$suffix(.txt)"]` |
| `$style(name)` | apply style | `["text", "$style(underlined)"]` |
| `$tag(name)` | tag values | `["text", "$tag(label)"]` |
| `$usage(msg)` | usage message | `["$usage(specify path)"]` |

### Traverse Modifiers (paths)

- `$gitdir` - .git folder
- `$gitworktree` - git worktree root
- `$parent([file, dir])` - first parent containing names
- `$tempdir` - temp directory
- `$userhomedir` - home directory
- `$userconfigdir` - config directory
- `$usercachedir` - cache directory
- `$xdgcachehome`, `$xdgconfighome` - XDG dirs

## Values Format

Values support description and style with `\t` delimiter:

```yaml
["value", "value\tdescription", "value\tdescription\tblue"]
```

Styles: `default`, `bold`, `dim`, `italic`, `underlined`, `blinking`, `inverse`, `hidden`, `strike`.

## Runnable Specs

Specs with `run` field become executable via shims.

### Alias (Command Array)

Aliases retain existing completion unless otherwise defined.

```yaml
name: gh-issues
run: ["gh", "issue", "--repo", "${REPO}"]
```

### Script (Macro with Environment Substitution)

```yaml
name: ls-remote
run: "$(git ls-remote --sort='${C_FLAG_SORT:-HEAD}' $@)"
flags:
  --sort=: field to sort on
completion:
  positionalany: ["$carapace.tools.git.LsRemoteRefs({url: '${C_ARG0}'})"]
```

### Shebang (Bash Script)

```yaml
name: myscript
run: |
  #!/usr/bin/env bash -x
  for arg in "$@"; do
    echo "${!arg}${C_FLAG_SUFFIX:-.suffix}"
  done
flags:
  -s, --suffix=: suffix to add
completion:
  positionalany: [one, two]
```

## Environment Variables

Variables replaced via `drone/envsubst`:

| Variable | Content |
|----------|---------|
| `${C_FLAG_<NAME>}` | flag value |
| `${C_ARG<n>}` | positional argument (0-indexed) |
| `${C_VALUE}` | current word being completed |
| `${C_PART<n>}` | parts during multipart completion |

### Defaults

```yaml
${C_FLAG_VERBOSE:-false}    # default if not set
${C_FLAG_SUFFIX//,/, }      # string replacement
```

## Example

```yaml
# yaml-language-server: $schema=https://carapace.sh/schemas/command.json
name: mytool
description: My custom tool
group: Development
parsing: interspersed
flags:
  -v, --verbose: verbose output
  --quiet: quiet output
  -f, --file=!: input file
  --format=: output format
persistentflags:
  --help: show help
exclusiveflags:
  - [verbose, quiet]
completion:
  flag:
    file: ["$files([.txt, .md])"]
    format: ["json\tJavaScript Object Notation\tblue underlined", "yaml\thuman-readable data serialization", "xml"]
  positional:
   - ["$files"]
documentation:
  command: "My tool does something useful"
  flag:
    file: "Path to input file (required)"
    format: "Output format (default: json)"
```

## Bridging External Completions

Use `parsing: disabled` and bridge macro:

```yaml
name: kubectl
description: Kubernetes CLI
parsing: disabled
completion:
  positionalany: ["$carapace.bridge.CarapaceBin([kubectl])"]
```

Generic bridge: `CarapaceBin`.
Framework bridges: `Argcomplete`, `Aws`, `Bash`, `Carapace`, `Clap`, `Click`, `Cobra`, `Complete`, `Gcloud`, `Inshellisense`, `JJ`, `Kingpin`, `Kitten`, `Urfavecli`, `UrfavecliV1`, `Yargs`.
Shell bridges: `Bash`, `Fish`, `Powershell`, `Zsh`.

## Testing

```bash
carapace _carapace           # reload specs
carapace --macro             # list all macros
carapace --macro <name>      # show macro details
carapace --macro <name> <TAB> # test macro
```
