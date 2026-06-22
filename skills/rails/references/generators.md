# Generators

Every built-in Rails generator with its flags, positional arguments, and special behavior. Used by both `rails generate` and `rails destroy`.

> **Source of truth**: `railties/lib/rails/generators/rails/` and component-specific generator directories in the Rails repo.

## Global Generator Options

Available on every generator:

| Flag | Alias | Type | Default | Description |
|------|-------|------|---------|-------------|
| `--pretend` | `-p` | bool | `false` | Run but make no changes |
| `--force` | `-f` | bool | `false` | Overwrite existing files |
| `--skip` | `-s` | bool | `false` | Skip existing files |
| `--quiet` | `-q` | bool | `false` | Suppress status output |
| `--no-color` | — | bool | `false` | Disable color output |
| `--skip-namespace` | — | bool | `false` | Skip namespace in file paths |
| `--skip-collision-check` | — | bool | `false` | Skip file collision check |
| `--help` | `-h` | — | — | Show generator help |

## Built-in Generators

### `model`

**Usage:** `bin/rails generate model NAME [field[:type][:index] ...] [options]`

**Positional:** `name` (required), `attributes` (array: `field[:type][:index]`)

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--parent` | string | `ApplicationRecord` | Parent model class |
| `--timestamps` | bool | — | Add timestamps |
| `--indexes` | bool | `true` | Indexes for references/belongs_to |
| `--primary-key-type` | string | — | PK type (e.g. `:uuid`) |
| `--database` / `--db` | string | — | Target database (multi-db) |
| `--migration` | bool | — | Create migration |

**Field types:** `string`, `text`, `integer`, `bigint`, `float`, `decimal`, `numeric`, `datetime`, `time`, `date`, `binary`, `boolean`, `references`, `belongs_to`, `digest` (password_digest), `token`. Appending `:index` after the type (e.g. `email:string:index`) creates a DB index. Appending `{null:false}` sets NOT NULL. Default type is `string` if omitted.

### `scaffold`

**Usage:** `bin/rails generate scaffold NAME [field[:type][:index] ...] [options]`

**Positional:** `name` (required), `attributes` (array: `field[:type][:index]`)

Generates model, migration, controller, views, routes, tests, jbuilder, helper. Triggers these hooks: `active_record:model` → `resource_route` → `scaffold_controller` → (`erb:scaffold`, `resource_route`, `test_unit:scaffold`, `helper`, `jbuilder`).

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--api` | bool | `false` | API-only (no views) |
| `--resource-route` | bool | — | Add resource route |
| `--actions` | array | `[]` | Controller actions |
| `--timestamps` | bool | — | Add timestamps |
| `--parent` | string | `ApplicationRecord` | Parent model class |
| `--indexes` | bool | `true` | Indexes for references |
| `--primary-key-type` | string | — | PK type |
| `--database` / `--db` | string | — | Database for multi-db |
| `--skip-collision-check` | bool | `false` | Skip collision check |
| `--skip-namespace` | bool | `false` | Skip namespace |

### `resource`

**Usage:** `bin/rails generate resource NAME [field[:type][:index] ...] [options]`

**Positional:** `name` (required), `attributes` (array: `field[:type][:index]`)

Like scaffold but generates no views. Creates model, migration, empty controller, routes, tests.

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--actions` | array | `[]` | Controller actions |
| `--timestamps` | bool | — | Add timestamps |
| `--parent` | string | `ApplicationRecord` | Parent model class |
| `--indexes` | bool | `true` | Indexes for references |
| `--primary-key-type` | string | — | PK type |
| `--database` / `--db` | string | — | Database for multi-db |
| `--migration` | bool | — | Create migration |

**Hooks:** `resource_controller` → `resource_route` → ORM hook (model + migration).

### `controller`

**Usage:** `bin/rails generate controller NAME [action action ...] [options]`

**Positional:** `name` (required), `actions` (array: method names)

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--skip-routes` | bool | `false` | Don't add routes |
| `--helper` | bool | — | Generate helper |
| `--parent` | string | `ApplicationController` | Parent controller class |
| `--skip-collision-check` | bool | `false` | Skip collision check |

Strips `_controller` suffix automatically. If actions provided and `--skip-routes` unset, inserts `get "controller/action"` routes.

**Hooks:** `template_engine` (ERB), `test_framework` (test_unit), `helper`.

### `scaffold_controller`

**Usage:** `bin/rails generate scaffold_controller NAME [field:type ...] [options]`

**Positional:** `name` (required), `attributes` (array: `field:type`)

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--helper` | bool | — | Generate helper |
| `--orm` | string | required | ORM for the controller |
| `--api` | bool | `false` | API controller (no views) |
| `--skip-routes` | bool | `false` | Don't add routes |
| `--skip-collision-check` | bool | `false` | Skip collision check |

When `--api` is used, skips ERB views and uses `api_controller.rb` template.

**Hooks:** `template_engine` (as `:scaffold`), `resource_route`, `test_framework` (as `:scaffold`), `helper` (as `:scaffold`).

### `migration`

**Usage:** `bin/rails generate migration NAME [field[:type][:index] ...] [options]`

**Positional:** `name` (required), `attributes` (array: `field[:type][:index]`)

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--timestamps` | bool | — | Add timestamps |
| `--primary-key-type` | string | — | PK type |
| `--database` / `--db` | string | — | Database for multi-db |

**Naming conventions determine the template:**

| Name pattern | Action | Template |
|-------------|--------|----------|
| `AddXxxToYyy` | Add columns | `add_column` |
| `RemoveXxxFromYyy` | Remove columns | `remove_column` |
| `CreateXxx` | Create table | `create_table` |
| `XxxJoinTable` | Join table | `create_join_table` |
| Other | Generic migration | plain `migration.rb` |

Migration filenames must only contain `[a-z0-9_]`.

### `job`

**Usage:** `bin/rails generate job NAME [options]`

**Positional:** `name` (required)

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--queue` | string | `default` | Queue name |
| `--parent` | string | `ApplicationJob` | Parent class |
| `--skip-collision-check` | bool | `false` | Skip collision check |

Strips `_job` suffix automatically. Creates `ApplicationJob` if missing.

**Hooks:** `test_framework`.

### `mailer`

**Usage:** `bin/rails generate mailer NAME [method method ...] [options]`

**Positional:** `name` (required), `actions` (array: method names)

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--skip-collision-check` | bool | `false` | Skip collision check |

Strips `_mailer` suffix automatically. Creates `ApplicationMailer` if missing. For each action, generates a view template (`app/views/<mailer>/<action>.html.erb`).

**Hooks:** `template_engine`, `test_framework`.

### `channel`

**Usage:** `bin/rails generate channel NAME [method method ...] [options]`

**Positional:** `name` (required), `actions` (array: method names)

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--assets` | bool | — | Generate JavaScript assets |
| `--skip-collision-check` | bool | `false` | Skip collision check |

Strips `_channel` suffix automatically. On first invocation creates shared files (`application_cable/channel.rb`, `application_cable/connection.rb`). If JavaScript is available, creates JS channel consumer and importmap/runtime setup.

**Hooks:** `test_framework`.

### `mailbox`

**Usage:** `bin/rails generate mailbox NAME [options]`

**Positional:** `name` (required)

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--skip-collision-check` | bool | `false` | Skip collision check |

Strips `_mailbox` suffix automatically. Creates `ApplicationMailbox` if missing.

**Hooks:** `test_framework`.

### `task`

**Usage:** `bin/rails generate task NAME [action action ...] [options]`

**Positional:** `name` (required), `actions` (array: action names)

No generator-specific flags. Creates `lib/tasks/<name>.rake` with a namespace block.

### `helper`

**Usage:** `bin/rails generate helper NAME [options]`

**Positional:** `name` (required)

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--skip-collision-check` | bool | `false` | Skip collision check |

Strips `_helper` suffix automatically. Creates `app/helpers/<name>_helper.rb`.

**Hooks:** `test_framework`.

### `system_test`

**Usage:** `bin/rails generate system_test NAME [options]`

**Positional:** `name` (required)

No generator-specific flags.

**Hooks:** `system_tests` (as `:system`).

### `integration_test`

**Usage:** `bin/rails generate integration_test NAME [options]`

**Positional:** `name` (required)

No generator-specific flags.

**Hooks:** `integration_tool` (as `:integration`).

### `jbuilder`

**Usage:** `bin/rails generate jbuilder NAME [field:type ...] [options]`

**Positional:** `name` (required), `attributes` (array: `field:type`)

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--timestamps` | bool | `true` | Include `created_at`/`updated_at` |

Creates `index.json.jbuilder`, `show.json.jbuilder`, `_<singular>.json.jbuilder`. Excludes `password`/`password_confirmation` if `digest` type present.

### `generator`

**Usage:** `bin/rails generate generator NAME [options]`

**Positional:** `name` (required)

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--namespace` | bool | `true` | Namespace under `lib/generators/<name>/` |
| `--skip-collision-check` | bool | `false` | Skip collision check |

Creates: `<name>_generator.rb`, `USAGE` file, `templates/` directory, test file.

**Hooks:** `test_framework`.

### `application_record`

**Usage:** `bin/rails generate application_record [options]`

No positional args. Creates `app/models/application_record.rb`. For engines, creates namespace-specific version.

### `benchmark`

**Usage:** `bin/rails generate benchmark NAME [report report ...] [options]`

**Positional:** `name` (required), `reports` (array, default: `["before", "after"]`)

Adds `benchmark-ips` gem to Gemfile. Creates `script/benchmarks/<name>.rb`.

### `active_record:migration` (alias for `migration`)

Same as `migration` — delegates to the ActiveRecord migration generator directly.

### `action_text:install`

**Usage:** `bin/rails generate action_text:install [options]`

No positional args. Installs Action Text (copies migration, creates `action_text.css`, etc.).

### `action_mailbox:install`

**Usage:** `bin/rails generate action_mailbox:install [options]`

No positional args. Installs Action Mailbox with its migrations.

### `active_record:authentication`

**Usage:** `bin/rails generate active_record:authentication [options]`

No positional args. Generates authentication model (Rails 8+), creates `User` model with `has_secure_password`.

## Generator Architecture Notes

- Every generator is a **Thor** class inheriting from `Rails::Generators::NamedBase` (has `name`) or `Rails::Generators::Base` (no `name`).
- Generators can **hook** into other generators (e.g., `scaffold` hooks into `scaffold_controller`, which hooks into `erb:scaffold`).
- Attributes with `references` or `belongs_to` type generate `belongs_to :name` in the model and a foreign key column.
- Attribute type `digest` generates `has_secure_password`.
- For field options appended with `{}` syntax, e.g. `name:string{default:"John"}`, Thor parses these as hash options.

## Edge Cases

- **`rails generate` without a generator name** lists all available generators.
- **`--skip-routes` on controller** with no actions given: no effect (no routes to skip).
- **`--force`** overwrites files without prompting.
- **`--pretend`** shows what would be created but writes nothing.
- **`--parent`** on model with a custom class and no `--database` skips migration generation.
- **Namespaced generators** work: `Admin::PostsController` creates `app/controllers/admin/posts_controller.rb`.
- **`rails destroy`** is the inverse of `rails generate` — same generator classes, same arguments.

## Related Skills

- For `rails generate` usage and global flags, see [cli-commands.md](cli-commands.md).
- For `rails new` (project generation), see [rails-new.md](rails-new.md).
