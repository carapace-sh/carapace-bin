package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(deleteCmd)
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete backups with the specified timestamp",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deleteCmd).Standalone()

	deleteCmd.Flags().StringP("destination", "d", "", "Backup mount point")
	deleteCmd.Flags().StringP("path", "p", "", "Specific path to delete")
	deleteCmd.Flags().StringP("timestamp", "t", "", "Timestamp of backup to delete")
}