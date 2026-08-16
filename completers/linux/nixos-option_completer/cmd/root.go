package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nixos-option",
	Short: "inspect a NixOS configuration option",
	Long:  "https://nixos.org/manual/nixos/stable/options.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("help", false, "Show usage information")
	rootCmd.Flags().Bool("xml", false, "Print the option in XML format")
}