package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var package_resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset the complete cache/build directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(package_resetCmd).Standalone()
	package_resetCmd.Flags().SetInterspersed(false)

	package_resetCmd.Flags().BoolP("help", "h", false, "Show help information")
	package_resetCmd.Flags().Bool("version", false, "Show the version")

	packageCmd.AddCommand(package_resetCmd)
}
