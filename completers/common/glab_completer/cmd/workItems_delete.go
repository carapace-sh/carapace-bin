package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/carapace-sh/carapace-jq/pkg/actions/tools/jq"
	"github.com/spf13/cobra"
)

var workItems_deleteCmd = &cobra.Command{
	Use:   "delete <iid>",
	Short: "Delete a work item in a project or group. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workItems_deleteCmd).Standalone()

	workItems_deleteCmd.Flags().StringP("group", "g", "", "Delete a work items from a group or subgroup.")
	workItems_deleteCmd.Flags().String("jq", "", "Filter JSON output with a jq expression.")
	workItems_deleteCmd.Flags().StringP("output", "F", "text", "Format output as: text, json.")
	workItems_deleteCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository. You can use either OWNER/REPO or GROUP/NAMESPACE/REPO. The full URL or Git URL is also accepted.")
	workItemsCmd.AddCommand(workItems_deleteCmd)

	carapace.Gen(workItems_deleteCmd).FlagCompletion(carapace.ActionMap{
		"group":  action.ActionGroups(workItems_deleteCmd),
		"jq":     jq.ActionFilters(),
		"output": carapace.ActionValues("text", "json"),
		"repo":   action.ActionRepo(workItems_deleteCmd),
	})

	carapace.Gen(workItems_deleteCmd).PositionalCompletion(
		action.ActionIssues(workItems_deleteCmd, "opened"),
	)
}
