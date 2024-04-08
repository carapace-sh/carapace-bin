package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/molecule"
	"github.com/spf13/cobra"
)

var initScenarioCmd = &cobra.Command{
	Use:   "scenario [flags] [SCENARIO_NAME]",
	Short: "Initializes a new scenario to use with Molecule",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(initScenarioCmd)

	initScenarioCmd.Flags().String("dependency-name", "", "Name of galaxy dependency to initialize")
	initScenarioCmd.Flags().StringP("driver-name", "d", "delegated", "Name of the driver to use")
	initScenarioCmd.Flags().Bool("help", false, "Show help message and exit")
	initScenarioCmd.Flags().String("provisioner-name", "ansible", "Name of provisioner to initialize")
	initScenarioCmd.Flags().StringP("role-name", "r", "", "Name of the role to create")
	initScenarioCmd.Flags().String("verifier-name", "ansible", "Name of verifier to initialize")

	carapace.Gen(initScenarioCmd).FlagCompletion(carapace.ActionMap{
		"driver-name":      molecule.ActionDrivers(),
		"provisioner-name": carapace.ActionExecutables(),
		"verifier-name":    carapace.ActionValues("ansible", "testinfra"),
	})

	initCmd.AddCommand(initScenarioCmd)
}
