package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_workingCopyCmd = &cobra.Command{
	Use:   "working-copy",
	Short: "Show information about the working copy state",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_workingCopyCmd).Standalone()

	debug_workingCopyCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	debugCmd.AddCommand(debug_workingCopyCmd)
}
