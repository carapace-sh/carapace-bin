package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tea"
	"github.com/spf13/cobra"
)

var organizationsCmd = &cobra.Command{
	Use:     "organizations",
	Short:   "List, create, delete organizations",
	GroupID: "ENTITIES",
	Aliases: []string{"organization", "org"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(organizationsCmd).Standalone()

	organizationsCmd.Flags().String("limit", "", "specify limit of items per page")
	organizationsCmd.Flags().StringP("login", "l", "", "Use a different Gitea Login. Optional")
	organizationsCmd.Flags().StringP("output", "o", "", "Output format. (simple, table, csv, tsv, yaml, json)")
	organizationsCmd.Flags().StringP("page", "p", "", "specify page, default is 1")
	organizationsCmd.Flags().StringP("remote", "R", "", "Discover Gitea login from remote. Optional")
	organizationsCmd.Flags().StringP("repo", "r", "", "Override local repository path or gitea repository slug to interact with. Optional")
	rootCmd.AddCommand(organizationsCmd)

	// TODO completion
	carapace.Gen(organizationsCmd).FlagCompletion(carapace.ActionMap{
		"login":  tea.ActionLogins(),
		"output": tea.ActionOutputFormats(),
		"remote": git.ActionRemotes(),
	})
}
