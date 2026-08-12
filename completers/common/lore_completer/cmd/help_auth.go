package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authentication commands",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_authCmd).Standalone()

	helpCmd.AddCommand(help_authCmd)
}
