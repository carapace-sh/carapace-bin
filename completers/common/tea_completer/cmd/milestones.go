package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tea"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var milestonesCmd = &cobra.Command{
	Use:     "milestones",
	Short:   "List and create milestones",
	GroupID: "ENTITIES",
	Aliases: []string{"milestone", "ms"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(milestonesCmd).Standalone()

	milestonesCmd.Flags().StringP("fields", "f", "", "Comma-separated list of fields to print. Available values:")
	milestonesCmd.Flags().String("limit", "", "specify limit of items per page")
	milestonesCmd.Flags().StringP("login", "l", "", "Use a different Gitea Login. Optional")
	milestonesCmd.Flags().StringP("output", "o", "", "Output format. (simple, table, csv, tsv, yaml, json)")
	milestonesCmd.Flags().StringP("page", "p", "", "specify page, default is 1")
	milestonesCmd.Flags().StringP("remote", "R", "", "Discover Gitea login from remote. Optional")
	milestonesCmd.Flags().StringP("repo", "r", "", "Override local repository path or gitea repository slug to interact with. Optional")
	milestonesCmd.Flags().String("state", "", "Filter by milestone state (all|open|closed)")
	rootCmd.AddCommand(milestonesCmd)

	//TODO completion
	carapace.Gen(milestonesCmd).FlagCompletion(carapace.ActionMap{
		"fields": tea.ActionMilestoneFields().UniqueList(","),
		"login":  tea.ActionLogins(),
		"output": tea.ActionOutputFormats(),
		"remote": git.ActionRemotes(),
		"state":  carapace.ActionValues("all", "open", "closed").StyleF(style.ForKeyword),
	})
}
