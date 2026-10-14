package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var volume_backup_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set volume backup properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_backup_setCmd).Standalone()

	volume_backup_setCmd.Flags().String("description", "", "New backup description (supported by --os-volume-api-version 3.9 or above)")
	volume_backup_setCmd.Flags().String("name", "", "New backup name (supported by --os-volume-api-version 3.9 or above)")
	volume_backup_setCmd.Flags().Bool("no-property", false, "Remove all properties from this backup (specify both --no-property and --property to remove the current properties before setting new properties)")
	volume_backup_setCmd.Flags().String("property", "", "Set a property on this backup (repeat option to set multiple values) (supported by --os-volume-api-version 3.43 or above)")
	volume_backup_setCmd.Flags().String("state", "", "New backup state (\"available\" or \"error\") (admin only) (This option simply changes the state of the backup in the database with no regard to actual status; exercise caution when using)")
	volume_backupCmd.AddCommand(volume_backup_setCmd)
}
