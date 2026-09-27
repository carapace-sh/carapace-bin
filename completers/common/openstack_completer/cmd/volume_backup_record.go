package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var volume_backup_recordCmd = &cobra.Command{
	Use:   "record",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_backup_recordCmd).Standalone()

	volume_backupCmd.AddCommand(volume_backup_recordCmd)
}
