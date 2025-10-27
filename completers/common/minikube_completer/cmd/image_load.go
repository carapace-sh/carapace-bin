package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/carapace-sh/carapace/pkg/condition"
	"github.com/spf13/cobra"
)

var image_loadCmd = &cobra.Command{
	Use:   "load IMAGE | ARCHIVE | -",
	Short: "Load an image into minikube",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_loadCmd).Standalone()

	image_loadCmd.Flags().Bool("daemon", false, "Cache image from docker daemon")
	image_loadCmd.Flags().Bool("overwrite", false, "Overwrite image even if same image:tag name exists")
	image_loadCmd.Flags().Bool("pull", false, "Pull the remote image (no caching)")
	image_loadCmd.Flags().Bool("remote", false, "Cache image from remote registry")
	imageCmd.AddCommand(image_loadCmd)

	carapace.Gen(image_loadCmd).PositionalCompletion(
		carapace.Batch(
			carapace.ActionFiles(),
			docker.ActionRepositoryTags().UnlessF(condition.CompletingPath),
		).ToA(),
	)
}
