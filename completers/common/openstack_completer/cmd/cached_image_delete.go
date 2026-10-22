package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openstack_completer/cmd/action"
	"github.com/spf13/cobra"
)

var cached_image_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete image(s) from cache",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cached_image_deleteCmd).Standalone()

	cached_imageCmd.AddCommand(cached_image_deleteCmd)

	carapace.Gen(cached_image_deleteCmd).PositionalAnyCompletion(action.ActionCachedImages(cached_image_deleteCmd))
}
