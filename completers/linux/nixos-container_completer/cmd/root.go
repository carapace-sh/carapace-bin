package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nixos-container",
	Short: "manage NixOS containers",
	Long:  "https://nixos.org/manual/nixos/stable/containers.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.PersistentFlags().Bool("help", false, "Show usage information")
	rootCmd.PersistentFlags().Bool("impure", false, "Allow impure builds")
	rootCmd.PersistentFlags().String("log-format", "", "Set the log format")
	rootCmd.PersistentFlags().StringSlice("option", nil, "Set a Nix configuration option")
	rootCmd.PersistentFlags().Bool("refresh", false, "Refresh the Nix cache")
}
