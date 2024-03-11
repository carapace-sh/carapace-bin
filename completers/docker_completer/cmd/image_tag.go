package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var image_tagCmd = &cobra.Command{
	Use:   "tag SOURCE_IMAGE[:TAG] TARGET_IMAGE[:TAG]",
	Short: "Create a tag TARGET_IMAGE that refers to SOURCE_IMAGE",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_tagCmd).Standalone()

	imageCmd.AddCommand(image_tagCmd)

	carapace.Gen(image_tagCmd).PositionalCompletion(
		docker.ActionRepositoryTags(),
		docker.ActionRepositoryTags(),
	)
}
