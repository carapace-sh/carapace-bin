package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_forge_listUsersCmd = &cobra.Command{
	Use:   "list-users",
	Short: "List authenticated forge accounts known to GitButler",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_forge_listUsersCmd).Standalone()

	config_forge_listUsersCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	config_forgeCmd.AddCommand(config_forge_listUsersCmd)
}
