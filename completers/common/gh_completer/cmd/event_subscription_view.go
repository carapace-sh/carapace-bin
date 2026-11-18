package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var event_subscription_viewCmd = &cobra.Command{
	Use:   "view <id>",
	Short: "View an event subscription (preview)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(event_subscription_viewCmd).Standalone()

	event_subscription_viewCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	event_subscription_viewCmd.Flags().StringSlice("json", nil, "Output JSON with the specified `fields`")
	event_subscription_viewCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template; see \"gh help formatting\"")
	event_subscriptionCmd.AddCommand(event_subscription_viewCmd)

	carapace.Gen(event_subscription_viewCmd).FlagCompletion(carapace.ActionMap{
		"json": action.ActionEventSubscriptionFields().UniqueList(","),
	})
}
