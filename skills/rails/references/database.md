# Database Tasks

All `rails db:*` subcommands. These are Rake tasks provided by Active Record, accessed through `bin/rails db:<task>`.

> **Source of truth**: `activerecord/lib/active_record/railties/databases.rake` in the Rails repo.

## Multi-DB Pattern

Many `db:*` commands support a `:<name>` suffix for multi-database applications. For example:
- `db:migrate` — migrate all databases
- `db:migrate:primary` — migrate the `primary` database only
- `db:rollback:primary` — rollback the `primary` database

The `:<name>` suffix works on: `create`, `drop`, `migrate` (and its sub-tasks), `rollback`, `schema:dump`, `schema:load`, `setup`, `reset`, `version`, `test:prepare`, `test:load_schema`, `test:purge`, `abort_if_pending_migrations`.

## Task Reference

### `db:create`

Create database for current environment. Reads `DATABASE_URL` or `config/database.yml`.

| Variant | Description |
|---------|-------------|
| `db:create` | Create current env database |
| `db:create:all` | Create all env databases |
| `db:create:<name>` | Create named database in current env |

### `db:drop`

Drop database for current environment.

| Variant | Description |
|---------|-------------|
| `db:drop` | Drop current env database |
| `db:drop:all` | Drop all env databases |
| `db:drop:<name>` | Drop named database in current env |

### `db:migrate`

Run pending migrations.

| Env var | Description |
|---------|-------------|
| `VERSION=x` | Migrate to specific version |
| `VERBOSE=false` | Suppress migration output |
| `SCOPE=blog` | Only migrations matching scope (for engines) |

| Variant | Description |
|---------|-------------|
| `db:migrate` | Migrate all databases |
| `db:migrate:<name>` | Migrate named database |
| `db:migrate:redo` | Rollback one then migrate up (`STEP=n`) |
| `db:migrate:reset` | Drop DB, recreate, migrate |
| `db:migrate:up` | Run `up` for `VERSION=x` (required) |
| `db:migrate:down` | Run `down` for `VERSION=x` (required) |
| `db:migrate:status` | Show migration status |

### `db:rollback`

Rollback the schema.

| Env var | Description |
|---------|-------------|
| `STEP=n` | Number of steps (default: 1) |

| Variant | Description |
|---------|-------------|
| `db:rollback` | Rollback one step |
| `db:rollback:<name>` | Rollback named database |
| `db:forward` | Push to next version (`STEP=n`) |

### `db:schema:dump`

Create `db/schema.rb` (or `db/structure.sql` via `config.active_record.schema_format = :sql`).

| Variant | Description |
|---------|-------------|
| `db:schema:dump` | Dump schema for current env |
| `db:schema:dump:<name>` | Dump for named database |

### `db:schema:load`

Load schema from `db/schema.rb` into database.

| Variant | Description |
|---------|-------------|
| `db:schema:load` | Load schema for current env |
| `db:schema:load:<name>` | Load for named database |

### `db:schema:cache:dump`

Create `db/schema_cache.yml`.

### `db:schema:cache:clear`

Clear `db/schema_cache.yml`.

### `db:setup`

Create databases, load schema, seed.

| Variant | Description |
|---------|-------------|
| `db:setup` | Setup current env |
| `db:setup:all` | Setup all envs |
| `db:setup:<name>` | Setup named database |

### `db:reset`

Drop, recreate, migrate, seed.

| Variant | Description |
|---------|-------------|
| `db:reset` | Reset current env |
| `db:reset:all` | Reset all envs |
| `db:reset:<name>` | Reset named database |

### `db:seed`

Load seed data from `db/seeds.rb`.

| Variant | Description |
|---------|-------------|
| `db:seed` | Load seeds |
| `db:seed:replant` | Truncate tables + load seeds |

### `db:prepare`

Create DB if not exists, or run migrations if it does. Used in CI.

### `db:fixtures:load`

Load test fixtures.

| Env var | Description |
|---------|-------------|
| `FIXTURES=x,y` | Specific fixtures |
| `FIXTURES_DIR=z` | Subdirectory in test/fixtures |
| `FIXTURES_PATH=p` | Alternative path (e.g. `spec/fixtures`) |

| Variant | Description |
|---------|-------------|
| `db:fixtures:load` | Load fixtures |
| `db:fixtures:identify` | Find fixture by LABEL or ID |

### `db:version`

Show current schema version number.

| Variant | Description |
|---------|-------------|
| `db:version` | Show version |
| `db:version:<name>` | Show version for named database |

### `db:encryption:init`

Generate Active Record encryption keys.

### `db:system:change`

**Usage:** `bin/rails db:system:change --to=VALUE`

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `--to` | enum | **yes** | Target database adapter |

**Valid `--to` values:**

| Value | Adapter | Gem |
|-------|---------|-----|
| `mysql` | MySQL | `mysql2` (~> 0.5) |
| `trilogy` | MySQL (Trilogy) | `trilogy` (~> 2.7) |
| `postgresql` | PostgreSQL | `pg` (~> 1.1) |
| `sqlite3` | SQLite3 | `sqlite3` (>= 2.1) |
| `mariadb-mysql` | MariaDB (mysql2) | `mysql2` (~> 0.5) |
| `mariadb-trilogy` | MariaDB (Trilogy) | `trilogy` (~> 2.7) |

Changes: `config/database.yml`, `Gemfile` gem swap, `Dockerfile` packages, `.devcontainer/` files.

### `test:db`

Reset test database and run `bin/rails test`.

### Other `db:*` Tasks

| Task | Description |
|------|-------------|
| `db:purge` | Empty the database |
| `db:purge:all` | Empty all databases |
| `db:truncate_all` | Truncate all tables in current env |
| `db:charset` | Retrieve charset for current env |
| `db:collation` | Retrieve collation for current env |
| `db:environment:set` | Set the environment value for the database |
| `db:abort_if_pending_migrations` | Raise error if pending migrations exist |
| `db:test:prepare` | Load test schema |
| `db:test:load_schema` | Recreate test DB from schema file |
| `db:test:purge` | Empty the test database |

### Engine Tasks

When running inside a Rails engine, these `app:db:*` tasks delegate to the host app:

`app:db:reset`, `app:db:migrate`, `app:db:migrate:up/down/redo/reset/status`, `app:db:create`, `app:db:create:all`, `app:db:drop`, `app:db:drop:all`, `app:db:fixtures:load`, `app:db:rollback`, `app:db:schema:dump`, `app:db:schema:load`, `app:db:seed`, `app:db:setup`, `app:db:version`, `app:db:test:prepare`.

## Edge Cases

- **`VERSION` is required** for `db:migrate:up` and `db:migrate:down`. Completion should require this value after those subcommands.
- **`STEP` defaults to 1** for `db:migrate:redo` and `db:rollback`. Can be any positive integer.
- **Multi-db `<name>` suffix** matches database config names in `config/database.yml`. Completion should list those keys.
- **`db:fixtures:load`** loads ALL fixtures by default; `FIXTURES=x,y` filters.
- **`db:system:change` mutates multiple files** — it's a significant change, not just config.
- **`db:prepare` is idempotent** — creates DB if absent, migrates if present. No separate create/migrate needed.

## Related Skills

- For `rails dbconsole`, see [cli-commands.md](cli-commands.md).
- For `rails test` subcommands including `test:db`, see [testing.md](testing.md).
- For `railties:install:migrations`, see [rake-tasks.md](rake-tasks.md).
