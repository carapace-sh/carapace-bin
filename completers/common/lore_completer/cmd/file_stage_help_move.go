package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var file_stage_help_moveCmd = &cobra.Command{
	Use:   "move",
	Short: "Move or rename a file or directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(file_stage_help_moveCmd).Standalone()

	file_stage_helpCmd.AddCommand(file_stage_help_moveCmd)
}
