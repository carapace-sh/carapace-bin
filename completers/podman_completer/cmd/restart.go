package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var restartCmd = &cobra.Command{
	Use:   "restart [options] CONTAINER [CONTAINER...]",
	Short: "Restart one or more containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(restartCmd).Standalone()

	restartCmd.Flags().BoolP("all", "a", false, "Restart all non-running containers")
	restartCmd.Flags().StringSlice("cidfile", []string{}, "Read the container ID from the file")
	restartCmd.Flags().StringSliceP("filter", "f", []string{}, "Filter output based on conditions given")
	restartCmd.Flags().BoolP("latest", "l", false, "Act on the latest container podman is aware of")
	restartCmd.Flags().Bool("running", false, "Restart only running containers")
	restartCmd.Flags().StringP("time", "t", "", "Seconds to wait for stop before killing the container")
	rootCmd.AddCommand(restartCmd)
}
