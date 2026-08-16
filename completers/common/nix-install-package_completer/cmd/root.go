package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nix-install-package",
	Short: "install a package from a binary cache",
	Long:  "https://nixos.org/manual/nix/stable/command-ref/nix-install-package.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("help", false, "Show usage information")
	rootCmd.Flags().Bool("non-interactive", false, "Do not ask for confirmation")
	rootCmd.Flags().StringP("profile", "p", "", "Path to profile")
	rootCmd.Flags().Bool("set", false, "Set profile to contain exactly one derivation")
	rootCmd.Flags().String("url", "", "URL of the package to install")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{})
}