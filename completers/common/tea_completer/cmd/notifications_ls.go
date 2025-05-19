package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tea"
	"github.com/spf13/cobra"
)

var notifications_lsCmd = &cobra.Command{
	Use:     "ls",
	Short:   "List notifications",
	Aliases: []string{"list"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notifications_lsCmd).Standalone()

	notifications_lsCmd.Flags().StringP("fields", "f", "", "Comma-separated list of fields to print. Available values:")
	notifications_lsCmd.Flags().String("limit", "", "specify limit of items per page")
	notifications_lsCmd.Flags().StringP("login", "l", "", "Use a different Gitea Login. Optional")
	notifications_lsCmd.Flags().BoolP("mine", "m", false, "Show notifications across all your repositories instead of the current repository only")
	notifications_lsCmd.Flags().StringP("output", "o", "", "Output format. (simple, table, csv, tsv, yaml, json)")
	notifications_lsCmd.Flags().StringP("page", "p", "", "specify page, default is 1")
	notifications_lsCmd.Flags().StringP("remote", "R", "", "Discover Gitea login from remote. Optional")
	notifications_lsCmd.Flags().StringP("repo", "r", "", "Override local repository path or gitea repository slug to interact with. Optional")
	notifications_lsCmd.Flags().StringP("states", "s", "", "Comma-separated list of notification states to filter by. Available values:")
	notifications_lsCmd.Flags().StringP("types", "t", "", "Comma-separated list of subject types to filter by. Available values:")
	notificationsCmd.AddCommand(notifications_lsCmd)

	// TODO completion
	carapace.Gen(notifications_lsCmd).FlagCompletion(carapace.ActionMap{
		"fields": tea.ActionNotificationFields().UniqueList(","),
		"login":  tea.ActionLogins(),
		"output": tea.ActionOutputFormats(),
		"remote": git.ActionRemotes(),
		"states": tea.ActionNotificationStates().UniqueList(","),
		"types":  tea.ActionNotificationTypes().UniqueList(","),
	})
}
