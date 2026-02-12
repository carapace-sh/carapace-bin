package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Log out",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_logoutCmd).Standalone()

	helpCmd.AddCommand(help_logoutCmd)
}
