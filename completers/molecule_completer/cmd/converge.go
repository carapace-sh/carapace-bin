package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/molecule_completer/cmd/action"
	"github.com/spf13/cobra"
)

var convergeCmd = &cobra.Command{
	Use:   "converge [flags]",
	Short: "Use the provisioner to configure instances",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(convergeCmd)

	convergeCmd.Flags().Bool("help", false, "Show help message and exit")
	convergeCmd.Flags().StringP("scenario-name", "s", "default", "Name of the scenario to target")

	carapace.Gen(convergeCmd).FlagCompletion(carapace.ActionMap{
		"scenario-name": action.ActionScenarios(),
	})

	rootCmd.AddCommand(convergeCmd)
}
