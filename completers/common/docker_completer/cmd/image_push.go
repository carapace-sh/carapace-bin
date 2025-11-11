package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var image_pushCmd = &cobra.Command{
	Use:   "push [OPTIONS] NAME[:TAG]",
	Short: "Upload an image to a registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_pushCmd).Standalone()

	image_pushCmd.Flags().BoolP("all-tags", "a", false, "Push all tags of an image to the repository")
	image_pushCmd.Flags().Bool("disable-content-trust", false, "Skip image verification (deprecated)")
	image_pushCmd.Flags().String("platform", "", "Push a platform-specific manifest as a single-platform image to the registry.")
	image_pushCmd.Flags().BoolP("quiet", "q", false, "Suppress verbose output")
	image_pushCmd.Flag("disable-content-trust").Hidden = true
	imageCmd.AddCommand(image_pushCmd)

	carapace.Gen(image_pushCmd).PositionalCompletion(
		docker.ActionRepositoryTags(),
	)
}
