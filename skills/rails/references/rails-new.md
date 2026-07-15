# `rails new` Command

Creating a new Rails application with `rails new APP_NAME [options]`. Under the hood this invokes `Rails::Generators::AppGenerator` (Thor-based).

> **Source of truth**: `railties/lib/rails/generators/rails/app/app_generator.rb` and the official guides.

## Usage

```
rails new APP_NAME [options]
```

**Positional Arguments:**
1. `APP_NAME` (required) — Name of the application. Used as directory name and Ruby module name.

## Database Options

| Flag | Values | Default | Description |
|------|--------|---------|-------------|
| `--database` / `-d` | `sqlite3`, `postgresql`, `mysql`, `oracle`, `sqlserver`, `jdbc`, `none` | `sqlite3` | Database adapter |

## JavaScript Options

| Flag | Values | Default | Description |
|------|--------|---------|-------------|
| `--javascript` / `-j` | `importmap`, `bun`, `webpack`, `esbuild`, `rollup` | `importmap` | JavaScript bundler |

## CSS Options

| Flag | Values | Default | Description |
|------|--------|---------|-------------|
| `--css` / `-c` | `tailwind`, `bootstrap`, `bulma`, `postcss`, `sass` | `tailwind` | CSS processor |

## Skip Flags (Component Exclusion)

These flags skip adding specific components to the generated app:

| Flag | What it skips |
|------|---------------|
| `--skip-action-mailer` | Action Mailer (email sending) |
| `--skip-action-mailbox` | Action Mailbox (inbound email) |
| `--skip-action-text` | Action Text (rich text / Trix) |
| `--skip-active-record` | Active Record (ORM) |
| `--skip-active-job` | Active Job (background jobs) |
| `--skip-active-storage` | Active Storage (file uploads) |
| `--skip-action-cable` | Action Cable (WebSockets) |
| `--skip-asset-pipeline` | Asset pipeline (JS/CSS/images) |
| `--skip-javascript` | JavaScript tooling/files |
| `--skip-hotwire` | Hotwire (Turbo + Stimulus) |
| `--skip-jbuilder` | Jbuilder (JSON templates) |
| `--skip-test` | Test framework (Minitest) |
| `--skip-system-test` | System tests (Capybara) |
| `--skip-bootsnap` | Bootsnap (boot caching) |
| `--skip-dev-gems` | Dev gems (web-console, listen) |
| `--skip-thruster` | Thruster (HTTP/2 proxy, Rails 8) |
| `--skip-rubocop` | RuboCop (linting) |
| `--skip-brakeman` | Brakeman (security scanner) |
| `--skip-ci` | GitHub Actions CI workflow |
| `--skip-kamal` | Kamal (deployment) |
| `--skip-solid` | Solid: Cache, Queue, Cable |
| `--skip-docker` | Dockerfile, bin/docker-entrypoint |
| `--skip-git` | Git init/commit, .gitignore |
| `--skip-keeps` | .keep files in empty dirs |
| `--skip-bundle` | `bundle install` auto-run |
| `--skip-collision-check` | Overwrite existing without checking |
| `--skip-decrypted-diffs` | Git filter for credential diffs |
| `--skip-namespace` | Namespacing (plugins/engines only) |

## Mode Flags

| Flag | Description |
|------|-------------|
| `--api` | API-only app (no views, helpers, assets) |
| `--minimal` | Minimal app (applies many `--skip-*` flags) |

## Generator Config

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--ruby` | string | — | Ruby shebang path |
| `--template` | string | — | Template file path or URL |
| `--rc` | string | `~/.railsrc` | Custom .railsrc path |
| `--no-rc` | bool | `false` | Skip .railsrc |

## Branch/Dev Flags

| Flag | Description |
|------|-------------|
| `--dev` | Use local Rails checkout |
| `--edge` | Use edge Rails branch |
| `--main` / `--master` | Use Rails main branch |
| `--devcontainer` | Add Dev Container config |

## `rails plugin new`

Same as `rails new` but for generating a plugin/engine:

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--name` | string | — | Internal module name |
| `--mountable` | bool | `false` | Mountable engine |
| `--full` | bool | `false` | Full engine (includes all layers) |
| (plus all `rails new` flags) | — | — | Same behavior |

## `.railsrc` File

Options from `~/.railsrc` are prepended to every `rails new` invocation. Format: one option per line (e.g. `--database=postgresql`). `--no-rc` skips loading.

## Edge Cases

- **`--skip-namespace` has no effect on `rails new`** (it's for `rails plugin new` only).
- **`--database=none`** means no database adapter is configured at all.
- **`--javascript` and `--css`** each install the corresponding build tooling (e.g. `--javascript=esbuild` adds `esbuild` to package.json and config).
- **`--devcontainer`** generates `.devcontainer/` with Docker Compose based on the app's database choice.
- **`--skip-solid`** became relevant in Rails 8 where Solid Cache, Solid Queue, and Solid Cable are defaults.
- **`--template`** can be a URL or local file path; it's evaluated as Ruby after all files are generated.
- **`--edge` and `--main` are mutually exclusive** — use one or the other.

## Related Skills

- For `rails generate` and `rails destroy`, see [generators.md](generators.md).
- For the main CLI commands, see [cli-commands.md](cli-commands.md).
- For `config/database.yml` format and adapter configurations, see [configuration.md](configuration.md).
