package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_stage_mergeCmd = &cobra.Command{
	Use:   "merge",
	Short: "Stage as a merge",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_stage_mergeCmd).Standalone()

	help_stageCmd.AddCommand(help_stage_mergeCmd)
}
