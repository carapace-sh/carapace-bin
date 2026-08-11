package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lock_statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Get lock status on file(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lock_statusCmd).Standalone()

	lock_statusCmd.Flags().String("branch", "", "Branch where lock was acquired")
	lock_statusCmd.Flags().BoolP("help", "h", false, "Print help")
	lockCmd.AddCommand(lock_statusCmd)
}
