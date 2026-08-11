package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_auth_clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear all stored authentication data",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_auth_clearCmd).Standalone()

	help_authCmd.AddCommand(help_auth_clearCmd)
}
