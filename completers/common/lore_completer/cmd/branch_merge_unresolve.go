package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var branch_merge_unresolveCmd = &cobra.Command{
	Use:   "unresolve",
	Short: "Marks the merge unresolved",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_merge_unresolveCmd).Standalone()

	branch_merge_unresolveCmd.Flags().BoolP("help", "h", false, "Print help")
	branch_merge_unresolveCmd.Flags().String("targets", "", "Path to a targets file")
	branch_mergeCmd.AddCommand(branch_merge_unresolveCmd)
}
