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

	workItems_createCmd.Flags().StringSlice("attach", nil, "(EXPERIMENTAL) Upload a file and reference it at the end of the description. Use \"-\" to read the file from standard input. Repeat the flag to attach multiple files.")
	workItems_createCmd.Flags().BoolP("confidential", "c", false, "Mark work item confidential.")
	workItems_createCmd.Flags().StringP("description", "d", "", "Description of the work item. Set to \"-\" to open an editor.")
	workItems_createCmd.Flags().String("description-file", "", "Read the work item description from a file. Use \"-\" to read from standard input.")
	workItems_createCmd.Flags().StringP("group", "g", "", "Create work items for a group or subgroup.")
	workItems_createCmd.Flags().String("jq", "", "Filter JSON output with a jq expression.")
	workItems_createCmd.Flags().StringP("output", "F", "text", "Format output as: text, json.")
	workItems_createCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository. You can use either OWNER/REPO or GROUP/NAMESPACE/REPO. The full URL or Git URL is also accepted.")
	workItems_createCmd.Flags().StringP("title", "t", "", "Add a title for the work item.")
	workItems_createCmd.Flags().StringP("type", "T", "", "Type of work item (epic, incident, issue, key_result, objective, requirement, task, test_case, ticket).")
	workItems_createCmd.MarkFlagRequired("type")
	workItemsCmd.AddCommand(workItems_createCmd)

	carapace.Gen(workItems_createCmd).FlagCompletion(carapace.ActionMap{
		"attach":           carapace.ActionFiles(),
		"description-file": carapace.ActionFiles(),
		"group":            action.ActionGroups(workItems_createCmd),
		"jq":               jq.ActionFilters(),
		"output":           carapace.ActionValues("text", "json"),
		"repo":             action.ActionRepo(workItems_createCmd),
	})
}
