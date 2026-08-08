package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var runner_deleteCmd = &cobra.Command{
	Use:   "delete <runner-id>",
	Short: "Delete a runner.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runner_deleteCmd).Standalone()

	runner_deleteCmd.Flags().BoolP("force", "f", false, "Skip confirmation prompt.")
	runnerCmd.AddCommand(runner_deleteCmd)
}
