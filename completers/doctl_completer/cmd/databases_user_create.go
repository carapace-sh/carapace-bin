package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var databases_user_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a database user",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databases_user_createCmd).Standalone()
	databases_user_createCmd.Flags().String("mysql-auth-plugin", "", "set auth mode for MySQL users")
	databases_userCmd.AddCommand(databases_user_createCmd)
}
