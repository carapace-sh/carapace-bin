# Building a Completion-Aware Lexer

How to build a parser/lexer for a DSL (domain-specific language) that integrates with carapace completion. This covers the two-layer architecture used by carapace-sh lexer projects (carapace-jq, carapace-jjlex, carapace-pnpm, carapace-ffmpeg, carapace-magick), where a pure-stdlib grammar package produces a `CompletionContext` and a thin actions package wires it to carapace.

> **When to read this**: You are building completion for a CLI argument value that has its own grammar (jq filter expressions, jj revsets, pnpm filter selectors, ffmpeg filtergraph strings). `ActionMultiParts` is not powerful enough because the grammar has nested constructs, context-dependent token expectations, or error recovery.

## When to Build a Lexer

Most flag value completion is simple — a fixed list, a file path, or a `KEY=VALUE` pair handled by `ActionMultiParts`. Build a dedicated lexer when:

| Situation | Why ActionMultiParts fails |
|-----------|---------------------------|
| Nested grammar (parentheses, brackets, string interpolation) | `ActionMultiParts` splits on one separator; nested delimiters need a real parser |
| Context-dependent completions | What to complete depends on what came before (e.g., function name vs. argument position) |
| Error recovery needed | Partial input is always syntactically incomplete; the parser must recover and report what's expected |
| Multiple token types at cursor | The cursor could be in a string, an identifier, an operator position — the parser disambiguates |

## Architecture Overview

Every carapace-sh lexer project has two layers:

```
┌──────────────────────────────────────────────────┐
│  Layer 1: Grammar Package (pkg/<grammar>/)       │
│  Pure Go stdlib — no carapace dependency          │
│                                                   │
│  Parse(input) → AST                               │
│  ParseForCompletion(input) → *CompletionContext   │
└──────────────┬───────────────────────────────────┘
               │ imports grammar package + carapace
┌──────────────▼───────────────────────────────────┐
│  Layer 2: Actions Package (pkg/actions/tools/<t>/)│
│  Wires CompletionContext to carapace.Action        │
│  Also: data actions (static lists, shell-out)      │
└──────────────────────────────────────────────────┘
```

**Why two layers?** The grammar package must not import carapace. This keeps it testable in isolation, reusable outside of completion, and keeps the dependency graph clean. The actions package is the only place that knows about carapace.

### Two Archetypes

| Archetype | Grammar input | Examples | carapace-bin integration |
|-----------|---------------|----------|--------------------------|
| **String-grammar** | `Parse(string)` | carapace-jq, carapace-jjlex, carapace-pnpm | Go module dependency; macros scanned by `go:generate` |
| **Argstream** | `Parse([]string, trailingSpace)` | carapace-ffmpeg, carapace-magick | `bridge.ActionCarapace("carapace-ffmpeg")` |

String-grammar lexers parse a single flag value (e.g., `--filter 'foo...'`). Argstream lexers parse the entire command-line argument array for CLIs that don't follow the traditional subcommand-tree model. Most new lexers will be string-grammar.

## Project Structure

A lexer project is a separate Go module in the carapace-sh ecosystem:

```
carapace-<tool>/
  go.mod                           # module github.com/carapace-sh/carapace-<tool>
  go.sum
  cmd/carapace-<tool>/             # Debug CLI
    main.go                        # package main → cmd.Execute()
    cmd/
      root.go                      # carapace.Gen(rootCmd).Standalone(); spec.Register(rootCmd)
      filter.go                    # "filter" subcommand: Parse(input) → JSON AST
      filter-complete.go           # "filter-complete" subcommand: ParseForCompletion(input) → JSON
  pkg/
    <grammar>/                     # Layer 1: grammar package (pure stdlib)
      scanner.go                   # Rune classification helpers
      span.go                      # Span, Pos types
      ast.go                       # AST node types
      parser.go                    # Full parser: Parse(input) (T, error)
      completion_parser.go         # compParser: ParseForCompletion(input) *CompletionContext
      completion.go               # CompletionContext, ExpectedToken, domain context structs
      format.go                    # AST pretty-printing (optional)
    actions/tools/<tool>/          # Layer 2: actions package (imports carapace)
      completion.go                # ActionXyz() — wraps ParseForCompletion in ActionCallback
      builtins.go                  # Static data actions (e.g., ActionBuiltins())
      uid.go                       # Uid() helper for macro UidF
  skills/<tool>/                   # Compound skill documenting the DSL
    SKILL.md
    references/
  man/<tool>/                      # Man page YAML (optional)
  AGENTS.md
```

For multi-grammar projects (e.g., carapace-jjlex has revset, fileset, template), repeat the `pkg/<grammar>/` package for each grammar. The actions package lives at `pkg/actions/tools/<tool>/` and imports all grammar packages.

## Layer 1: Grammar Package

### Scanner (scanner.go)

Character classification helpers. Keep these simple and self-contained:

```go
package jq

func isWhitespace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\r' || r == '\n' || r == '\x0c'
}

func isIdentifierStart(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || r == '_'
}

func isIdentifierPart(r rune) bool {
	return isIdentifierStart(r) || (r >= '0' && r <= '9')
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}
```

A keyword map distinguishes reserved words from identifiers:

```go
var jqKeywords = map[string]bool{
	"if": true, "then": true, "else": true, "elif": true, "end": true,
	"try": true, "catch": true, "reduce": true, "foreach": true, "as": true,
	"def": true, "import": true, "include": true, "module": true,
	"label": true, "break": true, "and": true, "or": true, "not": true,
}
```

### Span and Pos (span.go)

Source positions for error reporting and AST node locations:

```go
package jq

type Span struct {
	Start int
	End   int
}

type Pos struct {
	Offset int
	Line   int
	Column int
}
```

### AST (ast.go)

Design AST nodes as tagged unions with an enum kind discriminator. This keeps the AST serializable to JSON (useful for the debug CLI) and avoids interface complexity:

```go
type ExpressionKind int

const (
	KindIdentity ExpressionKind = iota
	KindField
	KindPipe
	KindBinary
	KindFunctionCall
	// ...
)

type Expression struct {
	Kind    ExpressionKind `json:"kind"`
	Span    Span           `json:"span"`
	payload any // unexported, kind-specific data
}
```

The `payload` field holds kind-specific data (e.g., child expressions for `KindPipe`, function name and args for `KindFunctionCall`). It is unexported and accessed via methods, keeping the JSON serialization clean (only `Kind` and `Span` appear in JSON output).

### Full Parser (parser.go)

A recursive-descent parser with `ParseError` carrying a span:

```go
type ParseError struct {
	Message string
	Span    Span
}

func (e *ParseError) Error() string { return e.Message }

type parser struct {
	input       string
	pos         int
	lastContent int // position after last non-whitespace content
}

// Parse parses a complete expression string into an AST.
func Parse(input string) (*Expression, error) {
	p := &parser{input: input}
	p.skipWhitespaceAndComments()
	expr, err := p.parseQuery()
	if err != nil {
		return nil, err
	}
	p.skipWhitespaceAndComments()
	if p.pos < len(p.input) {
		return nil, p.syntaxError("unexpected token")
	}
	return expr, nil
}
```

Helper methods for cursor movement:

```go
func (p *parser) peek() rune {
	if p.pos >= len(p.input) {
		return 0
	}
	r, _ := utf8.DecodeRuneInString(p.input[p.pos:])
	return r
}

func (p *parser) advance() rune {
	if p.pos >= len(p.input) {
		return 0
	}
	r, w := utf8.DecodeRuneInString(p.input[p.pos:])
	p.pos += w
	p.lastContent = p.pos
	return r
}

func (p *parser) syntaxError(msg string) *ParseError {
	return &ParseError{
		Message: msg,
		Span:    Span{Start: p.pos, End: min(p.pos+1, len(p.input))},
	}
}
```

### Completion Parser (completion_parser.go)

The key insight: **the completion parser is a separate struct from the full parser**. It re-implements the grammar with error recovery, tracking where the cursor is and what tokens are expected there.

```go
// ParseForCompletion parses a partial expression and returns a
// CompletionContext describing what is expected at the end of the input.
// Partial expressions are allowed — the parser recovers from errors to
// report what tokens would be valid at the cursor position.
func ParseForCompletion(input string) *CompletionContext {
	cursor := len(input)
	p := &compParser{
		input:  input,
		pos:    0,
		cursor: cursor,
		ctx:    &CompletionContext{},
	}
	p.skipWS()
	p.parseQuery()
	if len(p.ctx.ExpectedTokens) == 0 {
		p.ctx.ExpectedTokens = append(p.ctx.ExpectedTokens, ExpectedExpression)
	}
	p.ctx.ExpectedTokens = dedupTokens(p.ctx.ExpectedTokens)
	return p.ctx
}
```

The `compParser` struct carries cursor position and state stacks for nested constructs:

```go
type compParser struct {
	input  string
	pos    int
	cursor int
	ctx    *CompletionContext

	consumed    bool // at least one token consumed before cursor
	exprStarted bool // new expression started at cursor

	funcStack   []*funcState   // nested function calls
	objStack    []*objState    // nested object construction
	reduceStack []*reduceState // nested reduce/foreach
	ifStack     []*ifState     // nested if-then-elif-else-end
	parenDepth  int
}
```

**Error recovery strategy**: The completion parser never returns an error. Instead, when it encounters something unexpected, it records what *would* be valid at that position in `ctx.ExpectedTokens` and continues. The goal is to reach the cursor position with enough context to report what comes next.

Key patterns:
- **`beforeExpression()`**: Called when a new expression could start. Sets `ExpectedExpression` and marks `exprStarted`.
- **`afterExpression()`**: Called after a complete expression. Sets `ExpectedOperator`, `ExpectedPipe`, etc.
- **`atCursor()`**: Returns true when `p.pos >= p.cursor`. Used to decide whether to record expected tokens or continue parsing.
- **State stacks**: Track nesting (function arguments, object key/value positions, reduce sections, if sections) so the completion context can report what part of a construct the cursor is in.

### CompletionContext (completion.go)

The `CompletionContext` struct is the contract between the grammar package and the actions package. It describes what is expected at the completion position:

```go
// ExpectedToken represents a type of token expected at a completion position.
type ExpectedToken int

const (
	ExpectedExpression ExpectedToken = iota
	ExpectedOperator
	ExpectedPipe
	ExpectedColon
	ExpectedClosingParen
	ExpectedClosingBracket
	ExpectedClosingBrace
	ExpectedComma
	ExpectedKeyword // then, else, elif, end, catch, as
	ExpectedFormatName
	// ...
)

func (t ExpectedToken) String() string { /* ... */ }
func (t ExpectedToken) MarshalText() ([]byte, error) { return []byte(t.String()), nil }

// CompletionContext describes what is expected at the completion position.
type CompletionContext struct {
	ExpectedTokens []ExpectedToken `json:"expectedTokens"`

	ValidOperators []ValidOperator `json:"validOperators,omitempty"`
	ValidKeywords  []string        `json:"validKeywords,omitempty"`

	PartialIdent   string `json:"partialIdent,omitempty"`
	PartialString  string `json:"partialString,omitempty"`
	StringQuote    rune   `json:"stringQuote,omitempty"`
	InStringInterp bool   `json:"inStringInterp,omitempty"`

	Function *FunctionContext `json:"function,omitempty"`
	Object   *ObjectContext  `json:"object,omitempty"`
	Reduce   *ReduceContext  `json:"reduce,omitempty"`
	If       *IfContext      `json:"if,omitempty"`

	InFormat      bool   `json:"inFormat"`
	PartialFormat string `json:"partialFormat,omitempty"`
	AfterDot      bool   `json:"afterDot"`
}
```

Design principles for `CompletionContext`:
- **JSON tags on all fields** — the debug CLI serializes it to JSON for inspection and testing.
- **`ExpectedTokens` as the primary signal** — an enum of what token categories are valid at the cursor.
- **Domain-specific context structs** — `FunctionContext`, `ObjectContext`, `IfContext`, etc. carry the state needed for context-aware completion (e.g., "inside function `map` at argument index 0").
- **Partial fields** — `PartialIdent`, `PartialString` carry the incomplete text being typed, useful for prefix filtering.
- **`MarshalText()` on ExpectedToken** — renders as a human-readable string in JSON output.

## Layer 2: Actions Package

### The ActionCallback Wrapper (completion.go)

The actions package wraps `ParseForCompletion` inside `carapace.ActionCallback`. The action:

1. Calls `ParseForCompletion(c.Value)` to get the completion context
2. Splits the input into `typedPrefix` (already-typed portion) and `partialToken` (in-progress token)
3. Temporarily sets `c.Value = partialToken` so sub-actions filter on the partial
4. Dispatches to a sub-action based on `CompletionContext` fields
5. Re-attaches the prefix via `.Prefix(typedPrefix)`

```go
package jq

import (
	"strings"

	"github.com/carapace-sh/carapace"
	jqparser "github.com/carapace-sh/carapace-jq/pkg/jq"
)

func ActionFilters() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		expr := c.Value
		ctx := jqparser.ParseForCompletion(expr)

		// Split into already-typed prefix and in-progress partial token
		typedPrefix := ""
		partialToken := expr
		if lastSpace := strings.LastIndex(expr, " "); lastSpace >= 0 {
			typedPrefix = expr[:lastSpace+1]
			partialToken = expr[lastSpace+1:]
		}

		// For single-token expressions, don't strip the prefix
		if !strings.Contains(expr, " ") &&
			!hasExpected(ctx, jqparser.ExpectedExpression) &&
			!hasExpected(ctx, jqparser.ExpectedFormatName) {
			typedPrefix = expr
			partialToken = ""
		}

		c.Value = partialToken
		return actionForCompletionContext(ctx).Invoke(c).Prefix(typedPrefix).ToA()
	})
}
```

### Dispatching to Sub-Actions

A dispatcher function inspects `CompletionContext` fields and returns the appropriate carapace action. Priority matters — check specific contexts before generic ones:

```go
func actionForCompletionContext(ctx *jqparser.CompletionContext) carapace.Action {
	// String contexts — highest priority
	if ctx.StringQuote != 0 && !ctx.InStringInterp {
		return carapace.ActionValues("\"").NoSpace()
	}
	if ctx.InStringInterp {
		return actionForExpectedExpression(ctx)
	}
	if ctx.InFormat {
		return ActionFormats()
	}

	// Construct-specific contexts
	if ctx.Function != nil {
		return actionForExpectedExpression(ctx)
	}
	if ctx.Object != nil {
		if ctx.Object.InKey {
			return carapace.ActionMessage("object key")
		}
		if ctx.Object.InValue {
			return actionForExpectedExpression(ctx)
		}
	}

	// Generic fallback: expected expression → builtins, keywords, etc.
	if hasExpected(ctx, jqparser.ExpectedExpression) {
		return actionForExpectedExpression(ctx)
	}

	// Operators, keywords, closing tokens
	if hasExpected(ctx, jqparser.ExpectedOperator) {
		return actionForOperators(ctx)
	}
	if hasExpected(ctx, jqparser.ExpectedKeyword) {
		return actionForKeywords(ctx)
	}

	return carapace.ActionValues()
}
```

### Data Actions

Static data actions provide the candidate values (builtins, operators, format names). These follow the patterns from [action.md](action.md) — exported functions prefixed with `Action`, doc comments with examples, `UidF` for stable IDs:

```go
// ActionBuiltins completes jq builtin functions.
//
//	map (map values)
//	select (filter values)
//	walk (recursive update)
func ActionBuiltins() carapace.Action {
	return carapace.ActionValuesDescribed(
		"map", "map values",
		"select", "filter values",
		"walk", "recursive update",
		// ...
	).Tag("builtins").UidF(Uid("builtins"))
}
```

Dynamic data actions shell out to the target CLI for runtime data:

```go
// ActionWorkspacePackages completes workspace package names by running
// `pnpm list --json -r --depth -1`.
//
//	@scope/widget (1.0.0)
//	shared        (2.1.0)
func ActionWorkspacePackages() carapace.Action {
	return carapace.ActionExecCommandE("pnpm", "list", "--json", "-r", "--depth", "-1")(
		func(output []byte, err error) carapace.Action {
			var entries []pnpmListEntry
			if len(output) == 0 || json.Unmarshal(output, &entries) != nil {
				return carapace.ActionValues()
			}
			vals := make([]string, 0, len(entries)*2)
			for _, e := range entries {
				vals = append(vals, e.Name, e.Version)
			}
			return carapace.ActionValuesDescribed(vals...).Tag("workspace package").UidF(Uid("workspace-package"))
		},
	).Cache(0)
}
```

### Multi-Grammar Projects

For projects with multiple grammars (e.g., carapace-jjlex has revset, fileset, template), the actions package imports all grammar packages and provides one action per grammar:

```go
package jj

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-jjlex/pkg/revset"
	"github.com/carapace-sh/carapace-jjlex/pkg/fileset"
	"github.com/carapace-sh/carapace-jjlex/pkg/template"
)

// ActionRevsets completes jj revset expressions.
func ActionRevsets() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		ctx := revset.ParseForCompletion(c.Value)
		return actionForRevsetContext(ctx).Invoke(c).Prefix(/* ... */).ToA()
	})
}

// ActionFilesets completes jj fileset expressions.
func ActionFilesets() carapace.Action { /* ... */ }

// ActionTemplates completes jj template expressions.
func ActionTemplates() carapace.Action { /* ... */ }
```

Each grammar package is independent; the actions package is the shared integration point.

## Debug CLI (cmd/)

The debug CLI has two subcommands for testing the parser:

### filter — Full parse to JSON AST

```go
var filterCmd = &cobra.Command{
	Use:   "filter <expression>",
	Short: "Parse a filter expression",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		expression, err := jqparser.Parse(args[0])
		if err != nil {
			return err
		}
		m, _ := json.MarshalIndent(expression, "", "  ")
		fmt.Println(string(m))
		return nil
	},
}
```

### filter-complete — Completion context to JSON

```go
var filterCompleteCmd = &cobra.Command{
	Use:   "filter-complete <expression>",
	Short: "Get completion context for a filter expression",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := jqparser.ParseForCompletion(args[0])
		m, _ := json.MarshalIndent(ctx, "", "  ")
		fmt.Println(string(m))
		return nil
	},
}
```

Both subcommands get carapace completion themselves, so the debug CLI is self-completing:

```go
func init() {
	rootCmd.AddCommand(filterCmd)
	rootCmd.AddCommand(filterCompleteCmd)
	carapace.Gen(filterCmd).PositionalAnyCompletion(jq.ActionFilters())
	carapace.Gen(filterCompleteCmd).PositionalAnyCompletion(jq.ActionFilters())
}
```

## Integration with carapace-bin

### Go Module Dependency

carapace-bin adds the lexer project as a Go module dependency in its `go.mod`. The actions package (`pkg/actions/tools/<tool>/`) is imported by completers that need the lexer's completion actions:

```go
// completers/common/jq_completer/cmd/root.go
import (
	jqaction "github.com/carapace-sh/carapace-jq/pkg/actions/tools/jq"
)

carapace.Gen(rootCmd).PositionalAnyCompletion(jqaction.ActionFilters())
```

### Macro Scanning

carapace-bin's `cmd/carapace/main.go` has a `//go:generate` directive that scans external action packages for macros:

```go
//go:generate go run ../carapace-generate macros --code \
//   github.com/carapace-sh/carapace-bin/pkg/actions \
//   github.com/carapace-sh/carapace-bridge/pkg/actions \
//   github.com/carapace-sh/carapace-jjlex/pkg/actions \
//   github.com/carapace-sh/carapace-pnpm/pkg/actions \
//   --output ../../pkg/actions/actions_generated.go
```

Note that carapace-jq is **not** in this list — its actions are imported directly by completers (e.g., `gh`, `jq`) rather than exposed as macros. Add your lexer project's actions package path to this directive only if the actions should be available as macros (e.g., `tools.jj.Revs`). The generator discovers all exported `Action*` functions with doc comments and registers them. See [macro.md](macro.md) for macro naming and format.

After adding the package path, run `go generate ./cmd/...` to regenerate `actions_generated.go`.

### Bridge Integration (argstream archetype)

For argstream lexers (carapace-ffmpeg, carapace-magick), carapace-bin bridges to the standalone binary rather than importing it as a Go module:

```go
bridge.ActionCarapace("carapace-ffmpeg")
bridge.ActionCarapace("carapace-magick", "identify")
```

This is because argstream lexers need to process the entire argument array, which doesn't fit carapace-bin's per-flag-value completion model. The standalone binary handles the full argument stream and returns completions via the bridge protocol.

## Testing

Grammar packages are tested directly (no carapace dependency needed). Test both the full parser and the completion parser:

```go
func TestParseSimple(t *testing.T) {
	expr, err := jq.Parse(".foo | .bar")
	if err != nil {
		t.Fatal(err)
	}
	if expr.Kind != jq.KindPipe {
		t.Errorf("expected Pipe, got %s", expr.Kind)
	}
}

func TestCompletionContextAfterPipe(t *testing.T) {
	ctx := jq.ParseForCompletion(".foo | ")
	if !hasExpected(ctx, jq.ExpectedExpression) {
		t.Error("expected Expression after pipe")
	}
}
```

The actions package uses carapace's sandbox integration test framework:

```go
func TestActionFiltersEmpty(t *testing.T) {
	sandbox.Action(t, func() carapace.Action {
		return ActionFilters()
	})(func(s *sandbox.Sandbox) {
		s.Run("").Expect(
			carapace.Batch(
				carapace.ActionValuesDescribed(
					".", "the package in the current directory",
				).Tag("self"),
				carapace.ActionValuesDescribed(
					"{./", "glob group, e.g. {./apps/*,./packages/*}",
					"[", "packages changed since a git ref, e.g. [master]",
				).NoSpace().Tag("path glob"),
			).ToA(),
		)
	})
}
```

## Workflow Summary

1. **Research the DSL** — Write a compound skill documenting the language grammar (see existing skills in `carapace-jq/skills/jq/`, `carapace-jjlex/skills/jj/`). The skill is the knowledge base you implement from.
2. **Build the grammar package** (`pkg/<grammar>/`) — scanner, span, ast, parser, completion_parser, completion. Pure stdlib, no carapace import.
3. **Build the debug CLI** (`cmd/carapace-<tool>/`) — `filter` and `filter-complete` subcommands for interactive testing.
4. **Build the actions package** (`pkg/actions/tools/<tool>/`) — `ActionCallback` wrapper, dispatcher, data actions.
5. **Add to carapace-bin** — Add module dependency, add actions package path to `go:generate` directive, run `go generate ./cmd/...`.
6. **Wire completers** — Import the actions package in the relevant carapace-bin completer(s) and use in `PositionalCompletion` / `FlagCompletion`.

## Existing Examples

| Project | Grammars | Complexity | Good starting point |
|----------|----------|------------|---------------------|
| [carapace-jq](https://github.com/carapace-sh/carapace-jq) | jq filter | Medium (nested constructs, string interpolation, reduce/foreach) | Simplest single-grammar lexer |
| [carapace-pnpm](https://github.com/carapace-sh/carapace-pnpm) | pnpm filter selector | Low (flat grammar, relational modifiers) | Simplest grammar; good for learning the pattern |
| [carapace-jjlex](https://github.com/carapace-sh/carapace-jjlex) | revset, fileset, template | High (3 grammars, type-aware template completion) | Multi-grammar reference |
| [carapace-ffmpeg](https://github.com/carapace-sh/carapace-ffmpeg) | argstream, filtergraph, streamspec, mapvalue | High (argstream + 3 sub-grammars, probe) | Argstream archetype reference |
| [carapace-magick](https://github.com/carapace-sh/carapace-magick) | argstream, definevalue | Medium (argstream + 1 sub-grammar) | Simpler argstream example |

Start with carapace-jq or carapace-pnpm as a template for a new string-grammar lexer. Use carapace-ffmpeg as a template if you need an argstream lexer.
