package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pauseCmd = &cobra.Command{
	Use:   "pause [options] CONTAINER [CONTAINER...]",
	Short: "Pause all the processes in one or more containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pauseCmd).Standalone()

	pauseCmd.Flags().BoolP("all", "a", false, "Pause all running containers")
	pauseCmd.Flags().StringSlice("cidfile", []string{}, "Read the container ID from the file")
	pauseCmd.Flags().StringSliceP("filter", "f", []string{}, "Filter output based on conditions given")
	pauseCmd.Flags().BoolP("latest", "l", false, "Act on the latest container podman is aware of")
	rootCmd.AddCommand(pauseCmd)
}
