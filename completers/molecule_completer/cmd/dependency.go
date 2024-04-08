package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/molecule"
	"github.com/spf13/cobra"
)

var dependencyCmd = &cobra.Command{
	Use:   "dependency [flags]",
	Short: "Manage the role's dependencies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dependencyCmd)

	dependencyCmd.Flags().Bool("help", false, "Show help message and exit")
	dependencyCmd.Flags().StringP("scenario-name", "s", "default", "Name of the scenario to target")

	carapace.Gen(dependencyCmd).FlagCompletion(carapace.ActionMap{
		"scenario-name": molecule.ActionScenarios(),
	})

	rootCmd.AddCommand(dependencyCmd)
}
