# Configuration

Rails application configuration: the boot sequence, config files, `config.*` options, `config/database.yml`, environment-specific config, initializers, and the `config.load_defaults` versioning mechanism.

> **Source of truth**: `railties/lib/rails/application/configuration.rb`, `railties/lib/rails/application.rb`, and the official [Configuring Rails Applications](https://guides.rubyonrails.org/configuring.html) guide.

## Boot Sequence

Two entry points lead to the same boot sequence:

**Via `bin/rails` (CLI commands):**

```
bin/rails
  → config/boot.rb            (Bundler setup, Bootsnap)
    → rails/commands          (command resolution)
      → require APP_PATH      (loads config/application.rb)
        → config/environment.rb (Rails.application.initialize!)
```

**Via `config.ru` (Rack servers: Puma, Unicorn, Passenger):**

```
config.ru
  → require config/environment
    → config/application.rb
      → config/boot.rb
```

In both cases, `config/environment.rb` calls `Rails.application.initialize!`, which runs the initializer queue.

### `config/boot.rb`

Sets up Bundler and Bootsnap:

```ruby
ENV["BUNDLE_GEMFILE"] ||= File.expand_path("../Gemfile", __dir__)
require "bundler/setup"
require "bootsnap/setup"
```

### `config/application.rb`

Defines the `Application` class, loads all Rails frameworks, and sets app-wide config:

```ruby
require_relative "boot"
require "rails/all"

module MyApp
  class Application < Rails::Application
    config.load_defaults 7.2
    config.time_zone = "Central Time (US & Canada)"
    config.active_record.default_timezone = :utc
  end
end
```

### `config/environment.rb`

Boots the application:

```ruby
require_relative "application"
Rails.application.initialize!
```

### `config.ru`

Rack-compatible entry point used by Puma, Unicorn, etc.:

```ruby
require_relative "config/environment"
run Rails.application
Rails.application.load_server
```

## `config.load_defaults`

Loads default configuration values for a target Rails version and all prior versions. Newer versions override older ones for conflicting values.

```ruby
config.load_defaults 7.2
```

Key defaults by version target:

| Target | Notable Defaults |
|--------|-----------------|
| `8.1` | `config.action_controller.escape_json_responses: false`, `config.yjit: !Rails.env.local?` |
| `8.0` | `Regexp.timeout: 1`, `config.action_dispatch.strict_freshness: true` |
| `7.2` | `config.active_record.postgresql_adapter_decode_dates: true`, `config.yjit: true` |
| `7.1` | `config.active_support.cache_format_version: 7.1`, `config.add_autoload_paths_to_load_path: false` |
| `7.0` | `config.action_dispatch.cookies_serializer: :json`, `config.active_storage.variant_processor: :vips` |
| `6.1` | `config.action_dispatch.cookies_same_site_protection: :lax` |
| `6.0` | `config.action_dispatch.use_cookies_with_metadata: true` |
| `5.2` | `config.action_controller.default_protect_from_forgery: true` |
| `5.0` | `config.active_record.belongs_to_required_by_default: true` |

## Core `config.*` Options

### General Application

| Option | Values | Default | Description |
|--------|--------|---------|-------------|
| `config.load_defaults` | version string | — | Load defaults for target version |
| `config.enable_reloading` | bool | `true` (dev), `false` (prod) | Reload classes between requests |
| `config.eager_load` | bool | `false` (dev/test), `true` (prod) | Eager load all namespaces on boot |
| `config.cache_classes` | bool | `!enable_reloading` | Legacy alias for `!enable_reloading` |
| `config.consider_all_requests_local` | bool | `true` (dev/test), `false` (prod) | Show detailed errors in responses |
| `config.log_level` | `:debug`, `:info`, `:warn`, `:error`, `:fatal`, `:unknown` | `:debug` (dev), `:info` (prod) | Logger verbosity |
| `config.time_zone` | ActiveSupport::TimeZone name | `"UTC"` | Default app time zone |
| `config.encoding` | string | `UTF-8` | Application encoding |
| `config.cache_store` | `:memory_store`, `:file_store`, `:mem_cache_store`, `:redis_cache_store`, `:null_store` | varies by env | Cache backend |
| `config.session_store` | `:cookie_store`, `:cache_store`, `:mem_cache_store`, `:disabled` | `:cookie_store` | Session storage |
| `config.force_ssl` | bool | `false` | Force HTTPS for all requests |
| `config.hosts` | array of strings/regexps/IPAddr | varies | Host header validation |
| `config.filter_parameters` | array of symbols/strings/regexps | `[]` | Filter sensitive params from logs |
| `config.autoload_paths` | array of paths | `[]` | Autoload directories (rarely needed since Rails 6) |
| `config.autoload_once_paths` | array of paths | `[]` | Autoload directories not wiped per request |
| `config.exceptions_app` | Rack app | `ActionDispatch::PublicExceptions` | Exceptions handler |
| `config.asset_host` | URL string | — | Host for assets (CDN) |
| `config.public_file_server.enabled` | bool | `true` (dev), `false` (prod) | Serve static files from `public/` |

### Active Record

| Option | Values | Default | Description |
|--------|--------|---------|-------------|
| `config.active_record.default_timezone` | `:local`, `:utc` | `:utc` | Timezone for DB date/time values |
| `config.active_record.schema_format` | `:ruby`, `:sql` | `:ruby` | Schema dump format |
| `config.active_record.timestamped_migrations` | bool | `true` | Use timestamps in migration filenames |
| `config.active_record.primary_key_prefix_type` | `:table_name`, nil | nil | PK column naming convention |
| `config.active_record.belongs_to_required_by_default` | bool | `true` (since 5.0) | `belongs_to` required by default |

### Action Controller

| Option | Values | Default | Description |
|--------|--------|---------|-------------|
| `config.action_controller.perform_caching` | bool | `true` (prod), `false` (dev) | Enable controller caching |
| `config.action_controller.allow_forgery_protection` | bool | `true` (except test) | CSRF protection |
| `config.action_controller.default_protect_from_forgery` | bool | `true` (since 5.2) | Add CSRF protection by default |
| `config.action_controller.per_form_csrf_tokens` | bool | `true` (since 5.0) | Per-form CSRF tokens |
| `config.action_controller.action_on_unpermitted_parameters` | `:log`, `:raise` | — | Behavior for unpermitted params |

### Action Mailer

| Option | Values | Default | Description |
|--------|--------|---------|-------------|
| `config.action_mailer.delivery_method` | `:smtp`, `:sendmail`, `:file`, `:test` | `:smtp` | Mail delivery method |
| `config.action_mailer.perform_deliveries` | bool | `true` (false in test) | Whether mail is actually sent |
| `config.action_mailer.raise_delivery_errors` | bool | `false` | Raise on delivery failure |
| `config.action_mailer.default_options` | hash | — | Default mail options (e.g., `from:`) |
| `config.action_mailer.smtp_settings` | hash | — | SMTP config (`:address`, `:port`, `:domain`, etc.) |
| `config.action_mailer.show_previews` | bool | `false` | Show mailer previews (dev) |

### Active Storage

| Option | Values | Default | Description |
|--------|--------|---------|-------------|
| `config.active_storage.variant_processor` | `:mini_magick`, `:vips` | `:vips` (since 7.0) | Image variant processor |
| `config.active_storage.service` | `:local`, `:amazon`, `:google`, `:microsoft`, `:mirror` | `:local` | Storage service |
| `config.active_storage.queues.analysis` | symbol/nil | — | Queue for analysis jobs |
| `config.active_storage.queues.purge` | symbol/nil | — | Queue for purge jobs |

### Action Dispatch

| Option | Values | Default | Description |
|--------|--------|---------|-------------|
| `config.action_dispatch.cookies_serializer` | `:json`, `:marshal`, `:hybrid` | `:json` (since 7.0) | Cookie serializer |
| `config.action_dispatch.cookies_same_site_protection` | `:lax`, `:strict`, `:none` | `:lax` (since 6.1) | SameSite cookie attribute |
| `config.action_dispatch.default_headers` | hash | varies | Default HTTP response headers |
| `config.action_dispatch.show_exceptions` | `:all`, `:rescuable`, `:none` | `:all` (dev), `:none` (prod) | Exception display |
| `config.action_dispatch.ssl_default_redirect_status` | int | `308` | HTTP status for SSL redirects |

### Assets (Sprockets)

| Option | Values | Default | Description |
|--------|--------|---------|-------------|
| `config.assets.css_compressor` | `:yui`, `:sass` | — | CSS compressor |
| `config.assets.js_compressor` | `:terser`, `:closure`, `:uglifier`, `:yui` | — | JS compressor |
| `config.assets.precompile` | array | `[application.css, application.js]` | Additional precompiled assets |
| `config.assets.compile` | bool | `false` (prod) | Live compilation in production |
| `config.assets.debug` | bool | `true` (dev) | Disable concatenation |
| `config.assets.quiet` | bool | `true` (dev) | Suppress asset request logging |
| `config.assets.prefix` | string | `/assets` | URL prefix for assets |

## `config/database.yml`

Specifies database connection per environment. Supports ERB and URL format.

```yaml
development:
  adapter: sqlite3
  database: storage/development.sqlite3
  pool: 5
  timeout: 5000

production:
  url: <%= ENV['DATABASE_URL'] %>
```

### Adapter Configurations

| Adapter | Key Fields | Notes |
|---------|-----------|-------|
| `sqlite3` | `database`, `pool`, `timeout` | File path in `database` |
| `mysql2` | `database`, `encoding`, `pool`, `username`, `password`, `socket` | `encoding: utf8mb4` recommended |
| `postgresql` | `database`, `encoding`, `pool`, `host`, `username`, `password` | `encoding: unicode` common |
| `trilogy` | same as `mysql2` | MySQL via Trilogy adapter |

### Multi-DB Configuration

```yaml
development:
  primary:
    adapter: postgresql
    database: blog_development
  animals:
    adapter: sqlite3
    database: storage/animals_development.sqlite3
    migrations_paths: db/animals_migrate
```

The `:<name>` suffix on `db:*` commands targets a specific database config. See [database.md](database.md) for multi-db task details.

### `DATABASE_URL` Precedence

If both `config/database.yml` and `ENV['DATABASE_URL']` are set, the env var takes precedence for duplicate values. Non-duplicate values are merged.

## Environment-Specific Config

Files in `config/environments/` override `config/application.rb` settings per environment:

- `config/environments/development.rb`
- `config/environments/test.rb`
- `config/environments/production.rb`

Key differences by environment:

| Setting | Development | Test | Production |
|---------|------------|------|-----------|
| `config.enable_reloading` | `true` | `false` | `false` |
| `config.eager_load` | `false` | `false` | `true` |
| `config.consider_all_requests_local` | `true` | `true` | `false` |
| `config.log_level` | `:debug` | `:debug` | `:info` |
| `config.action_controller.perform_caching` | `false` | `false` | `true` |
| `config.cache_store` | `:memory_store` | `:memory_store` | `:file_store` |
| `config.public_file_server.enabled` | `true` | `true` | `false` |

Custom environments (e.g., `staging`) require a corresponding file in `config/environments/staging.rb`.

## `config/initializers/`

Files in `config/initializers/*.rb` are loaded alphabetically after all frameworks and gems are loaded.

Common initializers created by `rails new`:

| File | Purpose |
|------|---------|
| `filter_parameter_logging.rb` | Filter sensitive params from logs |
| `cors.rb` | CORS configuration |
| `inflections.rb` | Custom inflection rules |
| `assets.rb` | Asset pipeline configuration |
| `content_security_policy.rb` | CSP configuration |
| `permissions_policy.rb` | Permissions policy |
| `session_store.rb` | Session store configuration |
| `backtrace_silencers.rb` | Backtrace silencers |
| `generators.rb` | Generator configuration |

- Number prefixes control ordering: `01_critical.rb` runs before `02_normal.rb`
- Explicitly `require`-ing initializers causes double-loading
- For code that must run after gem initializers, use `config.after_initialize` blocks

## Initialization Hooks

| Hook | When It Runs |
|------|-------------|
| `config.before_configuration` | When `Application` class inherits from `Rails::Application`, before class body executes |
| `config.before_initialize` | Before the initialization process (via `:bootstrap_hook` initializer) |
| `config.after_initialize` | After all initializers complete (via `:finisher_hook` initializer) |
| `config.to_prepare` | After initializers, before eager loading; runs on every reload in dev, once in prod |
| `config.before_eager_load` | Directly before eager loading |

## `config/routes.rb`

Defines URL-to-controller/action mappings:

```ruby
Rails.application.routes.draw do
  resources :articles
  get "about", to: "pages#about"
  root "articles#index"
end
```

Loaded by the `add_routing_paths` initializer. Inspect with `rails routes` (see [cli-commands.md](cli-commands.md)).

## `config.generators` Block

Configures code generators used by `rails generate`:

```ruby
config.generators do |g|
  g.orm :active_record
  g.test_framework :test_unit
  g.template_engine :erb
  g.helper true
  g.system_tests :test_unit
  g.apply_rubocop_autocorrect_after_generate!
end
```

For generator subcommand details, see [generators.md](generators.md).

## `config.x` Namespace

Custom configuration accessible via `Rails.configuration.x`:

```ruby
config.x.payment_processing.schedule = :daily
Rails.configuration.x.payment_processing.schedule # => :daily
```

## Edge Cases

- **`config.cache_classes`** is deprecated in favor of `config.enable_reloading` (inverted logic). Both still work.
- **`config.autoload_paths`** should rarely be modified since Rails 6 with Zeitwerk. Use `app/` subdirectories instead.
- **`config.load_defaults` must match the Rails version** of the oldest environment in your deployment chain. Setting it to a newer version than the running Rails version raises an error.
- **`config/database.yml` with `DATABASE_URL`**: the env var wins for duplicate keys, merging non-overlapping values.
- **Custom environments** need a file in `config/environments/`; Rails does not create them automatically.
- **Initializers from gems** run interleaved with app initializers via topological sort; no guaranteed ordering between gem and app initializers.
- **`config.time_zone` vs `ENV['TZ']`**: `config.time_zone` sets the Rails-level `Time.zone` (display/conversion); `ENV['TZ']` sets the Ruby process system time zone (`Time.now`).

## Related Skills

- For `rails new` flags that set initial configuration, see [rails-new.md](rails-new.md).
- For CLI commands affected by config (`--environment`, `--database`), see [cli-commands.md](cli-commands.md).
- For database tasks and multi-db config, see [database.md](database.md).
- For Rails environment variables, see [environment.md](environment.md).
