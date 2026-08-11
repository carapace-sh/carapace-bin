package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var notification_subscribeCmd = &cobra.Command{
	Use:   "subscribe",
	Short: "Subscribe to events on the given repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notification_subscribeCmd).Standalone()

	notification_subscribeCmd.Flags().BoolP("help", "h", false, "Print help")
	notificationCmd.AddCommand(notification_subscribeCmd)

	carapace.Gen(notification_subscribeCmd).PositionalCompletion(
		carapace.ActionValues(),
	)
}
