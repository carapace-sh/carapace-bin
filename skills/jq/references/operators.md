# Operators

Arithmetic, comparison, boolean, and the alternative operator. jq operators are type-aware — they do different things depending on the types involved, but **never do implicit type conversion**.

> **Source of truth**: <https://jqlang.github.io/jq/manual/#builtin-operators-and-functions> (addition through alternative operator). For **control-flow operators** (`if`, `try`, `?`), see [control.md](control.md). For **assignment operators**, see [paths-assignment.md](paths-assignment.md).

## Arithmetic Operators

All arithmetic operators take two filters, apply both to the same input, and combine the results. Numbers use IEEE754 double precision (see [types.md](types.md)).

### Addition: `+`

What "adding" means depends on the types:

| Types | Behavior |
|-------|----------|
| numbers | Arithmetic addition |
| arrays | Concatenation into larger array |
| strings | Concatenation into larger string |
| objects | Merge — right side wins on key conflicts (for recursive merge use `*`) |
| null + anything | Returns the other value unchanged |
| anything + null | Returns the original value unchanged |

```jq
.a + 1                       # number + number
.a + .b                      # array + array (concatenate)
{a: 1} + {b: 2} + {a: 42}    # object merge (right wins → a: 42)
```

| Program | Input | Output |
|---------|-------|--------|
| `.a + 1` | `{"a": 7}` | `8` |
| `.a + .b` | `{"a": [1,2], "b": [3,4]}` | `[1,2,3,4]` |
| `.a + null` | `{"a": 1}` | `1` |
| `.a + 1` | `{}` | `1` |
| `{a: 1} + {b: 2} + {c: 3} + {a: 42}` | `null` | `{"a": 42, "b": 2, "c": 3}` |

### Subtraction: `-`

| Types | Behavior |
|-------|----------|
| numbers | Arithmetic subtraction |
| arrays | Remove all occurrences of second array's elements from first |

| Program | Input | Output |
|---------|-------|--------|
| `4 - .a` | `{"a":3}` | `1` |
| `. - ["xml", "yaml"]` | `["xml", "yaml", "json"]` | `["json"]` |

### Multiplication: `*`, Division: `/`, Modulo: `%`

| Operator | Numbers | Strings | Objects |
|----------|---------|---------|---------|
| `*` | Multiply | `"x" * 3` → `"xxx"` (repeat) | Recursive merge (deep) |
| `/` | Divide (error on 0) | `"a,b" / ","` → `["a","b"]` (split) | — |
| `%` | Modulo | — | — |

Object multiplication merges recursively: like `+` but if both contain the same key with object values, those are merged with the same strategy.

| Program | Input | Output |
|---------|-------|--------|
| `10 / . * 3` | `5` | `6` |
| `. / ", "` | `"a, b,c,d, e"` | `["a","b,c,d","e"]` |
| `{"k": {"a": 1, "b": 2}} * {"k": {"a": 0,"c": 3}}` | `null` | `{"k": {"a": 0, "b": 2, "c": 3}}` |
| `.[] | (1 / .)?` | `[1,0,-1]` | `1`, `-1` (0 raises error, suppressed by `?`) |

`"x" * 0` produces `""` (empty string).

### `abs`

Absolute value. Defined as `if . < 0 then - . else . end`.

```jq
abs                          # |-5| → 5
```

For floating-point absolute value, use `fabs`.

| Program | Input | Output |
|---------|-------|--------|
| `map(abs)` | `[-10, -1.1, -1e-1]` | `[10,1.1,1e-1]` |

### `length`

Length varies by type:

| Type | Length |
|------|--------|
| string | Number of Unicode codepoints |
| number | Absolute value |
| array | Number of elements |
| object | Number of key-value pairs |
| null | 0 |
| boolean | **Error** |

| Program | Input | Output |
|---------|-------|--------|
| `.[] | length` | `[[1,2], "string", {"a":2}, null, -5]` | `2`, `6`, `1`, `0`, `5` |

### `utf8bytelength`

Number of bytes to encode a string in UTF-8 (differs from `length` for non-ASCII):

| Program | Input | Output |
|---------|-------|--------|
| `utf8bytelength` | `"\u03bc"` | `2` |

## Comparison Operators

### Equality: `==`, `!=`

`==` produces `true` if both sides represent equivalent JSON values. This is **strict equality** — like JavaScript's `===`, not `==`. Strings are never equal to numbers. Object key ordering is irrelevant.

```jq
. == false                    # strict comparison
. == {"b": {"d": 4}, "a": 1}  # key order doesn't matter
```

| Program | Input | Output |
|---------|-------|--------|
| `. == false` | `null` | `false` |
| `. == {"b": {"d": (4 + 1e-20), "c": 3}, "a":1}` | `{"a":1, "b": {"c": 3, "d": 4}}` | `true` |
| `.[] == 1` | `[1, 1.0, "1", "banana"]` | `true`, `true`, `false`, `false` |

`!=` is the negation of `==`.

### Ordering: `>`, `>=`, `<=`, `<`

The sort order (same as used by `sort`):

```
null < false < true < numbers < strings < arrays < objects
```

- **Strings**: alphabetical by Unicode codepoint value
- **Arrays**: lexical order (element by element)
- **Objects**: compared by sorted keys first, then values key by key

| Program | Input | Output |
|---------|-------|--------|
| `. < 5` | `2` | `true` |

## Boolean Operators

### Truthiness

jq has a simple truthiness model: **`false` and `null` are false; everything else is true**. This differs from JavaScript/Python — empty strings, `0`, and empty arrays are all **true** in jq.

### `and`, `or`

Standard boolean operators. Operate on truthiness. If an operand produces multiple results, the operator produces a result for each.

```jq
42 and "a string"             # true (both truthy)
(true, false) or false        # true, false
(true, true) and (true, false)  # true, false, true, false (cartesian)
```

| Program | Input | Output |
|---------|-------|--------|
| `42 and "a string"` | `null` | `true` |
| `(true, false) or false` | `null` | `true`, `false` |
| `(true, true) and (true, false)` | `null` | `true`, `false`, `true`, `false` |

### `not`

A builtin function (not an operator), called as a filter:

```jq
.foo and .bar | not           # negate the boolean result
```

| Program | Input | Output |
|---------|-------|--------|
| `[true, false | not]` | `null` | `[false, true]` |

`and` and `or` only produce `true`/`false` — they are not for the Python/Ruby idiom of `value or default`. For that, use `//`.

## Alternative Operator: `//`

The `//` operator provides **defaults**. It produces all values from the left side that are neither `false` nor `null`. If the left side produces no values (or only `false`/`null`), it produces all values from the right side.

```jq
.foo // 42                    # .foo if truthy, else 42
empty // 42                   # → 42 (left produced nothing)
```

| Program | Input | Output |
|---------|-------|--------|
| `empty // 42` | `null` | `42` |
| `.foo // 42` | `{"foo": 19}` | `19` |
| `.foo // 42` | `{}` | `42` |
| `(false, null, 1) // 42` | `null` | `1` |
| `(false, null, 1) | . // 42` | `null` | `42`, `42`, `1` |

### Precedence Gotcha

`//` has higher precedence than `,` (comma). This means:

```jq
false, 1 // 2                 # parses as: false, (1 // 2) → false, 1
(false, null, 1) // 42        # left side is a generator of 3 values → 1
(false, null, 1) | . // 42    # left side of // is ., produces 42, 42, 1
```

The difference between `generator // default` and `generator | . // default`:
- `generator // default` — if the generator produces **any** truthy value, the default is not used
- `generator | . // default` — applies `// default` to **each** value individually

## References

- [jq manual — Addition](https://jqlang.github.io/jq/manual/#addition) through [Alternative operator](https://jqlang.github.io/jq/manual/#alternative-operator)
- For `if`/`then`/`else`/`elif`/`end`, see [control.md](control.md)
- For assignment operators (`=`, `|=`, `+=`), see [paths-assignment.md](paths-assignment.md)
- For number representation details, see [types.md](types.md)
