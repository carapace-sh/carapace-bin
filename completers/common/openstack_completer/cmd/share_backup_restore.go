package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_backup_restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Attempt to restore share backup",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_backup_restoreCmd).Standalone()

	share_backup_restoreCmd.Flags().String("target-share", "", "share to restore backup to.")
	share_backup_restoreCmd.Flags().Bool("wait", false, "Wait for restore conclusion")
	share_backupCmd.AddCommand(share_backup_restoreCmd)
}
