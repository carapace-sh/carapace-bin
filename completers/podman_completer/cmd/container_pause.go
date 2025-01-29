package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var container_pauseCmd = &cobra.Command{
	Use:   "pause [options] CONTAINER [CONTAINER...]",
	Short: "Pause all the processes in one or more containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_pauseCmd).Standalone()

	container_pauseCmd.Flags().BoolP("all", "a", false, "Pause all running containers")
	container_pauseCmd.Flags().StringSlice("cidfile", []string{}, "Read the container ID from the file")
	container_pauseCmd.Flags().StringSliceP("filter", "f", []string{}, "Filter output based on conditions given")
	container_pauseCmd.Flags().BoolP("latest", "l", false, "Act on the latest container podman is aware of")
	containerCmd.AddCommand(container_pauseCmd)
}
