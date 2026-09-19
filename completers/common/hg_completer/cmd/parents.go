package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/hg_completer/cmd/action"
	"github.com/spf13/cobra"
)

var parentsCmd = &cobra.Command{
	Use:     "parents",
	Short:   "show the parents of the working directory or revision (DEPRECATED)",
	GroupID: groups[group_change_navigation].ID,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(parentsCmd).Standalone()

	parentsCmd.Flags().StringP("rev", "r", "", "show parents of the specified revision")
	parentsCmd.Flags().String("style", "", "display using template map file (DEPRECATED)")
	parentsCmd.Flags().StringP("template", "T", "", "display with template")
	rootCmd.AddCommand(parentsCmd)

	carapace.Gen(parentsCmd).FlagCompletion(carapace.ActionMap{
		"rev": action.ActionRevisions(),
	})

	carapace.Gen(parentsCmd).PositionalAnyCompletion(
		action.ActionTrackedFiles(),
	)
}
