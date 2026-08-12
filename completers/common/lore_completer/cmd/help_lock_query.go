package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_lock_queryCmd = &cobra.Command{
	Use:   "query",
	Short: "Query the lock status given a branch, owner or path",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_lock_queryCmd).Standalone()

	help_lockCmd.AddCommand(help_lock_queryCmd)
}
