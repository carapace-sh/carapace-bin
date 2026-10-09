package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var package_configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manipulate configuration of the package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(package_configCmd).Standalone()
	package_configCmd.Flags().SetInterspersed(false)

	package_configCmd.Flags().BoolP("help", "h", false, "Show help information")
	package_configCmd.Flags().Bool("version", false, "Show the version")

	packageCmd.AddCommand(package_configCmd)
}
