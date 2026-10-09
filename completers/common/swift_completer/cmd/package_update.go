package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var package_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update package dependencies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(package_updateCmd).Standalone()
	package_updateCmd.Flags().SetInterspersed(false)

	package_updateCmd.Flags().BoolP("dry-run", "n", false, "Display the list of dependencies that can be updated")
	package_updateCmd.Flags().BoolP("help", "h", false, "Show help information")
	package_updateCmd.Flags().Bool("version", false, "Show the version")

	packageCmd.AddCommand(package_updateCmd)
}
