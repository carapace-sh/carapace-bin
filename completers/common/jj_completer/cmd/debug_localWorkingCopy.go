package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_localWorkingCopyCmd = &cobra.Command{
	Use:   "local-working-copy",
	Short: "Show information about the local working copy state",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_localWorkingCopyCmd).Standalone()

	debug_localWorkingCopyCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	debugCmd.AddCommand(debug_localWorkingCopyCmd)
}