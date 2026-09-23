package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/android"
	"github.com/spf13/cobra"
)

var removeUserCmd = &cobra.Command{
	Use:   "remove-user",
	Short: "Remove the user with the given identifier",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(removeUserCmd).Standalone()

	removeUserCmd.Flags().Bool("set-ephemeral-if-in-use", false, "mark the user as ephemeral if it is currently running")
	removeUserCmd.Flags().Bool("wait", false, "wait until user is removed")

	rootCmd.AddCommand(removeUserCmd)

	carapace.Gen(removeUserCmd).PositionalCompletion(
		android.ActionUsers(),
	)
}
