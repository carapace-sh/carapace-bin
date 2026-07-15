# Session Management

HTTPie's persistent session system — named sessions, anonymous sessions, read-only sessions, cookie persistence, and session upgrades.

## Overview

By default, every HTTPie request is independent. Sessions allow data (headers, auth, cookies) to persist between requests to the same host.

```bash
# Create a session with auth + custom header
$ http --session=user1 -a user1:password pie.dev/get X-Foo:Bar

# Reuse the session — auth + headers are automatically applied
$ http --session=user1 pie.dev/get
```

## What Is Persisted

| Data | Persisted? | Notes |
|------|-----------|-------|
| Custom HTTP headers | Yes | Except headers starting with `Content-` or `If-` |
| Authentication credentials | Yes | Including prompted passwords |
| Cookies (from `Set-Cookie`) | Yes | Server-sent cookies are stored |
| Cookies (from command line) | Yes | Manually specified via `Cookie:name=value` |

All session data is stored in **plain text JSON**. Anyone with file access can read credentials and cookies.

## Named Sessions

Named sessions are identified by a name string and scoped to a host:

```bash
# Create session "user1" for pie.dev
$ http --session=user1 -a user1:password pie.dev/get

# Create a different session for the same host
$ http --session=user2 -a user2:password pie.dev/get

# Reuse session "user1"
$ http --session=user1 pie.dev/get
```

### Storage Location

| Platform | Path |
|----------|------|
| Linux/macOS | `~/.config/httpie/sessions/<host>/<name>.json` |
| Windows | `%APPDATA%\httpie\sessions\<host>\<name>.json` |

## Anonymous Sessions

Instead of a name, specify a file path. This enables cross-host session reuse:

```bash
# Create an anonymous session
$ http --session=/tmp/session.json example.org

# Reuse it for a different host
$ http --session=/tmp/session.json admin.example.org
```

The path must contain at least one `/` (use `--session=./session.json`, not `--session=session.json`), otherwise HTTPie treats it as a named session.

You can also reference a previously created named session by its path:

```bash
$ http --session=~/.config/httpie/sessions/another.example.org/test.json example.org
```

## Read-Only Sessions

`--session-read-only` creates a session if it doesn't exist, but never updates it:

```bash
# Creates the session file
$ http --session-read-only=./ro-session.json pie.dev/headers Custom-Header:orig-value

# Does NOT update the session file
$ http --session-read-only=./ro-session.json pie.dev/headers Custom-Header:new-value
```

## Cookie Persistence

### Cookie Storage Priority

When multiple sources provide the same cookie, priority is:

1. **Response** (highest) — Cookies received via `Set-Cookie` header
2. **Command line** — Cookies specified via `Cookie:name=value`
3. **Session file** (lowest) — Pre-existing cookies in the session JSON

### Host-Based Cookie Policy

Cookies in session files have a `domain` field binding them to a host. A cookie with `"domain": "pie.dev"` is only sent to `pie.dev` and its subdomains.

Set `"domain": null` for unbound cookies that are sent to all hosts (including across redirect chains):

```json
{
    "cookies": [
        {
            "domain": null,
            "name": "unbound-cookie",
            "value": "send-me-to-any-host"
        }
    ]
}
```

### Cookie Expiration

- When the server expires a cookie, HTTPie removes it from the session file.
- When a cookie in a session file has expired, HTTPie removes it before sending the next request.

### Manual Cookie Format in Session File

```json
{
    "cookies": {
        "foo": {
            "expires": null,
            "path": "/",
            "secure": false,
            "value": "bar"
        }
    }
}
```

## Session Upgrade

HTTPie may change the session file format between versions. When an obsolete format is detected, HTTPie shows a warning.

### Upgrade All Sessions

```bash
$ httpie cli sessions upgrade-all
Upgraded 'api_auth' @ 'pie.dev' to v3.1.0
Upgraded 'login_cookies' @ 'httpie.io' to v3.1.0
```

### Upgrade a Single Named Session

```bash
$ httpie cli sessions upgrade pie.dev api_auth
Upgraded 'api_auth' @ 'pie.dev' to v3.1.0
```

### Upgrade an Anonymous Session

```bash
$ httpie cli sessions upgrade pie.dev ./session.json
Upgraded 'session.json' @ 'pie.dev' to v3.1.0
```

### `--bind-cookies`

Available for both `upgrade` and `upgrade-all`. Binds previously unbound cookies to the session's host:

```bash
$ httpie cli sessions upgrade pie.dev api_auth --bind-cookies
```

## Carapace Completion

The carapace completer currently completes `--session` and `--session-read-only` with `carapace.ActionFiles()` as a placeholder (session names are not yet dynamically completed). See `completers/common/http_completer/cmd/root.go:87-88`.

## References

- Source of truth: <https://httpie.io/docs/cli/sessions>
- Carapace completer: `completers/common/http_completer/cmd/root.go`
- Session config directory: see [config.md](config.md)

## Related Skills

- [auth.md](auth.md) — how auth credentials interact with sessions
- [config.md](config.md) — session storage directory location
