package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sbomCmd = &cobra.Command{
	Use:   "sbom",
	Short: "Generate a Software Bill of Materials (SBOM)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sbomCmd).Standalone()
	sbomCmd.Flags().String("omit", "", "omit dependency type")
	sbomCmd.Flags().Bool("package-lock-only", false, "ignore node-modules")
	sbomCmd.Flags().String("sbom-format", "", "SBOM format to use")
	sbomCmd.Flags().String("sbom-type", "library", "type of package described by the generated SBOM")
	addWorkspaceFlags(sbomCmd)

	carapace.Gen(sbomCmd).FlagCompletion(carapace.ActionMap{
		"omit":        carapace.ActionValues("dev", "optional", "peer"),
		"sbom-format": carapace.ActionValues("cyclonedx", "spdx"),
		"sbom-type":   carapace.ActionValues("library", "application", "framework"),
	})

	rootCmd.AddCommand(sbomCmd)
}
