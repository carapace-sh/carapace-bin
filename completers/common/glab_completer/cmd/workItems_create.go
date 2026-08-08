package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/carapace-sh/carapace-jq/pkg/actions/tools/jq"
	"github.com/spf13/cobra"
)

var workItems_createCmd = &cobra.Command{
	Use:   "create [flags]",
	Short: "Create work items in a project or group. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workItems_createCmd).Standalone()

	workItems_createCmd.Flags().BoolP("confidential", "c", false, "Mark work item confidential.")
	workItems_createCmd.Flags().StringP("description", "d", "", "Description of the work item. Set to \"-\" to open an editor.")
	workItems_createCmd.Flags().StringP("group", "g", "", "Create work items for a group or subgroup.")
	workItems_createCmd.Flags().String("jq", "", "Filter JSON output with a jq expression.")
	workItems_createCmd.Flags().StringP("output", "F", "", "Format output as: text, json.")
	workItems_createCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository. Can use either `OWNER/REPO` or `GROUP/NAMESPACE/REPO` format. Also accepts full URL or Git URL.")
	workItems_createCmd.Flags().StringP("title", "t", "", "Add a title for the work item.")
	workItems_createCmd.Flags().StringP("type", "T", "", "Type of work item (epic, incident, issue, key_result, objective, requirement, task, test_case, ticket).")
	workItems_createCmd.MarkFlagRequired("type")
	workItemsCmd.AddCommand(workItems_createCmd)

	carapace.Gen(workItems_createCmd).FlagCompletion(carapace.ActionMap{
		"jq":     jq.ActionFilters(),
		"output": carapace.ActionValues("text", "json"),
		"repo":   action.ActionRepo(workItems_createCmd),
	})
}
