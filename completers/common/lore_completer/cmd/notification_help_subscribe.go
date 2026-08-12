package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var notification_help_subscribeCmd = &cobra.Command{
	Use:   "subscribe",
	Short: "Subscribe to events on the given repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notification_help_subscribeCmd).Standalone()

	notification_helpCmd.AddCommand(notification_help_subscribeCmd)
}
