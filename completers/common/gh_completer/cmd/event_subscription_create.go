package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var event_subscription_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an event subscription (preview)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(event_subscription_createCmd).Standalone()

	event_subscription_createCmd.Flags().StringSlice("event", nil, "Event `type` to subscribe to; repeat or separate with commas for more than one")
	event_subscription_createCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	event_subscription_createCmd.Flags().StringSlice("json", nil, "Output JSON with the specified `fields`")
	event_subscription_createCmd.Flags().String("resource", "", "URL of the resource to subscribe to")
	event_subscription_createCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template; see \"gh help formatting\"")
	event_subscriptionCmd.AddCommand(event_subscription_createCmd)

	carapace.Gen(event_subscription_createCmd).FlagCompletion(carapace.ActionMap{
		"event": carapace.ActionValues(
			"com.github.pull_request.check_suite.status.changed.v1",
			"com.github.pull_request.head.changed.v1",
			"com.github.pull_request.merge_conflict.v1",
			"com.github.pull_request.review.submitted.v1",
			"com.github.pull_request.state.changed.v1",
		).UniqueList(","),
		"json":     action.ActionEventSubscriptionFields().UniqueList(","),
		"resource": action.ActionPullRequestUrls(event_subscription_createCmd, action.PullRequestOpts{Open: true, Closed: true, Merged: true}),
	})
}
