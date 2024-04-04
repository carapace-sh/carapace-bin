package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/molecule_completer/cmd/action"
	"github.com/spf13/cobra"
)

var sideEffectCmd = &cobra.Command{
	Use:   "side-effect [flags]",
	Short: "Use the provisioner to perform side-effects on the instances",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sideEffectCmd)

	sideEffectCmd.Flags().Bool("help", false, "Show help message and exit")
	sideEffectCmd.Flags().StringP("scenario-name", "s", "default", "Name of the scenario to target")

	carapace.Gen(sideEffectCmd).FlagCompletion(carapace.ActionMap{
		"scenario-name": action.ActionScenarios(),
	})

	rootCmd.AddCommand(sideEffectCmd)
}
