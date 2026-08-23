package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var unlistenCmd = &cobra.Command{
	Use:   "unlisten",
	Short: "Unlisten to notifications",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unlistenCmd).Standalone()
	unlistenCmd.Flags().Bool("all", false, "Unlisten to all notifications")
	unlistenCmd.Flags().String("id", "", "Notification ID")
	rootCmd.AddCommand(unlistenCmd)
}
