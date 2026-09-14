package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/nix_completer/cmd/common"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:     "search",
	Short:   "search for packages",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(searchCmd).Standalone()

	searchCmd.Flags().BoolP("exclude", "e", false, "Hide packages whose attribute path, name or description contain regex")
	rootCmd.AddCommand(searchCmd)

	common.AddEvaluationFlags(searchCmd)
	common.AddFlakeFlags(searchCmd)
	common.AddInterpretationFlags(searchCmd)
	common.AddLoggingFlags(searchCmd)

	carapace.Gen(searchCmd).FlagCompletion(carapace.ActionMap{})
	carapace.Gen(searchCmd).PositionalCompletion(nix.ActionInstallables())
}
