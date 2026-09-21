package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var package_addTargetDependencyCmd = &cobra.Command{
	Use:   "add-target-dependency",
	Short: "Add a new target dependency to the manifest",
	Args:  cobra.ExactArgs(2),
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(package_addTargetDependencyCmd).Standalone()
	package_addTargetDependencyCmd.Flags().SetInterspersed(false)

	package_addTargetDependencyCmd.Flags().BoolP("help", "h", false, "Show help information")
	package_addTargetDependencyCmd.Flags().String("package", "", "The package in which the dependency resides")
	package_addTargetDependencyCmd.Flags().Bool("version", false, "Show the version")

	packageCmd.AddCommand(package_addTargetDependencyCmd)
}
