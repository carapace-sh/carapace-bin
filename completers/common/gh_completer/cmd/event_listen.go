package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var event_listenCmd = &cobra.Command{
	Use:   "listen",
	Short: "Wait for events on a specific resource",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(event_listenCmd).Standalone()

	event_listenCmd.Flags().String("count", "", "Exit after this many events (default 1)")
	event_listenCmd.Flags().StringSlice("event", nil, "Event `type` to listen to; repeat or separate with commas for more than one")
	event_listenCmd.Flags().Bool("follow", false, "Keep listening for events until stopped")
	event_listenCmd.Flags().StringP("jq", "q", "", "Filter each event using a jq `expression`")
	event_listenCmd.Flags().String("resource", "", "URL of the pull request to listen to")
	event_listenCmd.Flags().String("subscription", "", "ID of an existing subscription to listen to")
	event_listenCmd.Flags().StringP("template", "t", "", "Format each event using a Go template; see \"gh help formatting\"")
	event_listenCmd.Flags().String("timeout", "0", "Exit with status 8 after this long, in seconds or as a duration such as 30s or 10m")
	eventCmd.AddCommand(event_listenCmd)

	carapace.Gen(event_listenCmd).FlagCompletion(carapace.ActionMap{
		"event": carapace.ActionValues(
			"com.github.pull_request.check_suite.status.changed.v1",
			"com.github.pull_request.head.changed.v1",
			"com.github.pull_request.merge_conflict.v1",
			"com.github.pull_request.review.submitted.v1",
			"com.github.pull_request.state.changed.v1",
		).UniqueList(","),
	})
}
