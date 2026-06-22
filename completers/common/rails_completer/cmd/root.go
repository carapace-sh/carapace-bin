package cmd

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rails_completer/cmd/action"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:                "rails",
	Short:              "Rails CLI",
	Long:               "https://guides.rubyonrails.org/command_line.html",
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("help", "h", false, "Show help")
	rootCmd.Flags().StringP("tasks", "T", "", "Filter tasks by pattern")

	carapace.Gen(rootCmd).PositionalCompletion(
		action.ActionSubcommands(),
	)

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) == 0 {
				return carapace.ActionValues()
			}

			cmd := c.Args[0]
			switch cmd {
			// Primary commands
			case "new":
				return carapace.ActionExecute(newCmd).Shift(1)
			case "plugin":
				return carapace.ActionExecute(pluginCmd).Shift(1)
			case "generate", "g", "destroy", "d":
				return carapace.ActionExecute(generateCmd).Shift(1)
			case "server", "s":
				return carapace.ActionExecute(serverCmd).Shift(1)
			case "console", "c":
				return carapace.ActionExecute(consoleCmd).Shift(1)
			case "runner", "r":
				return carapace.ActionExecute(runnerCmd).Shift(1)
			case "dbconsole", "db":
				return carapace.ActionExecute(dbconsoleCmd).Shift(1)
			case "routes":
				return carapace.ActionExecute(routesCmd).Shift(1)
			case "test", "t":
				return carapace.ActionExecute(testCmd).Shift(1)
			case "notes":
				return carapace.ActionExecute(notesCmd).Shift(1)
			case "query":
				return carapace.ActionExecute(queryCmd).Shift(1)
			case "devcontainer":
				return carapace.ActionExecute(devcontainerCmd).Shift(1)

			// Credentials subcommands
			case "credentials:edit":
				return carapace.ActionExecute(credentialsEditCmd).Shift(1)
			case "credentials:show":
				return carapace.ActionExecute(credentialsShowCmd).Shift(1)
			case "credentials:diff":
				return carapace.ActionExecute(credentialsDiffCmd).Shift(1)
			case "credentials:fetch":
				return carapace.ActionExecute(credentialsFetchCmd).Shift(1)

			// Encrypted subcommands
			case "encrypted:edit":
				return carapace.ActionExecute(encryptedEditCmd).Shift(1)
			case "encrypted:show":
				return carapace.ActionExecute(encryptedShowCmd).Shift(1)

			// db:* namespace
			case "db:system:change":
				return carapace.ActionExecute(dbSystemChangeCmd).Shift(1)
			case "db:version", "db:create", "db:drop", "db:prepare",
				"db:seed", "db:seed:replant", "db:setup", "db:reset", "db:abort_if_pending_migrations",
				"db:purge", "db:purge:all", "db:truncate_all", "db:charset", "db:collation",
				"db:environment:set", "db:forward", "db:rollback":
				return carapace.ActionExecute(dbCmd).Shift(1)

			// db:migrate subcommands
			case "db:migrate", "db:migrate:redo", "db:migrate:up", "db:migrate:down", "db:migrate:status", "db:migrate:reset":
				return carapace.ActionExecute(dbMigrateCmd).Shift(1)

			// db:schema subcommands
			case "db:schema:dump", "db:schema:load", "db:schema:cache:clear", "db:schema:cache:dump":
				return carapace.ActionExecute(dbSchemaCmd).Shift(1)

			// db:test subcommands
			case "db:test:prepare", "db:test:load_schema", "db:test:purge":
				return carapace.ActionExecute(dbTestTaskCmd).Shift(1)

			// db:fixtures subcommands
			case "db:fixtures:load", "db:fixtures:identify":
				return carapace.ActionExecute(dbFixturesCmd).Shift(1)

			// db:encryption subcommands
			case "db:encryption:init":
				return carapace.ActionExecute(dbEncryptionCmd).Shift(1)

			// db:create/drop variants
			case "db:create:all", "db:drop:all",
				"db:setup:all", "db:reset:all":
				return carapace.ActionExecute(dbCmd).Shift(1)

			// test subcommands
			case "test:all", "test:system", "test:functionals", "test:units",
				"test:generators", "test:models", "test:helpers", "test:channels",
				"test:controllers", "test:mailers", "test:integration", "test:jobs",
				"test:mailboxes", "test:db":
				return carapace.ActionExecute(testSubCmd).Shift(1)

			// assets subcommands
			case "assets:clean", "assets:clobber", "assets:environment",
				"assets:precompile", "assets:reveal", "assets:reveal:full":
				return carapace.ActionExecute(assetsCmd).Shift(1)

			// log/temp tasks
			case "log:clear", "tmp:clear", "tmp:create",
				"tmp:cache:clear", "tmp:sockets:clear", "tmp:pids:clear",
				"tmp:screenshots:clear", "tmp:storage:clear":
				return carapace.ActionExecute(rakeCmd).Shift(1)

			// time zones
			case "time:zones", "time:zones:all", "time:zones:local", "time:zones:us":
				return carapace.ActionExecute(rakeCmd).Shift(1)

			// Standalone commands (with --environment)
			case "boot", "initializers":
				return carapace.ActionExecute(standaloneEnvCmd).Shift(1)

			// Standalone commands (no special flags)
			case "about", "dev:cache", "middleware",
				"restart", "secret", "stats", "version":
				return carapace.ActionExecute(standaloneCmd).Shift(1)

			// Framework rake tasks
			case "zeitwerk:check", "yarn:install", "javascript:install",
				"railties:install:migrations", "app:template", "app:update",
				"app:templates:copy", "cache_digests:dependencies",
				"cache_digests:nested_dependencies",
				"importmap:install", "secrets:edit", "secrets:setup", "secrets:show":
				return carapace.ActionExecute(rakeCmd).Shift(1)

			// CSS install
			case "css:install", "css:install:tailwind", "css:install:bootstrap",
				"css:install:postcss", "css:install:sass":
				return carapace.ActionExecute(rakeCmd).Shift(1)

			// Action Mailbox
			case "action_mailbox:install", "action_mailbox:install:migrations",
				"action_mailbox:ingress:exim", "action_mailbox:ingress:postfix",
				"action_mailbox:ingress:qmail":
				return carapace.ActionExecute(rakeCmd).Shift(1)

			// Action Text
			case "action_text:install", "action_text:install:migrations":
				return carapace.ActionExecute(rakeCmd).Shift(1)

			// Active Storage
			case "active_storage:install", "active_storage:update":
				return carapace.ActionExecute(rakeCmd).Shift(1)

			// Stimulus
			case "stimulus:install", "stimulus:install:importmap", "stimulus:install:node":
				return carapace.ActionExecute(rakeCmd).Shift(1)

			// Turbo
			case "turbo:install", "turbo:install:importmap", "turbo:install:node",
				"turbo:install:redis", "turbo:install:bun":
				return carapace.ActionExecute(rakeCmd).Shift(1)

			// TailwindCSS
			case "tailwindcss:build", "tailwindcss:clobber",
				"tailwindcss:install", "tailwindcss:watch":
				return carapace.ActionExecute(rakeCmd).Shift(1)

			default:
				// Dynamic dispatch for prefixes
				if strings.HasPrefix(cmd, "db:") {
					return carapace.ActionExecute(dbCmd).Shift(1)
				}
				if strings.HasPrefix(cmd, "test:") {
					return carapace.ActionExecute(testSubCmd).Shift(1)
				}
				if strings.HasPrefix(cmd, "assets:") {
					return carapace.ActionExecute(assetsCmd).Shift(1)
				}
				// Catch-all for custom rake tasks
				return carapace.ActionExecute(rakeCmd).Shift(1)
			}
		}),
	)
}
