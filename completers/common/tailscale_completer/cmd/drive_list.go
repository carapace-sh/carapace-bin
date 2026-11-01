package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var drive_listCmd = &cobra.Command{
	Use:   "list",
	Short: "[ALPHA] List current shares",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(drive_listCmd).Standalone()

	driveCmd.AddCommand(drive_listCmd)
}
