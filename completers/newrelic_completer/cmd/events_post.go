package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var events_postCmd = &cobra.Command{
	Use:   "post",
	Short: "Post a custom event to New Relic",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(events_postCmd).Standalone()
	events_postCmd.Flags().StringP("event", "e", "{}", "a JSON-formatted event payload to post")
	eventsCmd.AddCommand(events_postCmd)
}
