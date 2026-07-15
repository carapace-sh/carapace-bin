# Plugin System

HTTPie's plugin architecture — four plugin types, the built-in plugin manager CLI, and plugin development.

## Plugin Types

| Type | Entry Point Group | Activated By | Purpose |
|------|-------------------|-------------|---------|
| AuthPlugin | `httpie.plugins.auth.v1` | `--auth-type` | Custom authentication mechanisms |
| TransportPlugin | `httpie.plugins.transport.v1` | URL prefix | Custom transport adapters (e.g., Unix sockets) |
| ConverterPlugin | `httpie.plugins.converter.v1` | Response MIME type | Convert binary response data to text |
| FormatterPlugin | `httpie.plugins.formatter.v1` | Content type / CLI options | Format response output (headers, body, metadata) |

The `PluginManager` class discovers plugins via Python entry points (`importlib_metadata.entry_points()`). If a plugin fails to load, a warning is displayed and HTTPie continues.

## AuthPlugin

Custom authentication activated via `--auth-type`:

| Attribute | Default | Description |
|-----------|---------|-------------|
| `auth_type` | | Value passed to `--auth-type` (e.g., `"ntlm"`, `"jwt"`) |
| `auth_require` | `True` | Credentials must be provided via `--auth` |
| `auth_parse` | `True` | Parse `-a` as `username:password` |
| `netrc_parse` | `False` | Allow fetching credentials from `.netrc` |
| `prompt_password` | `True` | Prompt for password if missing |
| `raw_auth` | | Raw value of `-a`, set before `get_auth()` |

Key method: `get_auth(self, username=None, password=None)` — returns a `requests.auth.AuthBase` instance.

## TransportPlugin

Custom transport adapters activated by URL prefix:

| Attribute | Description |
|-----------|-------------|
| `prefix` | URL prefix (e.g., `http+unix://`) |

Key method: `get_adapter(self)` — returns a `requests.adapters.BaseAdapter` instance.

Example: `httpie-unixsocket` enables Unix domain socket support:

```bash
$ http http+unix://%2Fvar%2Frun%2Fdocker.sock/info
```

## ConverterPlugin

Converts binary response data to text for terminal display:

| Method | Description |
|--------|-------------|
| `supports(cls, mime)` | Class method — returns `True` if plugin handles the MIME type |
| `convert(self, body)` | Takes raw bytes, returns `(new_content_type, textual_content)` |

## FormatterPlugin

Formats response output:

| Method | Description |
|--------|-------------|
| `format_headers(self, headers)` | Format response headers |
| `format_body(self, content, mime)` | Format response body |
| `format_metadata(self, metadata)` | Format response metadata |

Built-in formatters: `HeadersFormatter`, `JSONFormatter`, `XMLFormatter`, `ColorFormatter`.

## Built-in Plugins

HTTPie ships with these registered automatically:

| Type | Plugins |
|------|---------|
| Auth | `BasicAuthPlugin`, `DigestAuthPlugin`, `BearerAuthPlugin` |
| Formatter | `HeadersFormatter`, `JSONFormatter`, `XMLFormatter`, `ColorFormatter` |

## Plugin Manager CLI

Starting with HTTPie 3.0, a built-in plugin manager is available (beta):

### Install

```bash
$ httpie cli plugins install httpie-plugin
```

Install from a local path:

```bash
$ httpie cli plugins install /path/to/plugin
```

### List

```bash
$ httpie cli plugins list
httpie_plugin (1.0.2)
  httpie_plugin (httpie.plugins.auth.v1)
httpie_converter (1.0.0)
  httpie_iterm_converter (httpie.plugins.converter.v1)
```

### Upgrade

```bash
$ httpie cli plugins upgrade httpie-plugin
```

### Uninstall

```bash
$ httpie cli plugins uninstall httpie-plugin
```

Only uninstalls plugins installed via `httpie cli plugins install`.

## Known Auth Plugins

| Plugin | Auth Type | Repository |
|--------|-----------|------------|
| httpie-ntlm | ntlm | [httpie/httpie-ntlm](https://github.com/httpie/httpie-ntlm) |
| httpie-jwt-auth | jwt | [teracyhq/httpie-jwt-auth](https://github.com/teracyhq/httpie-jwt-auth) |
| httpie-oauth1 | oauth1 | [qcif/httpie-oauth1](https://github.com/qcif/httpie-oauth1) |
| httpie-aws-auth | aws | [httpie/httpie-aws-auth](https://github.com/httpie/httpie-aws-auth) |
| httpie-edgegrid | edgegrid | [akamai-open/httpie-edgegrid](https://github.com/akamai-open/httpie-edgegrid) |
| httpie-hmac-auth | hmac | [guardian/httpie-hmac-auth](https://github.com/guardian/httpie-hmac-auth) |
| httpie-api-auth | api-auth | [pd/httpie-api-auth](https://github.com/pd/httpie-api-auth) |
| httpie-negotiate | negotiate | [ndzou/httpie-negotiate](https://github.com/ndzou/httpie-negotiate) |
| requests-hawk | hawk | [mozilla-services/requests-hawk](https://github.com/mozilla-services/requests-hawk) |

## Plugin Development

1. Choose a plugin type (AuthPlugin, TransportPlugin, ConverterPlugin, FormatterPlugin)
2. Implement a Python class inheriting from the base class
3. Register an entry point in `setup.cfg` or `pyproject.toml`:

```ini
# setup.cfg
[options.entry_points]
httpie.plugins.auth.v1 =
    my_custom_auth = my_package.my_module:MyCustomAuthPlugin
```

```toml
# pyproject.toml
[project.entry-points."httpie.plugins.auth.v1"]
my_custom_auth = "my_package.my_module:MyCustomAuthPlugin"
```

4. Install into the same Python environment as HTTPie

Plugins are discovered automatically by the `PluginManager` on startup.

## Plugin Storage

Plugins are stored in the `plugins_dir` configured in `config.json`. See [config.md](config.md). By default, this is under the config directory.

## References

- Source of truth: <https://httpie.io/docs/cli/plugins>
- Auth plugins: <https://httpie.io/docs/cli/auth-plugins>
- Plugin manager: <https://httpie.io/docs/cli/plugin-manager>
- Source code: `httpie/plugins/base.py`, `httpie/plugins/manager.py`, `httpie/plugins/registry.py`

## Related Skills

- [auth.md](auth.md) — built-in auth types and how plugins extend them
- [config.md](config.md) — `plugins_dir` configuration
