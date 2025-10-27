package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var image_rmCmd = &cobra.Command{
	Use:     "rm [OPTIONS] IMAGE [IMAGE...]",
	Short:   "Remove one or more images",
	Aliases: []string{"rmi", "remove"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_rmCmd).Standalone()

	image_rmCmd.Flags().BoolP("force", "f", false, "Force removal of the image")
	image_rmCmd.Flags().Bool("no-prune", false, "Do not delete untagged parents")
	image_rmCmd.Flags().StringSlice("platform", nil, "Remove only the given platform variant. Formatted as \"os[/arch[/variant]]\" (e.g., \"linux/amd64\")")
	imageCmd.AddCommand(image_rmCmd)

	carapace.Gen(image_rmCmd).PositionalAnyCompletion(docker.ActionRepositoryTags())
}
