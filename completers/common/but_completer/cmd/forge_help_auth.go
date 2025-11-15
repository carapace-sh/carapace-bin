package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var forge_help_authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authenticate with your forge provider (at the moment, only GitHub is supported)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forge_help_authCmd).Standalone()

	forge_helpCmd.AddCommand(forge_help_authCmd)
}
