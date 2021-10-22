package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var databases_user_resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Resets a user's auth",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databases_user_resetCmd).Standalone()
	databases_userCmd.AddCommand(databases_user_resetCmd)
}
