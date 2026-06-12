# Help, Usage, and Version

How cobra generates help, usage, and version output, and how to customize it.

## Automatic Help

Cobra automatically adds:
- A `--help` (`-h`) flag on every command
- A `help` subcommand on the root when subcommands exist

These are always available:

```sh
app help           # show root help
app help create    # show help for 'create' subcommand
app create --help  # same as above
```

## Customizing Help

### Replace the Help Command

```go
rootCmd.SetHelpCommand(&cobra.Command{
    Use:   "help [command]",
    Short: "Custom help",
    Run:   func(cmd *cobra.Command, args []string) { /* custom logic */ },
})
```

### Replace the Help Function

```go
cmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
    fmt.Println("Custom help for", cmd.Name())
})
```

### Customize the Help Template

```go
cmd.SetHelpTemplate(`{{.Name}} - {{.Short}}

Usage:
  {{.UseLine}}

{{if .HasAvailableSubCommands}}Commands:{{range .Commands}}{{if .IsAvailableCommand}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}

{{if .HasAvailableFlags}}Flags:
{{.Flags().FlagUsages}}{{end}}
`)
```

Templates use Go's `text/template` with these built-in functions: `trim`, `trimRightSpace`, `trimTrailingWhitespaces`, `appendIfNotPresent`, `rpad`, `gt`, `eq`. Add custom functions with `cobra.AddTemplateFunc` or `cobra.AddTemplateFuncs`.

## Command Grouping in Help

Group subcommands under headings:

```go
rootCmd.AddGroup(&cobra.Group{
    ID:    "basic",
    Title: "Basic Commands:",
})
rootCmd.AddGroup(&cobra.Group{
    ID:    "advanced",
    Title: "Advanced Commands:",
})

listCmd.GroupID = "basic"
debugCmd.GroupID = "advanced"
```

Control where auto-generated commands appear:

```go
rootCmd.SetHelpCommandGroupId("basic")
rootCmd.SetCompletionCommandGroupId("advanced")
```

## Usage Messages

Usage is shown when the user provides invalid input (wrong flags, wrong number of args). The default help embeds usage in its output.

### Customize Usage Function

```go
cmd.SetUsageFunc(func(cmd *cobra.Command) error {
    fmt.Fprintf(cmd.ErrOrStderr(), "Usage: %s\n", cmd.UseLine())
    return nil
})
```

### Customize Usage Template

```go
cmd.SetUsageTemplate(`Usage: {{.UseLine}}

{{if .HasAvailableSubCommands}}Available Commands:{{range .Commands}}{{if .IsAvailableCommand}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}
`)
```

## Version Flag

Setting `Version` on a command enables `--version`:

```go
rootCmd.Version = "1.2.3"
```

### Customize Version Template

```go
cmd.SetVersionTemplate(`{{.Name}} version {{.Version}}
`)
```

## Error Prefix

The default error prefix is `Error:`. Customize it:

```go
cmd.SetErrPrefix("FATAL")
// Output: FATAL: something went wrong
```

## Silence Errors and Usage

```go
cmd.SilenceErrors = true   // don't print "Error: ..." to stderr
cmd.SilenceUsage = true    // don't print usage on error
```

These are useful when you want to handle errors yourself in `main()`:

```go
func main() {
    if err := rootCmd.Execute(); err != nil {
        // custom error handling
        os.Exit(1)
    }
}
```

## Template Variables

Available in help, usage, and version templates:

| Variable | Description |
|----------|-------------|
| `.Use` | Usage line |
| `.Name` | Command name |
| `.Short` | Short description |
| `.Long` | Long description |
| `.Example` | Example strings |
| `.Version` | Version string |
| `.Commands` | Subcommands |
| `.Flags` | Flag set |
| `.HasAvailableSubCommands` | Whether subcommands exist |
| `.HasAvailableFlags` | Whether flags exist |
| `.HasInheritedFlags` | Whether inherited flags exist |
| `.UseLine` | Full usage line with command path |
| `.CommandPath` | Full command path |

## Edge Cases

- **Help command conflicts**: If you define a `help` subcommand manually, cobra won't add the default one.
- **--help on non-runnable commands**: Returns `flag.ErrHelp`, which triggers the help function. This is why commands without `Run` show help.
- **SilenceUsage**: Only suppresses usage on error. Help output from `--help` is still shown.
- **Template binary size**: `text/template` adds to binary size. For minimal binaries, use `SetHelpFunc`/`SetUsageFunc` instead of templates.

## References

- [cobra source: command.go](https://github.com/spf13/cobra/blob/main/command.go) — help and usage methods
- [cobra source: cobra.go](https://github.com/spf13/cobra/blob/main/cobra.go) — template functions

## Related Skills

- For the internal template rendering mechanism, see the **cobra-dev** skill → `references/global-state.md`.
