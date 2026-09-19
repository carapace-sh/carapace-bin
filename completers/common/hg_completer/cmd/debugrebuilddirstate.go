package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/hg_completer/cmd/action"
	"github.com/spf13/cobra"
)

var debugrebuilddirstateCmd = &cobra.Command{
	Use:    "debugrebuilddirstate",
	Short:  "rebuild the dirstate as it would look like for the given revision",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugrebuilddirstateCmd).Standalone()

	debugrebuilddirstateCmd.Flags().Bool("minimal", false, "only rebuild files that are inconsistent with the working copy parent")
	debugrebuilddirstateCmd.Flags().StringP("rev", "r", "", "revision to rebuild to")
	rootCmd.AddCommand(debugrebuilddirstateCmd)

	carapace.Gen(debugrebuilddirstateCmd).FlagCompletion(carapace.ActionMap{
		"rev": action.ActionRevisions(),
	})
}
