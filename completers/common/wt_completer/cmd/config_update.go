package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update deprecated config settings",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_updateCmd).Standalone()

	config_updateCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	config_updateCmd.Flags().BoolP("yes", "y", false, "Skip confirmation prompt")
	configCmd.AddCommand(config_updateCmd)
}
