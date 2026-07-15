# Configuration

HTTPie's configuration system — config file location, `config.json` format, `default_options`, `plugins_dir`, environment variables, and config directory resolution.

## Config File Location

### Default Paths

| Platform | Default Path |
|----------|-------------|
| Linux/macOS | `$XDG_CONFIG_HOME/httpie/config.json` (defaults to `~/.config/httpie/config.json`) |
| Windows | `%APPDATA%\httpie\config.json` |
| Legacy | `~/.httpie/config.json` (if `~/.httpie` directory exists) |

The config file does not exist by default — you must create it manually.

### Resolution Order (Unix-like)

1. If `$HTTPIE_CONFIG_DIR` is set, use that directory
2. If `~/.httpie` exists (legacy), use `~/.httpie/config.json`
3. Otherwise, use `$XDG_CONFIG_HOME/httpie/config.json` (defaults to `~/.config/httpie/config.json`)

### Finding the Config Directory

```bash
$ http --debug 2>&1 | grep config_dir
config_dir: '/home/user/.config/httpie'
```

## Config File Format

The config file is a JSON object with two keys:

### `default_options`

An array of CLI options applied to every invocation:

```json
{
    "default_options": [
        "--style=fruity"
    ]
}
```

Any HTTPie option can be included. Un-set a default for a single invocation with `--no-OPTION`:

```bash
# If config has "--style=fruity", override for one request:
$ http --no-style example.org
```

### `plugins_dir`

The directory where plugins are installed. HTTPie needs read/write access:

```json
{
    "plugins_dir": "/home/user/.local/share/httpie/plugins"
}
```

By default, plugins are stored under the config directory.

## Environment Variables

| Variable | Purpose | Default |
|----------|---------|---------|
| `HTTPIE_CONFIG_DIR` | Override the config directory location | Platform-specific (see above) |
| `XDG_CONFIG_HOME` | Base directory for XDG config home | `~/.config` |
| `HTTP_PROXY` | Proxy for HTTP requests | No proxy |
| `HTTPS_PROXY` | Proxy for HTTPS requests | No proxy |
| `NO_PROXY` | Comma-separated hosts to exclude from proxying | All requests use proxy |
| `ALL_PROXY` | Fallback proxy for all protocols | No fallback |
| `REQUESTS_CA_BUNDLE` | Path to CA bundle for certificate verification | System default |

### Example: Custom Config Directory

```bash
$ export HTTPIE_CONFIG_DIR=/tmp/httpie
$ http pie.dev/get
```

## Session Storage

Named sessions are stored in the `sessions` subdirectory of the config directory:

- `~/.config/httpie/sessions/<host>/<name>.json`

See [sessions.md](sessions.md) for session details.

## Proxies

### Environment Variable Proxies

```bash
$ export HTTP_PROXY=http://10.10.1.10:3128
$ export HTTPS_PROXY=https://10.10.1.10:1080
$ export NO_PROXY=localhost,127.0.0.1
$ http example.org
```

### `--proxy` Flag

The `--proxy` flag takes precedence over environment variables and can be specified multiple times:

```bash
$ http --proxy=http:http://user:pass@host:3128 \
      --proxy=https:https://user:pass@host:1080 \
      example.org
```

SOCKS proxy support:

```bash
$ http --proxy=http:socks5://user:pass@host:port \
      --proxy=https:socks5://user:pass@host:port \
      example.org
```

## Redirects

| Setting | Behavior |
|---------|---------|
| Default | Redirects are NOT followed (only first response shown) |
| `--follow` (`-F`) | Follow 30x Location redirects |
| `--max-redirects=N` | Limit redirects (default: 30) |
| `--all` | Show intermediary redirect responses |
| `307`/`308` | Method and body are reused |
| Other 3xx | Body-less GET is used for the redirect |

## Carapace Completion

The carapace completer does not complete `--proxy` flag values (free-form input). The `--default-scheme` flag is completed with `http://` and `https://`.

## References

- Source of truth: <https://httpie.io/docs/cli/config>
- Environment variables: <https://httpie.io/docs/cli> (Environment variables section)
- Carapace completer: `completers/common/http_completer/cmd/root.go`

## Related Skills

- [sessions.md](sessions.md) — session storage directory
- [ssl.md](ssl.md) — `REQUESTS_CA_BUNDLE`
- [plugins.md](plugins.md) — `plugins_dir`
