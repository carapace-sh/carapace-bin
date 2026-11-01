package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var file_help_stage_moveCmd = &cobra.Command{
	Use:   "move",
	Short: "Move or rename a file or directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(file_help_stage_moveCmd).Standalone()

	file_help_stageCmd.AddCommand(file_help_stage_moveCmd)
}
