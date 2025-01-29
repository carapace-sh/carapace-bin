package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var container_startCmd = &cobra.Command{
	Use:   "start [options] CONTAINER [CONTAINER...]",
	Short: "Start one or more containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_startCmd).Standalone()

	container_startCmd.Flags().Bool("all", false, "Start all containers regardless of their state or configuration")
	container_startCmd.Flags().BoolP("attach", "a", false, "Attach container's STDOUT and STDERR")
	container_startCmd.Flags().String("detach-keys", "", "Select the key sequence for detaching a container")
	container_startCmd.Flags().StringSliceP("filter", "f", []string{}, "Filter output based on conditions given")
	container_startCmd.Flags().BoolP("interactive", "i", false, "Make STDIN available to the contained process")
	container_startCmd.Flags().BoolP("latest", "l", false, "Act on the latest container podman is aware of")
	container_startCmd.Flags().Bool("sig-proxy", false, "Proxy received signals to the process (default true if attaching, false otherwise)")
	containerCmd.AddCommand(container_startCmd)
}
