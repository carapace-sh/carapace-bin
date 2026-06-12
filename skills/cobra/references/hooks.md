# Lifecycle Hooks

How cobra's PreRun, PostRun, and PersistentPreRun/PersistentPostRun hooks work.

## Hook Execution Order

When a command runs, hooks execute in this order:

```
1. PersistentPreRun / PersistentPreRunE   (inherited, parent → child)
2. PreRun / PreRunE                       (current command only)
3. Run / RunE                              (main logic)
4. PostRun / PostRunE                     (current command only)
5. PersistentPostRun / PersistentPostRunE  (inherited, child → parent)
```

The `*E` variants return an `error`. If a hook returns an error, execution stops — subsequent hooks and `Run` are skipped.

## Persistent Hooks

`PersistentPreRun` and `PersistentPostRun` are **inherited** by child commands. When a child runs, the parent's persistent hooks execute.

### Default Behavior (EnableTraverseRunHooks = false)

Only the **first** persistent hook found in the command chain executes. If a child defines its own `PersistentPreRun`, the parent's is **not** called:

```
root (PersistentPreRun: A)
  └── sub (PersistentPreRun: B)
        └── leaf (no PersistentPreRun)

# Running "sub":
#   B runs (child's hook)
#   A does NOT run (parent's hook is skipped)

# Running "leaf":
#   B runs (inherited from sub)
#   A does NOT run
```

### Traverse Mode (EnableTraverseRunHooks = true)

**All** parents' persistent hooks execute, from outermost to innermost:

```go
cobra.EnableTraverseRunHooks = true
```

```
root (PersistentPreRun: A)
  └── sub (PersistentPreRun: B)

# Running "sub":
#   A runs (root's hook)
#   B runs (sub's hook)
```

## Hook Variants

| Hook | Inherited | Error-returning |
|------|-----------|-----------------|
| `PersistentPreRun` | Yes | No |
| `PersistentPreRunE` | Yes | Yes |
| `PreRun` | No | No |
| `PreRunE` | No | Yes |
| `Run` | — | No |
| `RunE` | — | Yes |
| `PostRun` | No | No |
| `PostRunE` | No | Yes |
| `PersistentPostRun` | Yes | No |
| `PersistentPostRunE` | Yes | Yes |

When both the error-returning and non-error-returning variant are defined on the same command, the `*E` variant takes precedence.

## Common Patterns

### Global Setup with PersistentPreRun

```go
var verbose bool

rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")

rootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
    if verbose {
        log.SetLevel(log.DebugLevel)
    }
}
```

All subcommands inherit the verbose flag and the pre-run hook.

### Viper Config Loading

```go
cobra.OnInitialize(initConfig)

func initConfig() {
    if cfgFile != "" {
        viper.SetConfigFile(cfgFile)
    } else {
        viper.AddConfigPath(home)
        viper.SetConfigName(".app")
    }
    viper.AutomaticEnv()
    viper.ReadInConfig()
}
```

`OnInitialize` registers functions that run before **every** command execution. See [references/viper-integration.md](references/viper-integration.md) for the full pattern.

### Error-Returning PreRun for Validation

```go
cmd := &cobra.Command{
    Use: "deploy ENV",
    PreRunE: func(cmd *cobra.Command, args []string) error {
        if !isValidEnv(args[0]) {
            return fmt.Errorf("invalid environment: %s", args[0])
        }
        return nil
    },
    RunE: func(cmd *cobra.Command, args []string) error {
        // deploy logic
        return nil
    },
}
```

## Execution Order Within execute()

The internal `execute()` method runs hooks in this precise order:

```
1. Print deprecation warning (if Deprecated is set)
2. InitDefaultHelpFlag, InitDefaultVersionFlag
3. ParseFlags
4. Check --help flag → return flag.ErrHelp
5. Check --version flag → print version and return
6. Check Runnable() → return flag.ErrHelp if not runnable
7. preRun() — runs PersistentPreRun chain, then PreRun
8. ValidateArgs
9. ValidateRequiredFlags
10. ValidateFlagGroups
11. Run / RunE
12. PostRun / PostRunE
13. postRun() — runs PersistentPostRun chain
```

## Edge Cases

- **Hooks only fire if Run/RunE is declared**: If a command has no `Run`/`RunE`, `PreRun` and `PostRun` do not execute. The command is treated as a "help topic" command.
- **PersistentPreRunE stops the chain**: If a parent's `PersistentPreRunE` returns an error, no further hooks or `Run` execute.
- **OnInitialize vs PersistentPreRun**: `OnInitialize` runs before the command tree is traversed. `PersistentPreRun` runs after traversal, on the resolved command. Use `OnInitialize` for global setup (Viper, logging), `PersistentPreRun` for per-command setup that needs access to flags.
- **OnFinalize**: Registered functions run after command execution completes (success or error).

## References

- [cobra source: command.go](https://github.com/spf13/cobra/blob/main/command.go) — `execute()` method
- [cobra source: cobra.go](https://github.com/spf13/cobra/blob/main/cobra.go) — `OnInitialize`, `OnFinalize`

## Related Skills

- For the full execute() flow and how Find/Traverse interacts with hooks, see the **cobra-dev** skill → `references/execute-flow.md`.
