package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/molecule_completer/cmd/action"
	"github.com/spf13/cobra"
)

var verifyCmd = &cobra.Command{
	Use:   "verify [flags]",
	Short: "Run automated tests against instances",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(verifyCmd)

	verifyCmd.Flags().Bool("help", false, "Show help message and exit")
	verifyCmd.Flags().StringP("scenario-name", "s", "default", "Name of the scenario to target")

	carapace.Gen(verifyCmd).FlagCompletion(carapace.ActionMap{
		"scenario-name": action.ActionScenarios(),
	})

	rootCmd.AddCommand(verifyCmd)
}
