package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cached_image_clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear all images from cache, queue or both",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cached_image_clearCmd).Standalone()

	cached_image_clearCmd.Flags().Bool("cache", false, "Clears all the cached images")
	cached_image_clearCmd.Flags().Bool("queue", false, "Clears all the queued images")
	cached_imageCmd.AddCommand(cached_image_clearCmd)
}
