package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var releaseCmd = &cobra.Command{
	Use:     "release <command>",
	Short:   "Manage releases",
	GroupID: "core",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(releaseCmd).Standalone()
	releaseCmd.AddGroup(
		&cobra.Group{ID: "general", Title: "General commands"},
		&cobra.Group{ID: "targeted", Title: "Targeted commands"},
	)

	releaseCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository using the `[HOST/]OWNER/REPO` format")
	rootCmd.AddCommand(releaseCmd)

	carapace.Gen(releaseCmd).FlagCompletion(carapace.ActionMap{
		"repo": action.ActionRepoOverride(issueCmd),
	})
}
