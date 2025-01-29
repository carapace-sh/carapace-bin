package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var container_waitCmd = &cobra.Command{
	Use:   "wait [options] CONTAINER [CONTAINER...]",
	Short: "Block on one or more containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_waitCmd).Standalone()

	container_waitCmd.Flags().StringSlice("condition", []string{}, "Condition to wait on")
	container_waitCmd.Flags().Bool("ignore", false, "Ignore if a container does not exist")
	container_waitCmd.Flags().StringP("interval", "i", "", "Time Interval to wait before polling for completion")
	container_waitCmd.Flags().BoolP("latest", "l", false, "Act on the latest container podman is aware of")
	containerCmd.AddCommand(container_waitCmd)
}
