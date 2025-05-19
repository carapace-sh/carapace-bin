package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var workflowCmd = &cobra.Command{
	Use:     "workflow <command>",
	Short:   "View details about GitHub Actions workflows",
	GroupID: "actions",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workflowCmd).Standalone()

	workflowCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository using the `[HOST/]OWNER/REPO` format")
	rootCmd.AddCommand(workflowCmd)

	carapace.Gen(workflowCmd).FlagCompletion(carapace.ActionMap{
		"repo": gh.ActionHostOwnerRepositories(),
	})
}
