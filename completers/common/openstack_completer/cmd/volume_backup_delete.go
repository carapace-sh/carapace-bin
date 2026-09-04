package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var volume_backup_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete volume backup(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_backup_deleteCmd).Standalone()

	volume_backup_deleteCmd.Flags().Bool("force", false, "Allow delete in state other than error or available")
	volume_backupCmd.AddCommand(volume_backup_deleteCmd)
}
