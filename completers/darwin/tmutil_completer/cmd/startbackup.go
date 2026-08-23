package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var startbackupCmd = &cobra.Command{
	Use:   "startbackup",
	Short: "begin a backup",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(startbackupCmd).Standalone()
	rootCmd.AddCommand(startbackupCmd)

	startbackupCmd.Flags().Bool("auto", false, "Run in a mode similar to system-scheduled backups")
	startbackupCmd.Flags().BoolP("block", "b", false, "Wait until the backup is finished")
	startbackupCmd.Flags().StringP("destination", "d", "", "Perform the backup to the specified destination ID")
	startbackupCmd.Flags().BoolP("rotation", "r", false, "Allow automatic destination rotation")
}