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

	rootCmd.Flags().StringP("flake", "F", "", "Specify the flake containing NixOS configuration")
	rootCmd.Flags().Bool("help", false, "Show usage information")
	rootCmd.Flags().StringP("include", "I", "", "Add an entry to the Nix expression search path")
	rootCmd.Flags().Bool("no-flake", false, "Do not imply --flake if /etc/nixos/flake.nix exists")
	rootCmd.Flags().BoolP("recursive", "r", false, "Print all values at or below the specified path recursively")
	rootCmd.Flags().Bool("show-trace", false, "Print eval trace")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"include": carapace.ActionDirectories(),
	})
}
