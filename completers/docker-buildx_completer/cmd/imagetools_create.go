package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var imagetools_createCmd = &cobra.Command{
	Use:   "create [OPTIONS] [SOURCE] [SOURCE...]",
	Short: "Create a new image based on source images",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagetools_createCmd).Standalone()
	imagetools_createCmd.Flags().Bool("append", false, "Append to existing manifest")
	imagetools_createCmd.Flags().Bool("dry-run", false, "Show final image instead of pushing")
	imagetools_createCmd.Flags().StringArrayP("file", "f", nil, "Read source descriptor from file")
	imagetools_createCmd.Flags().String("progress", "auto", "Set type of progress output (\"auto\", \"plain\", \"tty\"). Use plain to show container output")
	imagetools_createCmd.Flags().StringArrayP("tag", "t", nil, "Set reference for new image")
	imagetoolsCmd.AddCommand(imagetools_createCmd)

	carapace.Gen(imagetools_createCmd).FlagCompletion(carapace.ActionMap{
		"file":     carapace.ActionFiles(),
		"progress": carapace.ActionValues("auto", "plain", "tty"),
	})

	// TODO verify positional completion - is this correct?
	carapace.Gen(imagetools_createCmd).PositionalAnyCompletion(
		docker.ActionRepositoryTags(),
	)
}
