package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auth_help_infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Display identity information for the current user or specified user IDs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auth_help_infoCmd).Standalone()

	auth_helpCmd.AddCommand(auth_help_infoCmd)
}
