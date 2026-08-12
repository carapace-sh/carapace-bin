package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var notificationCmd = &cobra.Command{
	Use:   "notification",
	Short: "Notifications",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notificationCmd).Standalone()

	notificationCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(notificationCmd)
}
