# Integrating Carapace into Cobra-based CLI Applications

This guide covers integrating [carapace](https://github.com/carapace-sh/carapace) as a library into existing applications using [cobra](https://github.com/spf13/cobra).

## Basic Setup

### Installation

```bash
go get github.com/carapace-sh/carapace
```

### Minimal Integration

Add `carapace.Gen(cmd)` to your root command:

```go
import "github.com/carapace-sh/carapace"

rootCmd := &cobra.Command{
    Use:   "myapp",
    Short: "My application",
}
carapace.Gen(rootCmd)  // initializes carapace for the root command
rootCmd.Execute()
```

This registers the `_carapace` subcommand for spec generation: `myapp _carapace spec`

## Core Concepts

### PreRun — Modifying Command Structure

`PreRun` executes **before** argument parsing and allows dynamic modification of the command structure (adding/removing subcommands, flags, etc.).

Use cases:
- Dynamic subcommand registration (e.g., plugins loaded at runtime)
- Conditional command structure based on environment or configuration
- Adding commands for external tools/extensions with completions

```go
import "github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"

carapace.Gen(rootCmd).PreRun(func(cmd *cobra.Command, args []string) {
    // Discover and add extensions dynamically
    extensions := discoverExtensions() // your extension discovery logic
    for _, extension := range extensions {
        extCmd := &cobra.Command{
            Use:                extension.Name,
            Short:              extension.Description,
            DisableFlagParsing: true,
            Run:                func(cmd *cobra.Command, args []string) {},
        }

        // Bridge completions from external completer (carapace-bin)
        carapace.Gen(extCmd).PositionalAnyCompletion(
            bridge.ActionCarapaceBin("myapp-" + extension.Name),
        )
        cmd.AddCommand(extCmd)
    }
})
```

The above pattern enables:
- **Dynamic plugin discovery** at runtime
- **Delegating completions** to external completers via `bridge.ActionCarapaceBin`
- **Spec-based completions** when plugins provide their own carapace specs

### PreInvoke — Modifying Actions Before Execution

`PreInvoke` executes **after** argument parsing but **before** action invocation. It allows modification of completion actions based on parsed flags/arguments.

Primary use case: **changing the working directory** for completions (e.g., `git -C <dir>` support).

#### Simple Chdir — when flag always has a value

```go
carapace.Gen(rootCmd).PreInvoke(func(cmd *cobra.Command, flag *pflag.Flag, action carapace.Action) carapace.Action {
    return action.Chdir(cmd.Flag("chdir").Value.String())
})
```

#### Conditional Chdir — check if flag was actually set

```go
carapace.Gen(rootCmd).PreInvoke(func(cmd *cobra.Command, flag *pflag.Flag, action carapace.Action) carapace.Action {
    if f := cmd.Flag("chdir"); f.Changed {
        return action.Chdir(f.Value.String())
    }
    return action
})
```

#### Deferred Chdir with traverse.Flag — resolves flag value at invoke time

Prefer `ChdirF(traverse.Flag(...))` over `Chdir(flag.Value.String())` when the flag value
may change between PreInvoke and actual action invocation:

```go
import "github.com/carapace-sh/carapace/pkg/traverse"

carapace.Gen(rootCmd).PreInvoke(func(cmd *cobra.Command, flag *pflag.Flag, action carapace.Action) carapace.Action {
    return action.ChdirF(traverse.Flag(cmd.Flag("chdir")))
})
```

## Action System

### Action Types

Carapace provides rich action types in `carapace` package:

| Function | Purpose |
|----------|---------|
| `ActionValues(vals ...string)` | Static completion values |
| `ActionValuesDescribed(vals ...string)` | Values with descriptions (val, desc, val, desc, ...) |
| `ActionFiles(suffix ...string)` | Files with optional suffix filter |
| `ActionDirectories()` | Directories only |
| `ActionExecutables(dirs ...string)` | Executable commands (defaults to PATH) |
| `ActionCommands(cmd *cobra.Command)` | Subcommands of a command |
| `ActionCallback(func(c Context) Action)` | Dynamic completions |
| `ActionExecCommand(name string, arg ...string)` | Execute external command for completions |
| `ActionExecCommandE(name string, arg ...string)` | Like ExecCommand with error handling |
| `ActionExecute(cmd *cobra.Command)` | Delegate to a cobra subcommand's completions |
| `ActionMessage(msg string, args ...any)` | Display a help message (format string) |

### Flag Completion with ActionMap

Map completion actions to specific flags:

```go
carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
    "file":      carapace.ActionFiles(),
    "format":    carapace.ActionValues("json", "yaml", "xml"),
    "verbose":   carapace.ActionValues("true", "false"),
    "output-dir": carapace.ActionDirectories(),
})
```

### Positional Completion

```go
carapace.Gen(cmd).PositionalCompletion(
    carapace.ActionDirectories(),           // first positional: directories
    carapace.ActionFiles(),                 // second positional: files
)

carapace.Gen(cmd).PositionalAnyCompletion(
    carapace.ActionFiles(),                 // all remaining positions
)
```

### Action Modifiers

Actions can be modified with chained calls:

```go
carapace.ActionValues("one", "two", "three").
    Suffix(":").                      // add suffix to values
    Style(style.Blue).                // apply style
    Tag("numbers").                   // tag for filtering
    Usage("specify a number")         // usage hint
```

#### StyleF

`StyleF` applies a dynamic style function that can change based on the value and context:

```go
import "github.com/carapace-sh/carapace/pkg/style"

carapace.ActionValues("file.txt", "image.png", "data.json").
    StyleF(style.ForPathExt)
```

Common style functions in `github.com/carapace-sh/carapace/pkg/style`:

| Function | Purpose |
|----------|---------|
| `style.ForPath(path, sc)` | Style based on full path |
| `style.ForPathExt(path, sc)` | Style based on file extension |
| `style.ForExtension(path, sc)` | Style for extension only |
| `style.ForKeyword(s, sc)` | Style for keywords |
| `style.ForLogLevel(s, sc)` | Style for log levels |

Pre-defined styles include:
- Colors: `style.Red`, `style.Blue`, `style.Green`, etc.
- Bright variants: `style.BrightRed`, etc.
- Background colors: `style.BgRed`, etc.
- Effects: `style.Bold`, `style.Dim`, `style.Italic`, `style.Underlined`
- Combinations: `style.Of(style.Bold, style.Red)`

#### NoSpace

`NoSpace` disables space suffix for specific characters, useful when the completion value shouldn't be followed by a space:

```go
carapace.ActionValues(
    "one,",
    "two/",
    "three",
).NoSpace(',', '/')
```

This is particularly useful for:
- Path-like completions (e.g., `foo/` should not add a space after `/`)
- Values ending with punctuation (e.g., `value,`)
- Multi-part completions where the divider handles spacing

### Combining Actions

**Batch** — execute multiple actions in parallel:

```go
carapace.Batch(
    carapace.ActionValues("opt1", "opt2"),
    carapace.ActionExecutables(),
).ToA()
```

**MultiParts** — completions with multiple parts (e.g., paths):

```go
carapace.ActionMultiParts("/", func(c carapace.Context) carapace.Action {
    switch len(c.Parts) {
    case 0:
        return carapace.ActionDirectories().Suffix("/")
    default:
        return carapace.ActionFiles()
    }
})
```

Note: The divider is implicitly added to `NoSpace`, so the delimiter won't add a trailing space.

## Standalone Mode

For applications that don't use cobra's built-in `help` command or completion command, use `Standalone()` to prevent cobra defaults from interfering:

```go
carapace.Gen(rootCmd).Standalone()
```

This disables cobra's default completion subcommand, hides the implicit `help` command and `--help` flag, and sets the `carapace_standalone` annotation. Use this when the completer is a synthetic cobra command that mimics a CLI's interface rather than using the actual tool's cobra command tree.

## Directory Traversal with ChdirF

Use `ChdirF` with traversal functions for dynamic directory changes:

```go
import "github.com/carapace-sh/carapace/pkg/traverse"

carapace.ActionFiles().ChdirF(traverse.GitWorkTree)   // relative to git worktree root
carapace.ActionFiles().ChdirF(traverse.GitDir)        // relative to .git directory
carapace.ActionFiles().ChdirF(traverse.Flag(cmd.Flag("C")))  // relative to flag value
```

Available traversal functions:
| Function | Purpose |
|----------|---------|
| `traverse.Flag(f *pflag.Flag)` | Resolves to the current value of a flag |
| `traverse.GitWorkTree` | Git worktree root |
| `traverse.GitDir` | .git directory |
| `traverse.Parent(names ...string)` | First parent containing specific files/dirs |
| `traverse.UserHomeDir` | User home directory |
| `traverse.UserCacheDir` | User cache directory |
| `traverse.UserConfigDir` | User config directory |
| `traverse.XdgCacheHome` | XDG cache home |
| `traverse.XdgConfigHome` | XDG config home |

## Dynamic Completions with ActionCallback and ActionExecCommand

`ActionCallback` wraps a Go function that is invoked during completion, providing access to the `carapace.Context` (current value, positional args, environment):

```go
carapace.ActionCallback(func(c carapace.Context) carapace.Action {
    // c.Value - current partial value being completed
    // c.Args  - positional arguments of current subcommand
    // c.Env   - environment variables
    // c.Dir   - working directory
    if f := cmd.Flag("repo"); f.Changed {
        return carapace.ActionValues(f.Value.String())
    }
    return carapace.ActionMessage("--repo is required")
})
```

Carapace core actions (`ActionFiles`, `ActionDirectories`, etc.) defer execution implicitly. Custom action helpers that access runtime state **must** be wrapped in `ActionCallback` to prevent execution at init time. Wrap inside the helper so callers can use it directly:

```go
func actionModels() carapace.Action {
    return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
        // access runtime state (flags, config, filesystem) here
        return carapace.ActionValues("provider/model1", "provider/model2")
    })
}

carapace.Gen(runCmd).FlagCompletion(carapace.ActionMap{
    "model":       actionModels(),
    "small-model": actionModels(),
})
```

`ActionExecCommand` executes an external command and processes its stdout. It returns a curried function — call with the command first, then with the callback:

```go
carapace.ActionExecCommand("git", "branch", "--show-current")(func(output []byte) carapace.Action {
    return carapace.ActionValues(strings.Split(string(output), "\n")[0])
})
```

`ActionExecCommandE` is the error-handling variant:

```go
carapace.ActionExecCommandE("git", "remote")(func(output []byte, err error) carapace.Action {
    if err != nil {
        return carapace.ActionMessage(err.Error())
    }
    return carapace.ActionValues(strings.Split(string(output), "\n")...)
})
```

`ActionExecute` delegates completion to a **cobra subcommand** within the same process. Use `.Shift(n)` to skip positional args:

```go
carapace.ActionExecute(someSubCmd).Shift(1)
```

## carapace-bridge: External Completion Bridging

[carapace-bridge](https://github.com/carapace-sh/carapace-bridge) bridges external completion engines (bash, fish, zsh, argcomplete, etc.) for shells not natively supported.

### Installation

```bash
go get github.com/carapace-sh/carapace-bridge
```

### Usage

```go
import "github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"

// Bridge completions from carapace-bin
bridge.ActionCarapaceBin("gh")

// Bridge completions from a cobra-based command
bridge.ActionCobra("podman")

// Bridge completions from Python argcomplete
bridge.ActionArgcomplete("mycommand")
```

Available bridge functions (all return `carapace.Action`):
| Function | Completion Engine |
|----------|-----------------|
| `bridge.ActionCarapaceBin(cmd ...string)` | carapace-bin completers/specs |
| `bridge.ActionCarapace(cmd ...string)` | Carapace self-bridging |
| `bridge.ActionCobra(cmd ...string)` | Cobra native completions |
| `bridge.ActionBash(cmd ...string)` | Bash native completions |
| `bridge.ActionFish(cmd ...string)` | Fish native completions |
| `bridge.ActionZsh(cmd ...string)` | Zsh native completions |
| `bridge.ActionArgcomplete(cmd ...string)` | Python argcomplete |
| `bridge.ActionClap(cmd ...string)` | Rust clap |
| `bridge.ActionClick(cmd ...string)` | Python Click |
| `bridge.ActionYargs(cmd ...string)` | Node.js yargs |
| `bridge.ActionBridge(cmd ...string)` | Auto-detect via CARAPACE_BRIDGE env |

### Split and SplitP — Completing Command-String Flag Values

When a flag accepts a command string (e.g., `--editor`, `--command`, `TF_CLI_ARGS`), use `.Split()` or `.SplitP()` on the bridge action. These parse the current completion value as a shell command line using shlex, split it into tokens, and complete the **current token** within that command string:

```go
carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
    // --editor accepts a command like "vim --noplugin"
    "editor": bridge.ActionCarapaceBin().Split(),

    // --command accepts a subcommand with args like "terraform apply -var"
    "command": bridge.ActionCarapaceBin("terraform").Split(),
})
```

- **`.Split()`** — for simple command strings without pipelines
- **`.SplitP()`** — like Split but supports shell pipelines (`|`, `>`, `>>`); use when the flag value might contain pipes or redirects

Common patterns:
```go
// Environment variable that holds a command
bridge.ActionCarapaceBin().Split()

// Environment variable that holds args for a specific command
bridge.ActionCarapaceBin("go").Split()

// Filter command that might contain pipes
bridge.ActionCarapaceBin().SplitP()

// Git hook command with ! prefix
bridge.ActionCarapaceBin().Split().Prefix("!")
```

## carapace-spec: Exposing Actions for External Commands

[carapace-spec](https://github.com/carapace-sh/carapace-spec) allows exposing actions to other commands via spec registration.

### Installation

```bash
go get github.com/carapace-sh/carapace-spec
```

### Registering the Spec Subcommand

```go
import spec "github.com/carapace-sh/carapace-spec"

spec.Register(rootCmd)  // registers _carapace spec subcommand
```

Call `spec.Register` directly on the root command (not inside `PreRun`).
Other commands can then use these completions via `bridge.ActionCarapaceBin`.

### Custom Spec Macros

Extend the spec with custom macros using `spec.AddMacro`. Register macros **before** calling `spec.Register`:

```go
// No-argument macro (N = No argument)
spec.AddMacro("tools.custom.Dirs", spec.MacroN(func() carapace.Action {
    return carapace.ActionDirectories()
}))

// Single-argument macro (I = Input)
spec.AddMacro("tools.custom.Filter", spec.MacroI(func(s string) carapace.Action {
    return carapace.ActionValues().Filter(s)
}))

// Variadic-argument macro (V = Variadic)
spec.AddMacro("tools.custom.Multi", spec.MacroV(func(s ...string) carapace.Action {
    return carapace.ActionValues(s...)
}))

spec.Register(rootCmd)
```

Macro names use a dotted prefix where the first segment is the executable providing the macro:
- `$carapace.tools.git.Refs` — macro provided by the `carapace` executable (fully qualified, always works)
- `$_.Refs`, equivalent to `$git.Refs` when the current executable is `git` (resolves to its own macros, avoid unless explicitly asked for)
- `$files` — core macros have no prefix (provided by carapace-spec itself)

Use actions directly (call the Go function) when they are internal, core, or from an already imported package. Use `spec.ActionMacro` only for loosely coupled access to macros exposed by other commands. The carapace MCP tool `list_macros` lists any exposed macros known to it.

### Exposing Actions via spec.ActionMacro

Actions from carapace-bin (and carapace-bridge) can be exposed through the spec system using `spec.ActionMacro`:

```go
spec.AddMacro("tools.git.Refs", spec.MacroN(spec.ActionMacro("$carapace.bridge.ActionCarapaceBin(git)")))
```

## Flag Configurations

Cobra, spf13/pflag, and carapace-pflag provide several flag configuration options that affect completion behavior.

### Optional Argument Flags (Optarg)

Setting `NoOptDefVal` to a non-empty string (typically `" "`) makes a flag's argument optional — the flag can be used with or without a value:

```go
cmd.Flags().StringP("optarg", "o", "", "optional argument")
cmd.Flag("optarg").NoOptDefVal = " " // flag now accepts `--optarg` alone (no value required)
```

When `NoOptDefVal` is set, carapace marks the flag as optarg in the spec (appending `?` to the definition). Completion still provides the values, but the flag is valid without one.

A non-space `NoOptDefVal` (like `" "`) indicates the argument is truly optional at the flag level. The default value `"defaultValue"` is used when the flag is present without an argument:

```go
cmd.Flag("persistentFlag").NoOptDefVal = "defaultValue" // `--persistentFlag` uses "defaultValue" when no arg given
```

### Optarg Delimiter

`OptargDelimiter` changes the character that separates a flag from its value (default `=`). This is commonly used with non-POSIX flags where `:` or `/` separates the flag from its argument:

```go
cmd.Flag("D").NoOptDefVal = " "
cmd.Flag("D").OptargDelimiter = ':' // `-D:key=value` instead of `-D=key=value`
```

Completion splits the current value at the delimiter to determine what is being completed.

### Nargs — Multi-Value Flags

`Nargs` controls how many values a flag consumes:
- `0` (default): single value
- Positive integer (e.g., `2`, `3`): exactly that many values
- `-1`: any number of values (variadic)

Nargs requires the flag type to be a slice type (`StringSlice`, `IntSlice`, etc.):

```go
cmd.Flags().StringSlice("option", nil, "Set key-value option")
cmd.Flag("option").Nargs = 2 // consumes exactly 2 values: --option key value

cmd.Flags().StringSlice("files", nil, "Files to process")
cmd.Flag("files").Nargs = -1 // consumes any remaining values: --files a b c...
```

For completion, use `ActionCallback` with `c.Parts` to provide different completions per position:

```go
carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
    "option": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
        switch len(c.Parts) {
        case 0:
            return carapace.ActionValues("key1", "key2") // first value: key
        case 1:
            return carapace.ActionValues("val1", "val2") // second value: value
        default:
            return carapace.ActionValues()
        }
    }),
    "files": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
        return carapace.ActionFiles().Invoke(c).Filter(c.Parts...).ToA() // filter already-provided values
    }),
})
```

### Hidden Flags

Setting `Hidden = true` on a flag hides it from help text and completion:

```go
cmd.Flags().String("deprecated", "", "deprecated flag")
cmd.Flag("deprecated").Hidden = true // hidden from help and completion
```

Hidden flags are excluded from completion suggestions but still function at runtime. Use this for deprecated or alias flags that should still work but not be advertised.

### Mutually Exclusive Flags

`MarkFlagsMutuallyExclusive` ensures that only one flag in a group can be used at a time:

```go
cmd.Flags().Bool("json", false, "JSON output")
cmd.Flags().Bool("form", false, "form output")
cmd.Flags().Bool("multipart", false, "multipart output")
cmd.MarkFlagsMutuallyExclusive("json", "form", "multipart")
```

Carapace filters out already-set mutually exclusive flags from completion suggestions at runtime. When one flag in the group is set, the others are no longer offered as completions.

## Non-POSIX Flag Handling (carapace-pflag)

[carapace-pflag](https://github.com/carapace-sh/carapace-pflag) is a fork of spf13/pflag that adds non-POSIX flag modes. These are primarily useful for external completers that model tools with non-standard flag syntax.

### Installation

```bash
go get github.com/carapace-sh/carapace-pflag
```

```go
import pflag "github.com/carapace-sh/carapace-pflag"
```

### Flag Modes

carapace-pflag adds a `Mode` field to flags with three modes:

| Mode | Constant | Behavior | Spec definition |
|------|----------|----------|----------------|
| Default | `Default` | Standard POSIX: `-s, --name` | `-s, --name` |
| ShorthandOnly | `ShorthandOnly` | Only accessible via shorthand: `-s` | `-s` |
| NameAsShorthand | `NameAsShorthand` | Name also registered as shorthand: `-s, -name` | `-s, -name` |

### N-Suffix Methods (NameAsShorthand)

The `*N` methods create flags where the name is also registered as a shorthand, enabling single-dash long flags (non-POSIX style). This is like `*P` but the name uses a single dash prefix:

```go
// Standard POSIX: -v, --verbose
cmd.Flags().BoolP("verbose", "v", false, "verbose output")

// Non-POSIX: -v, -verbose (both single-dash)
cmd.Flags().BoolN("verbose", "v", false, "verbose output")
```

Available `*N` methods for all types: `BoolN`, `CountN`, `IntN`, `Int8N`, `Int16N`, `Int32N`, `Int64N`, `UintN`, `Uint8N`, `Uint16N`, `Uint32N`, `Uint64N`, `Float32N`, `Float64N`, `StringN`, `DurationN`, `StringSliceN`, `StringArrayN`, `IntSliceN`, `Int32SliceN`, `Int64SliceN`, `Float32SliceN`, `Float64SliceN`, `BoolSliceN`, `DurationSliceN`, `UintSliceN`, `IPN`, `IPMaskN`, `IPNetN`, `IPSliceN`, `StringToIntN`, `StringToInt64N`, `StringToStringN`, `BytesHexN`, `BytesBase64N`, `TimeN`, `FuncN`, `BoolFuncN`, `TextVarN`, `VarN` — plus their `Var` variants (e.g., `IntVarN`, `StringVarN`).

### S-Suffix Methods (ShorthandOnly)

The `*S` methods create flags that are only accessible via their shorthand. The name is not registered as a long flag — only the single-letter form works:

```go
// Only -a works, --all is NOT available
cmd.Flags().BoolS("all", "a", false, "list all entries")

// Only -t works, --time is NOT available
cmd.Flags().StringS("time", "t", "", "use specific time format")
```

Available `*S` methods mirror the `*N` methods: `BoolS`, `CountS`, `IntS`, `StringS`, `StringSliceS`, `StringArrayS`, etc.

### Non-POSIX Example

```go
import pflag "github.com/carapace-sh/carapace-pflag"

func init() {
    carapace.Gen(rootCmd).Standalone()

    // BoolN: -v and -verbose (both single-dash, non-POSIX)
    rootCmd.Flags().BoolN("verbose", "v", false, "verbose output")
    // CountN: -c is non-POSIX counter
    rootCmd.Flags().CountN("count", "c", "count flag")
    // StringS: only -d works, --delim-colon is NOT a valid flag
    rootCmd.Flags().StringS("delim-colon", "d", "", "colon-delimited optarg")
    // StringN: -o and -overlapping both work (single-dash)
    rootCmd.Flags().StringN("overlapping", "o", "", "overlapping shorthand")

    // Optarg with custom delimiter
    rootCmd.Flag("delim-colon").NoOptDefVal = " "
    rootCmd.Flag("delim-colon").OptargDelimiter = ':'

    // Nargs
    rootCmd.Flags().StringSlice("nargs-any", []string{}, "any number of values")
    rootCmd.Flag("nargs-any").Nargs = -1
    rootCmd.Flags().StringSlice("nargs-two", []string{}, "exactly two values")
    rootCmd.Flag("nargs-two").Nargs = 2

    rootCmd.Flags().SetInterspersed(false)

    carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
        "delim-colon": carapace.ActionValues("d1", "d2", "d3"),
        "nargs-any": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
            return carapace.ActionValues("na1", "na2", "na3").Invoke(c).Filter(c.Parts...).ToA()
        }),
        "nargs-two": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
            switch len(c.Parts) {
            case 0:
                return carapace.ActionValues("nt1", "nt2", "nt3")
            case 1:
                return carapace.ActionValues("nt4", "nt5", "nt6")
            default:
                return carapace.ActionValues()
            }
        }),
        "overlapping": carapace.ActionValues("o1", "o2", "o3"),
    })
}
```

## Complete Example

```go
package main

import (
    "os"

    "github.com/carapace-sh/carapace"
    "github.com/carapace-sh/carapace/pkg/traverse"
    "github.com/spf13/cobra"
    "github.com/spf13/pflag"
)

func main() {
    rootCmd := &cobra.Command{
        Use:   "myapp",
        Short: "My application with carapace completions",
    }

    rootCmd.PersistentFlags().StringP("chdir", "C", "", "change directory before running")

    carapace.Gen(rootCmd)

    // PreRun: dynamically add subcommands (e.g., discover plugins)
    carapace.Gen(rootCmd).PreRun(func(cmd *cobra.Command, args []string) {
        // Add dynamic subcommands here if needed
    })

    // PreInvoke: change directory for all completions based on --chdir flag
    carapace.Gen(rootCmd).PreInvoke(func(cmd *cobra.Command, flag *pflag.Flag, action carapace.Action) carapace.Action {
        return action.ChdirF(traverse.Flag(cmd.Flag("chdir")))
    })

    // Flag completion
    carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
        "chdir": carapace.ActionDirectories(),
        "file":  carapace.ActionFiles(),
        "format": carapace.ActionValues("json", "yaml", "toml"),
    })

    // Positional completion
    carapace.Gen(rootCmd).PositionalCompletion(
        carapace.ActionDirectories(),
        carapace.ActionFiles(),
    )

    if err := rootCmd.Execute(); err != nil {
        os.Exit(1)
    }
}
```

## Reference Repositories

- [carapace](https://github.com/carapace-sh/carapace) — core library (see `example/` subdirectory)
- [carapace-spec](https://github.com/carapace-sh/carapace-spec) — spec generation and action registration
- [carapace-bridge](https://github.com/carapace-sh/carapace-bridge) — external completion bridging
- [carapace-pflag](https://github.com/carapace-sh/carapace-pflag) — non-POSIX flag handling
