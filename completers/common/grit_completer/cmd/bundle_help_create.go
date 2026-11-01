package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bundle_help_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Write a bundle file from revision specs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bundle_help_createCmd).Standalone()

	bundle_helpCmd.AddCommand(bundle_help_createCmd)
}
