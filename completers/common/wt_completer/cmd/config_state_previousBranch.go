package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_state_previousBranchCmd = &cobra.Command{
	Use:   "previous-branch",
	Short: "Previous branch (for wt switch -)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_state_previousBranchCmd).Standalone()

	config_state_previousBranchCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	config_stateCmd.AddCommand(config_state_previousBranchCmd)
}
