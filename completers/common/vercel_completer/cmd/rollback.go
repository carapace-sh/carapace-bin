package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/vercel_completer/cmd/action"
	"github.com/spf13/cobra"
)

var rollbackCmd = &cobra.Command{
	Use:   "rollback",
	Short: "Quickly revert back to a previous deployment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rollbackCmd).Standalone()

	rollbackCmd.Flags().String("timeout", "", "Time to wait for rollback completion")
	rollbackCmd.Flags().Bool("yes", false, "Skip confirmation")

	rootCmd.AddCommand(rollbackCmd)

	carapace.Gen(rollbackCmd).PositionalCompletion(
		action.ActionDeployments(rollbackCmd),
	)
}
