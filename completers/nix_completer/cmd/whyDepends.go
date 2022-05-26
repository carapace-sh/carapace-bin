package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var whyDependsCmd = &cobra.Command{
	Use:     "why-depends",
	Short:   "show why a package has another package in its closure",
	GroupID: "infrequently used",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(whyDependsCmd).Standalone()

	whyDependsCmd.Flags().BoolP("all", "a", false, "Show all edges in the dependency graph leading from package to dependency")
	whyDependsCmd.Flags().Bool("precise", false, "For each edge in the dependency graph, show the files in the parent that cause the dependency")
	rootCmd.AddCommand(whyDependsCmd)

	addEvaluationFlags(whyDependsCmd)
	addFlakeFlags(whyDependsCmd)
	addLoggingFlags(whyDependsCmd)

	// TODO positional completion
}
