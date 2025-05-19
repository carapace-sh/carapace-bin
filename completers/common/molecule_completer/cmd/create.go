package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/molecule"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create [flags]",
	Short: "Use the provisioner to start the instances",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(createCmd)

	createCmd.Flags().StringP("driver-name", "d", "delegated", "Name of the driver to use")
	createCmd.Flags().Bool("help", false, "Show help message and exit")
	createCmd.Flags().StringP("scenario-name", "s", "default", "Name of the scenario to target")

	carapace.Gen(createCmd).FlagCompletion(carapace.ActionMap{
		"driver-name":   molecule.ActionDrivers(),
		"scenario-name": molecule.ActionScenarios(),
	})

	rootCmd.AddCommand(createCmd)
}
