# Format Strings and Escaping

The `@foo` syntax formats and escapes values — useful for building URLs, CSV, HTML, shell commands, and base64-encoded output. Format strings combine with string interpolation for powerful output generation.

> **Source of truth**: <https://jqlang.github.io/jq/manual/#format-strings-and-escaping>. For **string interpolation** `\(...)`, see [strings.md](strings.md).

## Available Formats

| Format | Escaping | Input Requirement |
|--------|----------|-------------------|
| `@text` | Calls `tostring` | Any |
| `@json` | Serializes as JSON | Any |
| `@html` | HTML/XML entity escaping | Any (string expected) |
| `@uri` | Percent-encoding for URIs | Any (string expected) |
| `@csv` | CSV row with quoted strings | Must be array |
| `@tsv` | TSV row with escape sequences | Must be array |
| `@sh` | POSIX shell quoting | String or array |
| `@base64` | Base64 encode (RFC 4648) | String |
| `@base64d` | Base64 decode (RFC 4648) | String |

## Format Details

### `@text`

Calls `tostring` — strings pass through, other types are JSON-encoded.

### `@json`

Serializes the input as a JSON text string. Differs from `@text`/`tostring` in that strings are JSON-quoted (with escaped quotes).

### `@html`

Maps characters to HTML entities:

| Character | Entity |
|-----------|--------|
| `<` | `&lt;` |
| `>` | `&gt;` |
| `&` | `&amp;` |
| `'` | `&apos;` |
| `"` | `&quot;` |

```jq
@html               # "This works if x < y" → "This works if x &lt; y"
```

### `@uri`

Percent-encodes all reserved URI characters as `%XX` sequences.

### `@csv`

Input must be an array. Renders as a CSV row: strings double-quoted, quotes escaped by repetition.

```jq
@csv                # ["a","b,c"] → "\"a\",\"b,c\""
```

### `@tsv`

Input must be an array. Renders as TSV (tab-separated). Fields separated by single tab (`0x09`). Special characters escaped:

| Character | Escape |
|-----------|--------|
| `\n` (LF, `0x0a`) | `\n` |
| `\r` (CR, `0x0d`) | `\r` |
| `\t` (TAB, `0x09`) | `\t` |
| `\` (backslash, `0x5c`) | `\\` |

### `@sh`

Escapes for POSIX shell use. If input is an array, output is space-separated quoted strings.

```jq
@sh                 # "O'Hara's Ale" → "'O'\''Hara'\''s Ale'"
@sh "echo \(.)"     # builds a shell command string
```

### `@base64`, `@base64d`

RFC 4648 base64 encode/decode. Note: if the decoded string is not UTF-8, results are undefined.

```jq
@base64             # "This is a message" → "VGhpcyBpcyBhIG1lc3NhZ2U="
@base64d            # "VGhpcyBpcyBhIG1lc3NhZ2U=" → "This is a message"
```

## Combining with String Interpolation

`@foo` can be followed by a string literal. The literal text is **not** escaped, but interpolations inside it **are** escaped:

```jq
@uri "https://www.google.com/search?q=\(.search)"
```

Input `{"search":"what is jq?"}`:

```
"https://www.google.com/search?q=what%20is%20jq%3F"
```

The `/`, `?`, `=` in the URL are not escaped (part of the literal), but the interpolated `what is jq?` is percent-encoded.

## Examples

| Program | Input | Output |
|---------|-------|--------|
| `@html` | `"This works if x < y"` | `"This works if x &lt; y"` |
| `@sh "echo \(.)"` | `"O'Hara's Ale"` | `"echo 'O'\''Hara'\''s Ale'"` |
| `@base64` | `"This is a message"` | `"VGhpcyBpcyBhIG1lc3NhZ2U="` |
| `@base64d` | `"VGhpcyBpcyBhIG1lc3NhZ2U="` | `"This is a message"` |

## References

- [jq manual — Format strings and escaping](https://jqlang.github.io/jq/manual/#format-strings-and-escaping)
- For string interpolation, see [strings.md](strings.md)
- For `tostring`/`tojson`, see [strings.md](strings.md) and [builtins.md](builtins.md)
