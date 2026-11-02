package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Backup your logs database to this file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_backupCmd).Standalone()

	logsCmd.AddCommand(logs_backupCmd)
}
