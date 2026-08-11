package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var branch_merge_abortCmd = &cobra.Command{
	Use:   "abort",
	Short: "Abort a merge process",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_merge_abortCmd).Standalone()

	branch_merge_abortCmd.Flags().BoolP("help", "h", false, "Print help")
	branch_merge_abortCmd.Flags().Bool("ignore-links", false, "Abort only the main repository merge, keeping link pin updates")
	branch_merge_abortCmd.Flags().String("link", "", "Abort only a specific linked repository merge at the given mount path")
	branch_mergeCmd.AddCommand(branch_merge_abortCmd)
}
