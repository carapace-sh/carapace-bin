package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var unmountCmd = &cobra.Command{
	Use:     "unmount [options] CONTAINER [CONTAINER...]",
	Short:   "Unmount working container's root filesystem",
	Aliases: []string{"umount"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unmountCmd).Standalone()

	unmountCmd.Flags().BoolP("all", "a", false, "Unmount all of the currently mounted containers")
	unmountCmd.Flags().BoolP("force", "f", false, "Force the complete unmount of the specified mounted containers")
	unmountCmd.Flags().BoolP("latest", "l", false, "Act on the latest container podman is aware of")
	rootCmd.AddCommand(unmountCmd)
}
