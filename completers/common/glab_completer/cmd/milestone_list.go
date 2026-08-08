package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var milestone_listCmd = &cobra.Command{
	Use:   "list [flags]",
	Short: "List milestones in a project or group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(milestone_listCmd).Standalone()

	milestone_listCmd.Flags().String("group", "", "The ID or URL-encoded path of the group.")
	milestone_listCmd.Flags().Bool("include-ancestors", false, "Include milestones from all parent groups.")
	milestone_listCmd.Flags().String("jq", "", "Filter JSON output with a jq expression.")
	milestone_listCmd.Flags().StringP("output", "F", "", "Format output as: text, json.")
	milestone_listCmd.Flags().StringP("page", "p", "", "Page number.")
	milestone_listCmd.Flags().StringP("per-page", "P", "", "Number of items to list per page.")
	milestone_listCmd.Flags().String("project", "", "The ID or URL-encoded path of the project.")
	milestone_listCmd.Flags().String("search", "", "Return only milestones with a title or description matching the provided string.")
	milestone_listCmd.Flags().Bool("show-id", false, "Show IDs in table output.")
	milestone_listCmd.Flags().String("state", "", "Return only 'active' or 'closed' milestones.")
	milestone_listCmd.Flags().String("title", "", "Return only the milestones having the given title.")
	milestoneCmd.AddCommand(milestone_listCmd)

	carapace.Gen(milestone_listCmd).FlagCompletion(carapace.ActionMap{
		"group":  action.ActionGroups(milestone_listCmd),
		"output": carapace.ActionValues("text", "json"),
		"state":  carapace.ActionValues("active", "closed"),
	})
}
