package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var container_cleanupCmd = &cobra.Command{
	Use:   "cleanup [options] CONTAINER [CONTAINER...]",
	Short: "Clean up network and mountpoints of one or more containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_cleanupCmd).Standalone()

	container_cleanupCmd.Flags().BoolP("all", "a", false, "Cleans up all containers")
	container_cleanupCmd.Flags().String("exec", "", "Clean up the given exec session instead of the container")
	container_cleanupCmd.Flags().BoolP("latest", "l", false, "Act on the latest container podman is aware of")
	container_cleanupCmd.Flags().Bool("rm", false, "After cleanup, remove the container entirely")
	container_cleanupCmd.Flags().Bool("rmi", false, "After cleanup, remove the image entirely")
	container_cleanupCmd.Flags().Bool("stopped-only", false, "Only cleanup when the container is in the stopped state")
	container_cleanupCmd.Flag("stopped-only").Hidden = true
	containerCmd.AddCommand(container_cleanupCmd)
}
