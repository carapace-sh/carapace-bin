package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var file_stageCmd = &cobra.Command{
	Use:   "stage",
	Short: "Stage changes for commit",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(file_stageCmd).Standalone()

	file_stageCmd.Flags().String("case", "", "Case change handling")
	file_stageCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	file_stageCmd.Flags().Bool("scan", false, "Walk the filesystem under the given paths to detect modified, added, and deleted files")
	file_stageCmd.Flags().String("targets", "", "Path to a targets file containing all the paths to all files")
	fileCmd.AddCommand(file_stageCmd)

	carapace.Gen(file_stageCmd).FlagCompletion(carapace.ActionMap{
		"case":    carapace.ActionValues("error", "keep", "rename"),
		"targets": carapace.ActionFiles(),
	})

	carapace.Gen(file_stageCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
