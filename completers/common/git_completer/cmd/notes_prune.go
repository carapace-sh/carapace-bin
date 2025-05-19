package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var notes_pruneCmd = &cobra.Command{
	Use:   "prune",
	Short: "Remove all notes for non-existing/unreachable objects",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notes_pruneCmd).Standalone()

	notes_pruneCmd.Flags().BoolP("dry-run", "n", false, "Do not remove anything, just report")
	notes_pruneCmd.Flags().BoolP("verbose", "v", false, "When merging notes, be more verbose")
	notesCmd.AddCommand(notes_pruneCmd)
}
