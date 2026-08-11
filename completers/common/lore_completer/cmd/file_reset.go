package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/lore_completer/cmd/action"
	"github.com/spf13/cobra"
)

var file_resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset changes to a path or file to the current revision, discarding your local changes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(file_resetCmd).Standalone()

	file_resetCmd.Flags().BoolP("help", "h", false, "Print help")
	file_resetCmd.Flags().String("last-merged-from", "", "If given, the files will be reset to the last point of merge from this branch, or the branch point from this branch if no merge has been performed")
	file_resetCmd.Flags().Bool("purge", false, "Delete untracked files")
	file_resetCmd.Flags().String("revision", "", "Revision to reset files to")
	file_resetCmd.Flags().String("targets", "", "Path to a targets file containing all the paths to all files")
	fileCmd.AddCommand(file_resetCmd)

	carapace.Gen(file_resetCmd).FlagCompletion(carapace.ActionMap{
		"last-merged-from": action.ActionBranches(file_resetCmd),
		"revision":         action.ActionRevisions(file_resetCmd),
		"targets":          carapace.ActionFiles(),
	})

	carapace.Gen(file_resetCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
