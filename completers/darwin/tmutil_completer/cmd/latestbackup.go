package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(latestbackupCmd)
}

var latestbackupCmd = &cobra.Command{
	Use:   "latestbackup",
	Short: "list this computer's latest completed backup",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(latestbackupCmd).Standalone()

	latestbackupCmd.Flags().StringP("destination", "d", "", "Destination volume to list backups from")
	latestbackupCmd.Flags().BoolP("mount", "m", false, "Attempt to mount backups")
	latestbackupCmd.Flags().BoolP("timestamp-only", "t", false, "Show only the backup timestamp")
}