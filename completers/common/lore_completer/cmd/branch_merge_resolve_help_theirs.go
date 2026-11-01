package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var branch_merge_resolve_help_theirsCmd = &cobra.Command{
	Use:   "theirs",
	Short: "Resolve using their changes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_merge_resolve_help_theirsCmd).Standalone()

	branch_merge_resolve_helpCmd.AddCommand(branch_merge_resolve_help_theirsCmd)
}
