package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/molecule"
	"github.com/spf13/cobra"
)

var destroyCmd = &cobra.Command{
	Use:   "destroy [flags]",
	Short: "Use the provisioner to destroy the instances",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(destroyCmd)

	destroyCmd.Flags().Bool("all", false, "Destroy all scenarios")
	destroyCmd.Flags().StringP("driver-name", "d", "delegated", "Name of the driver to use")
	destroyCmd.Flags().Bool("help", false, "Show help message and exit")
	destroyCmd.Flags().Bool("no-all", true, "Don't destroy all scenarios (default)")
	destroyCmd.Flags().Bool("no-parallel", true, "Disable parallel mode (default)")
	destroyCmd.Flags().Bool("parallel", false, "Enabe parallel mode")
	destroyCmd.Flags().StringP("scenario-name", "s", "default", "Name of the scenario to target")

	destroyCmd.MarkFlagsMutuallyExclusive("all", "scenario-name")
	destroyCmd.MarkFlagsMutuallyExclusive("all", "no-all")
	destroyCmd.MarkFlagsMutuallyExclusive("parallel", "no-parallel")

	carapace.Gen(destroyCmd).FlagCompletion(carapace.ActionMap{
		"driver-name":   molecule.ActionDrivers(),
		"scenario-name": molecule.ActionScenarios(),
	})

	rootCmd.AddCommand(destroyCmd)
}
