package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var discussionCmd = &cobra.Command{
	Use:     "discussion <command>",
	Short:   "Work with GitHub Discussions (preview)",
	GroupID: "core",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(discussionCmd).Standalone()
	discussionCmd.AddGroup(
		&cobra.Group{ID: "Targeted commands", Title: ""},
		&cobra.Group{ID: "General commands", Title: ""},
	)

	discussionCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository using the `[HOST/]OWNER/REPO` format")
	rootCmd.AddCommand(discussionCmd)

	carapace.Gen(discussionCmd).FlagCompletion(carapace.ActionMap{
		"repo": gh.ActionHostOwnerRepositories(),
	})
}
