package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lock_releaseCmd = &cobra.Command{
	Use:   "release",
	Short: "Release lock on file(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lock_releaseCmd).Standalone()

	lock_releaseCmd.Flags().String("branch", "", "Branch where lock was acquired")
	lock_releaseCmd.Flags().BoolP("help", "h", false, "Print help")
	lock_releaseCmd.Flags().String("owner", "", "Owner of the lock")
	lockCmd.AddCommand(lock_releaseCmd)
}
