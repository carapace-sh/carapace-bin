package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var notes_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List notes attached to commits",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notes_listCmd).Standalone()

	notes_listCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	notesCmd.AddCommand(notes_listCmd)
}
