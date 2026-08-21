package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cached_image_queueCmd = &cobra.Command{
	Use:   "queue",
	Short: "Queue image(s) for caching.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cached_image_queueCmd).Standalone()

	cached_imageCmd.AddCommand(cached_image_queueCmd)
}
