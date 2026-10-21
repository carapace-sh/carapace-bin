package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var removeBackupCmd = &cobra.Command{
	Use:   "remove-backup",
	Short: "remove a given backup",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(removeBackupCmd).Standalone()

	removeBackupCmd.Flags().Bool("keep-physical", false, "Only remove backups from nixops state")

	rootCmd.AddCommand(removeBackupCmd)

	carapace.Gen(removeBackupCmd).PositionalCompletion(
		carapace.ActionValues(), // backup id
	)
}
