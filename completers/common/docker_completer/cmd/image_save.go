package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var image_saveCmd = &cobra.Command{
	Use:   "save [OPTIONS] IMAGE [IMAGE...]",
	Short: "Save one or more images to a tar archive (streamed to STDOUT by default)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_saveCmd).Standalone()

	image_saveCmd.Flags().StringP("output", "o", "", "Write to a file, instead of STDOUT")
	imageCmd.AddCommand(image_saveCmd)

	carapace.Gen(image_saveCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionFiles(),
	})

	carapace.Gen(image_saveCmd).PositionalAnyCompletion(docker.ActionRepositoryTags())
}
