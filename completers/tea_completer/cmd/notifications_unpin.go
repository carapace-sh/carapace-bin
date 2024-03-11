package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tea"
	"github.com/spf13/cobra"
)

var notifications_unpinCmd = &cobra.Command{
	Use:   "unpin",
	Short: "Unpin all pinned or a specific notification",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notifications_unpinCmd).Standalone()

	notifications_unpinCmd.Flags().String("limit", "", "specify limit of items per page")
	notifications_unpinCmd.Flags().StringP("login", "l", "", "Use a different Gitea Login. Optional")
	notifications_unpinCmd.Flags().BoolP("mine", "m", false, "Show notifications across all your repositories instead of the current repository only")
	notifications_unpinCmd.Flags().StringP("output", "o", "", "Output format. (simple, table, csv, tsv, yaml, json)")
	notifications_unpinCmd.Flags().StringP("page", "p", "", "specify page, default is 1")
	notifications_unpinCmd.Flags().StringP("remote", "R", "", "Discover Gitea login from remote. Optional")
	notifications_unpinCmd.Flags().StringP("repo", "r", "", "Override local repository path or gitea repository slug to interact with. Optional")
	notifications_unpinCmd.Flags().StringP("states", "s", "", "Comma-separated list of notification states to filter by. Available values:")
	notificationsCmd.AddCommand(notifications_unpinCmd)

	// TODO completion
	carapace.Gen(notifications_unpinCmd).FlagCompletion(carapace.ActionMap{
		"login":  tea.ActionLogins(),
		"output": tea.ActionOutputFormats(),
		"remote": git.ActionRemotes(),
		"states": tea.ActionNotificationStates().UniqueList(","),
	})
}
