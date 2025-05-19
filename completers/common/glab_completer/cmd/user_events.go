package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var user_eventsCmd = &cobra.Command{
	Use:   "events",
	Short: "View user events.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(user_eventsCmd).Standalone()

	user_eventsCmd.Flags().BoolP("all", "a", false, "Get events from all projects.")
	user_eventsCmd.Flags().StringP("output", "F", "", "Format output as: 'text', 'json'.")
	user_eventsCmd.Flags().StringP("page", "p", "", "Page number.")
	user_eventsCmd.Flags().StringP("per-page", "P", "", "Number of items to list per page.")
	userCmd.AddCommand(user_eventsCmd)

	carapace.Gen(user_eventsCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionValues("text", "json"),
	})
}
