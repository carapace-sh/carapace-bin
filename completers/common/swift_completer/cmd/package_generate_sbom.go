package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/swift_completer/cmd/common"
	"github.com/spf13/cobra"
)

var package_generateSbomCmd = &cobra.Command{
	Use:   "generate-sbom",
	Short: "Generate a Software Bill of Materials (SBOM)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(package_generateSbomCmd).Standalone()
	package_generateSbomCmd.Flags().SetInterspersed(false)

	common.AddPackageFlags(package_generateSbomCmd)

	package_generateSbomCmd.Flags().BoolP("help", "h", false, "Show help information")
	package_generateSbomCmd.Flags().String("product", "", "The product to generate an SBOM for")
	package_generateSbomCmd.Flags().String("sbom-filter", "", "Filter the SBOM components and dependencies by products and/or packages")
	package_generateSbomCmd.Flags().String("sbom-output-dir", "", "The absolute or relative directory path to generate the SBOM(s) in")
	package_generateSbomCmd.Flags().StringArray("sbom-spec", nil, "Set the SBOM specification and generate an SBOM")
	package_generateSbomCmd.Flags().Bool("sbom-warning-only", false, "Treat SBOM generation errors as warnings")
	package_generateSbomCmd.Flags().Bool("version", false, "Show the version")

	packageCmd.AddCommand(package_generateSbomCmd)

	carapace.Gen(package_generateSbomCmd).FlagCompletion(carapace.ActionMap{
		"product": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return carapace.ActionExecCommand("swift", "package", "describe", "--type", "json")(func(output []byte) carapace.Action {
				return carapace.ActionValues()
			})
		}),
		"sbom-filter":     carapace.ActionValues("all", "product", "package"),
		"sbom-output-dir": carapace.ActionDirectories(),
		"sbom-spec":       carapace.ActionValues("cyclonedx", "spdx", "cyclonedx1", "spdx3"),
	})
}
