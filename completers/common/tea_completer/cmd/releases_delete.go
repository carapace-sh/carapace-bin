package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tea"
	"github.com/spf13/cobra"
)

var releases_deleteCmd = &cobra.Command{
	Use:     "delete",
	Short:   "Delete one or more releases",
	Aliases: []string{"rm"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(releases_deleteCmd).Standalone()

	releases_deleteCmd.Flags().BoolP("confirm", "y", false, "Confirm deletion (required)")
	releases_deleteCmd.Flags().Bool("delete-tag", false, "Also delete the git tag for this release")
	releases_deleteCmd.Flags().StringP("login", "l", "", "Use a different Gitea Login. Optional")
	releases_deleteCmd.Flags().StringP("output", "o", "", "Output format. (simple, table, csv, tsv, yaml, json)")
	releases_deleteCmd.Flags().StringP("remote", "R", "", "Discover Gitea login from remote. Optional")
	releases_deleteCmd.Flags().StringP("repo", "r", "", "Override local repository path or gitea repository slug to interact with. Optional")
	releasesCmd.AddCommand(releases_deleteCmd)

	// TODO completion
	carapace.Gen(releases_deleteCmd).FlagCompletion(carapace.ActionMap{
		"login":  tea.ActionLogins(),
		"output": tea.ActionOutputFormats(),
		"remote": git.ActionRemotes(),
	})
}
