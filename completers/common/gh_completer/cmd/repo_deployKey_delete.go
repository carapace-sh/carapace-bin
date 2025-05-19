package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var repo_deployKey_deleteCmd = &cobra.Command{
	Use:   "delete <key-id>",
	Short: "Delete a deploy key from a GitHub repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_deployKey_deleteCmd).Standalone()

	repo_deployKeyCmd.AddCommand(repo_deployKey_deleteCmd)

	carapace.Gen(repo_deployKey_deleteCmd).PositionalCompletion(
		action.ActionDeployKeys(repo_deployKey_deleteCmd),
	)
}
