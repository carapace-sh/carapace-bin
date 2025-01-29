package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var unpauseCmd = &cobra.Command{
	Use:   "unpause [options] CONTAINER [CONTAINER...]",
	Short: "Unpause the processes in one or more containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unpauseCmd).Standalone()

	unpauseCmd.Flags().BoolP("all", "a", false, "Unpause all paused containers")
	unpauseCmd.Flags().StringSlice("cidfile", []string{}, "Read the container ID from the file")
	unpauseCmd.Flags().StringSliceP("filter", "f", []string{}, "Filter output based on conditions given")
	unpauseCmd.Flags().BoolP("latest", "l", false, "Act on the latest container podman is aware of")
	rootCmd.AddCommand(unpauseCmd)
}
