# CLI Commands

The main rails CLI commands (server, console, runner, dbconsole, routes, credentials, encrypted, notes, query, devcontainer, and standalone commands). Every command has `--help` / `-h`.

> **Source of truth**: `railties/lib/rails/commands/` in the Rails repo.

## Aliases

| Alias | Command |
|-------|---------|
| `g` | `generate` |
| `d` | `destroy` |
| `c` | `console` |
| `s` | `server` |
| `db` | `dbconsole` |
| `r` | `runner` |
| `t` | `test` |

## `rails server` (`s`)

**Usage:** `bin/rails server [options]`

| Flag | Alias | Type | Default | Description |
|------|-------|------|---------|-------------|
| `--port` | `-p` | int | `3000` | Port to run on |
| `--binding` | `-b` | string | `localhost` (dev) / `0.0.0.0` (other) | IP to bind to |
| `--config` | `-c` | string | `config.ru` | Rackup config file |
| `--daemon` | `-d` | bool | `false` | Run as daemon |
| `--using` | `-u` | string | — | Rack server (thin/puma/webrick) |
| `--pid` | `-P` | string | `tmp/pids/server.pid` | PID file path |
| `--dev-caching` | `-C` | bool | `nil` | Enable dev mode caching |
| `--early-hints` | — | bool | `nil` | HTTP/2 early hints |
| `--log-to-stdout` | — | bool | `nil` | Log to stdout (enabled in dev when not daemonized) |
| `--environment` | `-e` | string | `development` | Rails environment |
| `--restart` | — | bool | `nil` | Internal (hidden): restart |
| `--help` | `-h` | — | — | Show help |

## `rails console` (`c`)

**Usage:** `bin/rails console [options] [-- args]`

| Flag | Alias | Type | Default | Description |
|------|-------|------|---------|-------------|
| `--sandbox` | `-s` | bool | `nil` | Rollback DB modifications on exit |
| `--skip-executor` | `-w` | bool | `false` | Don't wrap with Rails Executor |
| `--query-cache` | `-q` | bool | `false` | Enable AR query cache (ignored if `--skip-executor`) |
| `--environment` | `-e` | string | `development` | Rails environment |
| `--help` | `-h` | — | — | Show help |

Arguments after `--` are passed through to IRB.

## `rails runner` (`r`)

**Usage:** `bin/rails runner [options] [<'Some.ruby(code)'> | <filename.rb> | -]`

| Flag | Alias | Type | Default | Description |
|------|-------|------|---------|-------------|
| `--skip-executor` | `-w` | bool | `false` | Don't wrap with Rails Executor |
| `--environment` | `-e` | string | `development` | Rails environment |
| `--help` | `-h` | — | — | Show help |

**Positional arguments:**
1. `code_or_file` (required) — inline Ruby code, path to a `.rb` file, or `-` for stdin. Additional args after this go into `ARGV`.

## `rails dbconsole` (`db`)

**Usage:** `bin/rails dbconsole [options]`

| Flag | Alias | Type | Default | Description |
|------|-------|------|---------|-------------|
| `--include-password` | `-p` | bool | `false` | Auto-provide password from database.yml |
| `--mode` | — | enum | — | sqlite3 mode: `html`, `list`, `line`, `column` |
| `--header` | — | bool | — | Show headers (sqlite3) |
| `--database` / `--db` | — | string | — | Database config name (multi-db) |
| `--environment` | `-e` | string | `development` | Rails environment |
| `--help` | `-h` | — | — | Show help |

Automatically detects adapter (postgresql → psql, mysql2 → mysql, sqlite3 → sqlite3).

## `rails routes`

**Usage:** `bin/rails routes [options]`

| Flag | Alias | Type | Default | Description |
|------|-------|------|---------|-------------|
| `--controller` | `-c` | string | — | Filter by controller (e.g. `PostsController`) |
| `--grep` | `-g` | string | — | Grep by pattern (name/path/controller/action) |
| `--expanded` | `-E` | bool | `false` | Print expanded vertically |
| `--unused` | `-u` | bool | `false` | Print unused routes |
| `--help` | `-h` | — | — | Show help |

## `rails credentials`

**Usage:** `bin/rails credentials:SUBCOMMAND [options]`

### `credentials:edit`

Open decrypted credentials in `$VISUAL` or `$EDITOR`.

| Flag | Alias | Type | Default | Description |
|------|-------|------|---------|-------------|
| `--environment` | `-e` | string | — | Environment (e.g. `development` → `config/credentials/development.yml.enc`) |
| `--help` | `-h` | — | — | Show help |

### `credentials:show`

Show decrypted credentials.

| Flag | Alias | Type | Default | Description |
|------|-------|------|---------|-------------|
| `--environment` | `-e` | string | — | Environment scope |
| `--help` | `-h` | — | — | Show help |

### `credentials:diff`

Enroll/disenroll in decrypted git diffs.

| Flag | Alias | Type | Default | Description |
|------|-------|------|---------|-------------|
| `--enroll` | — | bool | `false` | Enroll in credentials diff via git |
| `--disenroll` | — | bool | `false` | Disenroll from credentials diff |
| `--help` | `-h` | — | — | Show help |

**Positional arguments:** `content_path` (optional).

### `credentials:fetch`

**Usage:** `bin/rails credentials:fetch PATH [options]`

| Flag | Alias | Type | Default | Description |
|------|-------|------|---------|-------------|
| `--environment` | `-e` | string | — | Environment scope |
| `--help` | `-h` | — | — | Show help |

**Positional arguments:** `PATH` (required) — key path in credentials, e.g. `secret_key_base`.

## `rails encrypted`

**Usage:** `bin/rails encrypted:SUBCOMMAND [options]`

### `encrypted:edit`

**Usage:** `bin/rails encrypted:edit FILE [options]`

| Flag | Alias | Type | Default | Description |
|------|-------|------|---------|-------------|
| `--key` | `-k` | string | `config/master.key` | Path to encryption key |
| `--help` | `-h` | — | — | Show help |

### `encrypted:show`

**Usage:** `bin/rails encrypted:show FILE [options]`

| Flag | Alias | Type | Default | Description |
|------|-------|------|---------|-------------|
| `--key` | `-k` | string | `config/master.key` | Path to encryption key |
| `--help` | `-h` | — | — | Show help |

## `rails notes`

**Usage:** `bin/rails notes [options]`

| Flag | Alias | Type | Default | Description |
|------|-------|------|---------|-------------|
| `--annotations` | `-a` | array | `FIXME,OPTIMIZE,TODO` | Annotation tags (comma-separated) |
| `--help` | `-h` | — | — | Show help |

## `rails query`

**Usage:** `bin/rails query [options] [EXPRESSION]`

A read-only Active Record REPL.

| Flag | Alias | Type | Default | Description |
|------|-------|------|---------|-------------|
| `--sql` | — | bool | `false` | Treat input as raw SQL |
| `--database` / `--db` | — | string | — | Database config name |
| `--page` | — | int | `1` | Page number |
| `--per` | — | int | `100` | Results per page (max 10000) |
| `--environment` | `-e` | string | `development` | Rails environment |
| `--help` | `-h` | — | — | Show help |

**Sub-commands (as expression):**
| Expression | Description |
|------------|-------------|
| `schema` | List all tables |
| `schema <table>` | Show table details |
| `models` | List all models with associations |
| `explain <expression>` | Show EXPLAIN output |

## `rails devcontainer`

**Usage:** `bin/rails devcontainer`

Generate a Dev Container setup based on current app configuration.

| Flag | Alias | Type | Default | Description |
|------|-------|------|---------|-------------|
| `--help` | `-h` | — | — | Show help |

No positional arguments.

## Standalone Commands

| Command | Description | Flags |
|---------|-------------|-------|
| `about` | Show Rails/Ruby/gem versions, env info | None |
| `boot` | Boot app and exit (debug boot issues) | `-e`/`--environment` |
| `dev:cache` | Toggle dev caching on/off | None |
| `middleware` | Print Rack middleware stack | None |
| `initializers` | Print initializers in invocation order | `-e`/`--environment` |
| `restart` | Touch `tmp/restart.txt` to signal server restart | None |
| `secret` | Generate cryptographically secure random hex string | None |
| `stats` | Report code statistics (KLOC) | None |
| `version` | Show Rails version | None |

## `rails destroy` (`d`)

**Usage:** `bin/rails destroy GENERATOR [args] [options]`

Inverse of `generate`. Uses the same generator classes. Same positional arguments and flags as the corresponding `generate` command (e.g. `destroy model User` removes files `generate model User` created).

## `rails plugin new`

**Usage:** `bin/rails plugin new NAME [options]`

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--rc` | string | `~/.railsrc` | Initialize with previous defaults |
| `--no-rc` | bool | `false` | Skip evaluating .railsrc |
| (plus all `rails new` flags) | — | — | See [rails-new.md](rails-new.md) |

## Edge Cases

- **`--environment`** is a string, not an enum. Common values: `development`, `test`, `production`, `staging`. Any value is accepted.
- **`--database` in dbconsole** refers to a key in `config/database.yml`, not an adapter name. Multi-db apps have multiple database configs.
- **`credentials:edit` without `--environment`** edits the primary `config/credentials.yml.enc`.
- **`encrypted`** works on any YAML file encrypted with a key, not just credentials.
- **`server --daemon`** is daemonized; `--log-to-stdout` has no effect when daemonized.
- **`query`** is read-only (wraps in a transaction that rolls back).

## Related Skills

- For `rails generate` and `rails destroy` generator arguments, see [generators.md](generators.md).
- For `rails new` (app creation), see [rails-new.md](rails-new.md).
- For `rails db:*` subcommands, see [database.md](database.md).
- For `rails test` subcommands, see [testing.md](testing.md).
- For `rails log:clear`, `tmp:*`, `time:zones`, `action_mailbox:*`, etc., see [rake-tasks.md](rake-tasks.md).
