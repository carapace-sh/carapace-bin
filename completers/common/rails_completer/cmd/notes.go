package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var notesCmd = &cobra.Command{
	Use:   "notes",
	Short: "Search code for annotation comments",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notesCmd).Standalone()

	notesCmd.Flags().StringSliceP("annotations", "a", []string{"FIXME", "OPTIMIZE", "TODO"}, "Annotation tags")
	notesCmd.Flags().BoolP("help", "h", false, "Show help")

	carapace.Gen(notesCmd).FlagCompletion(carapace.ActionMap{
		"annotations": carapace.ActionValues("FIXME", "OPTIMIZE", "TODO", "NOTE", "BUG"),
	})
}
