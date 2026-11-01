package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var notes_help_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add or replace a note on a revision",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notes_help_addCmd).Standalone()

	notes_helpCmd.AddCommand(notes_help_addCmd)
}
