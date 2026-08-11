package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_stage_moveCmd = &cobra.Command{
	Use:   "move",
	Short: "Move or rename a file or directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_stage_moveCmd).Standalone()

	help_stageCmd.AddCommand(help_stage_moveCmd)
}
