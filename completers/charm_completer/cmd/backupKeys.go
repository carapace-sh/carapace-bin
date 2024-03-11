package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backupKeysCmd = &cobra.Command{
	Use:   "backup-keys",
	Short: "Backup your Charm account keys",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backupKeysCmd).Standalone()

	backupKeysCmd.Flags().StringP("output", "o", "", "keys backup filepath")
	rootCmd.AddCommand(backupKeysCmd)

	carapace.Gen(backupKeysCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionFiles(),
	})
}
