package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var file_stage_mergeCmd = &cobra.Command{
	Use:   "merge",
	Short: "Stage as a merge",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(file_stage_mergeCmd).Standalone()

	file_stage_mergeCmd.Flags().BoolP("help", "h", false, "Print help")
	file_stage_mergeCmd.Flags().String("targets", "", "Path to a targets file containing all the paths to all files")
	file_stageCmd.AddCommand(file_stage_mergeCmd)
}
