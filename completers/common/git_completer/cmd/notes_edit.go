package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var notes_editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit the notes for a given object",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notes_editCmd).Standalone()

	notes_editCmd.Flags().Bool("allow-empty", false, "Allow an empty note object to be stored")
	notesCmd.AddCommand(notes_editCmd)

	carapace.Gen(notes_editCmd).PositionalCompletion(
		git.ActionRefs(git.RefOption{}.Default()),
	)
}
