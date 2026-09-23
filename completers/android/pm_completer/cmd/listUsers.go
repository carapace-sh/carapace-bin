package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var listUsersCmd = &cobra.Command{
	Use:   "users",
	Short: "Print all users",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listUsersCmd).Standalone()

	listCmd.AddCommand(listUsersCmd)

}
