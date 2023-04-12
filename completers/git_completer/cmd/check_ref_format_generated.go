package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var check_ref_formatCmd = &cobra.Command{
	Use:     "check-ref-format",
	Short:   "Ensures that a reference name is well formed",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_helper].ID,
}

func init() {
	carapace.Gen(check_ref_formatCmd).Standalone()

	rootCmd.AddCommand(check_ref_formatCmd)
}
