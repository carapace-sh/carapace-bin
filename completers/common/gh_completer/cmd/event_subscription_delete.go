package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var event_subscription_deleteCmd = &cobra.Command{
	Use:   "delete <id>",
	Short: "Delete an event subscription (preview)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(event_subscription_deleteCmd).Standalone()

	event_subscriptionCmd.AddCommand(event_subscription_deleteCmd)
}
