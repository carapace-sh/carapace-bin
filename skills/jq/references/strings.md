# String Functions

String manipulation builtins and string interpolation. For **regex-based string functions** (`test`, `match`, `sub`, `gsub`, `scan`, `split` with regex), see [regex.md](regex.md). For **format strings** (`@base64`, `@csv`, etc.), see [formats.md](formats.md).

> **Source of truth**: <https://jqlang.github.io/jq/manual/#builtin-operators-and-functions> (string-related entries).

## String Interpolation: `\(exp)`

Inside a string literal, `\(...)` evaluates the expression and interpolates the result:

```jq
"The input was \(.), which is one less than \(.+1)"
# input 42 → "The input was 42, which is one less than 43"
```

The expression inside `\(...)` is a full jq filter. The result is converted to a string (strings as-is, other types JSON-encoded).

| Program | Input | Output |
|---------|-------|--------|
| `"The input was \(.), which is one less than \(.+1)"` | `42` | `"The input was 42, which is one less than 43"` |

## JSON Conversion

### `tojson`, `fromjson`

`tojson` serializes as JSON text (strings get quotes — this differs from `tostring`). `fromjson` parses a JSON string.

```jq
tojson              # {"a":1} → "{\"a\":1}"
fromjson            # "{\"a\":1}" → {"a":1}
```

See [builtins.md](builtins.md) for the comparison table with `tostring`.

### `tostring`, `tonumber`

`tostring` — strings pass through; other types are JSON-encoded. `tonumber` — parses strings as numbers; numbers pass through.

See [builtins.md](builtins.md) for full details.

## Length

### `length`

For strings: number of Unicode codepoints (same as byte count only for pure ASCII).

### `utf8bytelength`

Number of bytes in UTF-8 encoding (differs from `length` for non-ASCII):

```jq
utf8bytelength       # "\u03bc" (μ) → 2
```

See [operators.md](operators.md) for `length` across all types.

## Splitting and Joining

### `split(str)` — string separator

Splits on a literal string separator:

```jq
split(", ")          # "a, b,c,d, e, " → ["a","b,c,d","e",""]
```

| Program | Input | Output |
|---------|-------|--------|
| `split(", ")` | `"a, b,c,d, e, "` | `["a","b,c,d","e",""]` |

For regex splitting, see [regex.md](regex.md) (`split(regex; flags)`).

### `join(str)`

Joins an array using the separator. Inverse of `split`. Numbers and booleans are converted to strings. `null` becomes empty string. Arrays and objects are not supported.

```jq
join(", ")           # ["a","b,c,d","e"] → "a, b,c,d, e"
```

| Program | Input | Output |
|---------|-------|--------|
| `join(", ")` | `["a","b,c,d","e"]` | `"a, b,c,d, e"` |
| `join(" ")` | `["a",1,2.3,true,null,false]` | `"a 1 2.3 true  false"` |

## Prefix and Suffix

### `startswith(str)`, `endswith(str)`

```jq
startswith("foo")   # true if input starts with "foo"
endswith("foo")     # true if input ends with "foo"
```

| Program | Input | Output |
|---------|-------|--------|
| `[.[]|startswith("foo")]` | `["fo", "foo", "barfoo", "foobar", "barfoob"]` | `[false, true, false, true, false]` |
| `[.[]|endswith("foo")]` | `["foobar", "barfoo"]` | `[false, true]` |

### `ltrimstr(str)`, `rtrimstr(str)`

Remove a prefix (`ltrimstr`) or suffix (`rtrimstr`) if present. No-op if the string doesn't start/end with the argument.

```jq
ltrimstr("foo")     # "foobar" → "bar", "barfoo" → "barfoo"
rtrimstr("foo")     # "barfoo" → "bar", "foobar" → "foobar"
```

| Program | Input | Output |
|---------|-------|--------|
| `[.[]|ltrimstr("foo")]` | `["fo", "foo", "barfoo", "foobar", "afoo"]` | `["fo","","barfoo","bar","afoo"]` |
| `[.[]|rtrimstr("foo")]` | `["fo", "foo", "barfoo", "foobar", "foob"]` | `["fo","","bar","foobar","foob"]` |

## Codepoint Operations

### `explode`, `implode`

`explode` converts a string to an array of codepoint numbers. `implode` is the inverse.

```jq
explode             # "foobar" → [102,111,111,98,97,114]
implode             # [65, 66, 67] → "ABC"
```

| Program | Input | Output |
|---------|-------|--------|
| `explode` | `"foobar"` | `[102,111,111,98,97,114]` |
| `implode` | `[65, 66, 67]` | `"ABC"` |

### `ascii_downcase`, `ascii_upcase`

Convert ASCII letters (a-z, A-Z) to the specified case. Non-ASCII characters are unchanged.

```jq
ascii_downcase      # "Hello World" → "hello world"
ascii_upcase        # "useful but not for é" → "USEFUL BUT NOT FOR é"
```

| Program | Input | Output |
|---------|-------|--------|
| `ascii_upcase` | `"useful but not for é"` | `"USEFUL BUT NOT FOR é"` |

## Searching in Strings

### `indices(s)`, `index(s)`, `rindex(s)`

`indices(s)` — all positions where `s` occurs. `index(s)` — first occurrence. `rindex(s)` — last occurrence.

| Program | Input | Output |
|---------|-------|--------|
| `indices(", ")` | `"a,b, cd, efg, hijk"` | `[3,7,12]` |
| `index(", ")` | `"a,b, cd, efg, hijk"` | `3` |
| `rindex(", ")` | `"a,b, cd, efg, hijk"` | `12` |

These also work on arrays — see [builtins.md](builtins.md).

## References

- [jq manual — String interpolation](https://jqlang.github.io/jq/manual/#string-interpolation-exp)
- [jq manual — Convert to/from JSON](https://jqlang.github.io/jq/manual/#convert-tofrom-json)
- For regex string functions (`test`, `match`, `sub`, `gsub`, `scan`), see [regex.md](regex.md)
- For format strings (`@base64`, `@csv`, `@sh`, etc.), see [formats.md](formats.md)
- For `tostring`/`tonumber`/`type`, see [builtins.md](builtins.md)
