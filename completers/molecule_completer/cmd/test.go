package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/molecule"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test [flags] [ANSIBLE_ARGS]...",
	Short: "Use the provisioner to test the role",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(testCmd)

	testCmd.Flags().Bool("all", false, "Test all scenarios")
	testCmd.Flags().String("destroy", "always", "The destroy strategy used to conclude the run")
	testCmd.Flags().StringP("driver-name", "d", "delegated", "Name of the driver to use")
	testCmd.Flags().Bool("help", false, "Show help message and exit")
	testCmd.Flags().Bool("no-all", true, "Don't test all scenarios (default)")
	testCmd.Flags().Bool("no-parallel", true, "Disable parallel mode (default)")
	testCmd.Flags().Bool("parallel", false, "Enabe parallel mode")
	testCmd.Flags().StringP("platform-name", "p", "", "Restrict testing to specified platform")
	testCmd.Flags().StringP("scenario-name", "s", "default", "Name of the scenario to target")

	testCmd.MarkFlagsMutuallyExclusive("all", "scenario-name")
	testCmd.MarkFlagsMutuallyExclusive("all", "no-all")
	testCmd.MarkFlagsMutuallyExclusive("parallel", "no-parallel")

	carapace.Gen(testCmd).FlagCompletion(carapace.ActionMap{
		"destroy":       carapace.ActionValues("always", "never").StyleF(style.ForKeyword),
		"driver-name":   molecule.ActionDrivers(),
		"scenario-name": molecule.ActionScenarios(),
	})

	rootCmd.AddCommand(testCmd)
}
