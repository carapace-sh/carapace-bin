package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_ai_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show current AI provider configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_ai_showCmd).Standalone()

	config_ai_showCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	config_aiCmd.AddCommand(config_ai_showCmd)
}
