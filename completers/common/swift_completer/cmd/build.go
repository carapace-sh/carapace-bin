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

	rootCmd.AddCommand(buildCmd)

}
