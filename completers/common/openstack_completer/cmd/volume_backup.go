package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var volume_backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_backupCmd).Standalone()

	volumeCmd.AddCommand(volume_backupCmd)
}
