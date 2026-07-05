# Variables and Functions

Variable bindings, destructuring patterns, function definitions, and scoping rules. Variables are an "advanced feature" in jq — often unnecessary because filters pass data through pipes, but essential for complex programs.

> **Source of truth**: <https://jqlang.github.io/jq/manual/#advanced-features>. For **passing variables from the command line** (`--arg`, `--argjson`), see [cli.md](cli.md). For **`reduce`/`foreach`** which use `as` bindings, see [control.md](control.md).

## Variable Binding: `... as $identifier | ...`

Bind a value to a variable for the rest of the pipeline. All variable names start with `$`.

```jq
expression as $variable | rest_of_pipeline
```

The expression `exp as $x | ...` means: for each value of `exp`, run the rest of the pipeline with `$x` set to that value. `as` functions like a foreach loop.

```jq
length as $array_length | add / $array_length    # averaging
```

### Why Variables Are Often Unnecessary

Since every filter has an input and output, data flows through pipes without explicit variables:

```jq
add / length           # no variable needed — add and length both get the array
```

Variables are useful when you need to reference a value multiple times or when the value comes from a different part of the input.

### Practical Example

```jq
.realnames as $names | .posts[] | {title, author: $names[.author]}
```

Input:
```json
{"posts": [{"title": "First post", "author": "anon"},
           {"title": "A well-written article", "author": "person1"}],
 "realnames": {"anon": "Anonymous Coward", "person1": "Person McPherson"}}
```

Output:
```json
{"title": "First post", "author": "Anonymous Coward"}
{"title": "A well-written article", "author": "Person McPherson"}
```

## Destructuring

Multiple variables can be declared in one `as` expression using pattern matching:

```jq
. as {realnames: $names, posts: [$first, $second]} | ...
. as [$first, $second] | ...
. as [$a, $b, {c: $c}] | ...
```

Array patterns bind to elements from index 0 onward. Missing elements bind to `null`.

| Program | Input | Output |
|---------|-------|--------|
| `.bar as $x | .foo | . + $x` | `{"foo":10, "bar":200}` | `210` |
| `. as $i | [.*2 | . as $i | $i, $i]` | `5` | `[10,5]` |
| `. as [$a, $b, {c: $c}] | $a + $b + $c` | `[2, 3, {"c": 4, "d": 5}]` | `9` |
| `.[] as [$a, $b] | {a: $a, b: $b}` | `[[0], [0, 1], [2, 1, 0]]` | `{"a":0,"b":null}`, `{"a":0,"b":1}`, `{"a":2,"b":1}` |

### Object Shorthand in Destructuring

`{$foo}` in a destructuring pattern binds the value at key `"foo"` to `$foo`:

```jq
. as {$a, $b, c: {$d, $e}} | ...    # binds $a, $b, $d, $e
```

## Destructuring Alternative Operator: `?//`

Handles inputs that can take one of several structural forms. Try each pattern in order; the first that matches without error wins.

```jq
.[] as [$id, $kind, $user_id, $ts] ?// {$id, $kind, $user_id, $ts} | ...
```

### Behavior

- Each alternative need not define all the same variables
- All named variables are available to the subsequent expression
- Variables not matched in the successful alternative are `null`
- If the subsequent expression errors, the next alternative is tried
- Errors in the final alternative are passed through

### Example

Input that wraps events in an array only when there are multiple:

```json
{"resources": [
  {"id": 1, "events": {"action": "create", "user_id": 1, "ts": 13}},
  {"id": 2, "events": [{"action": "create", "user_id": 1, "ts": 14},
                       {"action": "destroy", "user_id": 1, "ts": 15}]}
]}
```

Handle both forms:

```jq
.resources[] as {$id, $kind, events: {$user_id, $ts}}
  ?// {$id, $kind, events: [{$user_id, $ts}]}
  | {$user_id, $kind, $id, $ts}
```

| Program | Input | Output |
|---------|-------|--------|
| `.[] as {$a, $b, c: {$d, $e}} ?// {$a, $b, c: [{$d, $e}]} | {$a, $b, $d, $e}` | `[{"a": 1, "b": 2, "c": {"d": 3, "e": 4}}, {"a": 1, "b": 2, "c": [{"d": 3, "e": 4}]}]` | `{"a":1,"b":2,"d":3,"e":4}` (twice) |
| `.[] as [$a] ?// [$b] | if $a != null then error("err: \($a)") else {$a,$b} end` | `[[3]]` | `{"a":null,"b":3}` |

## Variable Scope

Variables are **lexically scoped**. The binding is visible to the right of the `as`, but not past closing parentheses:

```jq
.realnames as $names | (.posts[] | {title, author: $names[.author]})  # works
(.realnames as $names | .posts[]) | {title, author: $names[.author]}  # ERROR: $names not visible
```

There is no mutation — you can create a new binding with the same name, but the old one is not changed. The new binding shadows the old in its scope.

## Defining Functions: `def`

Give a filter a name. Functions can take arguments.

```jq
def increment: . + 1;

def map(f): [.[] | f];

def addvalue(f): f as $f | map(. + $f);
```

### Arguments Are Filters, Not Values

Function arguments are passed as **filters** (callbacks), not as values. The same argument can be referenced multiple times with different inputs:

```jq
def foo(f): f|f;
5 | foo(.*2)    # → 20 (f is .*2; first call input=5→10, second input=10→20)
```

This is crucial: `f` is run with whatever input it receives at each use site, not with the value at the call site.

### Value-Argument Shorthand

If you want value-argument behavior, bind a variable inside the function:

```jq
def addvalue(f): f as $f | map(. + $f);    # $f captures f's value once
```

Or use the `($f)` shorthand which does the same:

```jq
def addvalue($f): ...;    # equivalent to: def addvalue(f): f as $f | ...
```

With either, `addvalue(.foo)` adds the input's `.foo` to each element. But `addvalue(.[])` evaluates `map(. + $f)` once per value in `.` at the call site.

### Redefinition

Multiple definitions with the same name are allowed. Each re-definition replaces the previous one for the same arity, but only for references in functions (or the main program) **after** the re-definition.

| Program | Input | Output |
|---------|-------|--------|
| `def addvalue(f): . + [f]; map(addvalue(.[0]))` | `[[1,2],[10,20]]` | `[[1,2,1], [10,20,10]]` |
| `def addvalue(f): f as $x | map(. + $x); addvalue(.[0])` | `[[1,2],[10,20]]` | `[[1,2,1,2], [10,20,1,2]]` |

## Scoping Rules

Two types of symbols in jq: **value bindings** (variables) and **functions**. Both are lexically scoped — expressions can only refer to symbols defined "to the left" of them.

The one exception: **functions can refer to themselves** (for recursion).

```jq
.*3 as $times_three | [. + $times_three]    # $times_three visible here
(.*3 as $times_three | [. + $times_three]) | .   # $times_three NOT visible past the ')'
```

## Generator Utilities

These functions help work with generators (filters producing multiple outputs):

### `isempty(exp)`

True if `exp` produces no outputs.

| Program | Input | Output |
|---------|-------|--------|
| `isempty(empty)` | `null` | `true` |
| `isempty(.[])` | `[]` | `true` |
| `isempty(.[])` | `[1,2,3]` | `false` |

### `limit(n; expr)`

Up to `n` outputs from `expr`:

| Program | Input | Output |
|---------|-------|--------|
| `[limit(3; .[])]` | `[0,1,2,3,4,5,6,7,8,9]` | `[0,1,2]` |

### `skip(n; expr)`

All outputs from `expr` except the first `n`:

| Program | Input | Output |
|---------|-------|--------|
| `[skip(3; .[])]` | `[0,1,2,3,4,5,6,7,8,9]` | `[3,4,5,6,7,8,9]` |

### `first(expr)`, `last(expr)`, `nth(n; expr)`

Extract specific outputs from a generator expression.

| Program | Input | Output |
|---------|-------|--------|
| `[first(range(.)), last(range(.)), nth(5; range(.))]` | `10` | `[0,9,5]` |

### `first`, `last`, `nth(n)`

Extract from the array at `.` (no generator argument):

| Program | Input | Output |
|---------|-------|--------|
| `[range(.)] | [first, last, nth(5)]` | `10` | `[0,9,5]` |

## References

- [jq manual — Variable / Symbolic Binding Operator](https://jqlang.github.io/jq/manual/#variable-symbolic-binding-operator)
- [jq manual — Destructuring Alternative Operator](https://jqlang.github.io/jq/manual/#destructuring-alternative-operator)
- [jq manual — Defining Functions](https://jqlang.github.io/jq/manual/#defining-functions)
- For `reduce`/`foreach` (which use `as` bindings), see [control.md](control.md)
- For command-line variable passing (`--arg`), see [cli.md](cli.md)
