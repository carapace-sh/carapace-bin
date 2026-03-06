package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_state_defaultBranchCmd = &cobra.Command{
	Use:   "default-branch",
	Short: "Default branch detection and override",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_state_defaultBranchCmd).Standalone()

	config_state_defaultBranchCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	config_stateCmd.AddCommand(config_state_defaultBranchCmd)
}
