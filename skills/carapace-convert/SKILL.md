---
name: carapace-convert
description: >
  Use when user wants to convert a carapace YAML user spec to a native Go completer for carapace-bin.
  Covers the codegen tool, manual completion wiring, macro-to-Go mapping, and modifier translation.
  Triggers on: "convert spec", "spec to Go", "yaml to go completer", "carapace convert", "codegen",
  "generate Go completer", "spec to completer", "convert completion".
user-invocable: true
---

# Carapace Spec-to-Go Conversion Guide

Convert a carapace YAML user spec into a native Go completer that can be built into [carapace-bin](https://github.com/carapace-sh/carapace-bin).

## Overview

Conversion has two phases:

1. **Codegen** — generate the cobra command skeleton (flags, subcommands, groups) from the YAML spec
2. **Completion wiring** — manually translate the `completion` section into Go action code

The codegen tool does **not** produce completion code — it only generates the command structure. The `completion` section of the spec must be translated by hand using the macro and modifier mappings below.

## Step 1: Codegen

Use the carapace MCP `codegen` tool or the CLI equivalent:

```bash
carapace --codegen path/to/spec.yaml
```

This produces one or more `.go` files in a temp directory containing:
- Package declaration (`package cmd`)
- Cobra command structs (`Use`, `Short`, `GroupID`, `Aliases`, `Hidden`)
- Flag registrations (`String`, `Bool`, `CountP`, etc.)
- `Execute()` function (root command only)
- Subcommand wiring (`AddCommand`)

Copy the generated files into the carapace-bin completer directory (e.g. `cmd/traverse/`).

## Step 2: Add Standandalone Mode

In the `init()` function, add `carapace.Gen(rootCmd).Standalone()` so the completer runs as a standalone completion provider:

```go
func init() {
	carapace.Gen(rootCmd).Standalone()
	// ... flag definitions ...
}
```

## Step 3: Translate Completion Actions

Add `FlagCompletion`, `PositionalCompletion`, and/or `PositionalAnyCompletion` calls to `init()`. The mappings below translate YAML spec macros and modifiers to their Go equivalents.

### Core Action Macros

| YAML Macro | Go Equivalent |
|-----------|---------------|
| `$files` | `carapace.ActionFiles()` |
| `$files([.ext])` | `carapace.ActionFiles(".ext")` |
| `$directories` | `carapace.ActionDirectories()` |
| `$executables` | `carapace.ActionExecutables()` |
| `$message(text)` | `carapace.ActionMessage("text")` |
| `$values` | `carapace.ActionValues(...)` |
| `$valuesdescribed` | `carapace.ActionValuesDescribed(...)` |

### External Action Macros

External macros (prefixed with `carapace.`) reference actions from carapace-bin packages. Use the `list_macros` MCP tool or `carapace --macro` to find the Go import path and function name.

| YAML Macro | Go Equivalent |
|-----------|---------------|
| `$carapace.tools.git.Refs({tags: true})` | `git.ActionRefs(g.RefsOpts{Tags: true})` |
| `$carapace.tools.git.Tags` | `git.ActionTags()` |
| `$carapace.os.Users` | `os.ActionUsers()` |

Import the package from the `reference` field in the macro listing (e.g. `github.com/carapace-sh/carapace-bin/pkg/actions/tools/git`).

### Modifier Macros

Modifiers are applied to the preceding action using the `|||` delimiter in specs (generic modifier) or by chaining method calls in Go.

| YAML Modifier | Go Equivalent | Notes |
|--------------|---------------|-------|
| `$chdir(path)` | `.Chdir("path")` | Change to literal directory |
| `$chdir($traverse_macro)` | `.ChdirF(traverse.Xxx)` | See traverse mapping below |
| `$filter(pattern...)` | `.Filter(pattern...)` | |
| `$filterargs` | `.FilterArgs()` | |
| `$list(separator)` | `.List(separator)` | |
| `$multiparts(chars...)` | `.MultiParts(chars...)` | |
| `$nospace(chars)` | `.NoSpace(chars...)` | |
| `$prefix(prefix)` | `.Prefix(prefix)` | |
| `$retain(pattern...)` | `.Retain(pattern...)` | |
| `$shift(n)` | `.Shift(n)` | |
| `$split` | `.Split()` | |
| `$splitp` | `.SplitP()` | |
| `$suffix(suffix)` | `.Suffix(suffix)` | |
| `$suppress(pattern)` | `.Suppress(pattern)` | |
| `$style(style)` | `.Style(style)` | |
| `$tag(tag)` | `.Tag(tag)` | |
| `$uniquelist(separator)` | `.UniqueList(separator)` | |
| `$usage(text)` | `.Usage(text)` | |

### Traverse Macros

Inside `$chdir()`, the argument can be a traverse macro that resolves a directory at runtime. These map directly to functions in `github.com/carapace-sh/carapace/pkg/traverse`.

| YAML Macro | Go Equivalent | Resolves to |
|-----------|---------------|-------------|
| `$chdir($gitdir)` | `.ChdirF(traverse.GitDir)` | `GIT_DIR` env or `.git` dir |
| `$chdir($gitworktree)` | `.ChdirF(traverse.GitWorkTree)` | `GIT_WORK_TREE` env or parent of `.git` |
| `$chdir($parent([name1, name2]))` | `.ChdirF(traverse.Parent("name1", "name2"))` | Walks upward to find dir containing named files |
| `$chdir($tempdir)` | `.ChdirF(traverse.TempDir)` | `os.TempDir()` |
| `$chdir($usercachedir)` | `.ChdirF(traverse.UserCacheDir)` | `os.UserCacheDir()` |
| `$chdir($userconfigdir)` | `.ChdirF(traverse.UserConfigDir)` | `os.UserConfigDir()` |
| `$chdir($userhomedir)` | `.ChdirF(traverse.UserHomeDir)` | `os.UserHomeDir()` |
| `$chdir($xdgcachehome)` | `.ChdirF(traverse.XdgCacheHome)` | `$XDG_CACHE_HOME` or fallback to `UserCacheDir` |
| `$chdir($xdgconfighome)` | `.ChdirF(traverse.XdgConfigHome)` | `$XDG_CONFIG_HOME` or fallback to `UserConfigDir` |
| `$chdir($nixprofile)` | `.ChdirF(traverse.NixProfile)` | Nix profile directory |

### Flag Value References

In YAML specs, flag values are available as `${C_FLAG_<FLAGNAME>}` environment variables (uppercase, set via envsubst). In Go, use `traverse.Flag(cmd.Flag("flagname"))` to resolve the flag value at invoke time:

| YAML | Go |
|------|----|
| `$chdir(${C_FLAG_MYDIR})` | `.ChdirF(traverse.Flag(cmd.Flag("mydir")))` |
| `${C_FLAG_FORMAT}` in macro args | `cmd.Flag("format").Value.String()` in `ActionCallback` |

Prefer `ChdirF(traverse.Flag(...))` over `Chdir(flag.Value.String())` — it resolves the value lazily at invoke time.

### Generic vs Specific Modifier Application

In YAML, the `|||` delimiter controls modifier scope:

| YAML Syntax | Scope | Go Equivalent |
|-------------|-------|---------------|
| `["$files", "$chdir(/tmp)"]` | Generic — applies to **all** preceding actions | `carapace.ActionFiles().Chdir("/tmp")` |
| `["$files \|\|\| $chdir(/tmp)"]` | Specific — applies to **preceding value only** | `carapace.ActionFiles().Chdir("/tmp")` |

In Go, both translate to the same chained method calls since Go actions are composed by chaining.

### Static Values

Static completion values in specs translate to `ActionValues` or `ActionValuesDescribed`:

| YAML | Go |
|------|----|
| `["true", "false"]` | `carapace.ActionValues("true", "false")` |
| `["value\tdescription"]` | `carapace.ActionValuesDescribed("value", "description")` |
| `["value\tdesc\tblue"]` | `carapace.ActionStyledValuesDescribed("value", "desc", "blue")` |

### Subcommand Completions

For subcommands, create a separate `init()` function in the subcommand file:

```go
func init() {
	carapace.Gen(subCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "yaml"),
	})

	carapace.Gen(subCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)

	carapace.Gen(subCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
```

## Complete Example

Given this YAML spec:

```yaml
name: traverse
flags:
  --gitdir=: .git folder
  --gitworktree=: root of the working directory for a non-bare repository
  --parent=: first parent directory containing any of the given names/directories
  --tempdir=: default directory to use for temporary files
  --usercachedir=: root directory to use for user-specific cached data
  --userconfigdir=: default root directory to use for user-specific configuration data
  --userhomedir=: current user's home directory
  --xdgcachehome=: cache directory (fallback to UserCacheDir)
  --xdgconfighome=: home directory (fallback to UserConfigDir)
completion:
  flag:
    gitdir: ["$files", "$chdir($gitdir)"]
    gitworktree: ["$files", "$chdir($gitworktree)"]
    parent: ["$files", "$chdir($parent([go.mod, go.sum]))"]
    tempdir: ["$files", "$chdir($tempdir)"]
    usercachedir: ["$files", "$chdir($usercachedir)"]
    userconfigdir: ["$files", "$chdir($userconfigdir)"]
    userhomedir: ["$files", "$chdir($userhomedir)"]
    xdgcachehome: ["$files", "$chdir($xdgcachehome)"]
    xdgconfighome: ["$files", "$chdir($xdgconfighome)"]
```

The converted Go completer:

```go
package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/traverse"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "traverse",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("gitdir", "", ".git folder")
	rootCmd.Flags().String("gitworktree", "", "root of the working directory for a non-bare repository")
	rootCmd.Flags().String("parent", "", "first parent directory containing any of the given names/directories")
	rootCmd.Flags().String("tempdir", "", "default directory to use for temporary files")
	rootCmd.Flags().String("usercachedir", "", "root directory to use for user-specific cached data")
	rootCmd.Flags().String("userconfigdir", "", "default root directory to use for user-specific configuration data")
	rootCmd.Flags().String("userhomedir", "", "current user's home directory")
	rootCmd.Flags().String("xdgcachehome", "", "cache directory (fallback to UserCacheDir)")
	rootCmd.Flags().String("xdgconfighome", "", "home directory (fallback to UserConfigDir)")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"gitdir":       carapace.ActionFiles().ChdirF(traverse.GitDir),
		"gitworktree":  carapace.ActionFiles().ChdirF(traverse.GitWorkTree),
		"parent":       carapace.ActionFiles().ChdirF(traverse.Parent("go.mod", "go.sum")),
		"tempdir":      carapace.ActionFiles().ChdirF(traverse.TempDir),
		"usercachedir": carapace.ActionFiles().ChdirF(traverse.UserCacheDir),
		"userconfigdir": carapace.ActionFiles().ChdirF(traverse.UserConfigDir),
		"userhomedir":  carapace.ActionFiles().ChdirF(traverse.UserHomeDir),
		"xdgcachehome": carapace.ActionFiles().ChdirF(traverse.XdgCacheHome),
		"xdgconfighome": carapace.ActionFiles().ChdirF(traverse.XdgConfigHome),
	})
}
```

## Relationship to Other Skills

| Skill | When to use instead of this skill |
|-------|----------------------------------|
| **carapace-spec** | Writing YAML user specs — the *source* format for conversion |
| **carapace-macro** | Looking up macro signatures and formatting macro arguments in YAML |
| **carapace-action** | Creating/modifying Go actions that become macros in the converted completer |
| **carapace-scrape** | Generating YAML specs from CLI source code — typically the step *before* conversion |
| **carapace-integrate** | Integrating carapace into an existing cobra app — use when the app already exists in Go |
| **carapace-mcp** | Using the `codegen` and `list_macros` MCP tools during conversion |
