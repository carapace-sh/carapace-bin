package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bundle_unbundleCmd = &cobra.Command{
	Use:   "unbundle",
	Short: "Ingest a bundle pack into the current repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bundle_unbundleCmd).Standalone()

	bundle_unbundleCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	bundleCmd.AddCommand(bundle_unbundleCmd)

	carapace.Gen(bundle_unbundleCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
