package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var package_initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(package_initCmd).Standalone()
	package_initCmd.Flags().SetInterspersed(false)

	package_initCmd.Flags().Bool("disable-swift-testing", false, "Disable support for Swift Testing")
	package_initCmd.Flags().Bool("disable-xctest", false, "Disable support for XCTest")
	package_initCmd.Flags().Bool("enable-swift-testing", false, "Enable support for Swift Testing")
	package_initCmd.Flags().Bool("enable-xctest", false, "Enable support for XCTest")
	package_initCmd.Flags().BoolP("help", "h", false, "Show help information")
	package_initCmd.Flags().String("name", "", "Provide custom package name")
	package_initCmd.Flags().String("type", "", "Package type")
	package_initCmd.Flags().Bool("version", false, "Show the version")

	packageCmd.AddCommand(package_initCmd)

	carapace.Gen(package_initCmd).FlagCompletion(carapace.ActionMap{
		"type": carapace.ActionValuesDescribed(
			"library", "A package with a library",
			"executable", "A package with an executable",
			"tool", "A package with an executable that uses Swift Argument Parser",
			"build-tool-plugin", "A package that vends a build tool plugin",
			"command-plugin", "A package that vends a command plugin",
			"macro", "A package that vends a macro",
			"empty", "An empty package with a Package.swift manifest",
		),
	})
}
