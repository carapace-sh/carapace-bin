package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_state_markerCmd = &cobra.Command{
	Use:   "marker",
	Short: "Branch markers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_state_markerCmd).Standalone()

	config_state_markerCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	config_stateCmd.AddCommand(config_state_markerCmd)
}
