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
	rootCmd.AddCommand(cleanBackupsCmd)
}
