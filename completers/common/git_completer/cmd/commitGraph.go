package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var commitGraphCmd = &cobra.Command{
	Use:     "commit-graph",
	Short:   "Write and verify Git commit-graph files",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_manipulator].ID,
}

func init() {
	carapace.Gen(commitGraphCmd).Standalone()

	rootCmd.AddCommand(commitGraphCmd)
}
