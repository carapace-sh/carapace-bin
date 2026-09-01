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

	sbomCmd.Flags().BoolP("dev", "D", false, "Only include dev dependencies")
	sbomCmd.Flags().Bool("exclude-peers", false, "Exclude peer dependencies")
	sbomCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	sbomCmd.Flags().Bool("lockfile-only", false, "Only use lockfile data (skip reading from the store)")
	sbomCmd.Flags().Bool("no-optional", false, "Exclude optional dependencies")
	sbomCmd.Flags().Bool("optional", false, "Include optional dependencies")
	sbomCmd.Flags().String("out", "", "Write SBOM to a file instead of stdout. Use `%s` for the package name and `%v` for the version")
	sbomCmd.Flags().BoolP("prod", "P", false, "Only include production dependencies")
	sbomCmd.Flags().Bool("production", false, "Only include production dependencies")
	sbomCmd.Flags().String("sbom-authors", "", "Comma-separated list of SBOM authors (`CycloneDX` `metadata.authors`)")
	sbomCmd.Flags().String("sbom-format", "", "The SBOM output format (required)")
	sbomCmd.Flags().String("sbom-spec-version", "", "The `CycloneDX` specification version (`1.5`, `1.6`, or `1.7`; default: `1.7`). Only valid with `--sbom-format cyclonedx`")
	sbomCmd.Flags().String("sbom-supplier", "", "SBOM supplier name (`CycloneDX` `metadata.supplier`)")
	sbomCmd.Flags().String("sbom-type", "library", "The component type for the root package (default: library)")
	sbomCmd.Flags().Bool("split", false, "Generate a separate SBOM for each matched workspace package")
	sbomCmd.MarkFlagRequired("sbom-format")

	carapace.Gen(sbomCmd).FlagCompletion(carapace.ActionMap{
		"sbom-format": carapace.ActionValues("cyclonedx", "spdx"),
		"sbom-type":   carapace.ActionValues("library", "application"),
	})

	rootCmd.AddCommand(sbomCmd)
}
