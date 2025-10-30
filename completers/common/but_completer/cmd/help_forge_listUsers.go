package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_forge_listUsersCmd = &cobra.Command{
	Use:   "list-users",
	Short: "List authenticated forge accounts known to GitButler",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_forge_listUsersCmd).Standalone()

	help_forgeCmd.AddCommand(help_forge_listUsersCmd)
}
