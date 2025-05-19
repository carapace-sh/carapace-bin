package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/molecule"
	"github.com/spf13/cobra"
)

var matrixCmd = &cobra.Command{
	Use:   "matrix [flags] [SUBCOMMAND]",
	Short: "List matrix of steps used to test instances",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(matrixCmd)

	matrixCmd.Flags().Bool("help", false, "Show help message and exit")
	matrixCmd.Flags().StringP("scenario-name", "s", "default", "Name of the scenario to target")

	carapace.Gen(matrixCmd).FlagCompletion(carapace.ActionMap{
		"scenario-name": molecule.ActionScenarios(),
	})
	carapace.Gen(matrixCmd).PositionalCompletion(
		carapace.ActionValues(
			"check",
			"cleanup",
			"create",
			"destroy",
			"idempotence",
			"prepare",
			"side-effect",
			"syntax",
			"test",
			"verify",
		),
	)

	rootCmd.AddCommand(matrixCmd)
}
