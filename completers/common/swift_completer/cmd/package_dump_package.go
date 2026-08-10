package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var package_dumpPackageCmd = &cobra.Command{
	Use:   "dump-package",
	Short: "Print parsed Package.swift as JSON",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(package_dumpPackageCmd).Standalone()
	package_dumpPackageCmd.Flags().SetInterspersed(false)

	package_dumpPackageCmd.Flags().BoolP("help", "h", false, "Show help information")
	package_dumpPackageCmd.Flags().Bool("version", false, "Show the version")

	packageCmd.AddCommand(package_dumpPackageCmd)
}
