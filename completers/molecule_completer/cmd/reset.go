package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/molecule_completer/cmd/action"
	"github.com/spf13/cobra"
)

var resetCmd = &cobra.Command{
	Use:   "reset [flags]",
	Short: "Reset molecule temporary folders",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resetCmd)

	resetCmd.Flags().Bool("help", false, "Show help message and exit")
	resetCmd.Flags().StringP("scenario-name", "s", "default", "Name of the scenario to target")

	carapace.Gen(resetCmd).FlagCompletion(carapace.ActionMap{
		"scenario-name": action.ActionScenarios(),
	})

	rootCmd.AddCommand(resetCmd)
}
