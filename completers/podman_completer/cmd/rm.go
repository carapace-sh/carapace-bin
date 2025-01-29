package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm [options] CONTAINER [CONTAINER...]",
	Short: "Remove one or more containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rmCmd).Standalone()

	rmCmd.Flags().BoolP("all", "a", false, "Remove all containers")
	rmCmd.Flags().StringSlice("cidfile", []string{}, "Read the container ID from the file")
	rmCmd.Flags().Bool("depend", false, "Remove container and all containers that depend on the selected container")
	rmCmd.Flags().StringSlice("filter", []string{}, "Filter output based on conditions given")
	rmCmd.Flags().BoolP("force", "f", false, "Force removal of a running or unusable container")
	rmCmd.Flags().BoolP("ignore", "i", false, "Ignore errors when a specified container is missing")
	rmCmd.Flags().BoolP("latest", "l", false, "Act on the latest container podman is aware of")
	rmCmd.Flags().Bool("storage", false, "Remove container from storage library")
	rmCmd.Flags().StringP("time", "t", "", "Seconds to wait for stop before killing the container")
	rmCmd.Flags().BoolP("volumes", "v", false, "Remove anonymous volumes associated with the container")
	rmCmd.Flag("storage").Hidden = true
	rootCmd.AddCommand(rmCmd)
}
