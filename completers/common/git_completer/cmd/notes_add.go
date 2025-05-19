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

	notes_addCmd.Flags().Bool("allow-empty", false, "Allow an empty note object to be stored")
	notes_addCmd.Flags().StringP("file", "F", "", "Take the note message from the given file")
	notes_addCmd.Flags().StringP("message", "m", "", "Use the given note message")
	notes_addCmd.Flags().StringP("reedit-message", "c", "", "Like -C, but with -c the editor is invoked")
	notes_addCmd.Flags().StringP("reuse-message", "C", "", "Take the given blob object as the note message")
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
