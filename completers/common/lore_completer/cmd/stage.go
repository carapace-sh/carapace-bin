package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stageCmd = &cobra.Command{
	Use:   "stage",
	Short: "Stage changes for commit",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stageCmd).Standalone()

	stageCmd.Flags().String("case", "", "Case change handling")
	stageCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	stageCmd.Flags().Bool("scan", false, "Walk the filesystem under the given paths to detect modified, added, and deleted files")
	stageCmd.Flags().String("targets", "", "Path to a targets file containing all the paths to all files")
	rootCmd.AddCommand(stageCmd)
}
