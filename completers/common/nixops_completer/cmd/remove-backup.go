package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var RemoveBackupCmd = &cobra.Command{
	Use:   "remove-backup",
	Short: "Remove Backup",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(RemoveBackupCmd).Standalone()
	rootCmd.AddCommand(RemoveBackupCmd)
}
