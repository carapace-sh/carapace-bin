package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var distTagCmd = &cobra.Command{
	Use:     "dist-tag",
	Short:   "Manage a package's distribution tags",
	Aliases: []string{"dist-tags"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(distTagCmd).Standalone()

	distTagCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	distTagCmd.Flags().String("otp", "", "One-time password for registries that require two-factor authentication")
	distTagCmd.Flags().String("registry", "", "The base URL of the npm registry")
	rootCmd.AddCommand(distTagCmd)
}
