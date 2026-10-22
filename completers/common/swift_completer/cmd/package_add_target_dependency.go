package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/swift_completer/cmd/action"
	"github.com/spf13/cobra"
)

var package_addTargetDependencyCmd = &cobra.Command{
	Use:   "add-target-dependency",
	Short: "Add a new target dependency to the manifest",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(package_addTargetDependencyCmd).Standalone()
	package_addTargetDependencyCmd.Flags().SetInterspersed(false)

	package_addTargetDependencyCmd.Flags().BoolP("help", "h", false, "Show help information")
	package_addTargetDependencyCmd.Flags().String("package", "", "The package in which the dependency resides")
	package_addTargetDependencyCmd.Flags().Bool("version", false, "Show the version")

	packageCmd.AddCommand(package_addTargetDependencyCmd)

	carapace.Gen(package_addTargetDependencyCmd).PositionalCompletion(
		action.ActionTargetsAndProducts(),
		action.ActionTargets(),
	)
}
