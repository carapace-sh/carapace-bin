# carapace-bin

A collection of shell completers powered by [carapace](https://github.com/carapace-sh/carapace). This is **not** the carapace library itself — it aggregates completers, exposes shared actions as macros, and provides a runtime for spec-based completions.

For carapace library concepts (actions, macros, specs, integration, scraping), see the skills in `skills/`.

## Commands

| Command | Notes |
|---------|-------|
| `go generate ./cmd/...` | Regenerate all code (completer lists, macros, conditions). **Required after adding/modifying public actions or completers** |
| `go install -v -tags force_all ./cmd/carapace` | Build with all platform completers |
| `go install -v ./cmd/carapace` | Build for current platform only |
| `go test -v ./cmd/...` | Run tests |
| `gofmt -d -s .` | Format check (must produce empty diff) |
| `staticcheck ./...` | Static analysis |
| `go run ./cmd/carapace-lint completers/*/*/cmd/*.go` | Alphabetical ordering linter |
| `go run ./cmd/carapace-lint --fix-flags-order completers/*/*/cmd/*.go` | Auto-fix ordering |

No Makefile — orchestration is in `.github/workflows/go.yml`.

## MCP & Skills

This project provides a carapace MCP server. When working here, use it:

- **`list_macros`** / **`carapace_carapace_list_macros`** — Look up available macros, their signatures, and descriptions before referencing them in code or specs
- **`carapace_carapace_complete`** — Test completion output for any completer

The `skills/` directory contains detailed guides for carapace library concepts. **Use them** when working on the corresponding aspects:

| Skill | When to use |
|-------|-------------|
| `carapace-action` | Creating/modifying shared actions, naming, opts, Uid/QueryF, modifiers |
| `carapace-spec` | Writing YAML completion spec files |
| `carapace-macro` | Macro types, formatting, signature lookup, cross-executive macros |
| `carapace-scrape` | Generating specs from CLI source code (patch-and-container) |
| `carapace-integrate` | Integrating carapace into cobra CLIs (PreRun, PreInvoke, bridge) |

## Project Structure

```
cmd/
  carapace/              # Main binary: spec loader, macro dispatcher, completer invoker
  carapace-generate/     # Code generator (macros, completer lists, conditions, specs)
  carapace-lint/         # Custom linter: alphabetical flag/flag-action ordering
  carapace-parse/        # Spec parser
  carapace-shim/         # Shell shim manager
completers/              # Individual completers
  common/                # Cross-platform (some need networking)
  unix/                  # Unix-only
  linux/                 # Linux-only (~100+)
  darwin/                # macOS-only (3)
  bsd/                   # BSD variants
  windows/               # Windows-only
  android/               # Android/Termux
  bash/                  # Bash-only
pkg/
  actions/               # Shared actions exposed as macros
    actions_generated.go  # Macro map (GENERATED — do not edit)
    net/ os/ fs/ color/   # Domain categories
    tools/                # Tool-specific (one subpackage per tool)
  styles/                 # Project-specific style definitions
  conditions/             # Condition helpers (conditions_generated.go is GENERATED)
  env/                    # Environment variable helpers
  util/                   # Utility functions
  completer/              # Completer registry helpers
internal/condition/       # Internal condition logic
skills/                   # Crush skill definitions for carapace library concepts
docs/                     # mdBook documentation
.docker/                  # Docker Compose services for distro testing
```

## How This Project Differs from Normal Carapace Usage

### It's a Collection, Not a Single App

The `carapace` binary is a multiplexer — its root command uses **manual flag parsing** (`DisableFlagParsing: true`) and switches on the first argument to dispatch to subcommands, macro execution, spec loading, or completer invocation. This is fundamentally different from a typical cobra application.

### Completers Are Synthetic Cobra Commands

Each completer under `completers/` is a standalone Go binary that constructs a **synthetic cobra command tree** mimicking the target CLI's interface — not the actual tool's code. They use `carapace.Gen(rootCmd).Standalone()` to prevent cobra's default help/completion from interfering.

```
completers/common/bat_completer/
  main.go           # package main → cmd.Execute()
  cmd/
    root.go         # Synthetic cobra rootCmd + carapace completions
    cache.go        # Subcommand(s)
```

### Two Kinds of Actions: Shared vs Completer-Specific

| Aspect | `pkg/actions/` (shared) | `cmd/action/` (completer-specific) |
|--------|------------------------|-------------------------------------|
| Package | domain name (`git`, `net`) | `action` |
| Doc comments | Full format with examples | Typically none |
| Cobra awareness | No `*cobra.Command` params | Often takes `*cobra.Command` |
| Uid/QueryF | Yes (for caching) | Typically no |
| Exposed as macro | Yes | No |
| Reusability | Any completer can import | Only within its own completer |

Shared actions go in `pkg/actions/` when they're useful beyond a single completer. Completer-specific actions that depend on cobra flag state go in that completer's `cmd/action/` package.

### Code Generation Pipeline

`cmd/carapace/main.go` contains `//go:generate` directives running `carapace-generate`. Generated files:

| File | Generated from |
|------|---------------|
| `cmd/completers/completers_*.go` | Scans `completers/` directories per platform group |
| `pkg/actions/actions_generated.go` | `spec.ScanMacros()` discovers public actions from doc comments |
| `pkg/conditions/conditions_generated.go` | Scans condition functions |

**Always run `go generate ./cmd/...`** after adding/modifying public actions or completers.

The generation picks up:
- Function name → macro name (e.g. `ActionChanges` → `tools.git.Changes`)
- Doc comment first line → description
- Doc comment code footer → example
- Function signature → macro type (MacroN/MacroI/MacroV)
- Go package path → function identifier

### Platform Build Tags

Completers are partitioned by OS. Build tags control which are compiled:

- **Without `force_all`**: Only current OS group (e.g. Linux: `linux` + `unix` + `common`)
- **With `force_all`**: All groups included

Groups: `common`, `unix`, `linux`, `bsd`, `darwin`, `android`, `windows`, `freebsd`, `netbsd`, `openbsd`. The `release` tag is for goreleaser production builds.

### Shared Actions as Macros

Unlike normal carapace usage where actions are only used within the defining binary, this project exposes shared actions as macros so they're available in user specs and other completers. Macro registration happens at startup in `cmd/carapace/cmd/root.go`:

```go
for m, f := range actions.Macros {
    spec.AddMacro(m, f)
}
spec.Register(rootCmd)
```

In this project, macros use the `tools.<tool>.<Action>` naming pattern (e.g. `tools.git.Changes`). For macro type mapping and registration, see the `carapace-macro` and `carapace-action` skills.

### Bridge Actions as Completer Glue

This project uses `carapace-bridge` heavily to delegate completion to other tools. Common patterns:

```go
bridge.ActionCarapaceBin("kubectl")  // delegate to another carapace-bin completer
bridge.ActionCarapaceBin().Split()   // complete within a command-string flag value
```

For bridge details, see the `carapace-integrate` skill. Available bridges are registered as macros (e.g. `bridge.CarapaceBin`, `bridge.Bash`, `bridge.Argcomplete`).

### Uid / QueryF

- **`UidF`** creates a dynamic identifier per completion value (e.g. `git://local-branch/main` identifies a specific branch). Use it when the display value alone is insufficient, or when you need to embed additional context as query parameters
- **`QueryF`** identifies what kind of completion is being requested (e.g. `git://local-branches` indicates the current position seeks local git branches). This enables result updates with additional queries later

Some tool packages in `pkg/actions/tools/` define a package-level `Uid()` helper that constructs URL-based identifiers. Both `UidF` and `QueryF` are often chained with the same `Uid()` helper since they share the same URL construction logic.

```go
carapace.ActionExecCommand(...)(...).
    Tag("local branches").
    UidF(Uid("local-branch")).    // per-value identity: git://local-branch/main
    QueryF(Uid("local-branches")) // completion kind: git://local-branches
```

### Project-Specific Styles

Styles in `pkg/styles/` are project-level style constants:

```go
carapace.ActionValuesDescribed(vals...).Style(styles.Git.Branch)
carapace.ActionStyledValuesDescribed(vals...) // value, description, style triplets
```

## Lint Rules (carapace-lint)

Beyond standard Go tooling, two project-specific checks:

1. **Flag definitions must be alphabetically ordered** within contiguous blocks of `.Flags().(Bool|String|Float|Int|Uint|Count)()` calls
2. **Flag completion keys** in `carapace.ActionMap{}` blocks must be alphabetically ordered

Both auto-fixable with `--fix-flags-order`.

## staticcheck Configuration

`staticcheck.conf` enables all checks except ST1000, ST1003, SA1019 (deprecated func usage is common).

## Subcommand init() Ordering

Every subcommand file follows this exact init() order — violations cause subtle completion bugs:

1. `carapace.Gen(cmd).Standalone()` — **every** command gets this, not just root
2. Flag definitions (`.Flags().BoolP()`, etc.) — must be alphabetical within contiguous blocks
3. `parentCmd.AddCommand(cmd)` — register with parent
4. `carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{…})` — flag completion keys must be alphabetical
5. `carapace.Gen(cmd).PositionalCompletion(…)` / `PositionalAnyCompletion(…)`

## Completer-Specific Actions

Some completers have a `cmd/action/` directory for logic that needs `*cobra.Command` access (e.g. reading `--repo` flag value at completion time). These actions always wrap their logic in `carapace.ActionCallback(func(c carapace.Context) carapace.Action{…})` so flags are resolved lazily.

When a completer needs flag-aware actions, create `completers/<category>/<tool>_completer/cmd/action/action.go` with functions taking `*cobra.Command`. Do **not** put cobra-dependent logic in `pkg/actions/` — shared actions must not accept `*cobra.Command`.

## ActionMultiParts Pattern

Compound values (e.g. `KEY=VALUE`, `user@host`, `repo/branch`) use `carapace.ActionMultiParts(separator, callback)`. Inside the callback, switch on `len(c.Parts)` for each position. Use `.Suffix(sep)` to auto-append the separator. Nested `ActionMultiParts` handles compound formats like `KEY=VALUE,KEY2=VALUE2` (outer splits on `,`, inner on `=`).

## Testing

- **Sandbox integration tests**: `sandbox.Package(t, <import-path>)` + `s.Run(args).Expect(action)` — validates actual completion output
- **No per-completer tests**: Completers don't have `_test.go` files; validation is through the sandbox framework
- **CI runs**: `go test -v ./cmd/...` (not `./...` — there are few test files)

## Two-Tier Dispatch in the Main Binary

1. **os.Args level**: If `os.Args[1]` is `"carapace"` or `"_carapace"`, the binary enters the shim/snippet path (prints shell integration snippet). Otherwise falls through to cobra.
2. **Cobra level** (`DisableFlagParsing: true`): Manually switches on `args[0]` — `--macro`, `--condition`, `--list`, `--invoke` (default), etc.

## Gotchas

- **`go generate` is mandatory** after adding/modifying public actions — the generated macro map won't update otherwise
- **Flag alphabetical ordering** is enforced by carapace-lint, not Go tooling — CI will fail if you only run `go vet`/`staticcheck` locally
- **Build with `-tags force_all`** during development or you'll miss platform-specific compilation errors
- **The root command uses manual flag parsing** — it doesn't follow normal cobra patterns
- **`actions_generated.go` is auto-generated** — never edit directly; add public actions with proper doc comments instead
- **Opts `Default()` method** is called by the macro system when MacroI is invoked without arguments — missing this on boolean fields causes empty completion results
- **Spec YAML files live in user config dirs** (`~/.config/carapace/specs/`), not in this repo — this repo contains Go completers and the skills for generating specs
- **`Standalone()` must be called on every cobra.Command** in a completer, not just the root — skipping it on subcommands breaks their completion
- **`carapace-generate` is invoked via `go:generate` directives** in `cmd/carapace/main.go`, not by running `carapace-generate` directly
- **Windows shims are embedded** (`//go:embed` of `.exe` files in `cmd/carapace/cmd/shim/`) — changes to shim logic require rebuilding those binaries
- **The `pflag` replacement** (`go.mod` has `replace github.com/spf13/pflag => github.com/carapace-sh/carapace-pflag`) adds completion-aware flag parsing — don't reference upstream pflag behavior
