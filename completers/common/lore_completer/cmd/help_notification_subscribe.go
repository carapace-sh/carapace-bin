package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_notification_subscribeCmd = &cobra.Command{
	Use:   "subscribe",
	Short: "Subscribe to events on the given repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_notification_subscribeCmd).Standalone()

	help_notificationCmd.AddCommand(help_notification_subscribeCmd)
}
