package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auth_help_clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear all stored authentication data",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auth_help_clearCmd).Standalone()

	auth_helpCmd.AddCommand(auth_help_clearCmd)
}
