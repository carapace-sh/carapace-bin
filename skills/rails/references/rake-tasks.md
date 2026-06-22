# Rake Tasks

Miscellaneous Rake tasks accessible via `bin/rails <task>`. These are loaded from `railties/lib/rails/tasks/` and various component-specific task files.

> **Source of truth**: `railties/lib/rails/tasks/`, `activestorage/lib/tasks/`, `actionmailbox/lib/tasks/`, `actiontext/lib/tasks/`.

## Log Tasks

### `log:clear`

Truncate all log files in `log/` to zero bytes.

| Env var | Description |
|---------|-------------|
| `LOGS=test,development` | Only clear specific logs |

## Temp File Tasks

### `tmp:clear`

Clear all temp files (`tmp/cache`, `tmp/sockets`, `tmp/screenshots`).

| Variant | Description |
|---------|-------------|
| `tmp:create` | Create tmp dirs (cache, sockets, pids) |
| `tmp:cache:clear` | Clear `tmp/cache` |
| `tmp:sockets:clear` | Clear `tmp/sockets` |
| `tmp:pids:clear` | Clear `tmp/pids` |
| `tmp:screenshots:clear` | Clear `tmp/screenshots` (Capybara screenshots) |
| `tmp:storage:clear` | Clear `tmp/storage` |

## Time Zone Tasks

### `time:zones:all`

Print all time zones known to Active Support (with UTC offsets and Rails identifiers).

### `time:zones:local`

Print time zones matching the system's local UTC offset.

### `time:zones:us`

Print US time zones.

## Framework Tasks

### `app:template`

Apply a template to the application.

| Env var | Description |
|---------|-------------|
| `LOCATION=/path/to/template` | **Required.** Path or URL to template file |

### `app:templates:copy`

Copy all Rails templates to the application directory for customization.

## Yarn / JavaScript Tasks

### `yarn:install`

Install JavaScript dependencies from `package.json` via Yarn.

### `javascript:install`

Install JavaScript dependencies.

## Zeitwerk

### `zeitwerk:check`

Check the project's file structure for Zeitwerk autoloading compliance. Run in CI.

## Action Mailbox Tasks

### `action_mailbox:ingress:exim`

Relay an inbound email from Exim to Action Mailbox.

| Env var | Description |
|---------|-------------|
| `URL` | **Required.** The ingress URL |
| `INGRESS_PASSWORD` | **Required.** The shared password |

### `action_mailbox:ingress:postfix`

Relay an inbound email from Postfix to Action Mailbox.

| Env var | Description |
|---------|-------------|
| `URL` | **Required.** The ingress URL |
| `INGRESS_PASSWORD` | **Required.** The shared password |

### `action_mailbox:ingress:qmail`

Relay an inbound email from Qmail to Action Mailbox.

| Env var | Description |
|---------|-------------|
| `URL` | **Required.** The ingress URL |
| `INGRESS_PASSWORD` | **Required.** The shared password |

### `action_mailbox:install`

Install Action Mailbox migrations and configuration.

## Action Text Tasks

### `action_text:install`

Copy over the migration, stylesheet, and JavaScript files for Action Text / Trix.

## Active Storage Tasks

### `active_storage:install`

Copy over the migration needed for Active Storage.

### `active_storage:update`

Copy over migrations needed for upgrading Active Storage.

## Engine Migration Tasks

### `railties:install:migrations`

Copy missing migrations from Railties (e.g. engines) into the application's `db/migrate/`.

## Notes on Completion

- **`log:clear`** supports the `LOGS` env var for filtering. Common values: `development`, `test`, `production`.
- **`time:zones:`** subcommands are static: `all`, `local`, `us`. No dynamic list.
- **Action Mailbox ingress tasks** require both `URL` and `INGRESS_PASSWORD` env vars to be set.
- **`app:template`** requires `LOCATION` — can be a file path or URL.
- **`railties:install:migrations`** takes no args/flags. It reads migrations from engine gems.

## Related Skills

- For `rails test` subcommands including `test:db`, see [testing.md](testing.md).
- For `rails db:*` subcommands, see [database.md](database.md).
- For `rails secret`, `version`, `about`, `stats`, `middleware`, etc., see [cli-commands.md](cli-commands.md).
