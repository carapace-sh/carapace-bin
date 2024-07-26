package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storageUnlinkCmd = &cobra.Command{
	Use:   "storage:unlink",
	Short: "Delete existing symbolic links configured for the application",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storageUnlinkCmd).Standalone()

	rootCmd.AddCommand(storageUnlinkCmd)
}
