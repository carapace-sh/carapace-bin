package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var unpublishCmd = &cobra.Command{
	Use:   "unpublish",
	Short: "Removes a package from the registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unpublishCmd).Standalone()

	unpublishCmd.Flags().Bool("force", false, "Removes the package from the registry regardless of what version is currently published. Without this flag, pnpm refuses to unpublish an entire package")
	unpublishCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	unpublishCmd.Flags().String("otp", "", "One-time password for registries that require two-factor authentication")
	unpublishCmd.Flags().String("registry", "", "The base URL of the npm registry")
	rootCmd.AddCommand(unpublishCmd)
}
