package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var notes_help_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List notes attached to commits",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notes_help_listCmd).Standalone()

	notes_helpCmd.AddCommand(notes_help_listCmd)
}
