package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_aiCmd = &cobra.Command{
	Use:   "ai",
	Short: "View and configure AI provider settings",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_aiCmd).Standalone()

	config_aiCmd.Flags().Bool("global", false, "Configure global user git config")
	config_aiCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	config_aiCmd.Flags().Bool("local", false, "Configure local repository git config instead of global user config")
	configCmd.AddCommand(config_aiCmd)
}
