package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/molecule"
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
		"scenario-name": molecule.ActionScenarios(),
	})

	rootCmd.AddCommand(resetCmd)
}
