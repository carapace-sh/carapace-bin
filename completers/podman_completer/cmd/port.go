package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var portCmd = &cobra.Command{
	Use:   "port [options] CONTAINER [PORT]",
	Short: "List port mappings or a specific mapping for the container",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(portCmd).Standalone()

	portCmd.Flags().BoolP("all", "a", false, "Display port information for all containers")
	portCmd.Flags().BoolP("latest", "l", false, "Act on the latest container podman is aware of")
	rootCmd.AddCommand(portCmd)
}
