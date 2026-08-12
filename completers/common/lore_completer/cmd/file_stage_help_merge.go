package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var file_stage_help_mergeCmd = &cobra.Command{
	Use:   "merge",
	Short: "Stage as a merge",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(file_stage_help_mergeCmd).Standalone()

	file_stage_helpCmd.AddCommand(file_stage_help_mergeCmd)
}
