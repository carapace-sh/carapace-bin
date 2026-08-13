package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var alertsCmd = &cobra.Command{
	Use:   "alerts",
	Short: "List alert groups, inspect a group, or manage alert rules",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(alertsCmd).Standalone()

	alertsCmd.Flags().Bool("ai", false, "Print AI-focused sections")
	alertsCmd.Flags().BoolP("all", "a", false, "Show team-wide alerts")
	alertsCmd.Flags().StringP("format", "F", "", "Output format")
	alertsCmd.Flags().Bool("json", false, "Output as JSON")
	alertsCmd.Flags().String("limit", "", "Number of results")
	alertsCmd.Flags().StringP("project", "p", "", "Filter by project")
	alertsCmd.Flags().String("since", "", "Start of time range (ISO-8601)")
	alertsCmd.Flags().String("type", "", "Filter by alert type")
	alertsCmd.Flags().String("until", "", "End of time range (ISO-8601)")

	rootCmd.AddCommand(alertsCmd)
}
