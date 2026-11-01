package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ownerCmd = &cobra.Command{
	Use:     "owner",
	Short:   "Manage package owners on the registry",
	Aliases: []string{"owners"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ownerCmd).Standalone()

	ownerCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	ownerCmd.Flags().String("otp", "", "One-time password for registries that require two-factor authentication")
	ownerCmd.Flags().String("registry", "", "The base URL of the npm registry")
	rootCmd.AddCommand(ownerCmd)
}
