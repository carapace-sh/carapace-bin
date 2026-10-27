package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var notes_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove notes from one or more revisions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notes_removeCmd).Standalone()

	notes_removeCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	notesCmd.AddCommand(notes_removeCmd)

	carapace.Gen(notes_removeCmd).PositionalAnyCompletion(
		git.ActionRefs(git.RefOption{}.Default()),
	)
}
