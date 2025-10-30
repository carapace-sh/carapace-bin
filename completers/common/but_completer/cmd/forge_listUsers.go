package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var forge_listUsersCmd = &cobra.Command{
	Use:   "list-users",
	Short: "List authenticated forge accounts known to GitButler",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forge_listUsersCmd).Standalone()

	forge_listUsersCmd.Flags().BoolP("help", "h", false, "Print help")
	forgeCmd.AddCommand(forge_listUsersCmd)
}
