package cmd

import (
	"github.com/carapace-sh/carapace"
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
	searchCmd.Flags().Bool("json", false, "Produce output in JSON format, suitable for consumption by another program")
	rootCmd.AddCommand(searchCmd)

	addEvaluationFlags(searchCmd)
	addFlakeFlags(searchCmd)
	addLoggingFlags(searchCmd)

	carapace.Gen(searchCmd).FlagCompletion(carapace.ActionMap{
		"inputs-from": carapace.Batch(
			carapace.ActionDirectories(),
			nix.ActionFlakes(),
		).ToA(),
		"output-lock-file":    carapace.ActionFiles(),
		"reference-lock-file": carapace.ActionFiles("lock"),
	})
	carapace.Gen(searchCmd).PositionalCompletion(nix.ActionInstallables())
}
