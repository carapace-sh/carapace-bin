package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop [options] CONTAINER [CONTAINER...]",
	Short: "Stop one or more containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stopCmd).Standalone()

	stopCmd.Flags().BoolP("all", "a", false, "Stop all running containers")
	stopCmd.Flags().StringSlice("cidfile", []string{}, "Read the container ID from the file")
	stopCmd.Flags().StringSliceP("filter", "f", []string{}, "Filter output based on conditions given")
	stopCmd.Flags().BoolP("ignore", "i", false, "Ignore errors when a specified container is missing")
	stopCmd.Flags().BoolP("latest", "l", false, "Act on the latest container podman is aware of")
	stopCmd.Flags().StringP("time", "t", "", "Seconds to wait for stop before killing the container")
	rootCmd.AddCommand(stopCmd)
}
