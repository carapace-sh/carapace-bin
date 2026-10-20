package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var daemons_restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "[experimental] Restart project daemons",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(daemons_restartCmd).Standalone()

	daemons_restartCmd.Flags().BoolP("help", "h", false, "Print help")
	daemonsCmd.AddCommand(daemons_restartCmd)
}
