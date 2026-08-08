package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var milestone_createCmd = &cobra.Command{
	Use:   "create [flags]",
	Short: "Create a milestone in a project or group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(milestone_createCmd).Standalone()

	milestone_createCmd.Flags().String("description", "", "Description of the milestone.")
	milestone_createCmd.Flags().String("due-date", "", "Due date for the milestone. Expected in ISO 8601 format (2025-04-15T08:00:00Z).")
	milestone_createCmd.Flags().String("group", "", "The ID or URL-encoded path of the group.")
	milestone_createCmd.Flags().String("project", "", "The ID or URL-encoded path of the project.")
	milestone_createCmd.Flags().String("start-date", "", "Start date for the milestone. Expected in ISO 8601 format (2025-04-15T08:00:00Z).")
	milestone_createCmd.Flags().String("title", "", "Title of the milestone.")
	milestone_createCmd.MarkFlagRequired("title")
	milestoneCmd.AddCommand(milestone_createCmd)

	carapace.Gen(milestone_createCmd).FlagCompletion(carapace.ActionMap{
		"group": action.ActionGroups(milestone_createCmd),
	})
}
