package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var branch_merge_resolveCmd = &cobra.Command{
	Use:   "resolve",
	Short: "Resolves the merge",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_merge_resolveCmd).Standalone()

	branch_merge_resolveCmd.Flags().BoolP("help", "h", false, "Print help")
	branch_merge_resolveCmd.Flags().String("targets", "", "Path to a targets file")
	branch_mergeCmd.AddCommand(branch_merge_resolveCmd)
}
