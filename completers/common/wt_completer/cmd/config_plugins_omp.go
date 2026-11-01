package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_plugins_ompCmd = &cobra.Command{
	Use:   "omp",
	Short: "oh-my-pi activity hook",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_plugins_ompCmd).Standalone()

	config_plugins_ompCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	config_pluginsCmd.AddCommand(config_plugins_ompCmd)
}
