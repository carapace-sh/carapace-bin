package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_gc_commitGraphCmd = &cobra.Command{
	Use:   "commit-graph",
	Short: "Write a reachable commit-graph with changed-path Bloom filters",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_gc_commitGraphCmd).Standalone()

	help_gcCmd.AddCommand(help_gc_commitGraphCmd)
}
