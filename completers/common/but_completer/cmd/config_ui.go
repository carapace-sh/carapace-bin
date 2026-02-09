package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_uiCmd = &cobra.Command{
	Use:   "ui",
	Short: "View and configure UI preferences",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_uiCmd).Standalone()

	config_uiCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	configCmd.AddCommand(config_uiCmd)
}
