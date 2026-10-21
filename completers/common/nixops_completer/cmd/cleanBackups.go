package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanBackupsCmd = &cobra.Command{
	Use:   "clean-backups",
	Short: "clean up old backups",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanBackupsCmd).Standalone()

	cleanBackupsCmd.Flags().Int("keep", 0, "Number of backups to keep around")
	cleanBackupsCmd.Flags().Int("keep-days", 0, "Keep backups newer than N days")
	cleanBackupsCmd.Flags().Bool("keep-physical", false, "Only remove backups from nixops state")

	rootCmd.AddCommand(cleanBackupsCmd)
}
