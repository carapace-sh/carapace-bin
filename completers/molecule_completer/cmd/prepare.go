package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/molecule"
	"github.com/spf13/cobra"
)

var prepareCmd = &cobra.Command{
	Use:   "prepare [flags]",
	Short: "Use the provisioner to prepare instances into a starting state",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(prepareCmd)

	prepareCmd.Flags().StringP("driver-name", "d", "delegated", "Name of the driver to use")
	prepareCmd.Flags().BoolP("force", "f", false, "Enable force mode")
	prepareCmd.Flags().Bool("help", false, "Show help message and exit")
	prepareCmd.Flags().Bool("no-force", true, "Disable force mode (default)")
	prepareCmd.Flags().StringP("scenario-name", "s", "default", "Name of the scenario to target")

	prepareCmd.MarkFlagsMutuallyExclusive("force", "no-force")

	carapace.Gen(prepareCmd).FlagCompletion(carapace.ActionMap{
		"driver-name":   molecule.ActionDrivers(),
		"scenario-name": molecule.ActionScenarios(),
	})

	rootCmd.AddCommand(prepareCmd)
}
