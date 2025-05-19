package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var imagetools_inspectCmd = &cobra.Command{
	Use:   "inspect [OPTIONS] NAME",
	Short: "Show details of an image in the registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagetools_inspectCmd).Standalone()
	imagetools_inspectCmd.Flags().String("format", "", "Format the output using the given Go template")
	imagetools_inspectCmd.Flags().Bool("raw", false, "Show original, unformatted JSON manifest")
	imagetoolsCmd.AddCommand(imagetools_inspectCmd)

	// TODO verify positional completion - is this correct?
	carapace.Gen(imagetools_inspectCmd).PositionalCompletion(
		docker.ActionRepositoryTags(),
	)
}
