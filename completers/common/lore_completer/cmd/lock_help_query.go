package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lock_help_queryCmd = &cobra.Command{
	Use:   "query",
	Short: "Query the lock status given a branch, owner or path",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lock_help_queryCmd).Standalone()

	lock_helpCmd.AddCommand(lock_help_queryCmd)
}
