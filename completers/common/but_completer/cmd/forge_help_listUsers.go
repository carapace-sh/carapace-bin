package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var forge_help_listUsersCmd = &cobra.Command{
	Use:   "list-users",
	Short: "List authenticated forge accounts known to GitButler",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forge_help_listUsersCmd).Standalone()

	forge_helpCmd.AddCommand(forge_help_listUsersCmd)
}
