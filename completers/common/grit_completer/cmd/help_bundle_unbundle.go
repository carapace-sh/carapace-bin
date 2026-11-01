package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_bundle_unbundleCmd = &cobra.Command{
	Use:   "unbundle",
	Short: "Ingest a bundle pack into the current repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_bundle_unbundleCmd).Standalone()

	help_bundleCmd.AddCommand(help_bundle_unbundleCmd)
}
