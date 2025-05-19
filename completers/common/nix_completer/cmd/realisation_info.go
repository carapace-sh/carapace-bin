package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
	"github.com/spf13/cobra"
)

var realisation_infoCmd = &cobra.Command{
	Use:   "info",
	Short: "query information about one or several realisations",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(realisation_infoCmd).Standalone()

	realisation_infoCmd.Flags().Bool("json", false, "Produce output in JSON format")
	realisationCmd.AddCommand(realisation_infoCmd)

	addEvaluationFlags(realisation_infoCmd)
	addFlakeFlags(realisation_infoCmd)
	addLoggingFlags(realisation_infoCmd)

	carapace.Gen(realisation_infoCmd).FlagCompletion(carapace.ActionMap{
		"inputs-from": carapace.Batch(
			carapace.ActionDirectories(),
			nix.ActionFlakes(),
		).ToA(),
		"output-lock-file":    carapace.ActionFiles(),
		"reference-lock-file": carapace.ActionFiles("lock"),
	})
	carapace.Gen(realisation_infoCmd).PositionalAnyCompletion(nix.ActionInstallables())
}
