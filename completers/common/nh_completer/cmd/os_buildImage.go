package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var os_buildImageCmd = &cobra.Command{
	Use:   "build-image",
	Short: "Build a NixOS disk-image variant",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(os_buildImageCmd).Standalone()

	addOsRebuildFlags(os_buildImageCmd)
	os_buildImageCmd.Flags().String("image-variant", "", "Image variant")

	osCmd.AddCommand(os_buildImageCmd)
}
