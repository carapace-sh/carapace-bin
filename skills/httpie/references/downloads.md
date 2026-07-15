# Download Mode

HTTPie's wget-like download mode — saving response bodies to files, resuming interrupted downloads, and piping while downloading.

## Overview

The `--download` (`-d`) flag makes HTTPie act like `wget`: response headers are printed to stderr, a progress bar is shown, and the response body is saved to a file.

```bash
$ http --download https://github.com/httpie/cli/archive/master.tar.gz
```

Output:

```
HTTP/1.1 200 OK
Content-Disposition: attachment; filename=httpie-master.tar.gz
Content-Length: 257336
Content-Type: application/x-gzip

Downloading 251.30 kB to "httpie-master.tar.gz"
Done. 251.30 kB in 2.73862s (91.76 kB/s)
```

## Implied Behaviors

`--download` implies several other behaviors:

| Behavior | Reason |
|----------|--------|
| `--follow` | Redirects are always followed during downloads |
| `--check-status` | Error HTTP status codes result in non-zero exit |
| Exit code `1` | If the body is not fully downloaded |

`Accept-Encoding` **cannot** be set with `--download`.

## Filename Determination

The output filename is determined by three methods in priority order:

| Priority | Source | Example |
|----------|--------|---------|
| 1 (highest) | `--output, -o` flag | `--output=file.zip` |
| 2 | `Content-Disposition` header | `attachment; filename=report.pdf` |
| 3 (lowest) | URL + Content-Type fallback | Derived from URL path and response MIME type |

The initial URL is always used as the basis for the fallback filename, even after redirects.

Leading dots are stripped from server-provided filenames. To prevent data loss, HTTPie adds a unique numerical suffix when a file already exists (unless `--output` is explicitly specified).

## Resuming Downloads (`--continue`, `-c`)

Resume an interrupted download with `--continue` (`-c`). Requires `--output` (`-o`) and server `Range` request support:

```bash
$ http -dco file.zip example.org/file
```

`-dco` is shorthand for `--download --continue --output`.

If the server doesn't support `Range` requests and `206 Partial Content`, the entire file is re-downloaded.

## Piping While Downloading

Redirect the response body to another program while still showing headers and progress:

```bash
$ http -d https://github.com/httpie/cli/archive/master.tar.gz | tar zxf -
```

## Redirected Output (Without `--download`)

When output is redirected (piped to another command or saved to a file), HTTPie uses different defaults:

- No formatting or colors (unless `--pretty` is specified)
- Only the response body is printed (unless output options are set)
- Binary data is not suppressed

```bash
# Download a file via shell redirection
$ http pie.dev/image/png > image.png

# Pipe to ImageMagick, then re-upload
$ http octodex.github.com/images/original.jpg | convert - -resize 25% - | http example.org/Octocats

# Force formatting when piping
$ http --pretty=all --verbose pie.dev/get | less -R
```

Helper function for paged output:

```bash
function httpless {
    http --pretty=all --print=hb "$@" | less -R
}
```

## Conditional Body Download

The response body is only downloaded from the server if it's part of the output. Headers are always downloaded. This optimization saves bandwidth when you only need headers (e.g., `--headers`). This doesn't affect `--output` or `--download`.

## Carapace Completion

The carapace completer completes `--output` with `carapace.ActionFiles()`. The `--download` and `--continue` flags are boolean and do not require value completion.

## References

- Source of truth: <https://httpie.io/docs/cli/downloads>
- Carapace completer: `completers/common/http_completer/cmd/root.go`

## Related Skills

- [output.md](output.md) — output formatting and `--output`
- [scripting.md](scripting.md) — exit codes and `--check-status`
