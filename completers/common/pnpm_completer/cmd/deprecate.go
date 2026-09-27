package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deprecateCmd = &cobra.Command{
	Use:   "deprecate",
	Short: "Deprecates a version of a package in the registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deprecateCmd).Standalone()

	deprecateCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	deprecateCmd.Flags().String("otp", "", "One-time password for registries that require two-factor authentication")
	deprecateCmd.Flags().String("registry", "", "The base URL of the npm registry")
	rootCmd.AddCommand(deprecateCmd)
}
