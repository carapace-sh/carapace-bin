package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bundle_help_verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Check that a bundle file is well-formed",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bundle_help_verifyCmd).Standalone()

	bundle_helpCmd.AddCommand(bundle_help_verifyCmd)
}
