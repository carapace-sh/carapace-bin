package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_state_clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear all stored state",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_state_clearCmd).Standalone()

	config_state_clearCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	config_stateCmd.AddCommand(config_state_clearCmd)
}
