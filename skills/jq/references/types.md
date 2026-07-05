# Types and Values

jq's type system — the six JSON types, number representation, and how to construct arrays and objects.

> **Source of truth**: <https://jqlang.github.io/jq/manual/#types-and-values>. For **basic filter syntax** (`.`, `.foo`, `.[]`), see [filters.md](filters.md). For **type-checking builtins** (`type`, `numbers`, `strings`, etc.), see [builtins.md](builtins.md).

## The Six Types

jq supports the same types as JSON:

| Type | Literal | `type` output | Notes |
|------|---------|---------------|-------|
| **null** | `null` | `"null"` | Absence / default |
| **boolean** | `true`, `false` | `"boolean"` | Only two values |
| **number** | `42`, `3.14`, `1e10` | `"number"` | IEEE754 double internally |
| **string** | `"hello"` | `"string"` | UTF-8, JSON-quoted |
| **array** | `[1, 2, 3]` | `"array"` | Ordered, zero-indexed |
| **object** | `{"a": 1}` | `"object"` | String keys only |

Every value — including literals — is a filter that ignores its input and produces itself:

```jq
42                 # filter that always outputs 42
"hello"            # filter that always outputs "hello"
```

## Numbers

Numbers are internally represented as IEEE754 double precision. However, jq preserves the **original literal form** of number literals that are never mutated:

```jq
.                  # input 0.12345678901234567890123456789 → same output (no mutation)
. + 0              # same number, but now converted to double → precision may be lost
```

### Number Representation Rules

1. **Arithmetic triggers conversion** — any arithmetic operation converts to IEEE754 double
2. **Literal preservation** — if a number literal is never mutated, it reaches output in original form
3. **Leading minus triggers conversion** — `-5` is converted to double (the `-` is an operator, not part of the literal)
4. **Comparisons use big decimal** — comparisons use the untruncated representation if available

| Program | Input | Output |
|---------|-------|--------|
| `.` | `0.12345678901234567890123456789` | `0.12345678901234567890123456789` |
| `[., tojson]` | `12345678909876543212345` | `[12345678909876543212345,"12345678909876543212345"]` |
| `. < 0.12345678901234567890123456788` | `0.12345678901234567890123456789` | `false` |
| `map([., . == 1]) | tojson` | `[1, 1.000, 1.0, 100e-2]` | `"[[1,true],[1.000,true],[1.0,true],[1.00,true]]"` |
| `. as $big | [$big, $big + 1] | map(. > 10000000000000000000000000000000)` | `10000000000000000000000000000001` | `[true, false]` |

### Infinity and NaN

| Builtin | Returns |
|---------|---------|
| `infinite` | Positive infinity |
| `nan` | NaN (not a number) |
| `isinfinite` | `true` if input is infinite |
| `isnan` | `true` if input is NaN |
| `isfinite` | `true` if input is finite |
| `isnormal` | `true` if input is a normal number |

Division by zero raises an error (does not produce infinity). Most operations on infinities/NaNs do not raise errors.

```jq
infinite, nan | type              # "number", "number"
.[] | (infinite * .) < 0          # input [-1, 1] → true, false
```

## Array Construction: `[]`

`[]` collects all results of an expression into a single array:

```jq
[1, 2, 3]                    # literal array
[.user, .projects[]]         # collect results into array
[.[] | . * 2]                # map and collect
```

The key insight: `[1,2,3]` is the `[]` operator applied to the expression `1,2,3` (which produces three results via the comma operator). If a filter `X` produces four results, `[X]` produces a single array of four elements.

| Program | Input | Output |
|---------|-------|--------|
| `[.user, .projects[]]` | `{"user":"stedolan", "projects": ["jq", "wikiflow"]}` | `["stedolan", "jq", "wikiflow"]` |
| `[.[] | . * 2]` | `[1, 2, 3]` | `[2, 4, 6]` |

## Object Construction: `{}`

Construct objects (dictionaries). Keys are strings; values are any jq expression.

### Key Shorthand

If keys are identifier-like, quotes can be omitted:

```jq
{a: 42, b: 17}              # same as {"a": 42, "b": 17}
```

### Value Shorthand

`{user}` is shorthand for `{user: .user}` — the key name is also the field to read from input:

```jq
{user, title}               # same as {user: .user, title: .title}
```

### Variable Key Shorthand

`{$foo}` is shorthand for `{foo: $foo}` — the variable name becomes the key:

```jq
"f o o" as $foo | {$foo}     # → {"foo": "f o o"}
```

### Expression Keys

Non-literal keys (expressions, variable references) must be parenthesized:

```jq
{("a"+"b"): 59}             # key from expression
{(.user): .titles}          # key from field value
```

### Multiple Outputs

If a value expression produces multiple results, multiple objects are produced:

```jq
{user, title: .titles[]}     # produces one object per title
```

| Program | Input | Output |
|---------|-------|--------|
| `{user, title: .titles[]}` | `{"user":"stedolan","titles":["JQ Primer", "More JQ"]}` | `{"user":"stedolan", "title": "JQ Primer"}` then `{"user":"stedolan", "title": "More JQ"}` |
| `{(.user): .titles}` | `{"user":"stedolan","titles":["JQ Primer", "More JQ"]}` | `{"stedolan": ["JQ Primer", "More JQ"]}` |

### Variable Reference Keys

```jq
"f o o" as $foo | "b a r" as $bar | {$foo, $bar:$foo}
# → {"foo":"f o o","b a r":"f o o"}
```

Without a value (`$foo`), the variable name becomes the key and its value becomes the value. With a value (`$bar:$foo`), the left side uses the variable's **value** as the key and the right is the value.

## Type-Checking Selectors

These builtins select only inputs matching a type (produce the input if it matches, nothing otherwise):

| Builtin | Selects |
|---------|---------|
| `arrays` | arrays |
| `objects` | objects |
| `iterables` | arrays or objects |
| `booleans` | booleans |
| `numbers` | numbers |
| `normals` | normal numbers |
| `finites` | finite numbers |
| `strings` | strings |
| `nulls` | null |
| `values` | non-null values |
| `scalars` | non-iterables (booleans, numbers, strings, null) |

```jq
.[] | numbers               # only the numbers from an array
.[] | objects               # only the objects
```

| Program | Input | Output |
|---------|-------|--------|
| `.[]|numbers` | `[[],{},1,"foo",null,true,false]` | `1` |

See [builtins.md](builtins.md) for the full builtin reference.

## References

- [jq manual — Types and Values](https://jqlang.github.io/jq/manual/#types-and-values)
- For operators on these types, see [operators.md](operators.md)
- For type-checking builtins, see [builtins.md](builtins.md)
