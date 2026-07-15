# Authentication

HTTPie's authentication mechanisms — built-in types, .netrc, password prompting, and plugin-based auth.

## Built-in Auth Types

| Mechanism | `--auth-type` (`-A`) | `--auth` (`-a`) value | Description |
|-----------|---------------------|----------------------|-------------|
| Basic | `basic` (default) | `USER:PASS` | HTTP Basic authentication |
| Digest | `digest` | `USER:PASS` | HTTP Digest authentication |
| Bearer | `bearer` | `TOKEN` | Bearer token authentication |

### Basic Auth

```bash
$ http -a username:password pie.dev/basic-auth/username/password
```

If only the username is provided, HTTPie securely prompts for the password:

```bash
$ http -a username pie.dev/basic-auth/username/password
# HTTPie prompts: password:
```

An empty password can be sent without prompting:

```bash
$ http -a username: pie.dev/basic-auth/username/password
```

Basic auth also works via URL syntax:

```bash
$ http username:password@pie.dev/basic-auth/username/password
```

### Digest Auth

```bash
$ http -A digest -a username:password pie.dev/digest-auth/username/password
```

### Bearer Auth

```bash
$ https -A bearer -a my-token api.example.com/me
```

## .netrc Support

HTTPie automatically reads credentials from `~/.netrc` (or `~/._netrc` on Windows) unless `--ignore-netrc` is set.

A `.netrc` entry:

```
machine pie.dev
  login username
  password secret
```

```bash
# Credentials from .netrc are used automatically
$ http pie.dev/basic-auth/username/secret

# Disable .netrc lookup
$ http --ignore-netrc pie.dev/basic-auth/username/secret
```

## Password Prompting

When `--auth` is provided with a username but no password (`-a username`), HTTPie prompts for the password interactively. This prevents passwords from appearing in shell history or process lists.

Prompted passwords are persisted in session files if `--session` is used. See [sessions.md](sessions.md).

## Auth Plugins

Additional authentication mechanisms are available as plugins. See [plugins.md](plugins.md) for installation and development.

| Plugin | Auth Type | Description |
|--------|----------|-------------|
| httpie-ntlm | ntlm | NT LAN Manager authentication |
| httpie-jwt-auth | jwt | JSON Web Token authentication |
| httpie-oauth1 | oauth1 | OAuth 1.0a |
| httpie-aws-auth | aws | AWS / Amazon S3 authentication |
| httpie-edgegrid | edgegrid | Akamai EdgeGrid authentication |
| httpie-hmac-auth | hmac | HMAC authentication |
| httpie-api-auth | api-auth | API authentication |
| httpie-negotiate | negotiate | SPNEGO (GSS Negotiate) |
| requests-hawk | hawk | Hawk authentication |

Usage with a plugin:

```bash
$ http --auth-type=ntlm --auth=domain\\user:pass example.org
```

## Carapace Completion

The carapace completer provides flag completion for `--auth-type`:

```go
"auth-type": carapace.ActionValues("basic", "digest"),
```

Note: `bearer` is missing from the completer's `--auth-type` values in the current implementation (`completers/common/http_completer/cmd/root.go:75`).

## References

- Source of truth: <https://httpie.io/docs/cli> (Authentication section)
- Auth plugins list: <https://httpie.io/docs/cli/auth-plugins>
- Carapace completer: `completers/common/http_completer/cmd/root.go`

## Related Skills

- [sessions.md](sessions.md) — how auth credentials persist in sessions
- [plugins.md](plugins.md) — plugin installation and development
