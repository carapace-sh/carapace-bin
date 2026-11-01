package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bundle_verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Check that a bundle file is well-formed",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bundle_verifyCmd).Standalone()

	bundle_verifyCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	bundleCmd.AddCommand(bundle_verifyCmd)

	carapace.Gen(bundle_verifyCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
