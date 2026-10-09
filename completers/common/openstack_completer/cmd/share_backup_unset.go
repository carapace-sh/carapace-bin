package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_backup_unsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Unset share backup properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_backup_unsetCmd).Standalone()

	share_backup_unsetCmd.Flags().Bool("description", false, "Unset a description to the backup.")
	share_backup_unsetCmd.Flags().Bool("name", false, "Unset a name to the backup.")
	share_backupCmd.AddCommand(share_backup_unsetCmd)
}
