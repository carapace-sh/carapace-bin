package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var state_replaceProviderCmd = &cobra.Command{
	Use:   "replace-provider [options] FROM_PROVIDER_FQN TO_PROVIDER_FQN",
	Short: "Replace provider in the state",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(state_replaceProviderCmd).Standalone()

	state_replaceProviderCmd.Flags().BoolS("auto-approve", "auto-approve", false, "Skip interactive approval.")
	state_replaceProviderCmd.Flags().BoolS("ignore-remote-version", "ignore-remote-version", false, "A rare option used for the remote backend only.")
	state_replaceProviderCmd.Flags().BoolS("lock", "lock", false, "Don't hold a state lock during the operation.")
	state_replaceProviderCmd.Flags().StringS("lock-timeout", "lock-timeout", "", "Duration to retry a state lock.")
	stateCmd.AddCommand(state_replaceProviderCmd)

	state_replaceProviderCmd.Flag("lock-timeout").NoOptDefVal = " "

	// TODO positional completion
}
