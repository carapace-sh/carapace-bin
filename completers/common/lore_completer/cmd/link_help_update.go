package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var link_help_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the link to a new pin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(link_help_updateCmd).Standalone()

	link_helpCmd.AddCommand(link_help_updateCmd)
}
