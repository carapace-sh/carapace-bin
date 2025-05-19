package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
	"github.com/spf13/cobra"
)

var fmtCmd = &cobra.Command{
	Use:     "fmt",
	Short:   "reformat your code in the standard style",
	GroupID: "infrequently used",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fmtCmd).Standalone()

	rootCmd.AddCommand(fmtCmd)

	addEvaluationFlags(fmtCmd)
	addFlakeFlags(fmtCmd)
	addLoggingFlags(fmtCmd)

	carapace.Gen(fmtCmd).FlagCompletion(carapace.ActionMap{
		"inputs-from": carapace.Batch(
			carapace.ActionDirectories(),
			nix.ActionFlakes(),
		).ToA(),
		"output-lock-file":    carapace.ActionFiles(),
		"reference-lock-file": carapace.ActionFiles("lock"),
	})
	carapace.Gen(fmtCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
