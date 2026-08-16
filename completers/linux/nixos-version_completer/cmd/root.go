package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nixos-version",
	Short: "show the NixOS version",
	Long:  "https://nixos.org/manual/nixos/stable/options.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("hash", false, "Print the NixOS version hash")
	rootCmd.Flags().Bool("help", false, "Show usage information")
	rootCmd.Flags().Bool("revision", false, "Print the NixOS version revision")
}