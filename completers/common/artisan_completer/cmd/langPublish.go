package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var langPublishCmd = &cobra.Command{
	Use:   "lang:publish [--existing] [--force]",
	Short: "Publish all language files that are available for customization",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(langPublishCmd).Standalone()

	langPublishCmd.Flags().Bool("existing", false, "Publish and overwrite only the files that have already been published")
	langPublishCmd.Flags().Bool("force", false, "Overwrite any existing files")
	rootCmd.AddCommand(langPublishCmd)
}
