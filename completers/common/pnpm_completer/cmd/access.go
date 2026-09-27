package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var accessCmd = &cobra.Command{
	Use:   "access",
	Short: "Manage package access and visibility on the registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(accessCmd).Standalone()

	accessCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	accessCmd.Flags().Bool("json", false, "Output results in JSON format")
	accessCmd.Flags().String("otp", "", "One-time password for registries that require two-factor authentication")
	accessCmd.Flags().String("registry", "", "The base URL of the npm registry")
	rootCmd.AddCommand(accessCmd)
}
