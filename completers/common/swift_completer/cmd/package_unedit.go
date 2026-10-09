package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var package_uneditCmd = &cobra.Command{
	Use:   "unedit",
	Short: "Remove a package from editable mode",
	Args:  cobra.ExactArgs(1),
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(package_uneditCmd).Standalone()
	package_uneditCmd.Flags().SetInterspersed(false)

	package_uneditCmd.Flags().Bool("force", false, "Unedit the package even if it has uncommitted and unpushed changes")
	package_uneditCmd.Flags().BoolP("help", "h", false, "Show help information")
	package_uneditCmd.Flags().Bool("version", false, "Show the version")

	packageCmd.AddCommand(package_uneditCmd)
}
