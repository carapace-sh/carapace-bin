package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var notes_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add notes for a given object",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notes_addCmd).Standalone()

	notes_addCmd.Flags().Bool("allow-empty", false, "allow an empty note object to be stored")
	notes_addCmd.Flags().BoolP("edit", "e", false, "edit note message in editor")
	notes_addCmd.Flags().StringP("file", "F", "", "take the note message from the given file")
	notes_addCmd.Flags().StringP("message", "m", "", "use the given note message")
	notes_addCmd.Flags().StringP("reedit-message", "c", "", "like -C, but with -c the editor is invoked")
	notes_addCmd.Flags().StringP("reuse-message", "C", "", "take the given blob object as the note message")
	notes_addCmd.Flags().String("separator", "", "insert <paragraph-break> between paragraphs")
	notes_addCmd.Flags().Bool("stripspace", false, "remove unnecessary whitespace")
	notesCmd.AddCommand(notes_addCmd)

	carapace.Gen(notes_addCmd).FlagCompletion(carapace.ActionMap{
		"file":           carapace.ActionFiles(),
		"reedit-message": git.ActionRefs(git.RefOption{}.Default()),
		"reuse-message":  git.ActionRefs(git.RefOption{}.Default()),
	})

	carapace.Gen(notes_addCmd).PositionalCompletion(
		git.ActionRefs(git.RefOption{}.Default()),
	)
}
