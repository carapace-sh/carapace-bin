# Environment Variables

Rails framework environment variables recognized by Rails itself, Active Record, Action Dispatch, and related components. These control boot behavior, database connections, caching, host authorization, and credential decryption.

> **Source of truth**: `railties/lib/rails/application/configuration.rb`, `railties/lib/rails.rb`, and the official [Configuring Rails Applications](https://guides.rubyonrails.org/configuring.html) guide (Section 4: Rails Environment Settings).

## Framework Variables

| Variable | Purpose | Typical Values |
|----------|---------|----------------|
| `RAILS_ENV` | Defines the Rails environment | `development`, `test`, `production`, `staging` |
| `RACK_ENV` | Rack-level environment; Rails falls back to this if `RAILS_ENV` is unset | `development`, `production` |
| `RAILS_RELATIVE_URL_ROOT` | URL path prefix when deploying to a subdirectory | `/app1` |
| `RAILS_CACHE_ID` | Suffix for expanded cache keys, enabling multiple separate caches | any string |
| `RAILS_APP_VERSION` | Used alongside `RAILS_CACHE_ID` for expanded cache keys | any string |
| `RAILS_DEVELOPMENT_HOSTS` | Comma-separated hosts permitted in development by `HostAuthorization` middleware | `host1,host2` |
| `RAILS_MASTER_KEY` | Master key for decrypting credentials (alternative to `config/master.key` file) | raw key string |
| `SECRET_KEY_BASE` | Secret key base for Rails (used for HMAC session tokens, signed cookies) | random hex string |
| `DATABASE_URL` | Database connection URL (overrides `config/database.yml` for duplicate keys) | `postgresql://localhost/blog_development?pool=5` |
| `BUNDLE_GEMFILE` | Path to the Gemfile (set by `config/boot.rb`) | path to `Gemfile` |

## `RAILS_ENV`

The primary environment selector. Determines which `config/environments/*.rb` file is loaded.

```bash
RAILS_ENV=production bin/rails server
# equivalent to: bin/rails server -e production
```

If unset, Rails checks `RACK_ENV`, then defaults to `development`. The `--environment` / `-e` flag on CLI commands sets this variable internally. See [cli-commands.md](cli-commands.md) for which commands accept `--environment`.

## `RAILS_MASTER_KEY`

Provides the master key for decrypting `config/credentials.yml.enc` without needing the `config/master.key` file:

```bash
RAILS_MASTER_KEY=abc123... bin/rails console
```

When `config.require_master_key` is `true`, the app will not boot if neither `RAILS_MASTER_KEY` nor `config/master.key` is available.

## `SECRET_KEY_BASE`

Used as the input to the key derivation function for generating HMAC-based session tokens and signed/encrypted cookies. Set it in production via:

```bash
export SECRET_KEY_BASE=$(bin/rails secret)
```

Can also be stored in `config/credentials.yml.enc` under the `secret_key_base` key. The env var takes precedence over the credentials file.

## `DATABASE_URL`

A database connection URL that can replace or supplement `config/database.yml`:

```bash
export DATABASE_URL="postgresql://user:pass@localhost/blog_production?pool=5"
```

Precedence rules when both `DATABASE_URL` and `config/database.yml` are present:
- `DATABASE_URL` wins for duplicate keys
- Non-overlapping values from both sources are merged
- In `database.yml`, ERB can reference it: `url: <%= ENV['DATABASE_URL'] %>`

For multi-db, use named URLs or configure each database in `database.yml`. See [database.md](database.md) and [configuration.md](configuration.md) for multi-db patterns.

## `RAILS_DEVELOPMENT_HOSTS`

Additional comma-separated hosts permitted by the `ActionDispatch::HostAuthorization` middleware in development:

```bash
RAILS_DEVELOPMENT_HOSTS=preview.example.com,api.lvh.me bin/rails server
```

Without this, only `localhost`, `127.0.0.1`, and `::1` (plus `lvh.me` and subdomains) are allowed in development.

## `RAILS_RELATIVE_URL_ROOT`

Used by the routing system to recognize URLs when deploying to a subdirectory:

```bash
RAILS_RELATIVE_URL_ROOT=/app1 bin/rails server
```

Equivalent to setting `config.relative_url_root = "/app1"` in the application config.

## `RAILS_CACHE_ID` and `RAILS_APP_VERSION`

Both generate expanded cache keys, allowing multiple separate caches from the same application:

```bash
RAILS_CACHE_ID=build123 RAILS_APP_VERSION=v2.3 bin/rails server
```

Useful in deployment scenarios where multiple instances share a cache backend.

## `BUNDLE_GEMFILE`

Set by `config/boot.rb` to point to the application's Gemfile:

```ruby
ENV["BUNDLE_GEMFILE"] ||= File.expand_path("../Gemfile", __dir__)
```

Overriding it allows using an alternative Gemfile (e.g., `Gemfile.next`):

```bash
BUNDLE_GEMFILE=Gemfile.rails8 bin/rails server
```

## Server-Related Variables

These are used by `config/puma.rb` and the Rails server, not directly by the framework:

| Variable | Purpose |
|----------|---------|
| `RAILS_MAX_THREADS` | Max threads per Puma worker (used in `config/puma.rb`) |
| `RAILS_MIN_THREADS` | Min threads per Puma worker |
| `RAILS_SERVE_STATIC_FILES` | Whether to serve static files from `public/` (sets `config.public_file_server.enabled`) |
| `RAILS_LOG_TO_STDOUT` | Whether to log to stdout (used in `config/puma.rb` and production config) |
| `WEB_CONCURRENCY` | Number of Puma workers (common convention) |
| `PORT` | Port for Puma (used by `config/puma.rb`) |

## Action Mailbox Ingress Variables

Used by Action Mailbox ingress rake tasks (see [rake-tasks.md](rake-tasks.md)):

| Variable | Purpose |
|----------|---------|
| `URL` | The ingress URL for Exim/Postfix/Qmail relay |
| `INGRESS_PASSWORD` | Shared password for ingress authentication |

## Fixture Variables

Used by `db:fixtures:load` (see [database.md](database.md)):

| Variable | Purpose |
|----------|---------|
| `FIXTURES` | Comma-separated list of specific fixtures to load |
| `FIXTURES_DIR` | Subdirectory in `test/fixtures/` |
| `FIXTURES_PATH` | Alternative path (e.g., `spec/fixtures`) |

## Migration Variables

Used by `db:migrate` and related tasks (see [database.md](database.md)):

| Variable | Purpose |
|----------|---------|
| `VERSION` | Target migration version (required for `db:migrate:up`/`down`) |
| `STEP` | Number of migrations to rollback (default: `1`) |
| `SCOPE` | Filter migrations by scope (for engines) |
| `VERBOSE` | `false` to suppress migration output |

## Log Variables

| Variable | Purpose |
|----------|---------|
| `LOGS` | Comma-separated log files for `log:clear` task (see [rake-tasks.md](rake-tasks.md)) |

## Edge Cases

- **`RAILS_ENV` is a string, not an enum.** Any value is accepted; the corresponding `config/environments/<name>.rb` file must exist. Common values: `development`, `test`, `production`, `staging`.
- **`RACK_ENV` fallback**: If `RAILS_ENV` is unset, Rails checks `RACK_ENV`. If both are unset, defaults to `development`. `RACK_ENV` does not support all Rails environments (e.g., no `test` default).
- **`SECRET_KEY_BASE` vs credentials**: `ENV["SECRET_KEY_BASE"]` takes precedence over `credentials.secret_key_base` in `config/credentials.yml.enc`. The credentials file is only consulted as a fallback when the env var is not set.
- **`DATABASE_URL` + `database.yml`**: They merge, not override entirely. `DATABASE_URL` wins only for duplicate keys.
- **`RAILS_SERVE_STATIC_FILES`** is checked in `config/environments/production.rb` via `ENV["RAILS_SERVE_STATIC_FILES"].present?`. It is a convention, not a framework-level variable.
- **`RAILS_MAX_THREADS`** is a Puma config convention used in `config/puma.rb`, not read by Rails framework code directly.

## Related Skills

- For Rails configuration files and `config.*` options, see [configuration.md](configuration.md).
- For CLI commands that accept `--environment`, see [cli-commands.md](cli-commands.md).
- For database tasks using `VERSION`, `STEP`, `SCOPE`, see [database.md](database.md).
- For rake tasks using `URL`, `INGRESS_PASSWORD`, `LOGS`, see [rake-tasks.md](rake-tasks.md).
