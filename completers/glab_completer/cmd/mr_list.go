package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var mr_listCmd = &cobra.Command{
	Use:     "list [flags]",
	Short:   "List merge requests",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mr_listCmd).Standalone()

	mr_listCmd.Flags().BoolP("all", "A", false, "Get all merge requests")
	mr_listCmd.Flags().StringSliceP("assignee", "a", []string{}, "Get only merge requests assigned to users")
	mr_listCmd.Flags().String("author", "", "Filter merge request by Author <username>")
	mr_listCmd.Flags().BoolP("closed", "c", false, "Get only closed merge requests")
	mr_listCmd.Flags().BoolP("draft", "d", false, "Filter by draft merge requests")
	mr_listCmd.Flags().StringP("group", "g", "", "Get MRs from group and it's subgroups")
	mr_listCmd.Flags().StringSliceP("label", "l", []string{}, "Filter merge request by label <name>")
	mr_listCmd.Flags().BoolP("merged", "M", false, "Get only merged merge requests")
	mr_listCmd.Flags().StringP("milestone", "m", "", "Filter merge request by milestone <id>")
	mr_listCmd.Flags().StringSlice("not-label", []string{}, "Filter merge requests by not having label <name>")
	mr_listCmd.Flags().IntP("page", "p", 1, "Page number")
	mr_listCmd.Flags().IntP("per-page", "P", 30, "Number of items to list per page")
	mr_listCmd.Flags().StringSliceP("reviewer", "r", []string{}, "Get only merge requests with users as reviewer")
	mr_listCmd.Flags().String("search", "", "Filter by <string> in title and description")
	mr_listCmd.Flags().StringP("source-branch", "s", "", "Filter by source branch <name>")
	mr_listCmd.Flags().StringP("target-branch", "t", "", "Filter by target branch <name>")
	mrCmd.AddCommand(mr_listCmd)

	carapace.Gen(mr_listCmd).FlagCompletion(carapace.ActionMap{
		"assignee":      action.ActionProjectMembers(mr_listCmd).UniqueList(","),
		"author":        action.ActionUsers(mr_listCmd),
		"group":         action.ActionGroups(mr_listCmd),
		"label":         action.ActionLabels(mr_listCmd).UniqueList(","),
		"milestone":     action.ActionMilestones(mr_listCmd),
		"not-label":     action.ActionLabels(mr_listCmd).UniqueList(","),
		"reviewer":      action.ActionProjectMembers(mr_listCmd).UniqueList(","),
		"source-branch": action.ActionBranches(mr_listCmd),
		"target-branch": action.ActionBranches(mr_listCmd),
	})
}
