package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_state_hints_clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear hints (re-show on next trigger)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_state_hints_clearCmd).Standalone()

	config_state_hints_clearCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	config_state_hintsCmd.AddCommand(config_state_hints_clearCmd)

	// TODO complete names
}
