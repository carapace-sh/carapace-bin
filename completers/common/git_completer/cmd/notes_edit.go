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
	notes_editCmd.Flags().StringP("file", "F", "", "Take the note message from the given file")
	notes_editCmd.Flags().StringP("message", "m", "", "Use the given note message")
	notes_editCmd.Flags().StringP("reedit-message", "c", "", "Like -C, but with -c the editor is invoked")
	notes_editCmd.Flags().StringP("reuse-message", "C", "", "Take the given blob object as the note message")
	notes_editCmd.Flags().String("separator", "", "Insert <paragraph-break> between paragraphs")
	notes_editCmd.Flags().Bool("stripspace", false, "Remove unnecessary whitespace")
	notesCmd.AddCommand(notes_editCmd)

	carapace.Gen(notes_editCmd).FlagCompletion(carapace.ActionMap{
		"file":           carapace.ActionFiles(),
		"reedit-message": git.ActionRefs(git.RefOption{}.Default()),
		"reuse-message":  git.ActionRefs(git.RefOption{}.Default()),
	})

	carapace.Gen(notes_editCmd).PositionalCompletion(
		git.ActionRefs(git.RefOption{}.Default()),
	)
}
