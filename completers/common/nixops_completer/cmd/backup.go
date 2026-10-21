package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "make snapshots of persistent disks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backupCmd).Standalone()

	backupCmd.Flags().StringSlice("devices", nil, "Only backup the specified devices")
	backupCmd.Flags().StringSlice("exclude", nil, "Do not perform backup actions on the specified machines")
	backupCmd.Flags().Bool("force", false, "Start new backup even if previous is still running")
	backupCmd.Flags().Bool("freeze", false, "Freeze filesystems for non-root filesystems that support this")
	backupCmd.Flags().StringSlice("include", nil, "Perform backup actions on the specified machines only")

	rootCmd.AddCommand(backupCmd)
}
