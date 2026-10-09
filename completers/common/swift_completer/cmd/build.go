package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/swift_completer/cmd/common"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build sources into binary products",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(buildCmd).Standalone()
	buildCmd.Flags().SetInterspersed(false)

	common.AddPackageFlags(buildCmd)

	buildCmd.Flags().Bool("build-tests", false, "Build both source and test targets")
	buildCmd.Flags().Bool("disable-code-coverage", false, "Disable code coverage")
	buildCmd.Flags().Bool("disable-swift-testing", false, "Disable support for Swift Testing")
	buildCmd.Flags().Bool("disable-xctest", false, "Disable support for XCTest")
	buildCmd.Flags().Bool("enable-code-coverage", false, "Enable code coverage")
	buildCmd.Flags().Bool("enable-swift-testing", false, "Enable support for Swift Testing")
	buildCmd.Flags().Bool("enable-xctest", false, "Enable support for XCTest")
	buildCmd.Flags().BoolP("help", "h", false, "Show help information")
	buildCmd.Flags().Bool("no-static-swift-stdlib", false, "Do not link Swift stdlib statically")
	buildCmd.Flags().Bool("print-manifest-job-graph", false, "Write the command graph for the build manifest as a Graphviz file")
	buildCmd.Flags().Bool("print-pif-manifest-graph", false, "Write the PIF JSON sent to Swift Build as a Graphviz file")
	buildCmd.Flags().String("product", "", "Build the specified product")
	buildCmd.Flags().Bool("show-bin-path", false, "Print the binary output path")
	buildCmd.Flags().Bool("static-swift-stdlib", false, "Link Swift stdlib statically")
	buildCmd.Flags().String("target", "", "Build the specified target")
	buildCmd.Flags().Bool("version", false, "Show the version")

	buildCmd.Flags().String("sbom-filter", "", "Filter the SBOM components and dependencies by products and/or packages")
	buildCmd.Flags().String("sbom-output-dir", "", "The absolute or relative directory path to generate the SBOM(s) in")
	buildCmd.Flags().StringArray("sbom-spec", nil, "Set the SBOM specification and generate an SBOM")
	buildCmd.Flags().Bool("sbom-warning-only", false, "Treat SBOM generation errors as warnings")

	rootCmd.AddCommand(buildCmd)

	carapace.Gen(buildCmd).FlagCompletion(carapace.ActionMap{
		"product": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return carapace.ActionExecCommand("swift", "package", "describe", "--type", "json")(func(output []byte) carapace.Action {
				return carapace.ActionValues()
			})
		}),
		"sbom-filter":     carapace.ActionValues("all", "product", "package"),
		"sbom-output-dir": carapace.ActionDirectories(),
		"sbom-spec":       carapace.ActionValues("cyclonedx", "spdx", "cyclonedx1", "spdx3"),
		"target": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return carapace.ActionExecCommand("swift", "package", "describe", "--type", "json")(func(output []byte) carapace.Action {
				return carapace.ActionValues()
			})
		}),
	})
}
