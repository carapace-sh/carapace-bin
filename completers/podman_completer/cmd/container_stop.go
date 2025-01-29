package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var container_stopCmd = &cobra.Command{
	Use:   "stop [options] CONTAINER [CONTAINER...]",
	Short: "Stop one or more containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_stopCmd).Standalone()

	container_stopCmd.Flags().BoolP("all", "a", false, "Stop all running containers")
	container_stopCmd.Flags().StringSlice("cidfile", []string{}, "Read the container ID from the file")
	container_stopCmd.Flags().StringSliceP("filter", "f", []string{}, "Filter output based on conditions given")
	container_stopCmd.Flags().BoolP("ignore", "i", false, "Ignore errors when a specified container is missing")
	container_stopCmd.Flags().BoolP("latest", "l", false, "Act on the latest container podman is aware of")
	container_stopCmd.Flags().StringP("time", "t", "", "Seconds to wait for stop before killing the container")
	containerCmd.AddCommand(container_stopCmd)
}
