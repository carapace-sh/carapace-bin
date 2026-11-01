package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_plugins_piCmd = &cobra.Command{
	Use:   "pi",
	Short: "Pi activity extension",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_plugins_piCmd).Standalone()

	config_plugins_piCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	config_pluginsCmd.AddCommand(config_plugins_piCmd)
}
