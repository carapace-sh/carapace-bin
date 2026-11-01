package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_file_stageCmd = &cobra.Command{
	Use:   "stage",
	Short: "Stage changes for commit",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_file_stageCmd).Standalone()

	help_fileCmd.AddCommand(help_file_stageCmd)
}
