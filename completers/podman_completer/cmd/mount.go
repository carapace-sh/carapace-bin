package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mountCmd = &cobra.Command{
	Use:   "mount [options] [CONTAINER...]",
	Short: "Mount a working container's root filesystem",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mountCmd).Standalone()

	mountCmd.Flags().BoolP("all", "a", false, "Mount all containers")
	mountCmd.Flags().String("format", "", "Print the mounted containers in specified format (json)")
	mountCmd.Flags().BoolP("latest", "l", false, "Act on the latest container podman is aware of")
	mountCmd.Flags().Bool("no-trunc", false, "Do not truncate output")
	rootCmd.AddCommand(mountCmd)
}
