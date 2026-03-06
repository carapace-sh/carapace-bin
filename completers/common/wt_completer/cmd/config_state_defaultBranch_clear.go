package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_state_defaultBranch_clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear the default branch cache",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_state_defaultBranch_clearCmd).Standalone()

	config_state_defaultBranch_clearCmd.Flags().BoolP("help", "h", false, "Print help")
	config_state_defaultBranchCmd.AddCommand(config_state_defaultBranch_clearCmd)
}
