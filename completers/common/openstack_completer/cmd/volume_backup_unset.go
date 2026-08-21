package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var volume_backup_unsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Unset volume backup properties.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_backup_unsetCmd).Standalone()

	volume_backup_unsetCmd.Flags().String("property", "", "Property to remove from this backup (repeat option to unset multiple values) ")
	volume_backupCmd.AddCommand(volume_backup_unsetCmd)
}
