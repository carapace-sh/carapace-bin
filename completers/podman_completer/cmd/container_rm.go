package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var container_rmCmd = &cobra.Command{
	Use:   "rm [options] CONTAINER [CONTAINER...]",
	Short: "Remove one or more containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_rmCmd).Standalone()

	container_rmCmd.Flags().BoolP("all", "a", false, "Remove all containers")
	container_rmCmd.Flags().StringSlice("cidfile", []string{}, "Read the container ID from the file")
	container_rmCmd.Flags().Bool("depend", false, "Remove container and all containers that depend on the selected container")
	container_rmCmd.Flags().StringSlice("filter", []string{}, "Filter output based on conditions given")
	container_rmCmd.Flags().BoolP("force", "f", false, "Force removal of a running or unusable container")
	container_rmCmd.Flags().BoolP("ignore", "i", false, "Ignore errors when a specified container is missing")
	container_rmCmd.Flags().BoolP("latest", "l", false, "Act on the latest container podman is aware of")
	container_rmCmd.Flags().Bool("storage", false, "Remove container from storage library")
	container_rmCmd.Flags().StringP("time", "t", "", "Seconds to wait for stop before killing the container")
	container_rmCmd.Flags().BoolP("volumes", "v", false, "Remove anonymous volumes associated with the container")
	container_rmCmd.Flag("storage").Hidden = true
	containerCmd.AddCommand(container_rmCmd)
}
