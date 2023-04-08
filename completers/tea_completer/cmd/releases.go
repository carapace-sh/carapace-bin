package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/tea"
	"github.com/spf13/cobra"
)

var releasesCmd = &cobra.Command{
	Use:     "releases",
	Short:   "Manage releases",
	GroupID: "ENTITIES",
	Aliases: []string{"release", "r"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(releasesCmd).Standalone()

	releasesCmd.Flags().StringP("login", "l", "", "Use a different Gitea Login. Optional")
	releasesCmd.Flags().StringP("output", "o", "", "Output format. (simple, table, csv, tsv, yaml, json)")
	releasesCmd.Flags().StringP("remote", "R", "", "Discover Gitea login from remote. Optional")
	releasesCmd.Flags().StringP("repo", "r", "", "Override local repository path or gitea repository slug to interact with. Optional")
	rootCmd.AddCommand(releasesCmd)

	// TODO completion
	carapace.Gen(releasesCmd).FlagCompletion(carapace.ActionMap{
		"output": tea.ActionOutputFormats(),
		"remote": git.ActionRemotes(),
	})
}
