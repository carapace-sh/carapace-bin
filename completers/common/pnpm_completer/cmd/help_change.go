package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_changeCmd = &cobra.Command{
	Use:   "change",
	Short: "Record a change intent: which packages a change affects, the bump type for each, and a summary that becomes the changelog entry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_changeCmd).Standalone()

	helpCmd.AddCommand(help_changeCmd)
}
