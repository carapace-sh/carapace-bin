package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var container_portCmd = &cobra.Command{
	Use:   "port [options] CONTAINER [PORT]",
	Short: "List port mappings or a specific mapping for the container",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_portCmd).Standalone()

	container_portCmd.Flags().BoolP("all", "a", false, "Display port information for all containers")
	container_portCmd.Flags().BoolP("latest", "l", false, "Act on the latest container podman is aware of")
	containerCmd.AddCommand(container_portCmd)
}
