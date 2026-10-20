package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var daemons_tuiCmd = &cobra.Command{
	Use:   "tui",
	Short: "[experimental] Open the pitchfork dashboard",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(daemons_tuiCmd).Standalone()

	daemons_tuiCmd.Flags().BoolP("help", "h", false, "Print help")
	daemonsCmd.AddCommand(daemons_tuiCmd)
}
