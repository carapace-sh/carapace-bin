package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lock_help_statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Get lock status on file(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lock_help_statusCmd).Standalone()

	lock_helpCmd.AddCommand(lock_help_statusCmd)
}
