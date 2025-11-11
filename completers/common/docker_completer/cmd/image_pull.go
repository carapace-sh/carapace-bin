package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var image_pullCmd = &cobra.Command{
	Use:   "pull [OPTIONS] NAME[:TAG|@DIGEST]",
	Short: "Download an image from a registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_pullCmd).Standalone()

	image_pullCmd.Flags().BoolP("all-tags", "a", false, "Download all tagged images in the repository")
	image_pullCmd.Flags().Bool("disable-content-trust", false, "Skip image verification (deprecated)")
	image_pullCmd.Flags().String("platform", "", "Set platform if server is multi-platform capable")
	image_pullCmd.Flags().BoolP("quiet", "q", false, "Suppress verbose output")
	image_pullCmd.Flag("disable-content-trust").Hidden = true
	imageCmd.AddCommand(image_pullCmd)

	carapace.Gen(image_pullCmd).PositionalCompletion(docker.ActionRepositoryTags())
}
