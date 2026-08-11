package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lock_queryCmd = &cobra.Command{
	Use:   "query",
	Short: "Query the lock status given a branch, owner or path",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lock_queryCmd).Standalone()

	lock_queryCmd.Flags().String("branch", "", "Branch to query locks on")
	lock_queryCmd.Flags().BoolP("help", "h", false, "Print help")
	lock_queryCmd.Flags().String("owner", "", "Owner to query locks belonging to them")
	lock_queryCmd.Flags().String("path", "", "Path to query lock information on")
	lockCmd.AddCommand(lock_queryCmd)
}
