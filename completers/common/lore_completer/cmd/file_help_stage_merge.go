package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var file_help_stage_mergeCmd = &cobra.Command{
	Use:   "merge",
	Short: "Stage as a merge",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(file_help_stage_mergeCmd).Standalone()

	file_help_stageCmd.AddCommand(file_help_stage_mergeCmd)
}
