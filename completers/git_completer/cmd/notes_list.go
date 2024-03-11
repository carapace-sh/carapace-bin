package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var notes_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the notes object for a given object",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notes_listCmd).Standalone()

	notesCmd.AddCommand(notes_listCmd)

	carapace.Gen(notes_listCmd).PositionalCompletion(
		git.ActionRefs(git.RefOption{}.Default()),
	)
}
