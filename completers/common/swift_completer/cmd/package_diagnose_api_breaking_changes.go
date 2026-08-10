package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var package_diagnoseApiBreakingChangesCmd = &cobra.Command{
	Use:   "diagnose-api-breaking-changes",
	Short: "Diagnose API-breaking changes to Swift modules in a package",
	Args:  cobra.MinimumNArgs(1),
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(package_diagnoseApiBreakingChangesCmd).Standalone()
	package_diagnoseApiBreakingChangesCmd.Flags().SetInterspersed(false)

	package_diagnoseApiBreakingChangesCmd.Flags().String("baseline-dir", "", "The directory to use for the baseline")
	package_diagnoseApiBreakingChangesCmd.Flags().String("breakage-allowlist-path", "", "The path to a text file containing breaking changes which should be ignored")
	package_diagnoseApiBreakingChangesCmd.Flags().BoolP("help", "h", false, "Show help information")
	package_diagnoseApiBreakingChangesCmd.Flags().StringArray("products", nil, "The products to include in the comparison")
	package_diagnoseApiBreakingChangesCmd.Flags().Bool("regenerate-baseline", false, "Regenerate the baseline")
	package_diagnoseApiBreakingChangesCmd.Flags().StringArray("targets", nil, "The targets to include in the comparison")
	package_diagnoseApiBreakingChangesCmd.Flags().Bool("version", false, "Show the version")

	packageCmd.AddCommand(package_diagnoseApiBreakingChangesCmd)

	carapace.Gen(package_diagnoseApiBreakingChangesCmd).FlagCompletion(carapace.ActionMap{
		"baseline-dir":            carapace.ActionDirectories(),
		"breakage-allowlist-path": carapace.ActionFiles(),
	})
}
