package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var notes_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add or replace a note on a revision",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notes_addCmd).Standalone()

	notes_addCmd.Flags().Bool("force", false, "Replace an existing note")
	notes_addCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	notes_addCmd.Flags().StringP("message", "m", "", "Note text")
	notes_addCmd.MarkFlagRequired("message")
	notesCmd.AddCommand(notes_addCmd)

	carapace.Gen(notes_addCmd).PositionalCompletion(
		git.ActionRefs(git.RefOption{}.Default()),
	)
}
