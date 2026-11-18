package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var event_subscriptionCmd = &cobra.Command{
	Use:   "subscription <command>",
	Short: "Manage event subscriptions (preview)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(event_subscriptionCmd).Standalone()

	eventCmd.AddCommand(event_subscriptionCmd)
}
