package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
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

	carapace.Gen(whyDependsCmd).FlagCompletion(carapace.ActionMap{
		"inputs-from": carapace.Batch(
			carapace.ActionDirectories(),
			nix.ActionFlakes(),
		).ToA(),
		"output-lock-file":    carapace.ActionFiles(),
		"reference-lock-file": carapace.ActionFiles("lock"),
	})
	// TODO positional completion
}
