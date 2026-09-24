package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_state_hints_getCmd = &cobra.Command{
	Use:   "get",
	Short: "List hints that have been shown",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_state_hints_getCmd).Standalone()

	config_state_hints_getCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	config_state_hintsCmd.AddCommand(config_state_hints_getCmd)
}
