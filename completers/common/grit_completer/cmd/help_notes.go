package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_notesCmd = &cobra.Command{
	Use:   "notes",
	Short: "List, show, add, or remove Git notes on commits",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_notesCmd).Standalone()

	helpCmd.AddCommand(help_notesCmd)
}
