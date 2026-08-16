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

	rootCmd.Flags().Bool("help", false, "Show usage information")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValues("list", "create", "destroy", "start", "stop", "status", "update", "login", "root-login", "run", "show-ip", "show-host-key"),
	)
}