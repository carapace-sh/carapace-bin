---
name: carapace-spec
description: >
  Use when user needs to create a carapace user spec (YAML completion file). Covers schema,
  flags, subcommands, parsing modes, runnable specs, and environment variables.
  For macro formatting and lookup, use the carapace-macro skill instead.
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

Add to top of spec file for schema validation:

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

Macros are `$`-prefixed completion actions used in completion arrays. For full details on macro types, formatting, and how to look up available macros, see the **carapace-macro** skill.

### Quick Reference

```yaml
completion:
  positional:
    - ["$directories"]
    - ["$files([.go, .mod])"]
    - ["$carapace.tools.git.Refs({tags: true})"]
```

### Executing Commands

```yaml
run: "$sh(echo one; echo two)"
run: "$bash(echo -e 'one\\ntwo')"
run: "$pwsh(echo one`ntwo)"
```

## Modifiers

Change completion behavior with ` ||| ` delimiter. For full modifier reference, see the **carapace-macro** skill.

```yaml
# Generic: apply to all values
["$files", "$chdir(/tmp)"]

# Specific: apply to preceding value only
["$files ||| $chdir(/tmp)"]
```

## Values Format

Values support description and style with `\t` delimiter:

```yaml
["value", "value\tdescription", "value\tdescription\tblue"]
```

Styles can be combined using space as delimiter (adopted from Elvish):

```yaml
["value\tdescription\tblue underlined"]
```

### Foreground Colors

`black`, `red`, `green`, `yellow`, `blue`, `magenta`, `cyan`, `white`, `bright-*`

### Background Colors

Use `bg-` prefix: `bg-black`, `bg-red`, `bg-green`, `bg-yellow`, `bg-blue`, `bg-magenta`, `bg-cyan`, `bg-white`, `bg-bright-*`

### Extended Colors

- **XTerm 256**: `color<0-255>` (e.g., `color123`)
- **True Color (Hex)**: `#RRGGBB` (e.g., `#ff0000`)

### Attributes

`bold`, `dim`, `italic`, `underlined`, `blink`, `inverse`, `default`, `fg-default`, `bg-default`

### Negation/Toggle

`no-<attr>` (e.g., `no-bold`), `toggle-<attr>` (e.g., `toggle-italic`)

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

> For macro formatting details, see the **carapace-macro** skill.

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

> For macro formatting details, see the **carapace-macro** skill.

- Generic bridge: `CarapaceBin`.
- Framework bridges: `Argcomplete`, `Aws`, `Bash`, `Carapace`, `Clap`, `Click`, `Cobra`, `Complete`, `Gcloud`, `Inshellisense`, `JJ`, `Kingpin`, `Kitten`, `Urfavecli`, `UrfavecliV1`, `Yargs`.
- Shell bridges: `Bash`, `Fish`, `Powershell`, `Zsh`.

## Testing

```bash
carapace _carapace           # reload specs
```

> For macro lookup and formatting, see the **carapace-macro** skill.
