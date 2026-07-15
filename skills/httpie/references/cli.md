# CLI Flags and Invocation

Complete reference for HTTPie's command-line interface — flags, options, positional arguments, and the difference between `http` and `https`.

## Invocation

```
http  [flags] [METHOD] URL [REQUEST_ITEM ...]
https [flags] [METHOD] URL [REQUEST_ITEM ...]
```

### Positional Arguments

| Argument | Required | Description |
|----------|----------|-------------|
| `METHOD` | Optional | HTTP method (GET, POST, PUT, DELETE, PATCH, HEAD, OPTIONS, etc.). Defaults to GET if no data items are present, or POST if data items are present. |
| `URL` | Required | Target URL. Scheme defaults to `http://` for `http` and `https://` for `https`. See [url-syntax.md](url-syntax.md) for shortcuts. |
| `REQUEST_ITEM` | Optional | Key/value pairs for headers, params, data, and file uploads. See [request-items.md](request-items.md). |

### `http` vs `https`

HTTPie installs two executables that share all flags and syntax. The only difference is the default URL scheme:

| Command | Default Scheme |
|---------|---------------|
| `http` | `http://` |
| `https` | `https://` |

```bash
$ http example.org     # → http://example.org
$ https example.org    # → https://example.org
```

Override the default scheme with `--default-scheme`:

```bash
$ http --default-scheme=https example.org   # → https://example.org
```

## Content Type Options

| Flag | Short | Description |
|------|-------|-------------|
| `--json` | `-j` | (Default) Serialize data as JSON. Sets `Content-Type: application/json`. |
| `--form` | `-f` | Serialize data as form fields (`application/x-www-form-urlencoded`), or `multipart/form-data` if files are included. |
| `--multipart` | | Always use `multipart/form-data`, even without files. |
| `--boundary` | | Custom boundary string for `multipart/form-data` requests. |
| `--raw` | | Pass raw request data without processing. |

## Content Processing Options

| Flag | Short | Description |
|------|-------|-------------|
| `--compress` | `-x` | Compress request body with Deflate. Sets `Content-Encoding: deflate`. Use twice (`-xx`) to force compression even if ratio is negative. |

## Output Options

| Flag | Short | Description |
|------|-------|-------------|
| `--print` | `-p` | String specifying what to output: `H` (request headers), `B` (request body), `h` (response headers), `b` (response body), `m` (response metadata). |
| `--headers` | `-h` | Print only response headers (shortcut for `--print=h`). |
| `--body` | `-b` | Print only response body (shortcut for `--print=b`). |
| `--meta` | `-m` | Print only response metadata (shortcut for `--print=m`). |
| `--verbose` | `-v` | Show request and response (`--all --print=BHbh`). Use `-vv` for extra verbose (adds metadata). |
| `--all` | | Show intermediary requests/responses (redirects, auth attempts). |
| `--stream` | `-S` | Always stream the response body by line (like `tail -f`). |
| `--output` | `-o` | Save output to FILE instead of stdout. |
| `--download` | `-d` | Download body to a file (wget-like). See [downloads.md](downloads.md). |
| `--continue` | `-c` | Resume an interrupted download (requires `--output`). |
| `--quiet` | `-q` | Reduce output. Use `-qq` to suppress warnings too. |

For formatting options (--pretty, --style, --format-options, etc.), see [output.md](output.md).

## Network Options

| Flag | Short | Description |
|------|-------|-------------|
| `--offline` | | Build the request and print it but don't actually send it. |
| `--proxy` | | String mapping protocol to proxy URL (e.g., `http:http://foo.bar:3128`). Can specify multiple. |
| `--follow` | `-F` | Follow 30x Location redirects. |
| `--max-redirects` | | Maximum redirects to follow (default: 30). |
| `--max-headers` | | Maximum response headers to read (default: 0 = no limit). |
| `--timeout` | | Connection timeout in seconds (default: 0 = no timeout). |
| `--check-status` | | Exit with error status code if server returns 3xx/4xx/5xx. See [scripting.md](scripting.md). |
| `--path-as-is` | | Bypass dot segment (`/../` or `/./`) URL squashing. |
| `--chunked` | | Enable chunked transfer encoding (`Transfer-Encoding: chunked`). |
| `--ignore-stdin` | `-I` | Do not attempt to read stdin. |

For SSL/TLS options, see [ssl.md](ssl.md).

## Authentication Options

| Flag | Short | Description |
|------|-------|-------------|
| `--auth` | `-a` | Authentication credentials (`USER[:PASS]` or `TOKEN`). |
| `--auth-type` | `-A` | Auth mechanism: `basic` (default), `digest`, `bearer`, or plugin name. |
| `--ignore-netrc` | | Ignore credentials from `.netrc`. |

See [auth.md](auth.md) for full details.

## Session Options

| Flag | Description |
|------|-------------|
| `--session` | Create, or reuse and update a session. |
| `--session-read-only` | Create or read a session without updating it. |

See [sessions.md](sessions.md) for full details.

## Troubleshooting Options

| Flag | Description |
|------|-------------|
| `--help` | Show help message and exit. |
| `--manual` | Show full manual with pager. |
| `--version` | Show version and exit. |
| `--traceback` | Print exception traceback if one occurs. |
| `--default-scheme` | Default URL scheme if not specified in the URL. |
| `--debug` | Print debug information (includes `config_dir` location). |

## Un-setting Options

Every `--option` has a corresponding `--no-option` that reverts it to default. This applies to config defaults too:

```bash
# Override a default_options entry from config.json
$ http --no-style example.org
$ http --no-session example.org
```

## Carapace Completer

The `http` and `https` completers live in carapace-bin:

- `completers/common/http_completer/` — native Go completer
- `completers/common/https_completer/` — bridges to `http` via `bridge.ActionCarapaceBin("http")`

The completer uses shared actions from `pkg/actions/net/http/` for request header names, header values, media types, request methods, and status codes. Flag completions include `--auth-type` (basic, digest), `--ssl` (ssl2.3, tls1, tls1.1, tls1.2), `--pretty` (all, colors, format, none), `--verify` (yes, no), and file completions for `--cert`, `--cert-key`, `--output`.

The `ActionRequestItem()` function in the completer handles positional completion of request items by detecting the separator character and offering appropriate completions (header names, file paths, etc.).

## References

- Source of truth: <https://httpie.io/docs/cli>
- Carapace completer: `completers/common/http_completer/cmd/root.go`
- Shared HTTP actions: `pkg/actions/net/http/`

## Related Skills

- [request-items.md](request-items.md) — request item separators and syntax
- [output.md](output.md) — output formatting and printing
- [url-syntax.md](url-syntax.md) — URL shortcuts and defaults
