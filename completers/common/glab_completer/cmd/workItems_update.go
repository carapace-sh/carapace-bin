package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/carapace-sh/carapace-jq/pkg/actions/tools/jq"
	"github.com/spf13/cobra"
)

var workItems_updateCmd = &cobra.Command{
	Use:   "update <iid> [flags]",
	Short: "Update work items in a project or group. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workItems_updateCmd).Standalone()

	workItems_updateCmd.Flags().StringSliceP("assignee", "a", nil, "Update the work item assignee with the supplied GitLab usernames.")
	workItems_updateCmd.Flags().StringSlice("attach", nil, "(EXPERIMENTAL) Upload a file and reference it at the end of the description. Use \"-\" to read the file from standard input. Repeat the flag to attach multiple files.")
	workItems_updateCmd.Flags().StringP("description", "d", "", "Update the description for the work item.")
	workItems_updateCmd.Flags().String("description-file", "", "Read the work item description from a file. Use \"-\" to read from standard input.")
	workItems_updateCmd.Flags().String("duedate", "", "Update the due date for the work item.")
	workItems_updateCmd.Flags().StringP("group", "g", "", "Update work items for a group or subgroup.")
	workItems_updateCmd.Flags().String("jq", "", "Filter JSON output with a jq expression.")
	workItems_updateCmd.Flags().StringP("milestone", "m", "", "Update the work item milestone with the title or ID.")
	workItems_updateCmd.Flags().StringP("output", "F", "text", "Format output as: text, json.")
	workItems_updateCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository. You can use either OWNER/REPO or GROUP/NAMESPACE/REPO. The full URL or Git URL is also accepted.")
	workItems_updateCmd.Flags().String("startdate", "", "Update the start date for the work item.")
	workItems_updateCmd.Flags().StringP("title", "t", "", "Update the title for the work item.")
	workItems_updateCmd.Flags().StringP("weight", "w", "", "Update the weight value for the work item.")
	workItemsCmd.AddCommand(workItems_updateCmd)

	carapace.Gen(workItems_updateCmd).FlagCompletion(carapace.ActionMap{
		"assignee":         action.ActionProjectMembers(workItems_updateCmd).UniqueList(","),
		"attach":           carapace.ActionFiles(),
		"description-file": carapace.ActionFiles(),
		"group":            action.ActionGroups(workItems_updateCmd),
		"jq":               jq.ActionFilters(),
		"milestone":        action.ActionMilestones(workItems_updateCmd),
		"output":           carapace.ActionValues("text", "json"),
		"repo":             action.ActionRepo(workItems_updateCmd),
	})

	carapace.Gen(workItems_updateCmd).PositionalCompletion(
		action.ActionIssues(workItems_updateCmd, "opened"),
	)
}
