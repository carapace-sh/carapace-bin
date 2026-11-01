package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "rebuild container after configuration changes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(updateCmd).Standalone()

	updateCmd.Flags().String("config", "", "Inline NixOS configuration string")
	updateCmd.Flags().String("config-file", "", "Path to a NixOS configuration file")
	updateCmd.Flags().String("flake", "", "Flake reference")
	updateCmd.Flags().String("nixos-path", "", "Path to the NixOS source")

	carapace.Gen(updateCmd).FlagCompletion(carapace.ActionMap{
		"config-file": carapace.ActionFiles(".nix"),
		"nixos-path":  carapace.ActionDirectories(),
	})

	rootCmd.AddCommand(updateCmd)
}
