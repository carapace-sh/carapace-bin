package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var buildImageCmd = &cobra.Command{
	Use:   "build-image",
	Short: "Build a NixOS image from the configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(buildImageCmd).Standalone()
	rootCmd.AddCommand(buildImageCmd)
}
