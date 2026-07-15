# Scripting and Exit Codes

Using HTTPie in shell scripts, automation, and pipelines — exit codes, `--check-status`, stdin handling, and best practices.

## Exit Codes

Without `--check-status`, HTTPie exits with `0` unless a network or fatal error occurs. With `--check-status`, HTTP status codes map to exit codes:

| Exit Code | Meaning |
|-----------|---------|
| `0` | Success (no errors, no bad status codes) |
| `1` | General error (e.g., download body not fully downloaded) |
| `2` | Request timed out |
| `3` | HTTP 3xx Redirection (when `--follow` is not set) |
| `4` | HTTP 4xx Client Error |
| `5` | HTTP 5xx Server Error |
| `6` | Exceeded `--max-redirects=N` redirects |

## Shell Script Example

```bash
#!/bin/bash

if http --check-status --ignore-stdin --timeout=2.5 HEAD pie.dev/get &> /dev/null; then
    echo 'OK!'
else
    case $? in
        2) echo 'Request timed out!' ;;
        3) echo 'Unexpected HTTP 3xx Redirection!' ;;
        4) echo 'HTTP 4xx Client Error!' ;;
        5) echo 'HTTP 5xx Server Error!' ;;
        6) echo 'Exceeded --max-redirects=<n> redirects!' ;;
        *) echo 'Other Error!' ;;
    esac
fi
```

## Best Practices for Automation

### Use `--ignore-stdin`

In non-interactive contexts (cron jobs, scripts), stdin is not connected to a terminal. HTTPie assumes piped input contains a request body and will **hang** waiting for data. Always use `--ignore-stdin` in scripts, unless you are intentionally piping data:

```bash
# Correct — prevents hanging
$ http --ignore-stdin pie.dev/get

# Correct — piping data intentionally
$ echo '{"name": "John"}' | http pie.dev/post
```

### Use `--timeout`

Prevent unresponsive scripts by setting a connection timeout:

```bash
$ http --timeout=2.5 pie.dev/get
```

### Recommended Flags for Scripts

```bash
http --check-status --ignore-stdin --timeout=<seconds> <METHOD> <URL>
```

### Suppress All Output

Use `-qq` (quiet twice) with `--check-status` for completely silent operation:

```bash
$ http -qq --check-status pie.dev/post enjoy='the silence'
```

## Piping Data Into HTTPie

### From a File

```bash
$ http PUT pie.dev/put X-API-Token:123 < files/data.json
```

### From Another Command

```bash
$ grep '401 Unauthorized' /var/log/httpd/error_log | http POST pie.dev/post
```

### Bash Here String

```bash
$ http pie.dev/post <<<'{"name": "John"}'
```

### Chaining HTTP Requests

```bash
$ http GET https://api.github.com/repos/httpie/cli | http POST pie.dev/post
```

### Interactive Paste

```bash
$ cat | http POST pie.dev/post
<paste data>
^D
```

Data from stdin **cannot** be combined with data fields specified on the command line.

## `--raw` Flag

Pass raw request data without processing:

```bash
$ http --raw='{"name": "John"}' pie.dev/post
```

## Redirected Output Defaults

When output is piped or redirected, HTTPie changes defaults for convenience:

| Aspect | Terminal | Redirected |
|--------|----------|-----------|
| Formatting | Enabled (`--pretty=all`) | Disabled (`--pretty=none`) |
| Colors | Enabled | Disabled |
| Default content | Headers + body | Body only |
| Binary data | Suppressed | Not suppressed |

Override with `--pretty`:

```bash
$ http --pretty=all pie.dev/get | less -R
```

## Download Mode and Exit Codes

`--download` implies `--check-status` and `--follow`. If the body is not fully downloaded, HTTPie exits with code `1`. See [downloads.md](downloads.md).

## References

- Source of truth: <https://httpie.io/docs/cli/scripting>
- Carapace completer: `completers/common/http_completer/cmd/root.go`

## Related Skills

- [cli.md](cli.md) — full flag reference
- [downloads.md](downloads.md) — download mode and `--check-status`
- [output.md](output.md) — redirected output behavior
