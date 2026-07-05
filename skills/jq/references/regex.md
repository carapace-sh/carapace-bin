# Regular Expressions

jq uses the [Oniguruma](https://github.com/kkos/oniguruma) regex library with the "Perl NG" (Perl with named groups) flavor. Regex filters accept a regex string, optional flags, and work on string inputs.

> **Source of truth**: <https://jqlang.github.io/jq/manual/#regular-expressions>. For **non-regex string functions** (`split` with string, `ltrimstr`, etc.), see [strings.md](strings.md).

## Regex Filter Patterns

All regex filters follow these calling conventions:

```
STRING | FILTER(REGEX)
STRING | FILTER(REGEX; FLAGS)
STRING | FILTER([REGEX])
STRING | FILTER([REGEX, FLAGS])
```

- `STRING`, `REGEX`, and `FLAGS` are jq strings subject to string interpolation
- `REGEX` must evaluate to a valid regex string after interpolation
- `FILTER` is `test`, `match`, `capture`, `scan`, `split` (regex form), `sub`, `gsub`, or `splits`

### Escaping

Since the regex is a JSON string, regex metacharacters that need backslashes must be **double-escaped**:

```jq
test("a\\sb")       # regex \s (whitespace) — written as "\\s" in jq string
test("\\d+")        # regex \d (digit) — written as "\\d" in jq string
```

## Flags

A string of one or more flag characters:

| Flag | Name | Effect |
|------|------|--------|
| `g` | Global | Find all matches, not just the first |
| `i` | Case insensitive | Ignore case |
| `m` | Multi-line | `.` matches newlines |
| `n` | Ignore empty | Ignore empty matches |
| `p` | Both | `s` and `m` modes enabled |
| `s` | Single-line | `^` → `\A`, `$` → `\Z` |
| `l` | Longest | Find longest possible matches |
| `x` | Extended | Ignore whitespace and comments in regex |

```jq
test("a b c # spaces are ignored"; "ix")    # extended + case-insensitive
```

### Inline Flags

Some flags can be set within the regex itself:

```jq
test("(?i)te(?-i)st")     # case-insensitive for "te", sensitive for "st"
# matches: "test", "TEst", but NOT "teST" or "TEST"
```

## `test(val)`, `test(regex; flags)`

Returns `true`/`false` — whether the regex matches. Does not return match objects.

| Program | Input | Output |
|---------|-------|--------|
| `test("foo")` | `"foo"` | `true` |
| `.[] | test("a b c # spaces are ignored"; "ix")` | `["xabcd", "ABC"]` | `true`, `true` |

## `match(val)`, `match(regex; flags)`

Outputs an object for each match. The `g` flag produces multiple outputs (one per match).

### Match Object Fields

| Field | Type | Description |
|-------|------|-------------|
| `offset` | number | Offset in UTF-8 codepoints from start of input |
| `length` | number | Length in UTF-8 codepoints |
| `string` | string | The matched string |
| `captures` | array | Capturing group objects |

### Capture Group Object Fields

| Field | Type | Description |
|-------|------|-------------|
| `offset` | number | Offset (or `-1` if group didn't match) |
| `length` | number | Length |
| `string` | string | Captured string (or `null`) |
| `name` | string | Named group name (or `null`) |

| Program | Input | Output |
|---------|-------|--------|
| `match("(abc)+"; "g")` | `"abc abc"` | `{"offset":0,"length":3,"string":"abc","captures":[...]}`, `{"offset":4,"length":3,"string":"abc","captures":[...]}` |
| `match("foo")` | `"foo bar foo"` | `{"offset":0,"length":3,"string":"foo","captures":[]}` |
| `match(["foo", "ig"])` | `"foo bar FOO"` | `{"offset":0,...}`, `{"offset":8,...}` |
| `match("foo (?<bar123>bar)? foo"; "ig")` | `"foo bar foo foo  foo"` | *(two match objects with named capture `bar123`)* |

## `capture(val)`, `capture(regex; flags)`

Collects named captures into a JSON object — capture name as key, matched string as value.

```jq
capture("(?<a>[a-z]+)-(?<n>[0-9]+)")
```

| Program | Input | Output |
|---------|-------|--------|
| `capture("(?<a>[a-z]+)-(?<n>[0-9]+)")` | `"xyzzy-14"` | `{"a": "xyzzy", "n": "14"}` |

## `scan(regex)`, `scan(regex; flags)`

Emits a stream of non-overlapping matching substrings. If the regex has capturing groups, emits arrays of captured strings instead. Empty stream if no match.

```jq
[scan("c")]           # collect all matches into array
scan("(a+)(b+)")      # emits arrays: ["a","b"], ...
```

| Program | Input | Output |
|---------|-------|--------|
| `scan("c")` | `"abcdefabc"` | `"c"`, `"c"` |
| `scan("(a+)(b+)")` | `"abaabbaaabbb"` | `["a","b"]`, `["aa","bb"]`, `["aaa","bbb"]` |

## `split(regex; flags)`

Splits on regex matches. When called with a **single argument**, `split` splits on a literal string, not a regex (see [strings.md](strings.md)).

| Program | Input | Output |
|---------|-------|--------|
| `split(", *"; null)` | `"ab,cd, ef"` | `["ab","cd","ef"]` |

## `splits(regex)`, `splits(regex; flags)`

Same as `split` but returns a **stream** instead of an array.

| Program | Input | Output |
|---------|-------|--------|
| `splits(", *")` | `"ab,cd,   ef, gh"` | `"ab"`, `"cd"`, `"ef"`, `"gh"` |

## `sub(regex; tostring)`, `sub(regex; tostring; flags)`

Replace the **first** match with `tostring` after interpolation. Named captures are available as a JSON object to `tostring`:

```jq
sub("(?<x>[a-z]+)"; "Z\(.x)"; "g")    # reference capture group "x"
```

`tostring` can be a stream of strings, each producing a separate output.

| Program | Input | Output |
|---------|-------|--------|
| `sub("[^a-z]*(?<x>[a-z]+)"; "Z\(.x)"; "g")` | `"123abc456def"` | `"ZabcZdef"` |
| `[sub("(?<a>.)"; "\(.a|ascii_upcase)", "\(.a|ascii_downcase)")]` | `"aB"` | `["AB","aB"]` |

## `gsub(regex; tostring)`, `gsub(regex; tostring; flags)`

Like `sub` but replaces **all** non-overlapping matches. If the second argument is a stream, `gsub` produces a corresponding stream of output strings.

| Program | Input | Output |
|---------|-------|--------|
| `gsub("(?<x>.)[^a]*"; "+\(.x)-")` | `"Abcabc"` | `"+A-+a-"` |
| `[gsub("p"; "a", "b")]` | `"p"` | `["a","b"]` |

## References

- [jq manual — Regular expressions](https://jqlang.github.io/jq/manual/#regular-expressions)
- [Oniguruma syntax](https://github.com/kkos/oniguruma/blob/master/doc/RE)
- [Oniguruma syntax flavors](https://github.com/kkos/oniguruma/blob/master/doc/SYNTAX.md)
- For non-regex string functions, see [strings.md](strings.md)
- For string interpolation syntax `\(...)`, see [strings.md](strings.md)
