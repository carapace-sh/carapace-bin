package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/nix_completer/cmd/common"
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
	whyDependsCmd.Flags().Bool("derivation", false, "Operate on the store derivation rather than its outputs")
	whyDependsCmd.Flags().Bool("precise", false, "For each edge in the dependency graph, show the files in the parent that cause the dependency")
	rootCmd.AddCommand(whyDependsCmd)

	common.AddEvaluationFlags(whyDependsCmd)
	common.AddFlakeFlags(whyDependsCmd)
	common.AddInterpretationFlags(whyDependsCmd)
	common.AddLoggingFlags(whyDependsCmd)

	carapace.Gen(whyDependsCmd).FlagCompletion(carapace.ActionMap{})
	carapace.Gen(whyDependsCmd).PositionalAnyCompletion(nix.ActionInstallables())
}
