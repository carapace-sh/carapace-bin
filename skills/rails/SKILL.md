---
name: rails
description: >
  Use when working with the Rails CLI — commands, generators, database tasks,
  test subcommands, rake tasks, flags, and argument completion. Triggers on:
  "rails", "Ruby on Rails", "bin/rails", "rails new", "rails generate",
  "rails server", "rails console", "rails dbconsole", "rails runner", "rails test",
  "rails routes", "rails destroy", "rails db", "rails credentials",
  "rails middleware", "rails notes", "rails dev:cache", "rails restart",
  "rails secret", "rails stats", "rails about", "rails boot",
  "rails initializers", "rails plugin", "rails tmp", "rails version",
  "rails devcontainer", "rails encrypted", "rails zeitwerk", "rails yarn",
  "rails scaffold", "rails model", "rails migration", "rails controller",
  "rails resource", "rails job", "rails mailer", "rails channel",
  "railsrc", ".railsrc", "config/database.yml".
user-invocable: true
---

# Rails CLI In-Depth Reference

Complete reference for the Rails CLI surface — commands, generators, database tasks, rake tasks, flags, positional arguments, and shell completion targets. Built for creating a carapace completer.

## Data Flow

```
bin/rails COMMAND [args] [options]
  ├── rails new          → AppGenerator (Thor)
  ├── rails generate     → GenerateCommand → Generator subclasses
  ├── rails server       → ServerCommand (Rack/Puma)
  ├── rails console      → ConsoleCommand (IRB)
  ├── rails dbconsole    → DbconsoleCommand (psql/mysql/sqlite3)
  ├── rails runner       → RunnerCommand (eval code)
  ├── rails test         → TestCommand → Minitest
  ├── rails routes       → RoutesCommand (router inspection)
  ├── rails db:          → Rake tasks (Active Record)
  ├── rails credentials  → CredentialsCommand (encrypted YAML)
  ├── rails encrypted    → EncryptedCommand (generic encrypted files)
  ├── rails notes        → NotesCommand (annotation grep)
  ├── rails query        → QueryCommand (Active Record REPL)
  └── rails <task>       → Rake fallback (Thor delegate)
```

## Sub-Resources

Load the reference that matches your task. When in doubt, load multiple references.

| Keywords | Reference |
|----------|----------|
| rails new, AppGenerator, --database, --javascript, --css, --api, --minimal, --skip-*, --template, --rc, .railsrc, --dev, --edge, --main, --devcontainer, PluginGenerator | [references/rails-new.md](references/rails-new.md) |
| rails generate, rails destroy, generator subcommands, model, scaffold, resource, controller, migration, job, mailer, channel, mailbox, task, helper, jbuilder, system_test, integration_test, generator, application_record, benchmark, scaffold_controller, --skip-routes, --parent, --timestamps, --api, --queue, --assets, --actions, field:type:index, --pretend, --force, --skip, orm hook, template_engine hook, test_framework hook | [references/generators.md](references/generators.md) |
| rails server, rails console, rails runner, rails dbconsole, rails routes, middleware, initializers, about, stats, boot, dev:cache, restart, secret, version, encrypted, credentials, notes, query, devcontainer, --port, --binding, --sandbox, --environment, --database, --unused, --expanded, --controller, --grep, --enroll, --disenroll | [references/cli-commands.md](references/cli-commands.md) |
| rails db:*, database rake tasks, migrate, schema, seed, rollback, create, drop, setup, reset, prepare, version, encryption, truncate, charset, collation, fixtures, test:db, multi-db, db:system:change, --to, VERSION, STEP, SCOPE, FIXTURES, FIXTURES_DIR | [references/database.md](references/database.md) |
| rails test, test:system, test:all, test:functionals, test:units, test:generators, test:models, test:helpers, test:channels, test:controllers, test:mailers, test:integration, test:jobs, test:mailboxes, test:db, --name, --verbose, --seed, --warnings, --environment | [references/testing.md](references/testing.md) |
| rails log:clear, tmp:*, time:zones, yarn:install, zeitwerk:check, action_mailbox:*, action_text:install, active_storage:*, app:template, railties:install:migrations, javascript:install | [references/rake-tasks.md](references/rake-tasks.md) |

## Quick Guide

- **How do I complete `rails new` flags?** → [references/rails-new.md](references/rails-new.md)
- **How do I complete `rails generate <thing>`?** → [references/generators.md](references/generators.md)
- **What flags does `rails server` accept?** → [references/cli-commands.md](references/cli-commands.md)
- **What are all the `db:*` subcommands?** → [references/database.md](references/database.md)
- **How do I complete test paths and subcommands?** → [references/testing.md](references/testing.md)
- **What `--environment` values are valid?** → The string `environment` — value can be any string, typically `development`, `test`, `production`, or `staging`. Not enum-limited.
- **How does `--database` work in `rails new`?** → [references/rails-new.md](references/rails-new.md)
- **What field types does `generate model` accept?** → [references/generators.md](references/generators.md)
- **How does multi-db completion work for `db:*` tasks?** → [references/database.md](references/database.md)
- **What are the `rails credentials:*` subcommands?** → [references/cli-commands.md](references/cli-commands.md)
- **How does `db:system:change --to=` work?** → [references/database.md](references/database.md)
- **What are the aliases for rails commands?** → [references/cli-commands.md](references/cli-commands.md)
- **How do I complete `rails encrypted:edit/show`?** → [references/cli-commands.md](references/cli-commands.md)
