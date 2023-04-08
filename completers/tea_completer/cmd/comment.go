package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/tea"
	"github.com/spf13/cobra"
)

var commentCmd = &cobra.Command{
	Use:     "comment",
	Short:   "Add a comment to an issue / pr",
	GroupID: "ENTITIES",
	Aliases: []string{"c"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(commentCmd).Standalone()

	commentCmd.Flags().StringP("login", "l", "", "Use a different Gitea Login. Optional")
	commentCmd.Flags().StringP("output", "o", "", "Output format. (simple, table, csv, tsv, yaml, json)")
	commentCmd.Flags().StringP("remote", "R", "", "Discover Gitea login from remote. Optional")
	commentCmd.Flags().StringP("repo", "r", "", "Override local repository path or gitea repository slug to interact with. Optional")
	rootCmd.AddCommand(commentCmd)

	// TODO completion
	carapace.Gen(commentCmd).FlagCompletion(carapace.ActionMap{
		"output": tea.ActionOutputFormats(),
		"remote": git.ActionRemotes(),
	})
}
