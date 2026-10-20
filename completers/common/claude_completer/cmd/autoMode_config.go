package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoMode_configCmd = &cobra.Command{
	Use:   "config",
	Short: "Print the effective auto mode config as JSON: your settings where set, defaults otherwise",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoMode_configCmd).Standalone()

	autoMode_configCmd.Flags().BoolP("help", "h", false, "Display help for command")
	autoModeCmd.AddCommand(autoMode_configCmd)
}
