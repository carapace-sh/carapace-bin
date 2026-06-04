# Carapace Environment Variable Completion Guide

Define and customize environment variable name and value completions in [carapace](https://carapace.sh).

## Overview

Environment variable completion has two layers:

| Layer | Location | Mechanism |
|-------|----------|-----------|
| Built-in | `pkg/actions/env/*.go` | Go `init()` registration into `knownVariables` map |
| User YAML | `~/.config/carapace/variables/*.yaml` | Loaded at runtime, overrides built-in |

The public API is in `pkg/actions/env`:

| Function | Purpose |
|----------|---------|
| `ActionNames()` | Completes environment variable names (built-in + custom) |
| `ActionValues(varName)` | Completes values for a specific variable |
| `ActionNameValues(positional)` | Combined `NAME=VALUE` completion (flag or positional) |

## Go-Based Definitions

### Registration Pattern

Each file in `pkg/actions/env/` registers a group via `init()`:

```go
package env

func init() {
	knownVariables["docker"] = func() variables {
		return variables{
			Condition: conditions.ConditionPath("docker"),
			Variables: map[string]string{
				"DOCKER_HOST":          "Daemon socket to connect to",
				"DOCKER_BUILDKIT":      "Enable BuildKit for docker build",
			},
			VariableCompletion: map[string]carapace.Action{
				"DOCKER_HOST": carapace.ActionValues("tcp://", "unix://"),
			},
		}
	}
}
```

Key points:
- **Group key** is a logical name (e.g. `"git"`, `"docker"`, `"aws"`, `"common"`)
- **Factory function** — `knownVariables["key"] = func() variables { ... }` for lazy evaluation
- **`Condition`** — gates the entire group; `nil` means always active
- **`VariableCompletion`** keys must be a subset of `Variables` keys (enforced by test)

### Conditions

| Condition | Description |
|-----------|-------------|
| `nil` | Always active (e.g. `common`, `carapace`) |
| `conditions.ConditionPath("docker")` | Active when `docker` is on `$PATH` |
| `conditions.ConditionPath("aws")` | Active when `aws` is on `$PATH` |
| Build tags in filename | Platform gating (e.g. `common_unix.go`) — no runtime condition needed |

Available conditions from `pkg/conditions`:
- **`Path`** — executable exists in PATH
- **`Parent`** — parent directory contains given file/directory
- **`Os`** — current `runtime.GOOS` matches
- **`Arch`** — current `runtime.GOARCH` matches

### Completion Action Patterns

| Pattern | Example |
|---------|---------|
| Boolean (true/false) | `carapace.ActionValues("true", "false").StyleF(style.ForKeyword)` |
| Boolean (0/1) | `carapace.ActionValuesDescribed("0", "disabled", "1", "enabled").StyleF(style.ForKeyword)` |
| Styled boolean | `carapace.ActionStyledValuesDescribed("0", "disabled", style.Carapace.KeywordNegative, "1", "enabled", style.Carapace.KeywordPositive)` |
| Directories | `carapace.ActionDirectories()` |
| Files | `carapace.ActionFiles()` |
| Path list | `carapace.ActionDirectories().List(string(os.PathListSeparator))` |
| Comma list | `carapace.ActionValues(...).UniqueList(",")` |
| Tool-specific action | `aws.ActionProfiles()`, `git.ActionConfigValues(...)` |
| Bridge to other completer | `bridge.ActionCarapaceBin("go").Split()` |
| MultiParts (structured value) | `carapace.ActionMultiParts(";", ...)` |
| Batch (union) | `carapace.Batch(carapace.ActionFiles(), carapace.ActionValues("off").StyleF(...)).ToA()` |
| Static values | `carapace.ActionValues("off", "on", "auto")` |

### Helper Variables

Reuse common actions within a group:

```go
_bool := carapace.ActionValues("true", "false").StyleF(style.ForKeyword)
// ...
"NO_COLOR":  _bool,
"TERM_DUMB": _bool,
```

### Adding a New Tool Group

1. Create `pkg/actions/env/<tool>.go`
2. Register with `init()` using `knownVariables["<tool>"]`
3. Add `Condition` if the tool isn't always available
4. List all relevant env vars in `Variables` with descriptions
5. Add completion actions for vars where values are known/enumerable
6. Ensure every `VariableCompletion` key exists in `Variables`

## User YAML Definitions

Custom environment variable definitions live in `~/.config/carapace/variables/<group>.yaml`. They can define new variables, add completions, or override built-in definitions.

### YAML Format

```yaml
condition:
  - $Parent([.git])
variables:
  CUSTOM_VAR: description of the variable
  OTHER_VAR: another variable
completion:
  variable:
    CUSTOM_VAR:
      - "value1\tdescription1"
      - "value2\tdescription2"
    OTHER_VAR:
      - $carapace.tools.git.RemoteBranches
```

> **Note:** Values containing a tab character (`\t`) for description or style must be wrapped in double quotes, otherwise YAML treats `\t` as a literal backslash followed by `t` (two characters), and the style/description will not be applied.

### Fields

| Field | Type | Description |
|-------|------|-------------|
| `condition` | `[]string` | Condition macro strings — group is only active when all conditions pass |
| `variables` | `map[string]string` | Variable name → description |
| `completion.variable` | `map[string][]string` | Variable name → list of completion spec strings |

### Completion Spec Strings

Each entry in a variable's completion list is a spec string:

| Format | Meaning |
|--------|---------|
| `value` | Static value |
| `"value\tdescription"` | Value with description (tab-separated, must be double-quoted) |
| `"value\tdescription\tstyle"` | Value with description and style (tab-separated, must be double-quoted) |
| `$macroname` | Macro reference (e.g. `$files`, `$directories`) |
| `$carapace.tools.<tool>.<Action>` | Fully qualified macro |
| `\|\|\|` suffix | Action modifiers after `\|\|\|` (e.g. `$files \|\|\| $uniquelist(,)`) |

> **Important:** Values containing a tab character (`\t`) for description or style **must be wrapped in double quotes** (`"value\tdescription\tstyle"`). YAML does not interpret `\t` as a tab in unquoted strings — it remains as literal backslash-`t`, and the description/style will not be applied.

### Overriding Built-in Definitions

User YAML files **override** built-in completions. When a custom YAML defines a completion for a variable that also has a built-in completion, the custom one replaces it.

```yaml
# ~/.config/carapace/variables/proxy.yaml
variables:
  HTTPS_PROXY: override existing variable description
completion:
  variable:
    HTTPS_PROXY:
      - "https://localhost:8443\tdevelopment\tgreen"
      - "https://proxy.company:443\tproduction\tred"
```

> **Note:** Values containing a tab character (`\t`) for description or style must be wrapped in double quotes, otherwise YAML treats `\t` as a literal backslash followed by `t` (two characters), and the style/description will not be applied.

### Condition Examples

```yaml
# Only active inside a git repository
condition:
  - $Parent([.git])

# Only active on Linux
condition:
  - $Os([linux])

# Only active when docker is in PATH
condition:
  - $Path([docker])
```

## How Completion Flows Work

### Name Completion (`ActionNames`)

Merges two sources:
1. `actionKnownEnvironmentVariables()` — all groups from `knownVariables` where `Condition` passes
2. `actionCustomEnvironmentVariables()` — all YAML files from `~/.config/carapace/variables/`

Also includes OS-level environment variables (from `os.ActionEnvironmentVariables()`) styled in blue.

### Value Completion (`ActionValues`)

For a given variable name:
1. Look up built-in completion in `knownVariables`
2. Check user YAML files for overrides
3. If a custom YAML defines a completion for that variable, it replaces the built-in
4. If no completion is found at all, falls back to `carapace.ActionFiles()`

### Combined Name=Value (`ActionNameValues`)

Two modes:
- **Positional** (`positional=true`): First positional completes names, second completes values
- **Flag** (`positional=false`): `ActionMultiPartsN("=", 2, ...)` — first part is name (with `NoSpace`), second part is value
