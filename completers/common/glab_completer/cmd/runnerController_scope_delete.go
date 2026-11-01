package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var runnerController_scope_deleteCmd = &cobra.Command{
	Use:   "delete <controller-id> [flags]",
	Short: "Delete a scope from a runner controller. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runnerController_scope_deleteCmd).Standalone()

	runnerController_scope_deleteCmd.Flags().BoolP("force", "f", false, "Skip confirmation prompt.")
	runnerController_scope_deleteCmd.Flags().Bool("instance", false, "Remove an instance-level scope.")
	runnerController_scope_deleteCmd.Flags().StringSlice("runner", nil, "Remove a runner-level scope for the specified runner ID. Multiple IDs can be comma-separated or specified by repeating the flag.")
	runnerController_scopeCmd.AddCommand(runnerController_scope_deleteCmd)
}
