package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var build_nodeImageCmd = &cobra.Command{
	Use:   "node-image [kubernetes-source]",
	Short: "Build the node image",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(build_nodeImageCmd).Standalone()

	build_nodeImageCmd.Flags().String("arch", "", "architecture to build for, defaults to the host architecture")
	build_nodeImageCmd.Flags().String("base-image", "docker.io/kindest/base:v20260820-69b56db7", "name:tag of the base image to use for the build")
	build_nodeImageCmd.Flags().String("image", "kindest/node:latest", "name:tag of the resulting image to be built")
	build_nodeImageCmd.Flags().String("type", "", "optionally specify one of 'url', 'file', 'release', 'ci' or 'source' as the type of build")

	carapace.Gen(build_nodeImageCmd).FlagCompletion(carapace.ActionMap{
		"arch":       carapace.ActionValues("amd64", "arm64"),
		"base-image": docker.ActionRepositoryTags(),
		"image":      docker.ActionRepositoryTags(),
		"type":       carapace.ActionValues("ci", "file", "release", "source", "url"),
	})

	carapace.Gen(build_nodeImageCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			switch build_nodeImageCmd.Flag("type").Value.String() {
			case "file":
				return carapace.ActionFiles()
			case "source":
				return carapace.ActionDirectories()
			default:
				return carapace.ActionFiles()
			}
		}),
	)

	buildCmd.AddCommand(build_nodeImageCmd)
}
