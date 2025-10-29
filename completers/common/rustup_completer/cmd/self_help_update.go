package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var self_help_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Download and install updates to rustup",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(self_help_updateCmd).Standalone()

	self_helpCmd.AddCommand(self_help_updateCmd)
}
