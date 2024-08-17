package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var hookCmd = &cobra.Command{
	Use:     "hook",
	Short:   "Run git hooks",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_helper].ID,
}

func init() {
	carapace.Gen(hookCmd).Standalone()

	rootCmd.AddCommand(hookCmd)
}
