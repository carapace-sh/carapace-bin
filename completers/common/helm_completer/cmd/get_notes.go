package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/helm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var get_notesCmd = &cobra.Command{
	Use:   "notes",
	Short: "download the notes for a named release",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(get_notesCmd).Standalone()
	get_notesCmd.Flags().Int("revision", 0, "get the named release with revision")
	getCmd.AddCommand(get_notesCmd)

	carapace.Gen(get_notesCmd).PositionalCompletion(
		action.ActionReleases(get_notesCmd),
	)
}
