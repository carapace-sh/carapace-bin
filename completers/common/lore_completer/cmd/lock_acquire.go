package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lock_acquireCmd = &cobra.Command{
	Use:   "acquire",
	Short: "Acquire lock on file(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lock_acquireCmd).Standalone()

	lock_acquireCmd.Flags().String("branch", "", "Branch where lock is to be acquired")
	lock_acquireCmd.Flags().BoolP("help", "h", false, "Print help")
	lockCmd.AddCommand(lock_acquireCmd)
}
