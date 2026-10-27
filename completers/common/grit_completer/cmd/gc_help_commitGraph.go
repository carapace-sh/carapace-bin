package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gc_help_commitGraphCmd = &cobra.Command{
	Use:   "commit-graph",
	Short: "Write a reachable commit-graph with changed-path Bloom filters",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gc_help_commitGraphCmd).Standalone()

	gc_helpCmd.AddCommand(gc_help_commitGraphCmd)
}
