# Positional Arguments

How to validate positional arguments in cobra commands.

## The Args Field

Set the `Args` field on a `Command` to validate positional arguments before `Run`/`RunE` executes:

```go
cmd := &cobra.Command{
    Use:  "clone URL [NAME]",
    Args: cobra.ExactArgs(1),
    Run:  func(cmd *cobra.Command, args []string) { /* ... */ },
}
```

If `Args` is `nil`, cobra applies `legacyArgs` which:
- Root commands **with subcommands**: requires a valid subcommand (rejects unknown args)
- Root commands **without subcommands**: accepts arbitrary args
- Subcommands: accepts arbitrary args

## Built-in Validators

| Validator | Behavior |
|-----------|----------|
| `NoArgs` | Error if any positional args exist |
| `ArbitraryArgs` | Never errors (accepts any number) |
| `MinimumNArgs(n)` | Error if fewer than `n` args |
| `MaximumNArgs(n)` | Error if more than `n` args |
| `ExactArgs(n)` | Error if not exactly `n` args |
| `RangeArgs(min, max)` | Error if args not in `[min, max]` range |
| `OnlyValidArgs` | Error if any arg not in `ValidArgs` field |
| `NoDuplicateArgs` | Error if any arg value is repeated |

## Combining Validators with MatchAll

`MatchAll` runs multiple validators and returns the first error:

```go
cmd := &cobra.Command{
    Use:  "color COLOR",
    Args: cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
    ValidArgs: []string{"red", "green", "blue"},
    Run: func(cmd *cobra.Command, args []string) { /* ... */ },
}
```

This requires exactly 1 argument **and** that argument must be in `ValidArgs`.

## Custom Validators

Any function with signature `func(cmd *cobra.Command, args []string) error`:

```go
cmd := &cobra.Command{
    Use: "color COLOR",
    Args: func(cmd *cobra.Command, args []string) error {
        if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
            return err
        }
        if isValidColor(args[0]) {
            return nil
        }
        return fmt.Errorf("invalid color: %s", args[0])
    },
    Run: func(cmd *cobra.Command, args []string) { /* ... */ },
}
```

## ValidArgs and ArgAliases

`ValidArgs` provides static completion candidates for positional arguments. Each entry can include a tab-separated description:

```go
cmd.ValidArgs = []string{
    "red\tThe color red",
    "green\tThe color green",
    "blue\tThe color blue",
}
```

Or use the helper:

```go
cmd.ValidArgs = []string{
    cobra.CompletionWithDesc("red", "The color red"),
}
```

`ArgAliases` provides additional names that are accepted by `OnlyValidArgs` but not shown in completions:

```go
cmd.ArgAliases = []string{"r", "g", "b"}  // accepted but not suggested
```

## Validation Order

Argument validation happens **after** flag parsing but **before** `Run`/`RunE`:

```
ParseFlags → ValidateArgs → ValidateRequiredFlags → ValidateFlagGroups → Run/RunE
```

## Edge Cases

- **Default behavior**: If `Args` is nil, `legacyArgs` applies — this is a common source of confusion for root commands with subcommands.
- **OnlyValidArgs strips descriptions**: The validator splits on `\t` and only checks the value part, so tab-separated descriptions in `ValidArgs` don't interfere.
- **NoDuplicateArgs**: Checks for exact string matches — `"foo"` and `"FOO"` are considered different.
- **ExactValidArgs is deprecated**: Use `MatchAll(ExactArgs(n), OnlyValidArgs)` instead.

## References

- [cobra source: args.go](https://github.com/spf13/cobra/blob/main/args.go)

## Related Skills

- For completions (ValidArgsFunction, dynamic argument completion), see [references/completions.md](references/completions.md).
