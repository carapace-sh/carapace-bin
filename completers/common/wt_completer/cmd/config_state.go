package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_stateCmd = &cobra.Command{
	Use:   "state",
	Short: "Manage internal data and cache",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_stateCmd).Standalone()

	config_stateCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	configCmd.AddCommand(config_stateCmd)
}
