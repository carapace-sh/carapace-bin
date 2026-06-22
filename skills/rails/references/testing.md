# Test Commands

All `rails test:*` subcommands. Rails wraps Minitest with custom runner logic for default test paths, line-number filtering, and subcommand-based test directory selection.

> **Source of truth**: `railties/lib/rails/commands/test/test_command.rb` and `railties/lib/rails/test_unit/runner.rb`.

## `rails test` (`t`)

**Usage:** `bin/rails test [PATHS...] [options]`

Runs tests in default test directories (excluding `test/system`).

### Positional Arguments

Paths to test files or directories. Can be:

| Pattern | Example | Behavior |
|---------|---------|----------|
| File path | `test/models/user_test.rb` | Run that file |
| File:line | `test/models/user_test.rb:27` | Run test at that line |
| File:range | `test/models/user_test.rb:10-20` | Run tests in line range |
| Directory | `test/controllers` | Run all tests in that dir |
| Multiple | `test/controllers test/integration/login_test.rb` | Run all specified |

### Minitest Options

| Flag | Alias | Type | Description |
|------|-------|------|-------------|
| `--name` | `-n` | string/pattern | Filter by test name (regex) |
| `--verbose` | `-v` | bool | Verbose output |
| `--seed` | — | int | Random seed |
| `--no-plugins` | — | bool | Disable plugins |
| `--warnings` | `-w` | bool | Run with Ruby warnings enabled |
| `--environment` | `-e` | string | Run in specified environment |
| `--help` | `-h` | — | Show Minitest help |

## Test Subcommands

### `test:all`

Run all tests including system tests.

### `test:system`

Run system tests only (in `test/system/`).

### `test:functionals`

Run functional tests (`test/controllers/`, `test/mailers/`, `test/functional/`).

### `test:units`

Run unit tests (`test/models/`, `test/helpers/`, `test/unit/`).

### `test:generators`

Run generator tests (`test/lib/generators/`).

### `test:models`

Run model tests (`test/models/`).

### `test:helpers`

Run helper tests (`test/helpers/`).

### `test:channels`

Run channel tests (`test/channels/`).

### `test:controllers`

Run controller tests (`test/controllers/`).

### `test:mailers`

Run mailer tests (`test/mailers/`).

### `test:integration`

Run integration tests (`test/integration/`).

### `test:jobs`

Run job tests (`test/jobs/`).

### `test:mailboxes`

Run mailbox tests (`test/mailboxes/`).

### `test:db`

Reset test database and run `bin/rails test` (equivalent to `rails db:test:prepare && rails test`).

## Runner Architecture

- `Rails::TestUnit::Runner` manages test loading
- Default test file glob: `test/**/*_test.rb`
- System tests excluded from default `rails test` (use `rails test:system` or `rails test:all`)
- Line number filtering uses Minitest's `--name` option with a generated regex from the test file's AST
- `--environment` defaults to `test`

## Edge Cases

- **Line numbers** (`file:27`) match the test method definition line, not any arbitrary line. If the line doesn't match a test method, it runs the entire file.
- **`test:db`** is a convenience alias for `db:test:prepare && test`. Use it to avoid running migrations separately.
- **All subcommands** accept the same Minitest options as `rails test`.
- **No subcommand for system tests by default** — only `test:system` targets them.
- **`test:functionals` and `test:units`** are legacy terms kept for backwards compatibility; prefer the named subcommands (`test:controllers`, `test:models`, etc.).

## Related Skills

- For `rails test:db` which delegates to `db:test:prepare`, see [database.md](database.md).
- For the general CLI command structure, see [cli-commands.md](cli-commands.md).
