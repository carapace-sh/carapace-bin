package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var notes_help_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove notes from one or more revisions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notes_help_removeCmd).Standalone()

	notes_helpCmd.AddCommand(notes_help_removeCmd)
}
