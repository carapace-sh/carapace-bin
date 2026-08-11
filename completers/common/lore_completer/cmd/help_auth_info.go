package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_auth_infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Display identity information for the current user or specified user IDs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_auth_infoCmd).Standalone()

	help_authCmd.AddCommand(help_auth_infoCmd)
}
