# Request Items

How HTTPie uses positional key/value pairs to specify HTTP headers, URL parameters, data fields, and file uploads on the command line.

## Separators

Request items are key/value pairs separated by specific characters. The separator determines the item type:

| Separator | Type | Example | Description |
|-----------|------|---------|-------------|
| `:` | HTTP Headers | `X-API-Token:123` | Arbitrary HTTP request header |
| `==` | URL Parameters | `search==httpie` | Appends querystring parameter to URL (auto-URL-escaped) |
| `=` | Data Fields | `name=John` | Serialized as JSON (default), form-encoded (`--form`), or multipart (`--multipart`) |
| `:=` | Raw JSON Fields | `age:=29`, `married:=false` | Non-string JSON: Boolean, Number, Object, Array |
| `@` | File Upload Fields | `cv@~/CV.pdf` | Only with `--form` or `--multipart` |
| `=@` | Embed File Content | `description=@files/text.txt` | Reads file content as a data field value |
| `:=@` | Embed JSON File | `bookmarks:=@files/data.json` | Embeds file content as raw JSON |
| `:@` | Header from File | `X-Data:@files/text.txt` | Reads header value from file |
| `==@` | Query Param from File | `token==@files/text.txt` | Reads query parameter value from file |

### Separator Priority

When multiple separators could match, HTTPie uses the one that appears **earliest** in the string. The separators are checked in this order (first match wins by position, not by this list order):

```
==, =@, @, =, :=@, :=, :
```

The carapace completer's `determineSeparator()` function in `completers/common/http_completer/cmd/root.go` checks these separators by earliest index in the string.

## Escaping

Use `\` to escape characters that would otherwise be interpreted as separators:

```bash
# Sends data field with key "foo=" and value "bar"
$ http pie.dev/post foo\==bar
```

Items starting with `-` must be placed after `--` to prevent confusion with flags:

```bash
$ http pie.dev/post -- -name-starting-with-dash=foo
```

## Nested JSON Syntax

Data fields support nested JSON via bracket notation:

| Syntax | Resulting JSON |
|--------|---------------|
| `user[name]=John` | `{"user": {"name": "John"}}` |
| `user[age]:=29` | `{"user": {"age": 29}}` |
| `items[]=a` | `{"items": ["a"]}` |
| `items[]=a items[]=b` | `{"items": ["a", "b"]}` |
| `items[0]=first` | `{"items": ["first"]}` |
| `[0][key]=value` | `[{"key": "value"}]` (top-level array, omit starting key) |

Backslash escapes brackets when you need a literal `[` or `]` in the key name:

```bash
# Key is literally "foo[bar]"
$ http pie.dev/post foo\[bar\]:=1
```

## Content Type Behavior

The separator `=` behaves differently depending on the content type flag:

| Flag | Data field serialization |
|------|------------------------|
| `--json` (default) | JSON object: `name=John` → `{"name": "John"}` |
| `--form` | Form-encoded: `name=John` → `name=John` |
| `--multipart` | Multipart form data |
| `--raw` | Raw data, no request item processing |

## Combining stdin and Request Items

Data can be piped via stdin **or** specified via request items, but not both:

```bash
# This works — data from stdin
$ echo '{"name": "John"}' | http pie.dev/post

# This also works — data from request items
$ http pie.dev/post name=John

# This does NOT work — cannot combine both
$ echo '{"name": "John"}' | http pie.dev/post age:=29
```

## Carapace Completion

The carapace completer (`completers/common/http_completer/cmd/root.go`) provides intelligent completion for request items via `ActionRequestItem()`:

1. **No separator typed yet**: Offers separator characters (styled in blue) and HTTP header names (from `http.ActionRequestHeaderNames()`) with `:` suffix.
2. **Separator detected**: Uses `ActionMultiParts` with the detected separator:
   - `:` → completes header values via `http.ActionRequestHeaderValues()`
   - `=@`, `:=@`, `@` → completes file paths via `carapace.ActionFiles()`
   - `=`, `:=`, `==` → no value completion (free-form input)

The `ActionSeparators()` function provides the separator choices with descriptions:

| Value | Description |
|-------|-------------|
| `:` | HTTP headers |
| `==` | URL parameters to be appended to the request URI |
| `=` | Data fields to be serialized into a JSON object |
| `:=` | Non-string JSON data fields (only with --json, -j) |
| `@` | Form file fields (only with --form or --multipart) |
| `=@` | A data field like '=', but takes a file path and embeds its content |
| `:=@` | A raw JSON field like ':=', but takes a file path and embeds its content |

## References

- Source of truth: <https://httpie.io/docs/cli> (Request items section)
- Carapace completer: `completers/common/http_completer/cmd/root.go` — `ActionRequestItem()`, `determineSeparator()`, `ActionSeparators()`
- Shared header actions: `pkg/actions/net/http/http.go` — `ActionRequestHeaderNames()`, `ActionRequestHeaderValues()`

## Related Skills

- [cli.md](cli.md) — full flag reference
- [output.md](output.md) — how response output is formatted
