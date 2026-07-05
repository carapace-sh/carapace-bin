package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var reclaimCmd = &cobra.Command{
	Use:   "reclaim",
	Short: "Reclaim disk space by uninstalling inactive ports",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reclaimCmd).Standalone()
	reclaimCmd.Flags().Bool("disable-reminders", false, "Disable regular reclaim reminders")
	reclaimCmd.Flags().Bool("enable-reminders", false, "Enable regular reclaim reminders")
	rootCmd.AddCommand(reclaimCmd)
}
