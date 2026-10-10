package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mr_note_draftCmd = &cobra.Command{
	Use:   "draft <command> [flags]",
	Short: "Manage your pending review comments on a merge request. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mr_note_draftCmd).Standalone()

	mr_noteCmd.AddCommand(mr_note_draftCmd)
}
