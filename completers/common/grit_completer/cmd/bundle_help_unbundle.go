package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bundle_help_unbundleCmd = &cobra.Command{
	Use:   "unbundle",
	Short: "Ingest a bundle pack into the current repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bundle_help_unbundleCmd).Standalone()

	bundle_helpCmd.AddCommand(bundle_help_unbundleCmd)
}
