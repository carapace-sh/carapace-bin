package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var milestone_deleteCmd = &cobra.Command{
	Use:   "delete <id> [flags]",
	Short: "Delete a milestone from a project or group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(milestone_deleteCmd).Standalone()

	milestone_deleteCmd.Flags().String("group", "", "The ID or URL-encoded path of the group.")
	milestone_deleteCmd.Flags().String("project", "", "The ID or URL-encoded path of the project.")
	milestoneCmd.AddCommand(milestone_deleteCmd)

	carapace.Gen(milestone_deleteCmd).FlagCompletion(carapace.ActionMap{
		"group": action.ActionGroups(milestone_deleteCmd),
	})
}
