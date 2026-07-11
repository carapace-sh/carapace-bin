package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var showconfCmd = &cobra.Command{
	Use:   "showconf",
	Short: "Shows the current configuration of a given WireGuard interface, for use with setconf",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(showconfCmd).Standalone()
	rootCmd.AddCommand(showconfCmd)
	carapace.Gen(showconfCmd).PositionalCompletion(ActionInterfaces())
}