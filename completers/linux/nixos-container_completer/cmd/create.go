package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a new container",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(createCmd).Standalone()

	createCmd.Flags().Bool("auto-start", false, "Automatically start the container at boot")
	createCmd.Flags().String("bridge", "", "Attach the container to a host bridge interface")
	createCmd.Flags().String("config", "", "Inline NixOS configuration string")
	createCmd.Flags().String("config-file", "", "Path to a NixOS configuration file")
	createCmd.Flags().Bool("ensure-unique-name", false, "Append a unique suffix if the container name already exists")
	createCmd.Flags().String("flake", "", "Flake reference")
	createCmd.Flags().String("host-address", "", "Host-side IP address for the veth pair")
	createCmd.Flags().String("local-address", "", "Container-side IP address for the veth pair")
	createCmd.Flags().String("nixos-path", "", "Path to the NixOS source")
	createCmd.Flags().String("port", "", "Host port mapping")
	createCmd.Flags().String("system-path", "", "Path to a pre-built system derivation")
	createCmd.Flags().Bool("use-host-network", false, "Share the host's network namespace")

	carapace.Gen(createCmd).FlagCompletion(carapace.ActionMap{
		"config-file": carapace.ActionFiles(".nix"),
		"nixos-path":  carapace.ActionDirectories(),
		"system-path": carapace.ActionFiles(),
	})

	rootCmd.AddCommand(createCmd)
}
