package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var releaseCmd = &cobra.Command{
	Use:   "release <command>",
	Short: "Manage GitHub releases",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	releaseCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository using the `[HOST/]OWNER/REPO` format")
	rootCmd.AddCommand(releaseCmd)

	carapace.Gen(releaseCmd).FlagCompletion(carapace.ActionMap{
		"repo": action.ActionOwnerRepositories(issueCmd),
	})
}
