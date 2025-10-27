package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var image_inspectCmd = &cobra.Command{
	Use:   "inspect [OPTIONS] IMAGE [IMAGE...]",
	Short: "Display detailed information on one or more images",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_inspectCmd).Standalone()

	image_inspectCmd.Flags().StringP("format", "f", "", "Format output using a custom template:")
	image_inspectCmd.Flags().String("platform", "", "Inspect a specific platform of the multi-platform image.")
	imageCmd.AddCommand(image_inspectCmd)

	carapace.Gen(image_inspectCmd).PositionalAnyCompletion(docker.ActionRepositoryTags())
}
