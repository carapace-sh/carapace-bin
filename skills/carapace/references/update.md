# Incremental Completer Update Guide

Update an **existing** Go completer when a new version of the CLI tool adds, removes, or changes flags and subcommands. This is the most common real-world scenario — distinct from [references/convert.md](convert.md) which covers creating a completer from scratch.

## Overview

The workflow has three phases:

1. **Generate reference** — scrape the new spec and generate a cobra skeleton for comparison
2. **Identify changes** — diff the skeleton against the existing completer to find exactly what changed
3. **Apply surgically** — edit existing files to apply only the structural changes, preserving all custom completion logic

**The golden rule**: never bulk-copy codegen output over existing files. Codegen strips GroupIDs, completion actions, and completer-specific imports. Always apply changes surgically to existing files.

## Step 1: Scrape the New Spec

Generate a YAML spec from the new version of the CLI tool. See [references/scrape.md](scrape.md) for the full scraping workflow.

```bash
# If a scraper exists in the scrape project
docker build -t scrape-<tool> -f scrapers/<tool>/Dockerfile --build-arg VERSION=v1.2.3 scrapers/<tool>/
docker run --rm scrape-<tool> > /tmp/<tool>-spec.yaml
```

If no scraper exists and the tool uses cobra, you can also get a spec from the built binary:

```bash
carapace <tool> spec > /tmp/<tool>-spec.yaml
```

## Step 2: Generate Codegen Reference

Use the MCP `codegen` tool or the CLI equivalent to generate a cobra command skeleton from the spec:

```bash
# CLI
carapace --codegen /tmp/<tool>-spec.yaml

# MCP tool (returns list of generated filenames in /tmp/carapace-codegen-<hash>/)
mcp_carapace_codegen(path="/tmp/<tool>-spec.yaml")
```

The generated files contain:
- Cobra command structs (`Use`, `Short`, `GroupID`, `Aliases`, `Hidden`)
- Flag registrations (type, shorthand, description, default value)
- Group definitions
- Subcommand wiring (`AddCommand`)
- `Standalone()` calls

The generated files do **not** contain:
- `FlagCompletion`, `PositionalCompletion`, `PositionalAnyCompletion` — all completion actions are stripped
- Completer-specific imports (only `carapace` and `cobra`)
- Completer-specific `cmd/action/` packages

**This is a reference only.** The codegen output lives in `/tmp/carapace-codegen-<hash>/` and should never be copied directly into the completer directory.

## Step 3: Identify Structural Changes

Compare the codegen output to the existing completer to find exactly what changed. The categories of change are:

| Category | What to look for |
|----------|-----------------|
| New subcommands | Files in codegen that don't exist in the completer |
| Removed subcommands | Files in the completer that don't exist in codegen |
| New flags | Flag registrations in codegen not present in the completer |
| Removed flags | Flag registrations in the completer not present in codegen |
| Changed flag types | `String` → `StringSlice`, `Bool` → `BoolP`, etc. |
| Changed descriptions | One-line `Short` or flag description differs |
| New groups | `AddGroup` calls in codegen not present in the completer |

### Semi-automated diff script

```bash
CODEGEN_DIR=/tmp/carapace-codegen-XXXXXX
COMPLETER_DIR=completers/common/<tool>_completer/cmd

for f in "$CODEGEN_DIR"/*.go; do
    name=$(basename "$f")
    existing="$COMPLETER_DIR/$name"
    if [ ! -f "$existing" ]; then
        echo "NEW SUBCOMMAND: $name"
    else
        # Diff only the flag and structural sections (skip completion wiring)
        diff <(grep -E '^\t\S+Cmd\.|\.Flags\(\)|\.AddCommand|AddGroup|carapace\.Gen' "$f") \
             <(grep -E '^\t\S+Cmd\.|\.Flags\(\)|\.AddCommand|AddGroup|carapace\.Gen' "$existing") \
             && true
    fi
done

for f in "$COMPLETER_DIR"/*.go; do
    name=$(basename "$f")
    if [ ! -f "$CODEGEN_DIR/$name" ]; then
        echo "REMOVED SUBCOMMAND: $name"
    fi
done
```

### What to compare manually

The script above catches structural changes. Also check:

- **Flag descriptions** — codegen uses descriptions from the spec, which may differ from the completer's current descriptions. Run a broader diff and look for description changes.
- **Flag types** — a `String` → `StringSlice` change is a type change, not just a description change. Look for different `.Flags().*()` method calls.
- **Shorthand changes** — a flag gaining or losing a shorthand (e.g. `Bool` → `BoolP`).

## Step 4: Apply Changes Surgically

For each category of change, apply the minimal edit to the existing file.

### New flag

Insert the flag registration in the `init()` function, maintaining alphabetical order within the contiguous flag block. Also add the flag to the `FlagCompletion` map if it has a custom completion action.

```go
// Before (attach.go)
attachCmd.Flags().BoolP("stdin", "i", false, "Pass stdin to the container")

// After — new --detach-keys flag inserted alphabetically before --stdin
attachCmd.Flags().String("detach-keys", "", "Override the key sequence for detaching a container")
attachCmd.Flags().BoolP("stdin", "i", false, "Pass stdin to the container")
```

If the new flag has a known completion action, add it to the `FlagCompletion` map:

```go
carapace.Gen(attachCmd).FlagCompletion(carapace.ActionMap{
    "detach-keys": action.ActionDetachKeys(),
    "stdin":       carapace.ActionValues("true", "false"),
})
```

### Removed flag

Delete the flag registration line and its entry in the `FlagCompletion` map (if present). Removing a flag is rare but happens when upstream deprecates and removes an option.

### Changed flag type

Replace the flag registration call. Update the `FlagCompletion` entry if the type change affects completion (e.g. `String` → `StringSlice` may need `ActionValues` split on comma).

```go
// Before
waitCmd.Flags().String("for", "", "The condition to wait on")

// After — type changed from String to StringSlice
waitCmd.Flags().StringSlice("for", nil, "The condition to wait on")
```

### Changed description

Update only the description string. Do not touch the flag type, default value, or completion wiring.

```go
// Before
debugCmd.Flags().String("profile", "", "Legacy profiling mode")
// After
debugCmd.Flags().String("profile", "", "Profiling mode. Defaults to general")
```

### New subcommand

Create a new `.go` file following the existing completer's conventions. Use the codegen output as a starting point for the structure, then add completion wiring. The file must:

1. Declare the cobra command variable
2. Call `carapace.Gen(cmd).Standalone()` in `init()`
3. Define flags (alphabetically ordered)
4. Register with the parent command via `parentCmd.AddCommand(cmd)`
5. Add completion wiring (`FlagCompletion`, `PositionalCompletion`, etc.)

```go
package cmd

import (
    "github.com/carapace-sh/carapace"
    "github.com/spf13/cobra"
)

var kubercCmd = &cobra.Command{
    Use:   "kuberc",
    Short: "Manage kuberc configuration",
    GroupID: "management",
}

func init() {
    carapace.Gen(kubercCmd).Standalone()

    rootCmd.AddCommand(kubercCmd)
}
```

Check the codegen output for the correct `GroupID` value — the skeleton will have it if the spec defines groups.

### Removed subcommand

Delete the `.go` file and remove any references to the command variable from the parent's `AddCommand` call. Also remove any `cmd/action/` files that were specific to this subcommand.

### Subcommand promoted (moved in the tree)

When a subcommand moves (e.g. `alpha kuberc` → `kuberc`), this is both a removal and an addition:

1. Create the new file at the new location in the command tree
2. Copy over completion wiring from the old file, adjusting the parent command reference
3. Delete the old file
4. Update the old parent's `AddCommand` (remove the old reference) and the new parent's `AddCommand` (add the new reference)

### New flags on promoted subcommands

When a promoted subcommand also gains new flags, combine the operations: create the new file with the old completion wiring plus the new flags.

## Step 5: Update Man Docs

Use `carapace-man update` to refresh the spec structure in the man YAML files while preserving existing documentation:

```bash
# One-step update: generate spec, split with doc preservation, show diff
carapace-man update /tmp/<tool>-spec.yaml man/cmd/<tool>
```

This preserves all existing `documentation.command` and `documentation.flag` entries. It creates new files for new subcommands but does **not** delete files for removed subcommands.

After running `update`:

1. **Delete removed subcommand files** — manually remove YAML files for subcommands that no longer exist
2. **Add `documentation.command` entries** — for new subcommands and any existing ones missing docs. See the man-docs skill for standards.
3. **Review `documentation.flag` entries** — new flags may need flag docs if their short description is insufficient

## Step 6: Build, Lint, Test

```bash
# Build with all platform completers
go build -tags force_all -o /tmp/carapace-test ./cmd/carapace

# Lint flag ordering
go run ./cmd/carapace-lint completers/<category>/<tool>_completer/cmd/*.go

# Auto-fix ordering if needed
go run ./cmd/carapace-lint --fix-flags-order completers/<category>/<tool>_completer/cmd/*.go

# Test completion output
/tmp/carapace-test <tool> bash ''
/tmp/carapace-test <tool> bash '--flag '

# Run test suite
go test -v ./cmd/...
```

**Important**: The MCP `complete_command` tool uses the installed binary, not the freshly built one. Always test with the built binary directly.

## Anti-Patterns

| Anti-pattern | Why it fails |
|-------------|-------------|
| **Bulk-copying codegen output over existing files** | Strips all GroupIDs, completion actions, completer-specific imports, and `cmd/action/` references. Restoring them manually across 80+ files is extremely fragile. |
| **Regex-based merge scripts** | Go code has too much structural variation for regex merges to handle reliably — brace nesting, import grouping, and init() ordering all break easily. |
| **Using MCP `complete_command` before rebuilding** | The MCP server runs the installed binary, not your local build. Completion output will be stale until you build and install. |
| **Editing without reading the exact context first** | The `edit` tool requires exact string matching. Whitespace, indentation, and blank lines must match precisely. Always `view` the target lines before editing. |

## Relationship to Other Skills

| Skill | When to use instead of this skill |
|-------|----------------------------------|
| **references/convert.md** | Creating a completer from scratch (no existing completer exists) |
| **references/scrape.md** | Generating the YAML spec — the step before this workflow |
| **references/spec.md** | Writing YAML user specs (the source format) |
| **references/macro.md** | Looking up macro signatures and formatting macro arguments |
| **references/action.md** | Creating/modifying Go actions that become macros |
| **references/man.md** | Man page documentation format and UID resolution |
| **man-docs skill** | Documentation standards, `carapace-man update` workflow, `documentation.command`/`documentation.flag` guidelines |
