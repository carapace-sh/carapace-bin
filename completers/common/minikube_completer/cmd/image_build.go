package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var image_buildCmd = &cobra.Command{
	Use:   "build PATH | URL | -",
	Short: "Build a container image in minikube",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_buildCmd).Standalone()

	image_buildCmd.Flags().Bool("all", false, "Build image on all nodes.")
	image_buildCmd.Flags().StringSlice("build-env", nil, "Environment variables to pass to the build. (format: key=value)")
	image_buildCmd.Flags().StringSlice("build-opt", nil, "Specify arbitrary flags to pass to the build. (format: key=value)")
	image_buildCmd.Flags().StringP("file", "f", "", "Path to the Dockerfile to use (optional)")
	image_buildCmd.Flags().StringP("node", "n", "", "The node to build on. Defaults to the primary control plane.")
	image_buildCmd.Flags().Bool("push", false, "Push the new image (requires tag)")
	image_buildCmd.Flags().StringP("tag", "t", "", "Tag to apply to the new image (optional)")
	imageCmd.AddCommand(image_buildCmd)

	carapace.Gen(image_buildCmd).FlagCompletion(carapace.ActionMap{
		"file": carapace.ActionFiles(),
	})

	carapace.Gen(image_buildCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
