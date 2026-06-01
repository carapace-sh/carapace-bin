# carapace-bin

A collection of shell completers powered by [carapace](https://github.com/carapace-sh/carapace).

## Project Structure

```
cmd/carapace/                  # Main entry point
completers/                    # Individual completers
  common/                      # Completers that need networking
    gh_completer/
  unix/                        # Unix-only completers
  windows/                     # Windows-only completers
pkg/
  actions/                     # Shared actions (reusable across completers)
    actions.go                 # Package declaration
    actions_generated.go        # Macro registration map (generated)
    net/                       # Networking actions
    os/                        # OS-level actions
    fs/                        # Filesystem actions
    color/                     # Color actions
    tools/                     # Tool-specific actions (one subpackage per tool)
      git/
      gh/
      docker/
      helm/
      ...
  styles/                      # Style definitions
  conditions/                  # Condition helpers
```

## Shared Actions in `pkg/actions/`

Any action useful to external commands or user specs should be a fully documented public exposed action here.

### Package Layout

Each tool gets a subpackage under `pkg/actions/tools/`:

```
pkg/actions/tools/git/
  git.go          # Uid() helper, package-level shared types
  branch.go       # ActionLocalBranches, ActionRemoteBranches
  change.go       # ChangeOpts, ActionChanges, ActionRefChanges
```

Domain categories (not tool-specific) live directly under `pkg/actions/`:

```
pkg/actions/net/     # Ports, hosts, SSH
pkg/actions/os/      # Users, groups, locales
pkg/actions/fs/      # Filesystem types, extensions, archives
pkg/actions/color/   # Color palettes
pkg/actions/number/  # Ranges, versions
pkg/actions/time/    # Time formats
```

### Naming Convention

- **Package name**: matches the tool (`package git`, `package helm`, `package net`)
- **File naming**: group by entity (`branch.go`, `issue.go`, `release.go`)

### Exposed Actions (Macros)

Actions in `pkg/actions/` are exposed as macros via `pkg/actions/actions_generated.go`. The macro type is determined by the function signature. In this project, macros use the `tools.<tool>.<Action>` naming pattern (e.g. `tools.git.Changes`, `tools.ffmpeg.Codecs`).

Registration at startup in `cmd/carapace/cmd/root.go`:

```go
for m, f := range actions.Macros {
	spec.AddMacro(m, f)
}
spec.Register(rootCmd)
```

### Uid / QueryF

In this project, some tool packages define a package-level `Uid()` helper when multiple actions share common context (e.g. git needs `GIT_DIR` and `GIT_WORK_TREE` resolved for all its actions). This avoids repeating the same URL construction logic in each action.

```go
// pkg/actions/tools/git/git.go
func Uid(host string, opts ...string) func(s string, uc uid.Context) (*url.URL, error) {
	return func(s string, uc uid.Context) (*url.URL, error) {
		// resolves GIT_DIR, GIT_WORK_TREE from context
		// constructs URL: git://<host>/<value>?<opts>
	}
}
```

```go
carapace.ActionExecCommand(...)(...).
	Tag("local branches").
	UidF(Uid("local-branch")).    // e.g. git://local-branch/main
	QueryF(Uid("local-branches")) // e.g. git://local-branches
```

### Styles

Project-specific styles live in `pkg/styles/`:

```go
carapace.ActionValuesDescribed(vals...).Style(styles.Git.Branch)
carapace.ActionStyledValuesDescribed(vals...) // vals: value, description, style triplets
```

## Completer-Specific Actions in `cmd/action/`

Complex completers that need to adapt shared actions based on cobra flag state use a `cmd/action/` subpackage:

```go
package action

func ActionPackages(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return npm.ActionPackageSearch(cmd.Flag("registry").Value.String())
	})
}
```

Key differences from shared actions:

| Aspect | `pkg/actions/` (shared) | `cmd/action/` (completer-specific) |
|--------|------------------------|-------------------------------------|
| Package | domain name (e.g. `git`, `net`) | `action` |
| Doc comments | Full format with examples | Typically none |
| Cobra awareness | No `*cobra.Command` params | Often takes `*cobra.Command` |
| Tag/Uid | `.Tag()` and `.UidF()`/`.QueryF()` | `.Tag()` but typically no Uid |
| Reusability | Importable by any completer | Only within its own completer |

## Generated Code

`pkg/actions/actions_generated.go` is generated and contains the `Macros` map. After adding or modifying a public action in `pkg/actions/`, this file must be regenerated.

The generation picks up:
- Function name → macro name (e.g. `ActionChanges` → `tools.git.Changes`)
- Doc comment first line → description
- Doc comment code footer → example
- Function signature → macro type (MacroN/MacroI/MacroV)
- Go package path → function identifier
