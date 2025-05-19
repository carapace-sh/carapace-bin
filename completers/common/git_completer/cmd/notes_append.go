package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var notes_appendCmd = &cobra.Command{
	Use:   "append",
	Short: "Append to the notes of an existing object",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notes_appendCmd).Standalone()

	notes_appendCmd.Flags().Bool("allow-empty", false, "Allow an empty note object to be stored")
	notes_appendCmd.Flags().StringP("file", "F", "", "Take the note message from the given file")
	notes_appendCmd.Flags().StringP("message", "m", "", "Use the given note message")
	notes_appendCmd.Flags().StringP("reedit-message", "c", "", "Like -C, but with -c the editor is invoked")
	notes_appendCmd.Flags().StringP("reuse-message", "C", "", "Take the given blob object as the note message")
	notesCmd.AddCommand(notes_appendCmd)

	carapace.Gen(notes_appendCmd).FlagCompletion(carapace.ActionMap{
		"file":           carapace.ActionFiles(),
		"reedit-message": git.ActionRefs(git.RefOption{}.Default()),
		"reuse-message":  git.ActionRefs(git.RefOption{}.Default()),
	})

	carapace.Gen(notes_appendCmd).PositionalCompletion(
		git.ActionRefs(git.RefOption{}.Default()),
	)

}
