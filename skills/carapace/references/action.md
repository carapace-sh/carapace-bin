# Carapace Custom Actions Guide

Create and modify shell completion actions in [carapace](https://carapace.sh) using Go.

## Action Types

### Public Actions

Exported functions prefixed with `Action`. Reusable across commands, documented, and eligible for macro exposure.

```go
// ActionLocalBranches completes local branches
//
//	master (last commit msg)
//	another (last commit msg)
func ActionLocalBranches() carapace.Action {
	return carapace.ActionExecCommand("git", "branch", "--format", "%(refname:short)\n%(subject)")(
		func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			return carapace.ActionValuesDescribed(lines[:len(lines)-1]...)
		},
	).Tag("local branches")
}
```

### Private Actions

Unexported functions prefixed with `action`. Internal helpers called by public actions. No doc comments.

```go
func actionCodecs(opts CodecOpts, filter func(s string) bool) carapace.Action {
	return carapace.ActionExecCommand("ffmpeg", "-hide_banner", "-codecs")(
		func(output []byte) carapace.Action { /* ... */ },
	).Tag("codecs")
}
```

Private actions are typically:
- Shared execution/parsing logic extracted from multiple public actions
- Platform-specific implementations (e.g. `actionUsedPortsLsof` vs `actionUsedPortsNetstat`)
- Command execution wrappers with shared env/config

### Anonymous Actions

Defined inline in `ActionMap`, `PositionalCompletion`, or `PositionalAnyCompletion`. For simple, one-off completions.

```go
carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
	"sort": carapace.ActionValues("name", "none", "dirname"),
	"zoom": carapace.ActionValues("max", "fill"),
})

carapace.Gen(rootCmd).PositionalCompletion(
	carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if rootCmd.Flag("reference").Changed {
			return carapace.ActionFiles()
		}
		return fs.ActionFileModes()
	}),
)
```

Use anonymous actions when:
- The completion is a simple static list
- The logic is tightly coupled to a specific command's flag state
- The action will never be reused elsewhere

### Exposed Actions

Public actions registered as macros so they are available in YAML user specs. Must match `spec.MacroN`/`spec.MacroI`/`spec.MacroV` function signatures:

| Action signature | Macro type | Spec usage |
|-----------------|------------|------------|
| `func ActionFoo()` | `spec.MacroN` | `$_.Foo` or `$exe.Foo` |
| `func ActionFoo(opts FooOpts)` | `spec.MacroI` | `$_.Foo({Field: value})` |
| `func ActionFoo(arg string)` | `spec.MacroI` | `$_.Foo(arg)` |
| `func ActionFoo(args ...string)` | `spec.MacroV` | `$_.Foo(["a", "b"])` |

Register before `spec.Register`:

```go
spec.AddMacro("Refs", spec.MacroN(ActionRefs))
spec.AddMacro("Branches", spec.MacroI(ActionBranches))
spec.AddMacro("Sites", spec.MacroV(ActionSites))
spec.Register(rootCmd)
```

Any action useful to external commands or user specs should be a fully documented public exposed action.

## Naming

- **Prefer plural** for collections: `ActionUsers`, `ActionLocalBranches`, `ActionRepositories`, `ActionContainers`
- **Singular** for single-value completions: `ActionCurrentBranch`
- **Public actions**: `Action` prefix + PascalCase — `ActionChanges`, `ActionCodecs`
- **Private actions**: `action` prefix + PascalCase — `actionCodecs`, `actionUsedPortsLsof`
- **Opts structs**: singular prefix matching the action name + `Opts` suffix

| Action | Opts struct |
|--------|------------|
| `ActionChanges` | `ChangeOpts` |
| `ActionCodecs` | `CodecOpts` |
| `ActionReleases` | `ReleaseOpts` |
| `ActionContainers` | `ContainerOpts` |

## Opts Parameter

Since exposed actions must match the macro pattern, complex arguments are passed as a struct rather than multiple parameters.

### Opts with Default()

When an Opts struct is used with `spec.MacroI`, the macro system calls `Default()` when invoked without arguments. Implement `Default()` for boolean fields where the zero value would produce empty results:

```go
type CodecOpts struct {
	Attachment bool
	Audio      bool
	Data       bool
	Subtitle   bool
	Video      bool
}

func (o CodecOpts) Default() CodecOpts {
	// Enable all types by default so macro without args shows everything
	o.Attachment = true
	o.Audio = true
	o.Data = true
	o.Subtitle = true
	o.Video = true
	return o
}
```

```go
type ChangeOpts struct {
	Ignored  bool
	Staged   bool
	Unstaged bool
}

func (o ChangeOpts) Default() ChangeOpts {
	// Show staged and unstaged by default, but not ignored
	o.Staged = true
	o.Unstaged = true
	return o
}
```

### Opts without Default()

When all zero-values are sensible defaults (e.g. empty strings mean "no filter"), `Default()` is unnecessary:

```go
type ReleaseOpts struct {
	KubeContext string
	Namespace   string
}
// No Default() needed — empty strings mean "use default context/namespace"
```

### Opts Field Documentation

Document each member with field comments or group comments:

```go
type UnitFileOpts struct {
	// User restricts to user units
	User bool

	// type filters unit types
	Automount bool
	Mount     bool
	Service   bool

	// state filters unit states
	Enabled  bool
	Disabled bool
	Static   bool
}
```

### Simple String Arguments

For actions taking a single simple argument, use a plain parameter instead of an Opts struct:

```go
// ActionRemoteBranches completes remote branches
//
//	origin/master (last commit msg)
//	upstream/another (last commit msg)
func ActionRemoteBranches(remote string) carapace.Action { ... }
```

## Documentation

Every public action should have a doc comment with:

### 1. Description Line (required)

A single concise line. This appears in macro completions — keep it short:

```go
// ActionKnownPorts completes commonly used ports (opiniated)
```

### 2. Additional Explanation (optional)

Separated from the description by a blank `//` line. Explains behavior, defaults, or how arguments affect the action:

```go
// ActionRefDiffs completes changes between refs
// Accepts up to two refs
// 0: compare current workspace to HEAD
// 1: compare current workspace to given ref
// 2: compare first ref to second ref
```

### 3. Example Completions (required for exposed actions)

Two example completions indented with tab. Show the **display value** (without prefixes/suffixes), not the inserted value. If the action uses `ActionValuesDescribed`, append descriptions in round brackets:

```go
// ActionKnownPorts completes commonly used ports (opiniated)
//
//	80 (http)
//	443 (https)
func ActionKnownPorts() carapace.Action { ... }

// ActionLocalBranches completes local branches
//
//	master (last commit msg)
//	another (last commit msg)
func ActionLocalBranches() carapace.Action { ... }

// ActionDryRunModes completes dry run modes
//
//	none
//	server
func ActionDryRunModes() carapace.Action { ... }
```

## Common Action Patterns

### ActionValues — Static Values

```go
func ActionDryRunModes() carapace.Action {
	return carapace.ActionValues("none", "server", "client")
}
```

### ActionValuesDescribed — Values with Descriptions

```go
func ActionKnownPorts() carapace.Action {
	return carapace.ActionValuesDescribed(
		"21", "ftp",
		"22", "ssh",
		"80", "http",
		"443", "https",
	).Tag("known ports")
}
```

### ActionCallback — Dynamic Completions

Always wrap runtime state access in `ActionCallback` to prevent execution at init time:

```go
func ActionModels() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if f := cmd.Flag("repo"); f.Changed {
			return carapace.ActionValues(f.Value.String())
		}
		return carapace.ActionMessage("--repo is required")
	})
}
```

### ActionExecCommand — Execute External Command

```go
func ActionLocalBranches() carapace.Action {
	return carapace.ActionExecCommand("git", "branch", "--format", "%(refname:short)\n%(subject)")(
		func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			return carapace.ActionValuesDescribed(lines[:len(lines)-1]...)
		},
	)
}
```

### ActionExecCommandE — With Error Handling

```go
carapace.ActionExecCommandE("git", "remote")(func(output []byte, err error) carapace.Action {
	if err != nil {
		return carapace.ActionMessage(err.Error())
	}
	return carapace.ActionValues(strings.Split(string(output), "\n")...)
})
```

## Modifiers

### Tag

Group completion candidates with a descriptive tag. Use a plural noun matching the action:

```go
.Tag("local branches")
.Tag("known ports")
.Tag("codecs")
.Tag("units")
```

### Style

Apply visual differentiation:

```go
// Static style
carapace.ActionValuesDescribed(vals...).Style(styles.Git.Branch)

// Dynamic style based on value
carapace.ActionValues(vals...).StyleF(style.ForPathExt)

// Per-value style with ActionStyledValuesDescribed
carapace.ActionStyledValuesDescribed(
	"Up", "running", style.Carapace.KeywordPositive,
	"Exited", "stopped", style.Carapace.KeywordNegative,
)
```

### Suffix / Prefix

```go
carapace.ActionValues("option").Suffix("=")
carapace.ActionValues(":").Suffix("/").NoSpace('/')
carapace.ActionLocalChannels().Prefix("<").Suffix(">")
```

**Apply Suffix/Prefix/NoSpace at the call site, not in the action definition.** Reusable actions should return pure value lists without formatting modifiers. The call site knows the context — what delimiter follows, whether a space is needed — so it should chain the appropriate modifiers. This keeps actions composable: the same `ActionUsers()` can be used with `.Suffix(":")` in a multi-part `user:group` context, or without suffix in a standalone context.

```go
// Good — action returns pure values, call site adds suffix
return ActionUsers().Suffix(":")
ActionContainers().Suffix(":")

// Bad — suffix baked into definition limits reuse
func ActionUsers() carapace.Action {
    return carapace.ActionExecCommand(...)(...).Suffix(":") // don't do this
}
```

The exception is actions that are inherently context-specific (e.g. inline `ActionValues` inside `ActionMultiParts` where the suffix is part of the structure itself).

### Uid

Unique identifiers providing context for completion values:

- **Uid** — creates a static identifier from path segments
- **UidF** — creates a dynamic identifier per completion value (e.g. `git://local-branch/main` identifies a specific branch). Use it when the display value alone is insufficient, or when you need to embed additional context as query parameters
- **QueryF** — identifies what kind of completion is being requested (e.g. `git://local-branch` indicates the current position seeks local git branches). This enables result updates with additional queries later

```go
// Static UID
carapace.ActionValues("value").Uid("git", "branch")

// Dynamic UID per value
carapace.ActionValuesDescribed(vals...).UidF(func(s string, uc uid.Context) (*url.URL, error) {
	return Uid("local-branch")(s, uc)  // e.g. git://local-branch/main
})

// Query identifying the completion kind
.QueryF(Uid("local-branch"))  // e.g. git://local-branch
```

**Set Uid before Suffix/Prefix/NoSpace.** The UID identifies the completion values — modifiers like `.Suffix("=")` or `.Prefix("<")` change the display/insertion format but should not affect the identity. Chaining `.Suffix()` before `.Uid()` can leak suffix characters into the UID's value context.

```go
// Good — Uid on the base action, suffix after
carapace.ActionValuesDescribed(...).Uid("git", "object-filter").Suffix("=")

// Bad — suffix before Uid leaks "=" into the UID
carapace.ActionValuesDescribed(...).Suffix("=").Uid("git", "object-filter")
```

### Shift

Removes the first `n` positional arguments from `Context.Args`, shifting the view of what the action considers "already entered". This changes `len(c.Args)`, which determines which positional completion slot fires.

```go
// Bridge to another completer: shift off the routing arg so the bridge sees the "real" positionals
bridge.ActionCarapaceBin("tmux").Shift(1)

// Recurse into a subcommand: shift off the matched subcommand name
return ActionCommands(subCommand).Shift(1)
```

`Shift` modifies `Context.Args` in its callback. Since the **outermost** modifier fires first at invocation time, `Shift`'s position in the chain determines which `Context.Args` it sees:

```go
// Shift is innermost: FilterArgs fires first (captures original Args), then Shift shifts
// FilterArgs filters with original Args — Shift's c.Args modification is local to a's Invoke
a.Shift(1).FilterArgs()

// Shift is outermost: Shift fires first (captures original Args), then FilterArgs filters with shifted Args
// FilterArgs filters with shifted Args
a.FilterArgs().Shift(1)
```

| `n` value | Behavior |
|-----------|----------|
| `n < 0` | Returns `ActionMessage("invalid argument [ActionShift]: %v", n)` |
| `len(c.Args) < n` | `c.Args` becomes empty `[]string{}` |
| otherwise | `c.Args = c.Args[n:]` — slices off the first n elements |

### Suppress / Filter

```go
carapace.ActionValues("main", "develop").Suppress("main")
ActionAncestors(ctx.AttachedRevset).Filter("..", "::")
```

### Chaining Order

Modifiers are composed as nested callbacks — the **outermost** modifier fires **first**, then execution bubbles inward. Order matters when modifiers affect the same data:

| Chain | Effect |
|-------|--------|
| `a.Filter("x").Prefix("p/")` | Prefix fires first (outermost) → strips "p/" from `c.Value`, adds "p/" to results → Filter fires last (innermost) → removes "x" from original values → result: ["pa","pb"] |
| `a.Prefix("p/").Filter("x")` | Filter fires first (outermost) → invokes Prefix → Prefix adds "p/" to rawValues → Filter runs on now-prefixed values → removes "p/x" from results → result: ["pa","pb"] |
| `a.Filter("x").Suppress("err")` | Order doesn't matter (independent: values vs messages) |
| `a.Tag("t").Style("s")` | Order doesn't matter (independent fields) |

## Caching

Use `.Cache(timeout, keys...)` on actions that perform expensive operations (network calls, external commands). Cache is disk-based and shared across shell sessions.

### When to Add Cache Keys

Cache keys are **implicit** — only add them when the action produces **different results based on arguments or context**. Actions with no parameters or always returning the same results need no keys.

| Scenario | Keys needed | Example |
|----------|-------------|---------|
| No parameters, same results always | None | `ActionLocalBranches()` → `.Cache(24*time.Hour)` |
| Parameter changes results | Parameter as key | `ActionIssues(repo)` → `.Cache(1*time.Hour, key.String(repo))` |
| Opts change results | Opts cache key method | `ActionLabels(opts)` → `.Cache(24*time.Hour, opts.cacheKey())` |
| Results depend on a file | File stats key | `ActionTarContents(file)` → `.Cache(-1, key.FileStats(file))` |

### Timeout Values

| Timeout | When to use |
|--------|-------------|
| `5 * time.Minute` | Rapidly changing data (test servers) |
| `1 * time.Hour` | Moderately stable data (search results, extensions) |
| `24 * time.Hour` | Rarely changing data (labels, config keys, remote plugins) |
| `-1` | Never expire by time, only invalidate on key change (file-based sources) |

### Cache Key Helpers

Import `github.com/carapace-sh/carapace/pkg/cache/key`:

| Key function | Description |
|-------------|-------------|
| `key.String(s ...string)` | Simple string-based key from one or more strings |
| `key.FileStats(file)` | Key based on file path + size + modification time |
| `key.FileChecksum(file)` | Key based on file content SHA1 hash |
| `func() (string, error)` | Custom inline key function |

### Cache Key Method on Opts

When an Opts struct affects results, add a `cacheKey()` method returning `key.Key`:

```go
func (o RepoOpts) cacheKey() key.Key { return key.String(o.Host, o.Owner, o.Name) }
```

Use it in the `.Cache()` call:

```go
func ActionLabels(opts RepoOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		// ...
	}).Cache(24*time.Hour, opts.cacheKey())
}
```

For boolean Opts fields, convert them to strings in the key:

```go
func (o ResourceOpts) cacheKey() key.Key {
	return key.String(
		strconv.FormatBool(o.Analysis),
		strconv.FormatBool(o.Model),
	)
}
```

### Examples

No keys — same results regardless of context:

```go
func ActionLocalBranches() carapace.Action {
	return carapace.ActionExecCommand("git", "branch", "--format", "...")(func(output []byte) carapace.Action {
		// ...
	}).Cache(24 * time.Hour)
}
```

String parameter as key — results differ per argument:

```go
func ActionFormats(url string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("youtube-dl", "--list-formats", url)(func(output []byte) carapace.Action {
			// ...
		}).Cache(1*time.Hour, key.String(url))
	})
}
```

File stats as key — re-read when file changes:

```go
func ActionProjects(file string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("mvn", args...)(func(output []byte) carapace.Action {
			// ...
		}).Cache(-1, key.FileStats(file))
	})
}
```

Multiple keys — combine opts key with file key:

```go
}).Cache(24*time.Hour, opts.cacheKey(), key.FileStats(path))
```

## Combining Actions

### Batch

Execute multiple actions in parallel and merge results:

```go
return carapace.Batch(
	ActionLocalBranches(),
	ActionRemoteBranches(""),
).ToA()
```

### Batch with Conditional Logic

```go
batch := carapace.Batch()
if opts.LocalBookmarks {
	batch = append(batch, ActionLocalBookmarks())
}
if opts.RemoteBookmarks {
	batch = append(batch, ActionRemoteBookmarks(""))
}
return batch.ToA()
```

### Merge

Combine overlapping completions and deduplicate:

```go
carapace.Batch(actions...).Invoke(c).Merge().ToA()
```

### Invoke and ToA

Apply modifiers to batch results:

```go
ActionContainers().Invoke(c).Suffix(":").ToA().Style(styles.Docker.Container)
```

## Multi-Part Completions

### ActionMultiParts

For delimited inputs like `container:/path`:

```go
func ActionContainerPath() carapace.Action {
	return carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return ActionContainers().Suffix(":")
		default:
			return ActionContainerFiles(c.Parts[0])
		}
	})
}
```

### ActionMultiPartsN

Fixed number of parts:

```go
carapace.ActionMultiPartsN("=", 2, func(c carapace.Context) carapace.Action {
	switch len(c.Parts) {
	case 0:
		return carapace.ActionValues("option1", "option2")
	default:
		return carapace.ActionValues("value1", "value2")
	}
})
```

## Completer-Specific Action Wrappers

When integrating carapace into a cobra-based CLI, completer-specific actions often wrap shared actions by reading cobra flag state:

```go
package action

func ActionPackages(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return npm.ActionPackageSearch(cmd.Flag("registry").Value.String())
	})
}

func ActionConfigKeys(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if flag := cmd.Flag("global"); flag != nil && flag.Changed {
			return npm.ActionGlobalConfigKeys()
		}
		return npm.ActionLocalConfigKeys()
	})
}
```

These live in `cmd/action/` with package name `action`. They take `*cobra.Command` as first argument and delegate to shared actions in `pkg/actions/`.