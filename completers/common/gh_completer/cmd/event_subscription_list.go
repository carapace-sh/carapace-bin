package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var event_subscription_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List event subscriptions (preview)",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(event_subscription_listCmd).Standalone()

	event_subscription_listCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	event_subscription_listCmd.Flags().StringSlice("json", nil, "Output JSON with the specified `fields`")
	event_subscription_listCmd.Flags().StringP("limit", "L", "30", "Maximum number of subscriptions to list")
	event_subscription_listCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template; see \"gh help formatting\"")
	event_subscriptionCmd.AddCommand(event_subscription_listCmd)

	carapace.Gen(event_subscription_listCmd).FlagCompletion(carapace.ActionMap{
		"json": action.ActionEventSubscriptionFields().UniqueList(","),
	})
}
