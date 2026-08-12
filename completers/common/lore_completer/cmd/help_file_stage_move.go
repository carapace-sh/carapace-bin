package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_file_stage_moveCmd = &cobra.Command{
	Use:   "move",
	Short: "Move or rename a file or directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_file_stage_moveCmd).Standalone()

	help_file_stageCmd.AddCommand(help_file_stage_moveCmd)
}
