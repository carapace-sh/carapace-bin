package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stubPublishCmd = &cobra.Command{
	Use:   "stub:publish [--existing] [--force]",
	Short: "Publish all stubs that are available for customization",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stubPublishCmd).Standalone()

	stubPublishCmd.Flags().Bool("existing", false, "Publish and overwrite only the files that have already been published")
	stubPublishCmd.Flags().Bool("force", false, "Overwrite any existing files")
	rootCmd.AddCommand(stubPublishCmd)
}
