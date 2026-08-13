package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/vercel_completer/cmd/action"
	"github.com/spf13/cobra"
)

var activityCmd = &cobra.Command{
	Use:   "activity",
	Short: "List user activity events",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(activityCmd).Standalone()

	activityCmd.Flags().BoolP("all", "a", false, "Show all team events")
	activityCmd.Flags().StringP("format", "F", "", "Output format")
	activityCmd.Flags().Bool("json", false, "Output as JSON")
	activityCmd.Flags().String("limit", "", "Number of results per page")
	activityCmd.Flags().StringP("next", "N", "", "Show next page of results")
	activityCmd.Flags().StringP("project", "p", "", "Filter by project")
	activityCmd.Flags().String("since", "", "Show events after this date")
	activityCmd.Flags().String("type", "", "Filter by event type")
	activityCmd.Flags().String("until", "", "Show events before this date")

	rootCmd.AddCommand(activityCmd)

	carapace.Gen(activityCmd).FlagCompletion(carapace.ActionMap{
		"format":  carapace.ActionValues("plain", "json"),
		"project": action.ActionProjects(activityCmd),
		"type":    carapace.ActionValues("deployment", "project", "domain", "dns", "cert", "env"),
	})
}
