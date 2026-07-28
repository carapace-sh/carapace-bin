package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var drive_unshareCmd = &cobra.Command{
	Use:   "unshare",
	Short: "[ALPHA] Remove a share",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(drive_unshareCmd).Standalone()

	driveCmd.AddCommand(drive_unshareCmd)
}
