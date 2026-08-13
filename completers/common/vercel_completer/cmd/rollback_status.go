package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rollback_statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show the status of any current pending rollbacks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rollback_statusCmd).Standalone()

	rollback_statusCmd.Flags().String("timeout", "", "Time to wait for rollback completion")

	rollbackCmd.AddCommand(rollback_statusCmd)
}
