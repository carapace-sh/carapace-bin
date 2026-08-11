package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stage_moveCmd = &cobra.Command{
	Use:   "move",
	Short: "Move or rename a file or directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stage_moveCmd).Standalone()

	stage_moveCmd.Flags().BoolP("help", "h", false, "Print help")
	stageCmd.AddCommand(stage_moveCmd)
}
