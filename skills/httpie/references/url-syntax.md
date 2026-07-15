# URL Syntax and Defaults

HTTPie's URL handling — default schemes, localhost shortcuts, query parameters, path normalization, and custom protocol schemes.

## Default Scheme

The `http` command defaults to `http://` and the `https` command defaults to `https://`:

```bash
$ http example.org     # → http://example.org
$ https example.org    # → https://example.org
```

Override the default scheme with `--default-scheme`:

```bash
$ http --default-scheme=https example.org    # → https://example.org
```

The carapace completer completes `--default-scheme` with `http://` and `https://`.

## Localhost Shortcuts

HTTPie provides shortcuts for localhost URLs:

| Shortcut | Expands To |
|----------|-----------|
| `:/foo` | `http://localhost/foo` |
| `:3000/bar` | `http://localhost:3000/bar` |
| `:` | `http://localhost/` |
| `:8080` | `http://localhost:8080/` |

```bash
$ http :/foo            # → http://localhost/foo
$ http :3000/api/v1     # → http://localhost:3000/api/v1
$ http :                # → http://localhost/
```

## Keeping `://`

A space between the scheme and `://` is equivalent to using `://` directly:

```bash
$ https ://example.org   # Same as https://example.org
$ http ://example.org     # Same as http://example.org
```

## Query String Parameters

URL parameters can be specified via the `==` request item (see [request-items.md](request-items.md)) instead of embedding them in the URL:

```bash
$ http https://api.github.com/search/repositories q==httpie per_page==1
```

HTTPie auto-URL-escapes special characters in `==` parameter values.

Read a query parameter value from a file using `==@`:

```bash
$ http pie.dev/get token==@files/token.txt
```

## Path Normalization

By default, HTTPie squashes dot segments (`/../` and `/./`) in URLs:

```bash
$ http example.org/foo/../bar    # → http://example.org/bar
$ http example.org/foo/./bar     # → http://example.org/foo/bar
```

Disable this with `--path-as-is`:

```bash
$ http --path-as-is example.org/foo/../bar    # → http://example.org/foo/../bar
```

## Custom Protocol Schemes

Plugins can add custom URL scheme prefixes. For example, `httpie-unixsocket` adds `http+unix://`:

```bash
$ http http+unix://%2Fvar%2Frun%2Fdocker.sock/info
```

Create a convenience alias:

```bash
$ alias http-unix='http --default-scheme="http+unix"'
$ http-unix %2Fvar%2Frun%2Fdocker.sock/info
```

See [plugins.md](plugins.md) for transport plugin details.

## Carapace Completion

The carapace completer provides URL completion via `http.ActionUrls()` from `pkg/actions/net/http/url.go`. This action completes known hosts and ports as URLs. The first positional argument can be either an HTTP method or a URL:

```go
carapace.Gen(rootCmd).PositionalCompletion(
    carapace.Batch(
        http.ActionRequestMethods(),
        http.ActionUrls(),
    ).ToA(),
    carapace.ActionCallback(func(c carapace.Context) carapace.Action {
        // If first arg is a method, second is URL; otherwise request items
        if _, ok := requestMethods[c.Args[0]]; ok {
            return http.ActionUrls()
        }
        return ActionRequestItem()
    }),
)
```

## References

- Source of truth: <https://httpie.io/docs/cli> (URL syntax section)
- Carapace completer: `completers/common/http_completer/cmd/root.go`
- URL action: `pkg/actions/net/http/url.go`

## Related Skills

- [cli.md](cli.md) — full flag reference including `--default-scheme`
- [request-items.md](request-items.md) — `==` URL parameters
- [plugins.md](plugins.md) — custom transport schemes
