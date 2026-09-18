package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var undeprecateCmd = &cobra.Command{
	Use:   "undeprecate",
	Short: "Removes deprecation from a version of a package in the registry. Only works on already deprecated versions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(undeprecateCmd).Standalone()

	undeprecateCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	undeprecateCmd.Flags().String("otp", "", "One-time password for registries that require two-factor authentication")
	undeprecateCmd.Flags().String("registry", "", "The base URL of the npm registry")
	rootCmd.AddCommand(undeprecateCmd)
}
