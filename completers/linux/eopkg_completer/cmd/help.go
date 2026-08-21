package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var helpCmd = &cobra.Command{
	Use:     "help <subcommand?>",
	Aliases: []string{"?"},
	Short:   "display help topics, or help for the given subcommand",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(helpCmd).Standalone()

	rootCmd.AddCommand(helpCmd)

	carapace.Gen(helpCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return carapace.ActionValues(
				"add-repo", "autoremove", "blame", "build", "check", "clean",
				"configure-pending", "delete-cache", "delta", "disable-repo",
				"emerge", "enable-repo", "fetch", "graph", "help", "history",
				"index", "info", "install", "list-available", "list-components",
				"list-installed", "list-newest", "list-pending", "list-repo",
				"list-sources", "list-upgrades", "rebuild-db", "remove",
				"remove-orphans", "remove-repo", "search", "search-file",
				"update-repo", "upgrade",
			)
		}),
	)
}
