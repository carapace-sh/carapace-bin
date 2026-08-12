package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_file_stage_mergeCmd = &cobra.Command{
	Use:   "merge",
	Short: "Stage as a merge",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_file_stage_mergeCmd).Standalone()

	help_file_stageCmd.AddCommand(help_file_stage_mergeCmd)
}
