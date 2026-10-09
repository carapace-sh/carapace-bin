package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var package_addTargetCmd = &cobra.Command{
	Use:   "add-target",
	Short: "Add a new target to the manifest",
	Args:  cobra.ExactArgs(1),
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(package_addTargetCmd).Standalone()
	package_addTargetCmd.Flags().SetInterspersed(false)

	package_addTargetCmd.Flags().String("checksum", "", "The checksum for a remote binary target")
	package_addTargetCmd.Flags().StringArray("dependencies", nil, "A list of target dependency names")
	package_addTargetCmd.Flags().BoolP("help", "h", false, "Show help information")
	package_addTargetCmd.Flags().String("path", "", "The path to a local binary target")
	package_addTargetCmd.Flags().String("testing-library", "", "The testing library to use when generating test targets")
	package_addTargetCmd.Flags().String("type", "", "The type of target to add")
	package_addTargetCmd.Flags().String("url", "", "The URL for a remote binary target")
	package_addTargetCmd.Flags().Bool("version", false, "Show the version")

	packageCmd.AddCommand(package_addTargetCmd)

	carapace.Gen(package_addTargetCmd).FlagCompletion(carapace.ActionMap{
		"path":            carapace.ActionFiles(),
		"testing-library": carapace.ActionValues("xctest", "swift-testing", "none"),
		"type":            carapace.ActionValues("library", "executable", "test", "macro"),
	})
}
