package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var package_resolveCmd = &cobra.Command{
	Use:   "resolve",
	Short: "Resolve package dependencies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(package_resolveCmd).Standalone()
	package_resolveCmd.Flags().SetInterspersed(false)

	package_resolveCmd.Flags().String("branch", "", "The branch to resolve at")
	package_resolveCmd.Flags().BoolP("help", "h", false, "Show help information")
	package_resolveCmd.Flags().String("revision", "", "The revision to resolve at")
	package_resolveCmd.Flags().String("version", "", "The version to resolve at")

	packageCmd.AddCommand(package_resolveCmd)
}
