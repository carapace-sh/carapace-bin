package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var sshKey_deleteCmd = &cobra.Command{
	Use:   "delete [<key-id>]",
	Short: "Deletes a single SSH key specified by the ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sshKey_deleteCmd).Standalone()

	sshKey_deleteCmd.Flags().StringP("page", "p", "1", "Page number.")
	sshKey_deleteCmd.Flags().StringP("per-page", "P", "30", "Number of items to list per page.")
	sshKeyCmd.AddCommand(sshKey_deleteCmd)

	carapace.Gen(sshKey_deleteCmd).PositionalCompletion(
		action.ActionSshKeyIds(sshKey_deleteCmd),
	)
}
