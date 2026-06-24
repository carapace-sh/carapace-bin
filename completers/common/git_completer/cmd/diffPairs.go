package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/git_completer/cmd/common"
	"github.com/spf13/cobra"
)

var diffPairsCmd = &cobra.Command{
	Use:     "diff-pairs",
	Short:   "Diff from raw pair input",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_interrogator].ID,
}

func init() {
	carapace.Gen(diffPairsCmd).Standalone()

	common.AddDiffFlags(diffPairsCmd)
	rootCmd.AddCommand(diffPairsCmd)
}
