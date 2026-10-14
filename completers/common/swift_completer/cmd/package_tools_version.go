package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var package_toolsVersionCmd = &cobra.Command{
	Use:   "tools-version",
	Short: "Manipulate tools version of the current package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(package_toolsVersionCmd).Standalone()
	package_toolsVersionCmd.Flags().SetInterspersed(false)

	package_toolsVersionCmd.Flags().BoolP("help", "h", false, "Show help information")
	package_toolsVersionCmd.Flags().String("set", "", "Set tools version of package to the given value")
	package_toolsVersionCmd.Flags().Bool("set-current", false, "Set tools version of package to the current tools version in use")
	package_toolsVersionCmd.Flags().Bool("version", false, "Show the version")

	packageCmd.AddCommand(package_toolsVersionCmd)
}
