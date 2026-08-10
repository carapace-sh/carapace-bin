package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var package_cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Delete build artifacts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(package_cleanCmd).Standalone()
	package_cleanCmd.Flags().SetInterspersed(false)

	package_cleanCmd.Flags().BoolP("help", "h", false, "Show help information")
	package_cleanCmd.Flags().Bool("version", false, "Show the version")

	packageCmd.AddCommand(package_cleanCmd)
}
