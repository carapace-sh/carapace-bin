package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var milestone_editCmd = &cobra.Command{
	Use:   "edit <id> [flags]",
	Short: "Edit a milestone in a project or group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(milestone_editCmd).Standalone()

	milestone_editCmd.Flags().String("description", "", "Description of the milestone.")
	milestone_editCmd.Flags().String("due-date", "", "Due date for the milestone. Expected in ISO 8601 format (2025-04-15T08:00:00Z).")
	milestone_editCmd.Flags().String("group", "", "The ID or URL-encoded path of the group.")
	milestone_editCmd.Flags().String("project", "", "The ID or URL-encoded path of the project.")
	milestone_editCmd.Flags().String("start-date", "", "Start date for the milestone. Expected in ISO 8601 format (2025-04-15T08:00:00Z).")
	milestone_editCmd.Flags().String("state", "", "State of the milestone. Can be 'activate' or 'close'.")
	milestone_editCmd.Flags().String("title", "", "Title of the milestone.")
	milestoneCmd.AddCommand(milestone_editCmd)

	carapace.Gen(milestone_editCmd).FlagCompletion(carapace.ActionMap{
		"group": action.ActionGroups(milestone_editCmd),
		"state": carapace.ActionValues("activate", "close"),
	})
}
