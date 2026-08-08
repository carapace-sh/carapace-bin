package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-jq/pkg/actions/tools/jq"
	"github.com/spf13/cobra"
)

var schedule_listCmd = &cobra.Command{
	Use:   "list [flags]",
	Short: "List pipeline schedules in a project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(schedule_listCmd).Standalone()

	schedule_listCmd.Flags().String("jq", "", "Filter JSON output with a jq expression.")
	schedule_listCmd.Flags().StringP("output", "F", "", "Format output as: text, json.")
	schedule_listCmd.Flags().StringP("page", "p", "", "Page number.")
	schedule_listCmd.Flags().StringP("per-page", "P", "", "Number of items to list per page.")
	scheduleCmd.AddCommand(schedule_listCmd)

	carapace.Gen(schedule_listCmd).FlagCompletion(carapace.ActionMap{
		"jq":     jq.ActionFilters(),
		"output": carapace.ActionValues("text", "json"),
	})
}
