package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vendorPublishCmd = &cobra.Command{
	Use:   "vendor:publish [--existing] [--force] [--all] [--provider [PROVIDER]] [--tag [TAG]]",
	Short: "Publish any publishable assets from vendor packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vendorPublishCmd).Standalone()

	vendorPublishCmd.Flags().Bool("all", false, "Publish assets for all service providers without prompt")
	vendorPublishCmd.Flags().Bool("existing", false, "Publish and overwrite only the files that have already been published")
	vendorPublishCmd.Flags().Bool("force", false, "Overwrite any existing files")
	vendorPublishCmd.Flags().String("provider", "", "The service provider that has assets you want to publish")
	vendorPublishCmd.Flags().String("tag", "", "One or many tags that have assets you want to publish")
	rootCmd.AddCommand(vendorPublishCmd)
}
