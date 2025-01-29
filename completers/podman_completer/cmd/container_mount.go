package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var container_mountCmd = &cobra.Command{
	Use:   "mount [options] [CONTAINER...]",
	Short: "Mount a working container's root filesystem",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_mountCmd).Standalone()

	container_mountCmd.Flags().BoolP("all", "a", false, "Mount all containers")
	container_mountCmd.Flags().String("format", "", "Print the mounted containers in specified format (json)")
	container_mountCmd.Flags().BoolP("latest", "l", false, "Act on the latest container podman is aware of")
	container_mountCmd.Flags().Bool("no-trunc", false, "Do not truncate output")
	containerCmd.AddCommand(container_mountCmd)
}
