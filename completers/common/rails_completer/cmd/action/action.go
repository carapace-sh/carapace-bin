package action

import (
	"strings"
	"time"

	"github.com/carapace-sh/carapace"
)

// ActionSubcommands completes rails subcommands using colon-separated namespaces
func ActionSubcommands() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.Batch(
					ActionCommands(),
					ActionNamespaces().Suffix(":").Usage("namespace"),
				).ToA().NoSpace(':')
			case 1:
				return subcommandsForNamespace(c.Parts[0])
			case 2:
				return subcommandsForNamespace(c.Parts[0] + ":" + c.Parts[1])
			case 3:
				return subcommandsForNamespace(c.Parts[0] + ":" + c.Parts[1] + ":" + c.Parts[2])
			default:
				prefix := strings.Join(c.Parts[:len(c.Parts)-1], ":")
				return subcommandsForNamespace(prefix)
			}
		})
	})
}

// subcommandsForNamespace returns completions for a given namespace prefix.
func subcommandsForNamespace(namespace string) carapace.Action {
	switch namespace {
	case "db":
		return carapace.Batch(
			ActionDbTasks(),
			ActionDbMultiDbTasks().Suffix(":"),
			ActionDbNamespaces().Suffix(":"),
		).ToA().NoSpace(':')
	case "db:migrate":
		return carapace.Batch(
			ActionDbMigrateTasks().Suffix(":"),
			ActionDatabaseConfigs(),
		).ToA().NoSpace(':')
	case "db:schema":
		return carapace.Batch(
			ActionDbSchemaTasks().Suffix(":"),
			ActionDbSchemaNamespaces().Suffix(":"),
		).ToA().NoSpace(':')
	case "db:schema:cache":
		return ActionDbSchemaCacheTasks()
	case "db:schema:dump":
		return ActionDatabaseConfigs()
	case "db:schema:load":
		return ActionDatabaseConfigs()
	case "db:migrate:redo", "db:migrate:up", "db:migrate:down", "db:migrate:status", "db:migrate:reset":
		return ActionDatabaseConfigs()
	case "db:rollback":
		return ActionDatabaseConfigs()
	case "db:forward":
		return ActionDatabaseConfigs()
	case "db:version":
		return ActionDatabaseConfigs()
	case "db:purge":
		return ActionDatabaseConfigs()
	case "db:test":
		return carapace.Batch(
			ActionDbTestTasks().Suffix(":"),
			ActionDatabaseConfigs(),
		).ToA().NoSpace(':')
	case "db:test:prepare", "db:test:load_schema", "db:test:purge":
		return ActionDatabaseConfigs()
	case "db:abort_if_pending_migrations":
		return ActionDatabaseConfigs()
	case "db:seed":
		return ActionDbSeedTasks()
	case "db:fixtures":
		return ActionDbFixtureTasks()
	case "db:create":
		return carapace.Batch(
			carapace.ActionValues("all").Tag("db create tasks"),
			ActionDatabaseConfigs(),
		).ToA()
	case "db:drop":
		return carapace.Batch(
			carapace.ActionValues("all").Tag("db drop tasks"),
			ActionDatabaseConfigs(),
		).ToA()
	case "db:setup":
		return carapace.Batch(
			carapace.ActionValues("all").Tag("db setup tasks"),
			ActionDatabaseConfigs(),
		).ToA()
	case "db:reset":
		return carapace.Batch(
			carapace.ActionValues("all").Tag("db reset tasks"),
			ActionDatabaseConfigs(),
		).ToA()
	case "db:encryption":
		return carapace.ActionValues("init").Tag("db encryption tasks")
	case "db:system":
		return carapace.ActionValues("change").Tag("db system tasks")
	case "db:environment":
		return carapace.ActionValues("set").Tag("db environment tasks")
	case "assets":
		return carapace.Batch(
			ActionAssetTasks(),
			ActionAssetNamespaces().Suffix(":"),
		).ToA().NoSpace(':')
	case "assets:reveal":
		return ActionAssetRevealTasks()
	case "test":
		return ActionTestSubcommands()
	case "log":
		return ActionLogTasks()
	case "tmp":
		return carapace.Batch(
			ActionTmpTasks(),
			ActionTmpNamespaces().Suffix(":").Usage("namespace"),
		).ToA().NoSpace(':')
	case "tmp:cache", "tmp:sockets", "tmp:pids", "tmp:screenshots", "tmp:storage":
		return ActionTmpSubdirTasks()
	case "time":
		return carapace.ActionValues("zones").Suffix(":").Tag("time namespaces").NoSpace(':')
	case "time:zones":
		return ActionTimezoneTasks()
	case "credentials":
		return ActionCredentialTasks()
	case "encrypted":
		return ActionEncryptedTasks()
	case "dev":
		return ActionDevTasks()
	case "action_mailbox":
		return carapace.Batch(
			ActionMailboxTasks().Suffix(":"),
		).ToA().NoSpace(':')
	case "action_mailbox:ingress":
		return ActionMailboxIngressTasks()
	case "action_mailbox:install":
		return ActionMailboxInstallTasks()
	case "action_text":
		return carapace.ActionValues("install").Suffix(":").Tag("action text namespaces").NoSpace(':')
	case "action_text:install":
		return ActionTextInstallTasks()
	case "active_storage":
		return ActionStorageTasks()
	case "app":
		return carapace.Batch(
			ActionAppTasks(),
			ActionAppNamespaces().Suffix(":"),
		).ToA().NoSpace(':')
	case "app:templates":
		return ActionAppTemplatesTasks()
	case "cache_digests":
		return ActionCacheDigestTasks()
	case "css":
		return ActionCssNamespaces().Suffix(":").NoSpace(':')
	case "css:install":
		return ActionCssInstallTasks()
	case "importmap":
		return ActionImportmapTasks()
	case "javascript":
		return ActionJavascriptTasks()
	case "railties":
		return ActionRailtiesNamespaces().Suffix(":").NoSpace(':')
	case "railties:install":
		return ActionRailtiesInstallTasks()
	case "secrets":
		return ActionSecretsTasks()
	case "stimulus":
		return ActionStimulusNamespaces().Suffix(":").NoSpace(':')
	case "stimulus:install":
		return ActionStimulusInstallTasks()
	case "tailwindcss":
		return ActionTailwindcssTasks()
	case "turbo":
		return ActionTurboNamespaces().Suffix(":").NoSpace(':')
	case "turbo:install":
		return ActionTurboInstallTasks()
	case "yarn":
		return ActionYarnTasks()
	case "zeitwerk":
		return ActionZeitwerkTasks()
	default:
		return carapace.ActionExecCommand("rails", "-T")(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			var leaves []string
			namespaceSet := make(map[string]bool)
			var namespaces []string
			prefix := namespace + ":"
			for _, line := range lines {
				if !strings.HasPrefix(line, "rails ") {
					continue
				}
				task := strings.Fields(line)[1]
				suffix, ok := strings.CutPrefix(task, prefix)
				if !ok {
					continue
				}
				if before, after, found := strings.Cut(suffix, ":"); found {
					_ = after
					if !namespaceSet[before] {
						namespaceSet[before] = true
						namespaces = append(namespaces, before)
					}
				} else {
					leaves = append(leaves, suffix)
				}
			}
			if len(namespaces) > 0 || len(leaves) > 0 {
				return carapace.Batch(
					carapace.ActionValues(leaves...).Tag("dynamic tasks"),
					carapace.ActionValues(namespaces...).Suffix(":").Tag("dynamic namespaces"),
				).ToA().NoSpace(':')
			}
			return carapace.ActionValues()
		}).Cache(5 * time.Minute)
	}
}
