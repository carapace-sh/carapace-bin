package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/docker_completer/cmd/action"
	"github.com/spf13/cobra"
)

var image_tagCmd = &cobra.Command{
	Use:   "tag",
	Short: "Create a tag TARGET_IMAGE that refers to SOURCE_IMAGE",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_tagCmd).Standalone()

	imageCmd.AddCommand(image_tagCmd)

	rootAlias(image_tagCmd, func(cmd *cobra.Command, isAlias bool) {
		carapace.Gen(cmd).PositionalCompletion(
			action.ActionRepositoryTags(),
			action.ActionRepositoryTags(),
		)
	})
}
