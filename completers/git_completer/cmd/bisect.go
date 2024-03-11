package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var bisectCmd = &cobra.Command{
	Use:     "bisect",
	Short:   "Use binary search to find the commit that introduced a bug",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_main].ID,
}

func init() {
	carapace.Gen(bisectCmd).Standalone()

	rootCmd.AddCommand(bisectCmd)

	carapace.Gen(bisect_skipCmd).PositionalAnyCompletion(
		git.ActionRefRanges(git.RefOption{}.Default()),
	)
}
