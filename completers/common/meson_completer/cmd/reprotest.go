package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var reprotestCmd = &cobra.Command{
	Use:   "reprotest",
	Short: "Test if project builds reproducibly",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reprotestCmd).Standalone()

	reprotestCmd.Flags().Bool("intermediaries", false, "Check intermediate files.")
	applyCommonFlagCompletion(reprotestCmd)
	rootCmd.AddCommand(reprotestCmd)
}
