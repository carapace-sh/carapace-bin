package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repo_deployKey_addCmd = &cobra.Command{
	Use:   "add <key-file>",
	Short: "Add a deploy key to a GitHub repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_deployKey_addCmd).Standalone()

	repo_deployKey_addCmd.Flags().BoolP("allow-write", "w", false, "Allow write access for the key")
	repo_deployKey_addCmd.Flags().StringP("title", "t", "", "Title of the new key")
	repo_deployKeyCmd.AddCommand(repo_deployKey_addCmd)

	carapace.Gen(repo_deployKey_addCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
