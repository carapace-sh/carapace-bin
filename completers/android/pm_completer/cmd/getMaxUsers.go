package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var getMaxUsersCmd = &cobra.Command{
	Use:   "get-max-users",
	Short: "Return the maximum number of users supported",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(getMaxUsersCmd).Standalone()

	rootCmd.AddCommand(getMaxUsersCmd)

}
