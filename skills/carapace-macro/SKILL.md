---
name: carapace-macro
description: >
  Use when user needs to look up, format, or understand carapace macros — how to retrieve available
  macros from the MCP tool, how macro arguments are formatted in user specs (YAML) and in Go
  (spec.ActionMacro), and the three macro types (MacroN/MacroI/MacroV). Triggers on: "macro format",
  "carapace macro", "list macros", "macro signature", "ActionMacro", "MacroN", "MacroI", "MacroV".
user-invocable: true
---

# Carapace Macro Format Guide

Macros are completion actions identified by a `$`-prefixed name, optionally taking arguments in parentheses. They are used in user specs (YAML) and in Go code via `spec.ActionMacro`.

## Retrieving Available Macros

Use the carapace MCP tool `list_macros` to get all available macros with their names, descriptions, and argument signatures.

Each entry has:

| Field | Description |
|-------|-------------|
| `name` | Fully qualified macro name (e.g. `carapace.tools.git.Refs`) |
| `description` | What the macro completes |
| `signature` | Argument signature string — see [Signature Format](#signature-format) below |

The CLI equivalent:

```bash
carapace --macro             # list all macros
carapace --macro <name>      # show macro details
carapace --macro <name> <TAB> # test macro completion
```

## Signature Format

The `signature` field from `list_macros` encodes the argument type:

| Signature | Macro type | Meaning |
|-----------|------------|---------|
| `""` (empty string) | MacroN | No arguments |
| `""` (quoted empty) | MacroI | Single string argument |
| `false` | MacroI | Single bool argument |
| `{field: value, ...}` | MacroI | Single struct argument |
| `[""]` | MacroV | Variadic string arguments |
| `[false]` | MacroV | Variadic bool arguments |

## Three Macro Types

### MacroN — No Arguments

Action takes no parameters.

| Context | Format |
|---------|--------|
| User spec | `$macroname` |
| Go register | `spec.MacroN(func() carapace.Action { ... })` |
| Go use | `spec.ActionMacro("$carapace.tools.git.Tags")` |

```yaml
completion:
  positional:
    - ["$directories"]
    - ["$carapace.tools.git.Tags"]
```

### MacroI — Single Argument

Action takes one parameter: a string, bool, or struct.

| Argument type | Signature | Spec format | Go use |
|--------------|-----------|-------------|--------|
| `string` | `""` | `$macroname(arg)` | `spec.ActionMacro("$macroname(\"arg\")")` |
| `bool` | `false` | `$macroname(true)` | `spec.ActionMacro("$macroname(true)")` |
| struct | `{Field: val}` | `$macroname({Field: val})` | `spec.ActionMacro("$macroname({Field: val})")` |

When invoked without arguments, a struct argument's `Default()` method is called if it exists. Use this for boolean fields where zero values produce empty results.

> **Note:** For single string arguments in user specs, quotes are omitted (`$macroname(arg)`) for ease of use. This keeps exec macros simple.

```yaml
# String argument (quotes optional in user specs for ease of use)
completion:
  positional:
    - ["$files([.go, .mod])"]
    - ["$carapace.tools.git.RemoteBranches(origin)"]

# Bool argument
completion:
  positional:
    - ["$carapace.shell.Functions(true)"]

# Struct argument
completion:
  positional:
    - ["$carapace.tools.git.Refs({localbranches: true, tags: true})"]
```

Go registration:
```go
spec.AddMacro("Refs", spec.MacroI(ActionRefs))
spec.AddMacro("Tags", spec.MacroN(ActionTags))
```

### MacroV — Variadic Arguments

Action takes a variadic parameter (`...string`, `...bool`, etc.). Arguments are passed as a YAML flow sequence.

| Argument type | Signature | Spec format |
|--------------|-----------|-------------|
| `...string` | `[""]` | `$macroname(["a", "b"])` |
| `...bool` | `[false]` | `$macroname([true, false])` |

```yaml
completion:
  positional:
    - ["$files"]
    - ["$carapace.tools.git.RefDiffs([\"HEAD~1\", \"HEAD~3\"])"]
```

Go registration:
```go
spec.AddMacro("RefDiffs", spec.MacroV(ActionRefDiffs))
```

## Macro Name Prefixes

Macro names use a dotted prefix to indicate their source:

| Prefix | Source | Example |
|--------|--------|---------|
| (none) | Core macros from carapace-spec | `$files`, `$directories`, `$message` |
| `<exe>.` | Fully qualified — the executable providing the macro | `$carapace.tools.git.Refs`, `$git.Refs`, `$kubectl.Refs` |
| `$_.` | Self-referencing (current executable) | `$_.Refs` |

### Fully Qualified vs Short Form

- `$carapace.tools.git.Refs` — fully qualified via the `carapace` executable, always works
- `$git.Refs` — fully qualified via the `git` executable (when `git` exposes this macro)
- `$_.Refs` — short form, resolves to the current executable's own macros (avoid unless explicitly asked for)

### Cross-Executable Macros

When a macro name contains a dot and doesn't match the current executable prefix, carapace delegates to that executable:

```yaml
# Delegates to `gh _carapace macro tools.gh.Labels`
["$carapace.tools.gh.Labels({host: \"github.com\"})"]
```

In Go, this is handled by `spec.ActionMacro` automatically:
```go
spec.ActionMacro("$carapace.tools.gh.Labels({host: \"github.com\"})")
```

## Using Macros in User Specs

In YAML specs, macros appear inside completion arrays prefixed with `$`:

```yaml
completion:
  flag:
    format: ["json", "yaml"]                         # static values
    file: ["$files([.txt, .md])"]                     # macro with string list arg
    branch: ["$carapace.tools.git.Refs({tags: true})"] # macro with struct arg
  positional:
    - ["$directories"]                                # macro without args
    - ["$carapace.tools.docker.Containers"]           # external macro
  positionalany: ["$files"]
```

## Using Macros in Go

### spec.ActionMacro

`spec.ActionMacro` resolves a macro string at runtime. Use it for loosely coupled access to macros exposed by other commands:

```go
// Reference a macro from carapace-bin
spec.ActionMacro("$carapace.tools.git.Refs")
spec.ActionMacro("$carapace.tools.git.Refs({tags: true})")
```

### Registering Custom Macros

Register macros before `spec.Register`:

```go
// No-argument macro
spec.AddMacro("tools.custom.Dirs", spec.MacroN(func() carapace.Action {
    return carapace.ActionDirectories()
}))

// Single-argument macro (struct)
spec.AddMacro("tools.custom.Filter", spec.MacroI(func(s string) carapace.Action {
    return carapace.ActionValues().Filter(s)
}))

// Variadic-argument macro
spec.AddMacro("tools.custom.Multi", spec.MacroV(func(s ...string) carapace.Action {
    return carapace.ActionValues(s...)
}))

spec.Register(rootCmd)
```

### Exposing External Actions via spec.ActionMacro

Bridge actions from carapace-bin into the spec system:

```go
spec.AddMacro("tools.git.Refs", spec.MacroN(
    spec.ActionMacro("$carapace.bridge.ActionCarapaceBin(git)"),
))
```

## Action Signature to Macro Type Mapping

The Go action signature determines which macro type to use:

| Action signature | Macro type | Spec usage |
|-----------------|------------|------------|
| `func ActionFoo()` | `spec.MacroN` | `$_.Foo` or `$exe.Foo` |
| `func ActionFoo(opts FooOpts)` | `spec.MacroI` | `$_.Foo({Field: value})` |
| `func ActionFoo(arg string)` | `spec.MacroI` | `$_.Foo("arg")` |
| `func ActionFoo(arg bool)` | `spec.MacroI` | `$_.Foo(true)` |
| `func ActionFoo(args ...string)` | `spec.MacroV` | `$_.Foo(["a", "b"])` |
| `func ActionFoo(args ...bool)` | `spec.MacroV` | `$_.Foo([true, false])` |

## Core Macros (No Prefix)

These are built into carapace-spec:

| Macro | Type | Purpose |
|-------|------|---------|
| `$directories` | MacroN | Complete directories |
| `$files` | MacroV | Complete files, optional suffix filter |
| `$executables` | MacroV | Complete executables, optional directory filter |
| `$message` | MacroI | Show an error/info message |
| `$spec` | MacroI | Embed another spec |
| `$sh`, `$bash`, `$zsh`, `$fish`, `$pwsh` | MacroI | Execute shell command for completions |

## Modifiers (Applied with `|||`)

Modifiers are special macros that transform the preceding action. Use ` ||| ` to apply a modifier to only the preceding value, or list it without `|||` to apply to all values in the batch:

```yaml
# Apply to all values in the list
["$files", "$chdir(/tmp)"]

# Apply only to the preceding value
["$files ||| $chdir(/tmp)"]
```

Modifier macros are also listed by `list_macros` and follow the same type rules (MacroN/MacroI/MacroV).

### Common Modifier Macros

| Modifier | Type | Purpose |
|----------|------|---------|
| `$chdir(<dir>)` | MacroI | Change working directory |
| `$filter([val])` | MacroV | Remove values |
| `$retain([val])` | MacroV | Keep only these values |
| `$list(,)` | MacroI | Join with delimiter |
| `$multiparts([/])` | MacroV | Split by delimiter |
| `$nospace(/)` | MacroI | No space after character |
| `$prefix(pre)` | MacroI | Prepend to values |
| `$suffix(suf)` | MacroI | Append to values |
| `$style(name)` | MacroI | Apply style |
| `$tag(name)` | MacroI | Tag values |
| `$usage(msg)` | MacroI | Usage message |
| `$suppress(val)` | MacroI | Suppress specific values |
| `$shift(n)` | MacroI | Remove first `n` positional args from `Context.Args` |

`$shift` reindexes which positional completion slot fires by slicing `Context.Args`. After `$shift(n)`, `Context.Args` becomes `Context.Args[n:]`, so `len(Context.Args)` decreases by `n`. Use it when bridging or recursing where the first arg(s) are routing arguments that shouldn't count as positionals:

```yaml
# Rightmost runs first: $shift(1) (outermost) fires first, shifts Args,
# then $filterargs (innermost) runs last and sees shifted Args
["$carapace.tools.git.Refs", "$filterargs", "$shift(1)"]

# Leftmost runs first: $filterargs (outermost) fires first, captures original Args,
# then $shift(1) (innermost) shifts — $filterargs already captured original Args
["$carapace.tools.git.Refs", "$shift(1)", "$filterargs"]

# Specific form: $shift scoped to one macro via |||
["$carapace.tools.git.Refs ||| $shift(1)"]
```

| `$split` | MacroN | Split command string |
| `$splitp` | MacroN | Split command string with pipes |
| `$filterargs` | MacroN | Filter already-provided positional args |
| `$uniquelist(delim)` | MacroI | Deduplicate delimited list |

### Modifier Chaining Order

The **outermost** modifier fires **first** at invocation time, then bubbles inward. Order matters when modifiers affect the same data (e.g., `$prefix`/`$suffix` change the value text that `$filter`/`$retain` match against):

```yaml
# Rightmost runs first: $prefix adds "p/" to all values, then $filter([x]) removes "p/x" from results
["$files ||| $filter([x]) ||| $prefix(p/)"]

# Leftmost runs first: $filter removes "x" first, then $prefix adds "p/" to remaining ("p/a", "p/b")
["$files ||| $prefix(p/) ||| $filter([x])"]

# Order doesn't matter when modifiers target independent fields
["$files ||| $tag(local) ||| $style(blue)"]
```

### Traverse Modifier Macros (Directory Paths)

| Modifier | Type | Purpose |
|----------|------|---------|
| `$gitdir` | MacroN | Change to .git directory |
| `$gitworktree` | MacroN | Change to git worktree root |
| `$parent([name, ...])` | MacroV | Change to first parent containing names |
| `$tempdir` | MacroN | Change to temp directory |
| `$userhomedir` | MacroN | Change to home directory |
| `$userconfigdir` | MacroN | Change to config directory |
| `$usercachedir` | MacroN | Change to cache directory |
| `$xdgcachehome` | MacroN | Change to XDG cache home |
| `$xdgconfighome` | MacroN | Change to XDG config home |
| `$nixprofile` | MacroN | Change to nix profile directory |

## Environment Variables in Macros

Macros in user specs support environment variable substitution using `drone/envsubst`:

| Variable | Content |
|----------|---------|
| `${C_FLAG_<NAME>}` | Flag value |
| `${C_ARG<n>}` | Positional argument (0-indexed) |
| `${C_VALUE}` | Current word being completed |
| `${C_PART<n>}` | Parts during multipart completion |

```yaml
# Use flag value as macro argument
completion:
  positionalany: ["$carapace.tools.git.LsRemoteRefs({url: '${C_ARG0}'})"]

# Default value if not set
${C_FLAG_SORT:-HEAD}
```
