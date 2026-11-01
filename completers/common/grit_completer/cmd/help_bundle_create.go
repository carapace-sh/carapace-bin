package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_bundle_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Write a bundle file from revision specs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_bundle_createCmd).Standalone()

	help_bundleCmd.AddCommand(help_bundle_createCmd)
}
