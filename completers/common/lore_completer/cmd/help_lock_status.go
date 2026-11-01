package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_lock_statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Get lock status on file(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_lock_statusCmd).Standalone()

	help_lockCmd.AddCommand(help_lock_statusCmd)
}
