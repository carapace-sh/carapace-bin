package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var notes_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove the notes for given objects",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notes_removeCmd).Standalone()

	notes_removeCmd.Flags().Bool("ignore-missing", false, "Do not consider it an error to request removing non existing notes")
	notes_removeCmd.Flags().Bool("stdin", false, "Also read the object names to remove notes from the standard input")
	notesCmd.AddCommand(notes_removeCmd)

	carapace.Gen(notes_removeCmd).PositionalAnyCompletion(
		git.ActionRefs(git.RefOption{}.Default()),
	)
}
