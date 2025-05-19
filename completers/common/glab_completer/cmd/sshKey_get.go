package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var sshKey_getCmd = &cobra.Command{
	Use:   "get <key-id>",
	Short: "Returns a single SSH key specified by the ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sshKey_getCmd).Standalone()

	sshKey_getCmd.Flags().StringP("page", "p", "", "Page number.")
	sshKey_getCmd.Flags().StringP("per-page", "P", "", "Number of items to list per page.")
	sshKeyCmd.AddCommand(sshKey_getCmd)

	carapace.Gen(sshKey_getCmd).PositionalCompletion(
		action.ActionSshKeyIds(sshKeyCmd),
	)
}
