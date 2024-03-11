package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tea"
	"github.com/spf13/cobra"
)

var notificationsCmd = &cobra.Command{
	Use:     "notifications",
	Short:   "Show notifications",
	GroupID: "HELPERS",
	Aliases: []string{"notification", "n"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notificationsCmd).Standalone()

	notificationsCmd.Flags().StringP("fields", "f", "", "Comma-separated list of fields to print. Available values:")
	notificationsCmd.Flags().String("limit", "", "specify limit of items per page")
	notificationsCmd.Flags().StringP("login", "l", "", "Use a different Gitea Login. Optional")
	notificationsCmd.Flags().BoolP("mine", "m", false, "Show notifications across all your repositories instead of the current repository only")
	notificationsCmd.Flags().StringP("output", "o", "", "Output format. (simple, table, csv, tsv, yaml, json)")
	notificationsCmd.Flags().StringP("page", "p", "", "specify page, default is 1")
	notificationsCmd.Flags().StringP("remote", "R", "", "Discover Gitea login from remote. Optional")
	notificationsCmd.Flags().StringP("repo", "r", "", "Override local repository path or gitea repository slug to interact with. Optional")
	notificationsCmd.Flags().StringP("states", "s", "", "Comma-separated list of notification states to filter by. Available values:")
	notificationsCmd.Flags().StringP("types", "t", "", "Comma-separated list of subject types to filter by. Available values:")
	rootCmd.AddCommand(notificationsCmd)

	// TODO completion
	carapace.Gen(notificationsCmd).FlagCompletion(carapace.ActionMap{
		"fields": tea.ActionNotificationFields().UniqueList(","),
		"login":  tea.ActionLogins(),
		"output": tea.ActionOutputFormats(),
		"remote": git.ActionRemotes(),
		"states": tea.ActionNotificationStates().UniqueList(","),
		"types":  tea.ActionNotificationTypes().UniqueList(","),
	})
}
