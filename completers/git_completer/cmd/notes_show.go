package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var notes_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show the notes for a given object",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notes_showCmd).Standalone()

	notesCmd.AddCommand(notes_showCmd)

	carapace.Gen(notes_showCmd).PositionalCompletion(
		git.ActionRefs(git.RefOption{}.Default()),
	)
}
