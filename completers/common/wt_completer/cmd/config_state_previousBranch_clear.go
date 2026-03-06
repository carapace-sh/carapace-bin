package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_state_previousBranch_clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear the previous branch",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_state_previousBranch_clearCmd).Standalone()

	config_state_previousBranch_clearCmd.Flags().BoolP("help", "h", false, "Print help")
	config_state_previousBranchCmd.AddCommand(config_state_previousBranch_clearCmd)
}
