package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_backup_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set share backup properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_backup_setCmd).Standalone()

	share_backup_setCmd.Flags().String("description", "", "Set a description to the backup.")
	share_backup_setCmd.Flags().String("name", "", "Set a name to the backup.")
	share_backup_setCmd.Flags().String("status", "", "Assign a status to the backup(Admin only).")
	share_backupCmd.AddCommand(share_backup_setCmd)
}
