package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finishCmd = &cobra.Command{
	Use:   "finish",
	Short: "Merge and delete a feature branch",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finishCmd).Standalone()

	finishCmd.Flags().Bool("help", false, "show help")
	finishCmd.Flags().Bool("squash", false, "Squash merge on finish")

	rootCmd.AddCommand(finishCmd)
}
