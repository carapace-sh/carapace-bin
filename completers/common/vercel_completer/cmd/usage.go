package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var usageCmd = &cobra.Command{
	Use:   "usage",
	Short: "Show billing usage for the current billing period or a custom date range",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(usageCmd).Standalone()

	usageCmd.Flags().String("breakdown", "", "Show usage breakdown (daily, weekly, monthly)")
	usageCmd.Flags().StringP("format", "F", "", "Output format")
	usageCmd.Flags().String("from", "", "Start date (YYYY-MM-DD)")
	usageCmd.Flags().String("group-by", "", "Group usage (project, region)")
	usageCmd.Flags().Bool("json", false, "Output as JSON")
	usageCmd.Flags().String("to", "", "End date (YYYY-MM-DD)")

	rootCmd.AddCommand(usageCmd)

	carapace.Gen(usageCmd).FlagCompletion(carapace.ActionMap{
		"breakdown": carapace.ActionValues("daily", "weekly", "monthly"),
		"group-by":  carapace.ActionValues("project", "region"),
	})
}
