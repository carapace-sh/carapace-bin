package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debugpathcompleteCmd = &cobra.Command{
	Use:    "debugpathcomplete",
	Short:  "FILESPEC...",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugpathcompleteCmd).Standalone()

	debugpathcompleteCmd.Flags().BoolP("added", "a", false, "show only added files")
	debugpathcompleteCmd.Flags().BoolP("full", "f", false, "complete an entire path")
	debugpathcompleteCmd.Flags().BoolP("normal", "n", false, "show only normal files")
	debugpathcompleteCmd.Flags().BoolP("removed", "r", false, "show only removed files")
	rootCmd.AddCommand(debugpathcompleteCmd)
}
