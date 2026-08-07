package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
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
	workItems_updateCmd.Flags().StringP("description", "d", "", "Update the description for the work item.")
	workItems_updateCmd.Flags().String("duedate", "", "Update the due date for the work item.")
	workItems_updateCmd.Flags().StringP("group", "g", "", "Update work items for a group or subgroup.")
	workItems_updateCmd.Flags().String("jq", "", "Filter JSON output with a jq expression.")
	workItems_updateCmd.Flags().StringP("milestone", "m", "", "Update the work item milestone with the title or ID.")
	workItems_updateCmd.Flags().StringP("output", "F", "", "Format output as: text, json.")
	workItems_updateCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository. Can use either `OWNER/REPO` or `GROUP/NAMESPACE/REPO` format. Also accepts full URL or Git URL.")
	workItems_updateCmd.Flags().String("startdate", "", "Update the start date for the work item.")
	workItems_updateCmd.Flags().StringP("title", "t", "", "Update the title for the work item.")
	workItems_updateCmd.Flags().StringP("weight", "w", "", "Update the weight value for the work item.")
	workItemsCmd.AddCommand(workItems_updateCmd)

	carapace.Gen(workItems_updateCmd).FlagCompletion(carapace.ActionMap{
		"repo": action.ActionRepo(workItems_updateCmd),
	})
}
