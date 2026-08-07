package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var runnerController_deleteCmd = &cobra.Command{
	Use:   "delete <id> [flags]",
	Short: "Delete a runner controller. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runnerController_deleteCmd).Standalone()

	runnerController_deleteCmd.Flags().BoolP("force", "f", false, "Skip confirmation prompt.")
	runnerControllerCmd.AddCommand(runnerController_deleteCmd)
}
