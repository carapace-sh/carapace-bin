package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_forge_authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authenticat with your forge provider (at the moment, only GitHub is supported)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_forge_authCmd).Standalone()

	help_forgeCmd.AddCommand(help_forge_authCmd)
}
