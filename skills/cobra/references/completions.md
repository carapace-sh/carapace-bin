# Shell Completions

How to add shell completions to cobra applications. Cobra supports bash, zsh, fish, and PowerShell.

## Static Completions with ValidArgs

For fixed argument choices:

```go
cmd := &cobra.Command{
    Use:       "color",
    ValidArgs: []string{"red", "green", "blue"},
    Run:       func(cmd *cobra.Command, args []string) {},
}
```

With descriptions (tab-separated):

```go
cmd.ValidArgs = []string{
    cobra.CompletionWithDesc("red", "The color red"),
    cobra.CompletionWithDesc("green", "The color green"),
}
```

## Dynamic Completions with ValidArgsFunction

For completions computed at runtime:

```go
cmd := &cobra.Command{
    Use: "color",
    ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
        return []string{"red", "green", "blue"}, cobra.ShellCompDirectiveNoFileComp
    },
    Run: func(cmd *cobra.Command, args []string) {},
}
```

### CompletionFunc Signature

```go
type CompletionFunc = func(cmd *cobra.Command, args []string, toComplete string) ([]Completion, ShellCompDirective)
```

- `cmd` — the command being completed
- `args` — already-provided positional arguments
- `toComplete` — the partial word being typed
- Returns: completion candidates + a directive

## Flag Completions with RegisterFlagCompletionFunc

Register a completion function for a specific flag:

```go
cmd.RegisterFlagCompletionFunc("output", func(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
    return []string{"json", "yaml", "xml"}, cobra.ShellCompDirectiveNoFileComp
})
```

Flag completion functions are stored in a global `map[*pflag.Flag]CompletionFunc` protected by a `sync.RWMutex`.

## ShellCompDirective

Controls shell behavior after completions are returned:

| Directive | Value | Behavior |
|-----------|-------|----------|
| `ShellCompDirectiveDefault` | 0 | Shell performs default behavior (file completion if no matches) |
| `ShellCompDirectiveError` | 1 | Error occurred, ignore completions |
| `ShellCompDirectiveNoSpace` | 2 | Don't add a space after the completion |
| `ShellCompDirectiveNoFileComp` | 4 | Don't fall back to file completion |
| `ShellCompDirectiveFilterFileExt` | 8 | Completions are file extension filters (e.g., `*.go`) |
| `ShellCompDirectiveFilterDirs` | 16 | Only complete directory names |
| `ShellCompDirectiveKeepOrder` | 32 | Preserve the order of completions (don't sort) |

Directives are a bitmask — combine with `|`:

```go
return completions, cobra.ShellCompDirectiveNoFileComp | cobra.ShellCompDirectiveNoSpace
```

## Built-in Completion Helpers

```go
// Disable file completion (e.g., for enum-like flags)
cobra.NoFileCompletions

// Fixed/static completions
cobra.FixedCompletions([]string{"json", "yaml"}, cobra.ShellCompDirectiveNoFileComp)
```

## Active Help

Active Help displays contextual messages during completion. These appear as informational hints, not as completion candidates:

```go
ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
    var comps []cobra.Completion
    if len(args) == 0 {
        comps = cobra.AppendActiveHelp(comps, "You must provide a color name")
    }
    comps = append(comps, "red", "green", "blue")
    return comps, cobra.ShellCompDirectiveNoFileComp
}
```

### Disabling Active Help

Users can disable Active Help via environment variables:

```sh
# Disable for all cobra programs
export COBRA_ACTIVE_HELP=0

# Disable for a specific program
export MYAPP_ACTIVE_HELP=0
```

Program-specific env var format: `<PROGRAM>_ACTIVE_HELP` where the program name is uppercased and non-alphanumeric characters become `_`.

### GetActiveHelpConfig

Check the active help configuration at runtime:

```go
config := cobra.GetActiveHelpConfig(cmd)
if config == "0" {
    // active help is disabled
}
```

## CompletionOptions

Control the auto-generated `completion` subcommand:

```go
rootCmd.CompletionOptions = cobra.CompletionOptions{
    DisableDefaultCmd:   true,   // don't create 'completion' command
    DisableNoDescFlag:   true,   // don't add '--no-descriptions' flag
    DisableDescriptions: true,   // always strip descriptions
    HiddenDefaultCmd:    true,   // hide 'completion' from help
}
```

Set a default directive for all completions:

```go
rootCmd.CompletionOptions.SetDefaultShellCompDirective(cobra.ShellCompDirectiveNoFileComp)
```

## Generating Completion Scripts

### At Runtime (Recommended)

The `completion` subcommand generates scripts:

```sh
myapp completion bash > /etc/bash_completion.d/myapp
myapp completion zsh > "${fpath[1]}/_myapp"
myapp completion fish > ~/.config/fish/completions/myapp.fish
myapp completion powershell | Out-File -Encoding utf8 -FilePath MyApp.ps1
```

### Programmatically

```go
// Bash V2 (recommended)
rootCmd.GenBashCompletionV2(os.Stdout, true)  // true = include descriptions

// Zsh
rootCmd.GenZshCompletion(os.Stdout)

// Fish
rootCmd.GenFishCompletion(os.Stdout, true)  // true = include descriptions

// PowerShell
rootCmd.GenPowerShellCompletionWithDesc(os.Stdout)
```

### To File

```go
rootCmd.GenBashCompletionFileV2("/etc/bash_completion.d/myapp", true)
rootCmd.GenZshCompletionFile("/path/to/_myapp")
rootCmd.GenFishCompletionFile("~/.config/fish/completions/myapp.fish", true)
rootCmd.GenPowerShellCompletionFileWithDesc("MyApp.ps1")
```

## Completion Flow

When a user presses TAB, the shell script calls the program with the hidden `__complete` command:

```
myapp __complete cmd subcmd --flag val
```

The program:
1. Finds the target command
2. Checks if completing a flag value → calls `RegisterFlagCompletionFunc`
3. Checks if completing a flag name → suggests available flags
4. Checks if completing a subcommand → suggests subcommand names
5. Falls back to `ValidArgs` or `ValidArgsFunction`
6. Outputs completions (one per line) + directive as `:<integer>`

## Environment Variables

| Variable | Purpose |
|----------|---------|
| `COBRA_ACTIVE_HELP` | Disable active help globally (`0` = off) |
| `<PROG>_ACTIVE_HELP` | Disable active help per-program |
| `COBRA_COMPLETION_DESCRIPTIONS` | Control description display at runtime |
| `<PROG>_COMPLETION_DESCRIPTIONS` | Per-program description control |
| `BASH_COMP_DEBUG_FILE` | Debug file for bash completion |

## Edge Cases

- **Completion command and arbitrary args**: If the root has no real subcommands, the `completion` command is only kept if it's actually being called. This prevents breaking programs that accept arbitrary arguments.
- **Flag value with = syntax**: `--flag=value` completions work — cobra detects the `=` and completes the value portion.
- **Already-set flags**: Flags that are already set are excluded from completion unless they accept multiple values (Slice, Array, stringTo types).
- **Required flags in completion**: Required flags are listed first in flag name completion.
- **Bash V1 vs V2**: V1 generates a large standalone script. V2 delegates to the Go binary via `__complete`. Always use V2 for new applications.

## References

- [cobra shell completions docs](https://github.com/spf13/cobra/blob/main/site/content/completions/_index.md)
- [cobra active help docs](https://github.com/spf13/cobra/blob/main/site/content/active_help.md)

## Related Skills

- For the internal completion engine (`__complete`, `getCompletions`, directive handling), see the **cobra-dev** skill → `references/completion-engine.md`.
- For per-shell script generation internals (bash V1/V2, zsh, fish, powershell), see the **cobra-dev** skill → `references/shell-completions.md`.
