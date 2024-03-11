package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var notesCmd = &cobra.Command{
	Use:     "notes",
	Short:   "Add or inspect object notes",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_main].ID,
}

func init() {
	carapace.Gen(notesCmd).Standalone()

	notesCmd.Flags().String("ref", "", "use notes from <notes-ref>")
	rootCmd.AddCommand(notesCmd)
}
