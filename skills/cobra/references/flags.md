# Flags

How to define, organize, and validate flags in cobra applications. Cobra uses [pflag](https://github.com/spf13/pflag) for POSIX-compliant flag parsing.

## Flag Types

### Persistent Flags

Available to the command **and all its subcommands**. Use for global options like `--verbose`:

```go
rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
```

### Local Flags

Available **only to the specific command** where they are defined:

```go
localCmd.Flags().StringVarP(&Source, "source", "s", "", "Source directory to read from")
```

### Local Non-Persistent Flags

Flags that are local and **not inherited** by children. These are local flags that are not also persistent:

```go
cmd.LocalNonPersistentFlags()  // rarely needed directly
```

## Flag Registration Patterns

### Basic Types

```go
// Bool
cmd.Flags().BoolP("debug", "d", false, "enable debug mode")
cmd.Flags().BoolVarP(&debug, "debug", "d", false, "enable debug mode")

// String
cmd.Flags().StringP("output", "o", "json", "output format")
cmd.Flags().StringVarP(&output, "output", "o", "json", "output format")

// Int
cmd.Flags().IntP("port", "p", 8080, "port number")
cmd.Flags().IntVarP(&port, "port", "p", 8080, "port number")
```

The `VarP` variants bind to a variable. The `P` suffix adds a shorthand.

### Count Flags (Repeated)

For `-v`, `-vv`, `-vvv` verbosity patterns:

```go
var verbose int
rootCmd.PersistentFlags().CountVarP(&verbose, "verbose", "v", "verbosity level")
```

Each occurrence increments the counter: `-vv` → `verbose == 2`.

### Slice/Array Flags

```go
var files []string

// StringArrayVarP: each --input creates a separate entry
cmd.Flags().StringArrayVarP(&files, "input", "i", []string{}, "input files")

// StringSliceVarP: comma-separated OR repeated
// --input=a,b --input=c → ["a", "b", "c"]
cmd.Flags().StringSliceVarP(&files, "input", "i", []string{}, "input files")
```

| Type | `--flag=a,b --flag=c` | Behavior |
|------|----------------------|----------|
| `StringArray` | `["a,b", "c"]` | Each occurrence is a separate entry (commas are literal) |
| `StringSlice` | `["a", "b", "c"]` | Comma-separated values are split |

## Required Flags

```go
rootCmd.Flags().StringVarP(&Region, "region", "r", "", "AWS region (required)")
rootCmd.MarkFlagRequired("region")

// For persistent flags:
rootCmd.PersistentFlags().StringVarP(&Region, "region", "r", "", "AWS region (required)")
rootCmd.MarkPersistentFlagRequired("region")
```

Cobra validates required flags before running the command. Missing required flags produce an error.

## Flag Groups

### Required Together

All flags in the group must be set, or none of them:

```go
rootCmd.Flags().StringVarP(&user, "username", "u", "", "Username")
rootCmd.Flags().StringVarP(&pass, "password", "p", "", "Password")
rootCmd.MarkFlagsRequiredTogether("username", "password")
```

| `--username` | `--password` | Result |
|-------------|-------------|--------|
| ✗ | ✗ | ✓ (neither set — ok) |
| ✓ | ✗ | ✗ (error: missing password) |
| ✓ | ✓ | ✓ (both set — ok) |

### Mutually Exclusive

At most one flag from the group can be set:

```go
rootCmd.Flags().BoolVar(&json, "json", false, "JSON output")
rootCmd.Flags().BoolVar(&yaml, "yaml", false, "YAML output")
rootCmd.MarkFlagsMutuallyExclusive("json", "yaml")
```

| `--json` | `--yaml` | Result |
|----------|----------|--------|
| ✗ | ✗ | ✓ (neither set — ok) |
| ✓ | ✗ | ✓ (one set — ok) |
| ✓ | ✓ | ✗ (error: mutually exclusive) |

### One Required

At least one flag from the group must be set. Can combine with mutually exclusive for "exactly one":

```go
rootCmd.Flags().BoolVar(&json, "json", false, "JSON output")
rootCmd.Flags().BoolVar(&yaml, "yaml", false, "YAML output")
rootCmd.MarkFlagsOneRequired("json", "yaml")
rootCmd.MarkFlagsMutuallyExclusive("json", "yaml")  // exactly one
```

### Flag Group Rules

- Both local and persistent flags can participate in groups
- A group is only enforced on commands where **every flag in the group is defined**
- A flag can belong to **multiple groups**
- A group can contain **any number of flags**

## TraverseChildren

By default, cobra only parses local flags on the **target command**. To also parse local flags on parent commands during traversal:

```go
command := cobra.Command{
    Use:              "print [OPTIONS] [COMMANDS]",
    TraverseChildren: true,
}
```

This enables the `Traverse` resolution path instead of `Find`. See the **cobra-dev** skill for the internal difference.

## Flag Annotations

Cobra uses pflag annotations for special behaviors:

```go
// Mark a flag as filename (for shell completion)
cmd.MarkFlagFilename("config")

// Mark a flag as directory
cmd.MarkFlagDirname("output")

// Mark a flag as custom (for bash V1 completion)
cmd.MarkFlagCustom("output", "__complete_output")
```

## Accessing Flag Values

```go
// After ParseFlags (which happens before Run)
val, err := cmd.Flags().GetString("name")
val, err := cmd.Flags().GetBool("verbose")
val, err := cmd.Flags().GetInt("port")

// Check if a flag was explicitly set (not default)
changed := cmd.Flags().Changed("name")
```

## Edge Cases

- **Persistent flag shadowing**: If a child defines a local flag with the same name as a parent's persistent flag, the local flag takes precedence on the child.
- **Flag parsing disabled**: `DisableFlagParsing = true` skips all flag parsing — args are passed raw to `Run`.
- **FParseErrWhitelist**: Controls which flag-parse errors are silently ignored:

```go
cmd.FParseErrWhitelist = cobra.FParseErrWhitelist{
    UnknownFlags: true,  // ignore unknown flag errors
}
```

- **SortFlags**: pflag sorts flags by default. Disable with `cmd.Flags().SortFlags = false` to preserve registration order in help output.

## References

- [pflag documentation](https://github.com/spf13/pflag)
- [cobra user guide: flags](https://github.com/spf13/cobra/blob/main/site/content/user_guide.md)

## Related Skills

- For the internal flag resolution mechanism (mergePersistentFlags, ParseFlags, flag set hierarchy), see the **cobra-dev** skill → `references/flag-resolution.md`.
- For flag group validation internals (annotation storage, processFlagForGroupAnnotation), see the **cobra-dev** skill → `references/flag-groups-internals.md`.
- For pflag non-POSIX extensions, see the **carapace-dev** skill → `references/pflag.md`.
