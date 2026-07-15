---
name: httpie
description: >
  Use when working with HTTPie (http/https CLI) — the command-line HTTP client for the
  API era. Covers request items syntax (headers, params, data fields, file uploads),
  authentication (basic, digest, bearer, plugins), session management, output formatting
  and printing, SSL/TLS configuration, download mode, configuration files, environment
  variables, plugin system, exit codes, scripting, and URL syntax. Triggers on:
  "httpie", "http command", "https command", "http --json", "http --form",
  "http --download", "http --session", "http --auth", "http --auth-type",
  "http --print", "http --verbose", "http --follow", "http --offline",
  "http --check-status", "http --proxy", "http --ssl", "http --cert",
  "http --verify", "http --format-options", "http --pretty", "http --style",
  "http --stream", "http --boundary", "http --multipart", "http --raw",
  "httpie config", "HTTPIE_CONFIG_DIR", "httpie cli plugins", "request items",
  "data fields", "URL parameters", "HTTP headers", "raw JSON fields",
  "file upload fields", "named session", "anonymous session", "read-only session",
  "Content-Disposition", "httpie-unixsocket", "httpie plugins", "AuthPlugin",
  "TransportPlugin", "ConverterPlugin", "FormatterPlugin".
user-invocable: true
---

# HTTPie (`http`/`https`) In-Depth Reference

Comprehensive reference for [HTTPie](https://httpie.io/) — a command-line HTTP client designed for the API era. Covers the full CLI surface, request items syntax, authentication, sessions, output formatting, SSL/TLS, downloads, configuration, plugins, scripting, and the carapace-bin completer.

## Data Flow

```
http [flags] [METHOD] URL [REQUEST_ITEM ...]
  → parse flags & default_options from config.json
    → resolve URL (default scheme, localhost shortcuts, path normalization)
      → build request (method, headers, body from request items)
        → apply auth (basic/digest/bearer/plugin, .netrc, session)
          → send request (proxy, SSL/TLS, redirects, timeout)
            → receive response
              → format output (pretty, style, print options, stream)
                → stdout / file (--output) / download (--download)
```

## Sub-Resources

Load the reference that matches your task. When in doubt, load multiple references.

| Keywords | Reference |
|----------|----------|
| invocation, command line, flags, options, --json, -j, --form, -f, --multipart, --boundary, --raw, --compress, -x, --verbose, -v, --all, --quiet, -q, --debug, --traceback, --help, --version, --offline, --ignore-stdin, -I, --timeout, --max-headers, --max-redirects, --check-status, --path-as-is, --chunked, http vs https, METHOD, GET, POST, PUT, DELETE, PATCH, HEAD, positional arguments | [references/cli.md](references/cli.md) |
| request items, request item, separator, header, :, URL parameter, ==, data field, =, raw JSON, :=, file upload, @, embed file, =@, :=@, :@, ==@, escaping, backslash, --, nested JSON, key[subkey], array, escaping brackets, Content-Type, form data, JSON object | [references/request-items.md](references/request-items.md) |
| authentication, auth, --auth, -a, --auth-type, -A, basic, digest, bearer, token, .netrc, --ignore-netrc, password prompt, auth plugins, NTLM, OAuth, JWT, HMAC, AWS, EdgeGrid, SPNEGO, Hawk | [references/auth.md](references/auth.md) |
| session, --session, --session-read-only, named session, anonymous session, session file, cookie persistence, Set-Cookie, cookie domain, unbound cookie, session upgrade, sessions upgrade-all, --bind-cookies, session storage, session JSON format | [references/sessions.md](references/sessions.md) |
| output, --print, -p, --headers, -h, --body, -b, --meta, -m, --verbose, -v, --all, --stream, -S, --pretty, --style, -s, --sorted, --unsorted, --format-options, --response-charset, --response-mime, --output, -o, binary data, redirected output, terminal output, format options, headers.sort, json.format, json.indent, json.sort_keys, xml.format, xml.indent, Pie theme, color style | [references/output.md](references/output.md) |
| SSL, TLS, --verify, --ssl, --ciphers, --cert, --cert-key, --cert-key-pass, certificate verification, CA bundle, REQUESTS_CA_BUNDLE, protocol version, ssl2.3, tls1, tls1.1, tls1.2, tls1.3, OpenSSL cipher list | [references/ssl.md](references/ssl.md) |
| download, --download, -d, --continue, -c, --output, -o, Content-Disposition, filename, resume, Range, 206 Partial Content, piping download, wget-like, progress bar | [references/downloads.md](references/downloads.md) |
| configuration, config file, config.json, default_options, plugins_dir, HTTPIE_CONFIG_DIR, XDG_CONFIG_HOME, legacy config, ~/.httpie, config directory resolution, --no-OPTION, unsetting options, http --debug, config_dir | [references/config.md](references/config.md) |
| plugins, plugin manager, httpie cli plugins, install, list, upgrade, uninstall, AuthPlugin, TransportPlugin, ConverterPlugin, FormatterPlugin, PluginManager, entry points, httpie-unixsocket, httpie-ntlm, httpie-jwt-auth, httpie-oauth, httpie-aws-auth, httpie-edgegrid, httpie-hmac-auth, httpie-api-auth, httpie-negotiate, requests-hawk, plugin development | [references/plugins.md](references/plugins.md) |
| scripting, exit codes, --check-status, shell script, automation, --ignore-stdin, --timeout, piping, stdin, redirected input, --raw, here string, piping output, best practices, cron | [references/scripting.md](references/scripting.md) |
| URL, URL syntax, default scheme, --default-scheme, localhost shortcut, :/foo, :3000, ://, query string, ==, path normalization, path-as-is, dot segment, URL escaping, http+unix | [references/url-syntax.md](references/url-syntax.md) |

## Quick Guide

- **What flags and options does `http` support?** → [references/cli.md](references/cli.md)
- **How do I specify headers, data fields, query params, or file uploads?** → [references/request-items.md](references/request-items.md)
- **How do I authenticate with basic, digest, or bearer token?** → [references/auth.md](references/auth.md)
- **How do sessions work and how do I persist cookies/headers across requests?** → [references/sessions.md](references/sessions.md)
- **How do I control output formatting, colors, and what gets printed?** → [references/output.md](references/output.md)
- **How do I configure SSL/TLS certificates and protocol versions?** → [references/ssl.md](references/ssl.md)
- **How do I download files with `http`?** → [references/downloads.md](references/downloads.md)
- **Where is the config file and what can I configure?** → [references/config.md](references/config.md)
- **How do I install, list, or develop HTTPie plugins?** → [references/plugins.md](references/plugins.md)
- **What exit codes does HTTPie use and how do I use it in scripts?** → [references/scripting.md](references/scripting.md)
- **How do URL shortcuts like `:/foo` or `:3000` work?** → [references/url-syntax.md](references/url-syntax.md)
- **What is the difference between `http` and `https` commands?** → [references/cli.md](references/cli.md)
- **How do I send nested JSON or arrays via request items?** → [references/request-items.md](references/request-items.md)
- **How do I pipe data into HTTPie via stdin?** → [references/scripting.md](references/scripting.md)
- **How do I resume an interrupted download?** → [references/downloads.md](references/downloads.md)
- **How do I set default options for every invocation?** → [references/config.md](references/config.md)

## Cross-Project References

- For carapace completion integration (the `http`/`https` completers in `completers/common/http_completer/` and `completers/common/https_completer/`), see the **carapace-dev** skill and the project's `AGENTS.md`. The completer uses shared actions from `pkg/actions/net/http/` for request headers, methods, media types, and status codes.
- For the carapace YAML spec format and spec-based completions, see the **carapace** skill → `references/spec.md`.
- For shell quoting rules that affect how request items are passed on the command line, see the **bash**, **zsh**, or **fish** skills as appropriate.
- For HTTP protocol fundamentals (status codes, headers, methods, content negotiation), consult RFC 7230/7231 or the MDN Web Docs — this skill covers HTTPie's CLI surface, not HTTP itself.
