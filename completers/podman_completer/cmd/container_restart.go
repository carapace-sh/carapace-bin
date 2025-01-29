package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var container_restartCmd = &cobra.Command{
	Use:   "restart [options] CONTAINER [CONTAINER...]",
	Short: "Restart one or more containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_restartCmd).Standalone()

	container_restartCmd.Flags().BoolP("all", "a", false, "Restart all non-running containers")
	container_restartCmd.Flags().StringSlice("cidfile", []string{}, "Read the container ID from the file")
	container_restartCmd.Flags().StringSliceP("filter", "f", []string{}, "Filter output based on conditions given")
	container_restartCmd.Flags().BoolP("latest", "l", false, "Act on the latest container podman is aware of")
	container_restartCmd.Flags().Bool("running", false, "Restart only running containers")
	container_restartCmd.Flags().StringP("time", "t", "", "Seconds to wait for stop before killing the container")
	containerCmd.AddCommand(container_restartCmd)
}
