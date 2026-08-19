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

	rootCmd.Flags().Bool("auto-start", false, "Automatically start the container at boot")
	rootCmd.Flags().String("bridge", "", "Attach the container to a host bridge interface")
	rootCmd.Flags().String("config", "", "Inline NixOS configuration string")
	rootCmd.Flags().String("config-file", "", "Path to a NixOS configuration file")
	rootCmd.Flags().Bool("ensure-unique-name", false, "Append a unique suffix if the container name already exists")
	rootCmd.Flags().String("flake", "", "Flake reference")
	rootCmd.Flags().Bool("help", false, "Show usage information")
	rootCmd.Flags().String("host-address", "", "Host-side IP address for the veth pair")
	rootCmd.Flags().Bool("impure", false, "Allow impure builds")
	rootCmd.Flags().String("local-address", "", "Container-side IP address for the veth pair")
	rootCmd.Flags().String("log-format", "", "Set the log format")
	rootCmd.Flags().String("nixos-path", "", "Path to the NixOS source")
	rootCmd.Flags().StringSlice("option", nil, "Set a Nix configuration option")
	rootCmd.Flags().String("port", "", "Host port mapping")
	rootCmd.Flags().Bool("refresh", false, "Refresh the Nix cache")
	rootCmd.Flags().String("system-path", "", "Path to a pre-built system derivation")
	rootCmd.Flags().Bool("use-host-network", false, "Share the host's network namespace")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValues("list", "create", "destroy", "restart", "start", "stop", "terminate", "status", "update", "login", "root-login", "run", "show-ip", "show-host-key"),
	)
}
