package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nixos-install",
	Short: "install NixOS",
	Long:  "https://nixos.org/manual/nixos/stable/installation.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("chroot", false, "Chroot into the target root")
	rootCmd.Flags().Bool("help", false, "Show usage information")
	rootCmd.Flags().String("root", "", "Root directory to install to")
	rootCmd.Flags().Bool("show-trace", false, "Show the stack trace on evaluation errors")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"root": carapace.ActionDirectories(),
	})
}