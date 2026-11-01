package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var notes_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show the note for a revision",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notes_showCmd).Standalone()

	notes_showCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	notesCmd.AddCommand(notes_showCmd)

	carapace.Gen(notes_showCmd).PositionalCompletion(
		git.ActionRefs(git.RefOption{}.Default()),
	)
}
