package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var incident_viewCmd = &cobra.Command{
	Use:     "view <id>",
	Short:   "Display the title, body, and other information about an incident.",
	Aliases: []string{"show"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(incident_viewCmd).Standalone()

	incident_viewCmd.Flags().BoolP("comments", "c", false, "Show incident comments and activities.")
	incident_viewCmd.Flags().StringP("output", "F", "", "Format output as: text, json.")
	incident_viewCmd.Flags().StringP("page", "p", "", "Page number.")
	incident_viewCmd.Flags().StringP("per-page", "P", "", "Number of items to list per page.")
	incident_viewCmd.Flags().BoolP("system-logs", "s", false, "Show system activities and logs.")
	incident_viewCmd.Flags().BoolP("web", "w", false, "Open incident in a browser. Uses the default browser, or the browser specified in the $BROWSER variable.")
	incidentCmd.AddCommand(incident_viewCmd)

	// TODO positional completion
}
