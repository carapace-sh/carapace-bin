package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var notes_copyCmd = &cobra.Command{
	Use:   "copy",
	Short: "Copy the notes for the first object onto the second object",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notes_copyCmd).Standalone()

	notes_copyCmd.Flags().BoolP("force", "f", false, "Overwrite existing notes")
	notes_copyCmd.Flags().Bool("stdin", false, "Also read the object names to remove notes from the standard input")
	notesCmd.AddCommand(notes_copyCmd)

	carapace.Gen(notes_copyCmd).PositionalCompletion(
		git.ActionRefs(git.RefOption{}.Default()),
		git.ActionRefs(git.RefOption{}.Default()),
	)
}
