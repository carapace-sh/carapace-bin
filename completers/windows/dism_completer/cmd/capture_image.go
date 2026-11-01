package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var CaptureImageCmd = &cobra.Command{
	Use:   "Capture-Image",
	Short: "capture an image from a directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(CaptureImageCmd).Standalone()
	rootCmd.AddCommand(CaptureImageCmd)
}
