package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_backup_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete one or more share backups",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_backup_deleteCmd).Standalone()

	share_backup_deleteCmd.Flags().Bool("wait", false, "Wait for share backup deletion")
	share_backupCmd.AddCommand(share_backup_deleteCmd)
}
