package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_bundle_verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Check that a bundle file is well-formed",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_bundle_verifyCmd).Standalone()

	help_bundleCmd.AddCommand(help_bundle_verifyCmd)
}
