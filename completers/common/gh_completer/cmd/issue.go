package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var issueCmd = &cobra.Command{
	Use:     "issue <command>",
	Short:   "Manage issues",
	GroupID: "core",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(issueCmd).Standalone()
	issueCmd.AddGroup(
		&cobra.Group{ID: "Targeted commands", Title: ""},
		&cobra.Group{ID: "General commands", Title: ""},
	)

	issueCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository using the `[HOST/]OWNER/REPO` format")
	rootCmd.AddCommand(issueCmd)

	carapace.Gen(issueCmd).FlagCompletion(carapace.ActionMap{
		"repo": gh.ActionHostOwnerRepositories(),
	})
}
