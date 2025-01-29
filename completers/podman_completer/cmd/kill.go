package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var killCmd = &cobra.Command{
	Use:   "kill [options] CONTAINER [CONTAINER...]",
	Short: "Kill one or more running containers with a specific signal",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(killCmd).Standalone()

	killCmd.Flags().BoolP("all", "a", false, "Signal all running containers")
	killCmd.Flags().StringSlice("cidfile", []string{}, "Read the container ID from the file")
	killCmd.Flags().BoolP("latest", "l", false, "Act on the latest container podman is aware of")
	killCmd.Flags().StringP("signal", "s", "", "Signal to send to the container")
	rootCmd.AddCommand(killCmd)
}
