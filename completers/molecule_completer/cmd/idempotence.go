package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/molecule"
	"github.com/spf13/cobra"
)

var idempotenceCmd = &cobra.Command{
	Use:   "idempotence [flags]",
	Short: "Use the provisioner to configure instances and determine idempotence",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(idempotenceCmd)

	idempotenceCmd.Flags().Bool("help", false, "Show help message and exit")
	idempotenceCmd.Flags().StringP("scenario-name", "s", "default", "Name of the scenario to target")

	carapace.Gen(idempotenceCmd).FlagCompletion(carapace.ActionMap{
		"scenario-name": molecule.ActionScenarios(),
	})

	rootCmd.AddCommand(idempotenceCmd)
}
