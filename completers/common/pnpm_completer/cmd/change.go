package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var changeCmd = &cobra.Command{
	Use:   "change",
	Short: "Record a change intent: which packages a change affects, the bump type for each, and a summary that becomes the changelog entry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(changeCmd).Standalone()

	changeCmd.Flags().String("bump", "", "Bump type for the named packages: none, patch, minor, major. \"none\" records an explicit decline — the change needs no release")
	changeCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	changeCmd.Flags().String("summary", "", "The summary for the changelog entry. Runs non-interactively when given together with package names")
	rootCmd.AddCommand(changeCmd)
}
