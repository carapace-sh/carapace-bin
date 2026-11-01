package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backupStatusCmd = &cobra.Command{
	Use:   "backup-status",
	Short: "get status of backups",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backupStatusCmd).Standalone()

	backupStatusCmd.Flags().StringSlice("exclude", nil, "Do not perform backup actions on the specified machines")
	backupStatusCmd.Flags().StringSlice("include", nil, "Perform backup actions on the specified machines only")
	backupStatusCmd.Flags().Bool("latest", false, "Show status of latest backup only")
	backupStatusCmd.Flags().Bool("wait", false, "Wait until backup is finished")

	rootCmd.AddCommand(backupStatusCmd)

	carapace.Gen(backupStatusCmd).PositionalCompletion(
		carapace.ActionValues(), // backup id
	)
}
