# Output Formatting and Printing

How HTTPie controls response output — what to print, how to format it, color themes, and the difference between terminal and redirected output.

## Print Options (`--print`, `-p`)

The `--print` flag takes a string of characters specifying what to output:

| Character | Stands For |
|-----------|------------|
| `H` | Request headers |
| `B` | Request body |
| `h` | Response headers |
| `b` | Response body |
| `m` | Response metadata (elapsed time) |

```bash
# Print only response headers and body
$ http --print=hb pie.dev/get

# Print request and response headers
$ http --print=Hh pie.dev/get
```

Characters can be combined in any order. The `--print` value is treated as a unique list.

### Shortcuts

| Flag | Equivalent | Description |
|------|-----------|-------------|
| `--headers` (`-h`) | `--print=h` | Response headers only |
| `--body` (`-b`) | `--print=b` | Response body only |
| `--meta` (`-m`) | `--print=m` | Response metadata only |
| `--verbose` (`-v`) | `--all --print=BHbh` | Request + response (with `--all` for intermediaries) |
| `--verbose` (`-vv`) | `--all --print=mBHbh` | Extra verbose (adds metadata) |

### `--all`

Shows intermediary requests/responses (redirects, auth attempts) in addition to the final response:

```bash
$ http --all --follow pie.dev/redirect/3
```

### `--stream` (`-S`)

Always stream the response body by line, behaving like `tail -f`:

```bash
$ http --stream pie.dev/stream/5
```

## Pretty Printing (`--pretty`)

Controls output processing:

| Value | Description |
|-------|-------------|
| `all` | Apply both formatting and colors (default for terminal) |
| `colors` | Apply only colors |
| `format` | Apply only formatting |
| `none` | No processing (default for redirected output) |

```bash
$ http --pretty=none pie.dev/get > raw.json
$ http --pretty=format pie.dev/get    # format but no colors
```

## Color Styles (`--style`, `-s`)

Controls the color theme. The default is `auto`.

### Available Styles

```bash
$ http --style=auto pie.dev/get
$ http --style=monokai pie.dev/get
$ http --style=fruity pie.dev/get
```

The carapace completer (`ActionStyles()` in `completers/common/http_completer/cmd/root.go`) completes these style values: `abap`, `algol`, `algol_nu`, `arduino`, `auto`, `autumn`, `borland`, `bw`, `colorful`, `default`, `emacs`, `friendly`, `fruity`, `igor`, `inkpot`, `lovelace`, `manni`, `material`, `monokai`, `murphy`, `native`, `paraiso-dark`, `paraiso-light`, `pastie`, `perldoc`, `rainbow_dash`, `rrt`, `sas`, `solarized`, `solarized-dark`, `solarized-light`, `stata`, `stata-dark`, `stata-light`, `tango`, `trac`, `vim`, `vs`, `xcode`, `zenburn`.

The official docs list additional Pie themes: `pie`, `pie-dark`, `pie-light`.

## Format Options (`--format-options`)

Fine-grained control over output formatting. Multiple options are comma-separated:

```bash
$ http --format-options=headers.sort:true,json.indent:2 pie.dev/get
```

| Option | Default | Description |
|--------|---------|-------------|
| `headers.sort` | `true` | Sort response headers alphabetically |
| `json.format` | `true` | Pretty-print JSON |
| `json.indent` | `4` | Number of spaces for JSON indentation |
| `json.sort_keys` | `true` | Sort JSON object keys |
| `xml.format` | `true` | Pretty-print XML |
| `xml.indent` | `2` | Number of spaces for XML indentation |

### Shortcuts

| Flag | Effect |
|------|--------|
| `--sorted` | Re-enable all sorting (`headers.sort:true,json.sort_keys:true`) |
| `--unsorted` | Disable all sorting (`headers.sort:false,json.sort_keys:false`) |

### Response Override Options

| Flag | Description |
|------|-------------|
| `--response-charset` | Override response encoding for display (e.g., `utf8`, `big5`) |
| `--response-mime` | Override response MIME type for coloring/formatting |

## Terminal vs Redirected Output

HTTPie uses different defaults depending on whether output goes to a terminal or is redirected:

| Aspect | Terminal | Redirected (pipe/file) |
|--------|----------|----------------------|
| Formatting | Formatted (`--pretty=all`) | No formatting (`--pretty=none`) |
| Colors | Enabled | Disabled |
| Default print | Response headers + body | Response body only |
| Binary data | Suppressed (note shown) | Not suppressed |

Override redirected defaults with `--pretty`:

```bash
# Force formatting when piping
$ http --pretty=all pie.dev/get | less -R
```

### Binary Data

When output goes to a terminal and the response body is binary, HTTPie shows:

```
NOTE: binary data not shown in terminal
```

The connection is closed as soon as binary content is detected. Redirect output to a file to capture binary data:

```bash
$ http pie.dev/image/png > image.png
```

## Response Metadata

The metadata (`m` print character) currently includes the total elapsed time (seconds between opening the connection and receiving the last byte). Color-coded with Pie themes.

## Output to File (`--output`, `-o`)

```bash
$ http --output=response.json pie.dev/get
```

Saves the response body to a file instead of stdout. Combined with `--download` for wget-like behavior. See [downloads.md](downloads.md).

## Quiet Mode (`--quiet`, `-q`)

Reduces output:

| Usage | Effect |
|-------|--------|
| `-q` | Suppress non-essential output |
| `-qq` | Also suppress warnings |

```bash
$ http -qq --check-status pie.dev/post enjoy='the silence'
```

## Carapace Completion

The carapace completer provides:

- `ActionPrintOptions()` — completes `H`, `B`, `h`, `b` (response metadata `m` not yet included) as a unique list
- `ActionStyles()` — completes all style values
- `ActionFormatOptions()` — nested `ActionMultiParts` completing format option keys and boolean values
- `--pretty` — completes `all`, `colors`, `format`, `none`

See `completers/common/http_completer/cmd/root.go` for implementation details.

## References

- Source of truth: <https://httpie.io/docs/cli> (Output options section)
- Carapace completer: `completers/common/http_completer/cmd/root.go` — `ActionPrintOptions()`, `ActionStyles()`, `ActionFormatOptions()`

## Related Skills

- [cli.md](cli.md) — full flag reference
- [downloads.md](downloads.md) — download mode and `--output`
