package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/molecule"
	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check [flags]",
	Short: "Use the provisioner to perform a Dry-Run",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(checkCmd)

	checkCmd.Flags().Bool("help", false, "Show help message and exit")
	checkCmd.Flags().Bool("no-parallel", true, "Disable parallel mode (default)")
	checkCmd.Flags().Bool("parallel", false, "Enabe parallel mode")
	checkCmd.Flags().StringP("scenario-name", "s", "default", "Name of the scenario to target")

	checkCmd.MarkFlagsMutuallyExclusive("parallel", "no-parallel")

	carapace.Gen(checkCmd).FlagCompletion(carapace.ActionMap{
		"scenario-name": molecule.ActionScenarios(),
	})

	rootCmd.AddCommand(checkCmd)
}
