# Basic Filters

The fundamental jq filter syntax — identity, field access, indexing, slicing, iteration, piping, and grouping. These are the building blocks every jq program is composed from.

> **Source of truth**: <https://jqlang.github.io/jq/manual/#basic-filters>. For **type construction** (arrays, objects), see [types.md](types.md). For **operators**, see [operators.md](operators.md).

## The Filter Model

Every jq expression is a **filter**: it takes one input and produces zero or more outputs. Even literals like `"hello"` or `42` are filters — they ignore the input and produce a constant value. This uniformity is key to understanding jq:

```
input → [ filter ] → output(s)
```

Filters compose:
- **Pipe** `|` — feed output(s) of left into input of right
- **Comma** `,` — concatenate output streams (left's outputs, then right's)
- **Operators** `+`, `*`, etc. — feed the same input to both sides and combine

When a filter produces multiple outputs (like `.[]`), piping into another filter runs that filter **once per output** — a cartesian product.

## Identity: `.`

The simplest filter. Takes input, produces the same value as output.

```jq
.                          # identity — pretty-prints input
```

| Program | Input | Output |
|---------|-------|--------|
| `.` | `"Hello, world!"` | `"Hello, world!"` |
| `.` | `0.12345678901234567890123456789` | `0.12345678901234567890123456789` |

`.` is the default program when none is given. Its main use is validating and pretty-printing JSON.

**Note on numbers**: jq stores number literals in their original form. Arithmetic operations convert to IEEE754 double precision, potentially losing precision. A literal that is never mutated reaches output in its original form. See [types.md](types.md) for details.

## Object Identifier-Index: `.foo`, `.foo.bar`

Access a field of an object by key. Returns `null` if the key is absent.

```jq
.foo              # value at key "foo", or null
.foo.bar          # equivalent to .foo | .bar
```

| Program | Input | Output |
|---------|-------|--------|
| `.foo` | `{"foo": 42, "bar": "less interesting data"}` | `42` |
| `.foo` | `{"notfoo": true, "alsonotfoo": false}` | `null` |

### Identifier-like Keys Only

The `.foo` shorthand only works for keys that are **identifier-like**: alphanumeric characters and underscore, not starting with a digit. For other keys, use the bracket syntax:

```jq
."foo$"            # dot-quote syntax for special characters
.["foo$"]          # equivalent bracket syntax
.["foo::bar"]      # works (".foo::bar" does not)
.["foo.bar"]       # works (".foo.bar" means .foo|.bar)
```

## Optional Object Identifier-Index: `.foo?`

Like `.foo`, but suppresses errors when `.` is not an object (produces no output instead of error).

```jq
.foo?              # .foo if . is an object, else nothing
```

| Program | Input | Output |
|---------|-------|--------|
| `.foo?` | `{"foo": 42}` | `42` |
| `.foo?` | `[1,2]` | *(no output, no error)* |
| `[.foo?]` | `[1,2]` | `[]` |

The `?` operator is shorthand for `try` — see [control.md](control.md).

## Object Index: `.[<string>]`

Look up a field using a string expression. This is the general form of `.foo`:

```jq
.["foo"]           # same as .foo
.[$key]            # key from a variable
.["a"+"b"]         # key from an expression
```

## Array Index: `.[<number>]`

Index into an array. Zero-based. Negative indices count from the end.

```jq
.[0]               # first element
.[2]               # third element
.[-1]              # last element
.[-2]              # second-to-last
```

| Program | Input | Output |
|---------|-------|--------|
| `.[0]` | `[{"name":"JSON"}, {"name":"XML"}]` | `{"name":"JSON"}` |
| `.[2]` | `[{"name":"JSON"}, {"name":"XML"}]` | `null` |
| `.[-2]` | `[1,2,3]` | `2` |

Out-of-bounds indices return `null`, not an error.

## Array/String Slice: `.[<number>:<number>]`

Return a subarray or substring. Indices are zero-based, start inclusive, end exclusive.

```jq
.[2:4]             # elements at index 2, 3
.[:3]              # first 3 elements (start omitted = 0)
.[-2:]             # last 2 elements (end omitted = length)
.[10:15]           # 5 elements starting at index 10
```

| Program | Input | Output |
|---------|-------|--------|
| `.[2:4]` | `["a","b","c","d","e"]` | `["c", "d"]` |
| `.[2:4]` | `"abcdefghi"` | `"cd"` |
| `.[:3]` | `["a","b","c","d","e"]` | `["a", "b", "c"]` |
| `.[-2:]` | `["a","b","c","d","e"]` | `["d", "e"]` |

Either index may be negative (counts from end) or omitted (refers to start or end of array).

## Array/Object Value Iterator: `.[]`

Produce all values of an array or object as separate outputs. This is a **generator** — it produces multiple results, not a single array.

```jq
.[]                # all elements of array, or all values of object
.foo[]             # equivalent to .foo | .[]
```

| Program | Input | Output |
|---------|-------|--------|
| `.[]` | `[{"name":"JSON"}, {"name":"XML"}]` | `{"name":"JSON"}` then `{"name":"XML"}` |
| `.[]` | `[]` | *(no output)* |
| `.foo[]` | `{"foo":[1,2,3]}` | `1`, `2`, `3` |
| `.[]` | `{"a": 1, "b": 1}` | `1`, `1` |

The iterator is the primary mechanism for iteration in jq. Piping `.[]` into another filter runs that filter for each value — this replaces explicit loops.

### `.[]?`

Like `.[]`, but suppresses errors when `.` is not an array or object:

```jq
.foo[]?            # equivalent to .foo | .[]?
```

## Comma: `,`

Feed the same input to both filters, concatenate their outputs in order. This is one way to construct generators.

```jq
.foo, .bar         # outputs .foo, then .bar
```

| Program | Input | Output |
|---------|-------|--------|
| `.foo, .bar` | `{"foo": 42, "bar": "else", "baz": true}` | `42`, `"else"` |
| `.user, .projects[]` | `{"user":"stedolan", "projects": ["jq", "wikiflow"]}` | `"stedolan"`, `"jq"`, `"wikiflow"` |
| `.[4,2]` | `["a","b","c","d","e"]` | `"e"`, `"c"` |

The comma operator has **higher precedence than `|`**, so `true, false | not` means `(true, false) | not`. Use parentheses to change grouping. For example, `.a | .b, .c` means `.a | (.b, .c)` because the pipe feeds into the comma expression.

## Pipe: `|`

Feed the output(s) of the left filter into the input of the right filter. Similar to Unix shell pipe.

```jq
.[] | .name        # get .name of each element
.a.b.c             # equivalent to .a | .b | .c
```

| Program | Input | Output |
|---------|-------|--------|
| `.[] | .name` | `[{"name":"JSON"}, {"name":"XML"}]` | `"JSON"`, `"XML"` |

If the left side produces multiple results, the right side runs **for each one**. This is a cartesian product:

```jq
[1,2] | .[] | [., .]      # → [1,1], [2,2]
```

`.` in a pipeline refers to the value at that stage:

```jq
.a | . | .b         # same as .a.b (the middle . is .a's output)
```

## Parenthesis

Grouping operator, same as in any programming language:

```jq
(. + 2) * 5         # add 2 to input, then multiply by 5
```

| Program | Input | Output |
|---------|-------|--------|
| `(. + 2) * 5` | `1` | `15` |

## Recursive Descent: `..`

Recursively descends `.`, producing every value at every depth. Equivalent to `recurse` (zero-argument).

```jq
..                  # produces ., then all sub-values recursively
.. | .a?           # all values of key "a" at any depth
```

| Program | Input | Output |
|---------|-------|--------|
| `.. | .a?` | `[[{"a":1}]]` | `1` |

`..a` does **not** work — use `.. | .a` or `.. | .a?` instead.

Particularly useful with `path(EXP)` and the `?` operator for finding values at arbitrary depths. See [paths-assignment.md](paths-assignment.md) for `path()` and [builtins.md](builtins.md) for `recurse`.

## References

- [jq manual — Basic filters](https://jqlang.github.io/jq/manual/#basic-filters)
- For constructing arrays and objects, see [types.md](types.md)
- For all operators, see [operators.md](operators.md)
- For `try`/`catch` and the `?` operator, see [control.md](control.md)
