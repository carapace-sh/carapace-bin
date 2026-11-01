package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var drive_renameCmd = &cobra.Command{
	Use:   "rename",
	Short: "[ALPHA] Rename a share",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(drive_renameCmd).Standalone()

	driveCmd.AddCommand(drive_renameCmd)
}
