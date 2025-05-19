package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var state_pushCmd = &cobra.Command{
	Use:   "push [options] PATH",
	Short: "Update remote state from a local state file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(state_pushCmd).Standalone()

	state_pushCmd.Flags().BoolS("force", "force", false, "Write the state even if lineages don't match or the remote serial is higher")
	state_pushCmd.Flags().BoolS("lock", "lock", false, "Don't hold a state lock during the operation")
	state_pushCmd.Flags().StringS("lock-timeout", "lock-timeout", "", "Duration to retry a state lock")
	stateCmd.AddCommand(state_pushCmd)

	state_pushCmd.Flag("lock-timeout").NoOptDefVal = " "

	carapace.Gen(state_pushCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
