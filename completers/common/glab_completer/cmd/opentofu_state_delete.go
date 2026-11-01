package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opentofu_state_deleteCmd = &cobra.Command{
	Use:   "delete <state> [<serial>] [flags]",
	Short: "Delete a state or a specific version of a state.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opentofu_state_deleteCmd).Standalone()

	opentofu_state_deleteCmd.Flags().BoolP("force", "f", false, "Force delete the state without prompting.")
	opentofu_stateCmd.AddCommand(opentofu_state_deleteCmd)
}
