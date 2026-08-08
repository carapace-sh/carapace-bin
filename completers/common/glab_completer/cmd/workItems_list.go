package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/carapace-sh/carapace-jq/pkg/actions/tools/jq"
	"github.com/spf13/cobra"
)

var workItems_listCmd = &cobra.Command{
	Use:     "list [flags]",
	Short:   "List work items in a project or group. (EXPERIMENTAL)",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workItems_listCmd).Standalone()

	workItems_listCmd.Flags().String("after", "", "Fetch items after this cursor (for pagination)")
	workItems_listCmd.Flags().StringP("group", "g", "", "List work items for a group or subgroup.")
	workItems_listCmd.Flags().String("jq", "", "Filter JSON output with a jq expression.")
	workItems_listCmd.Flags().StringP("output", "F", "", "Format output as: text, json.")
	workItems_listCmd.Flags().StringP("per-page", "P", "", "Number of items to list per page (max 100)")
	workItems_listCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository. Can use either `OWNER/REPO` or `GROUP/NAMESPACE/REPO` format. Also accepts full URL or Git URL.")
	workItems_listCmd.Flags().String("state", "", "Filter by state: opened, closed, all.")
	workItems_listCmd.Flags().StringSliceP("type", "t", nil, "Filter by work item type (epic, issue, task, etc.) Multiple types can be comma-separated or specified by repeating the flag.")
	workItemsCmd.AddCommand(workItems_listCmd)

	carapace.Gen(workItems_listCmd).FlagCompletion(carapace.ActionMap{
		"group":  action.ActionGroups(workItems_listCmd),
		"jq":     jq.ActionFilters(),
		"output": carapace.ActionValues("text", "json"),
		"repo":   action.ActionRepo(workItems_listCmd),
		"state":  carapace.ActionValues("opened", "closed", "all"),
	})
}
