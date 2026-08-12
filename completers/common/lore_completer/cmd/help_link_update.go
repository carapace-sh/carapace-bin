package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_link_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the link to a new pin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_link_updateCmd).Standalone()

	help_linkCmd.AddCommand(help_link_updateCmd)
}
