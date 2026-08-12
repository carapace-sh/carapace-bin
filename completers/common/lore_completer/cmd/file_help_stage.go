package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var file_help_stageCmd = &cobra.Command{
	Use:   "stage",
	Short: "Stage changes for commit",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(file_help_stageCmd).Standalone()

	file_helpCmd.AddCommand(file_help_stageCmd)
}
