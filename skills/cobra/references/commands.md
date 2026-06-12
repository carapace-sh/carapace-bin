# Creating and Organizing Commands

How to define, organize, and wire up cobra commands and subcommands.

## The Command Struct

Every cobra CLI is built from `*cobra.Command` values. The key user-facing fields:

```go
cmd := &cobra.Command{
    Use:   "app [flags] FILE",  // usage line (command name + syntax)
    Short: "Brief one-line description",  // shown in parent's help
    Long:  `Long multi-line description shown in 'app --help'`,  // shown in this command's help
    Example: "app --verbose input.txt",  // example usage strings
    Aliases: []string{"alt"},  // alternative names
    SuggestFor: []string{"remove"},  // suggest this command for these names
    GroupID: "group-name",  // group in parent's help output
    Deprecated: "use 'app new-cmd' instead",  // deprecation message (non-empty = deprecated)
    Hidden: true,  // hide from help listing
    Version: "1.2.3",  // enables --version flag
    Annotations: map[string]string{
        cobra.CommandDisplayNameAnnotation: "kubectl myplugin",  // for plugin display names
    },
    Run: func(cmd *cobra.Command, args []string) { /* main logic */ },
    RunE: func(cmd *cobra.Command, args []string) error { /* main logic with error */ },
}
```

### Use Field

The `Use` string defines the command name (first word) and the argument syntax. The name is everything before the first space:

```go
Use: "clone URL [REPOSITORY]"  // Name() == "clone"
```

### Run vs RunE

- `Run` — for commands that cannot fail; panics on errors are unhandled
- `RunE` — returns an `error`; cobra prints the error and (unless `SilenceErrors`) shows usage

Always prefer `RunE` for production code.

## Project Structure

### Minimal Structure

```
▾ app/
  ▾ cmd/
      root.go
  main.go
```

`main.go`:

```go
package main

import "app/cmd"

func main() {
    if err := cmd.Execute(); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}
```

`cmd/root.go`:

```go
package cmd

var rootCmd = &cobra.Command{
    Use:   "app",
    Short: "My application",
    Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
    return rootCmd.Execute()
}
```

### Larger Applications

Organize subcommands in nested directories matching the command tree:

```
▾ app/
  ▾ cmd/
      root.go
      ▾ sub1/
          sub1.go
          ▾ sub2/
              sub2.go
              leafA.go
              leafB.go
  main.go
```

Wire up in `init()`:

```go
// root.go
func init() { rootCmd.AddCommand(sub1Cmd) }

// sub1.go
func init() { sub1Cmd.AddCommand(sub2Cmd) }

// sub2.go
func init() { sub2Cmd.AddCommand(leafACmd, leafBCmd) }
```

This avoids cyclic references because each level only imports its direct children.

### Using cobra-cli Generator

```sh
go install github.com/spf13/cobra-cli@latest
cobra-cli init appname
cobra-cli add serve
cobra-cli add config
```

Generates the project scaffold with `cmd/` directory, `root.go`, and per-command files.

## Adding Subcommands

```go
rootCmd.AddCommand(subCmd)
```

`AddCommand` sets the parent, updates max-length caches for help alignment, propagates the global normalization function, and panics if a command is added to itself.

### Removing Subcommands

```go
rootCmd.RemoveCommand(subCmd)
```

## Root Command Patterns

### Require a Subcommand

Omit `Run`/`RunE` on the root command. When the user types just `app`, cobra shows the help message.

### Execute as Root

Always call `Execute()` on the root command. If called on a subcommand, cobra internally redirects to the root:

```go
func (c *Command) ExecuteC() (cmd *Command, err error) {
    if c.HasParent() {
        return c.Root().ExecuteC()
    }
    // ...
}
```

### Context Propagation

Use `ExecuteContext` to pass a context through the entire command tree:

```go
ctx := context.Background()
rootCmd.ExecuteContext(ctx)
```

The context is stored on the root and propagated to the resolved command. Access it in `Run`/`RunE` via `cmd.Context()`.

## Command Introspection

| Method | Returns |
|--------|---------|
| `Name()` | Command name (first word of `Use`) |
| `CommandPath()` | Full path: `root sub1 sub2` |
| `UseLine()` | Full usage line: `root sub1 [flags]` |
| `HasParent()` | Whether this command has a parent |
| `Root()` | The root command |
| `Parent()` | Direct parent |
| `Commands()` | Slice of direct children |
| `HasSubCommands()` | Whether any children exist |
| `IsAvailableCommand()` | Whether this command is usable (not deprecated/hidden) |
| `IsAdditionalHelpTopicCommand()` | Whether this is a help topic command (no `Run`, `Long` != `Short`) |
| `CalledAs()` | The name/alias actually used to invoke this command |

## Plugin Support

For kubectl-style plugins (e.g., `kubectl-myplugin` invoked as `kubectl myplugin`):

```go
rootCmd := &cobra.Command{
    Use: "kubectl-myplugin",
    Annotations: map[string]string{
        cobra.CommandDisplayNameAnnotation: "kubectl myplugin",
    },
}
```

`CommandDisplayNameAnnotation` overrides how the command appears in help and error messages.

## Command Grouping

Group subcommands in help output:

```go
rootCmd.AddGroup(&cobra.Group{
    ID:    "management",
    Title: "Management Commands:",
})

// Assign subcommands to the group
listCmd.GroupID = "management"
getCmd.GroupID = "management"

// Control where auto-generated help/completion commands appear
rootCmd.SetHelpCommandGroupId("management")
rootCmd.SetCompletionCommandGroupId("management")
```

## Edge Cases

- **Omitting Run/RunE**: Commands without `Run` or `RunE` show help when invoked. This is the standard pattern for root commands that require subcommands.
- **Deprecated commands**: Setting `Deprecated` to a non-empty string prints the deprecation message but still executes the command.
- **Hidden commands**: `Hidden = true` removes the command from help listings but it remains callable.
- **Aliases**: Aliases are exact matches — they don't create separate commands. `CalledAs()` reports which name was used.
- **AddCommand panic**: Adding a command to itself panics. Adding the same command to multiple parents overwrites the previous parent.

## References

- [cobra user guide](https://github.com/spf13/cobra/blob/main/site/content/user_guide.md)
- [cobra source: command.go](https://github.com/spf13/cobra/blob/main/command.go)

## Related Skills

- For the internal Command struct fields and execute flow, see the **cobra-dev** skill → `references/command-struct.md` and `references/execute-flow.md`.
- For flag details, see [references/flags.md](references/flags.md).
