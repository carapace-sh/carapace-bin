package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var branch_merge_restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restart the merge, resetting the current merge state",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_merge_restartCmd).Standalone()

	branch_merge_restartCmd.Flags().BoolP("help", "h", false, "Print help")
	branch_merge_restartCmd.Flags().String("targets", "", "Path to a targets file")
	branch_mergeCmd.AddCommand(branch_merge_restartCmd)
}
