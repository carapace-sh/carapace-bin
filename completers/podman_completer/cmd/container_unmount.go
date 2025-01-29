package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var container_unmountCmd = &cobra.Command{
	Use:     "unmount [options] CONTAINER [CONTAINER...]",
	Short:   "Unmount working container's root filesystem",
	Aliases: []string{"umount"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_unmountCmd).Standalone()

	container_unmountCmd.Flags().BoolP("all", "a", false, "Unmount all of the currently mounted containers")
	container_unmountCmd.Flags().BoolP("force", "f", false, "Force the complete unmount of the specified mounted containers")
	container_unmountCmd.Flags().BoolP("latest", "l", false, "Act on the latest container podman is aware of")
	containerCmd.AddCommand(container_unmountCmd)
}
