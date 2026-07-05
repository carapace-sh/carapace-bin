package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var MountImageCmd = &cobra.Command{
	Use:   "Mount-Image",
	Short: "mount an image to a directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(MountImageCmd).Standalone()
	rootCmd.AddCommand(MountImageCmd)
}
