package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/vercel_completer/cmd/action"
	"github.com/spf13/cobra"
)

var logsCmd = &cobra.Command{
	Use:     "logs",
	Aliases: []string{"log"},
	Short:   "Display request logs for a project",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logsCmd).Standalone()

	logsCmd.Flags().StringP("branch", "b", "", "Filter by Git branch")
	logsCmd.Flags().String("deployment", "", "Filter to a specific deployment ID or URL")
	logsCmd.Flags().String("environment", "", "Filter by environment (production or preview)")
	logsCmd.Flags().BoolP("expand", "x", false, "Show full log message")
	logsCmd.Flags().BoolP("follow", "f", false, "Stream live runtime logs")
	logsCmd.Flags().BoolP("json", "j", false, "Output as JSON Lines")
	logsCmd.Flags().String("level", "", "Filter by log level (error, warning, info, fatal)")
	logsCmd.Flags().StringP("limit", "n", "", "Maximum number of results (default: 100)")
	logsCmd.Flags().String("project", "", "Project name or ID")
	logsCmd.Flags().StringP("query", "q", "", "Advanced search query")
	logsCmd.Flags().String("request-id", "", "Filter by request ID")
	logsCmd.Flags().String("since", "", "Start time (ISO format or relative)")
	logsCmd.Flags().String("source", "", "Filter by source (serverless, edge-function, etc.)")
	logsCmd.Flags().String("status-code", "", "Filter by HTTP status code")
	logsCmd.Flags().String("until", "", "End time (ISO format or relative)")

	rootCmd.AddCommand(logsCmd)

	carapace.Gen(logsCmd).FlagCompletion(carapace.ActionMap{
		"environment": action.ActionEnvironments(),
		"level":       carapace.ActionValues("error", "warning", "info", "fatal"),
		"project":     action.ActionProjects(logsCmd),
		"source":      carapace.ActionValues("serverless", "edge-function", "static"),
	})

	carapace.Gen(logsCmd).PositionalCompletion(
		action.ActionDeployments(logsCmd),
	)
}
