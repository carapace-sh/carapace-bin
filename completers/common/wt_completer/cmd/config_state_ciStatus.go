package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_state_ciStatusCmd = &cobra.Command{
	Use:   "ci-status",
	Short: "CI status cache",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_state_ciStatusCmd).Standalone()

	config_state_ciStatusCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	config_stateCmd.AddCommand(config_state_ciStatusCmd)
}
