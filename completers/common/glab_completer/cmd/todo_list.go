package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-jq/pkg/actions/tools/jq"
	"github.com/spf13/cobra"
)

var todo_listCmd = &cobra.Command{
	Use:     "list [flags]",
	Short:   "List your to-do items.",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(todo_listCmd).Standalone()

	todo_listCmd.Flags().StringP("action", "a", "", "Filter by action: assigned, mentioned, build_failed, marked, approval_required, directly_addressed.")
	todo_listCmd.Flags().String("jq", "", "Filter JSON output with a jq expression.")
	todo_listCmd.Flags().StringP("output", "F", "", "Format output as: text, json.")
	todo_listCmd.Flags().StringP("page", "p", "", "Page number.")
	todo_listCmd.Flags().StringP("per-page", "P", "", "Number of items to list per page.")
	todo_listCmd.Flags().StringP("state", "s", "", "Filter by state: pending, done, all.")
	todo_listCmd.Flags().StringP("type", "t", "", "Filter by target type: Issue, MergeRequest.")
	todoCmd.AddCommand(todo_listCmd)

	carapace.Gen(todo_listCmd).FlagCompletion(carapace.ActionMap{
		"action": carapace.ActionValues("assigned", "mentioned", "build_failed", "marked", "approval_required", "directly_addressed"),
		"jq":     jq.ActionFilters(),
		"output": carapace.ActionValues("text", "json"),
		"state":  carapace.ActionValues("pending", "done", "all"),
		"type":   carapace.ActionValues("Issue", "MergeRequest"),
	})
}
