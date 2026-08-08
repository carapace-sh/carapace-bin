package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/carapace-sh/carapace-jq/pkg/actions/tools/jq"
	"github.com/spf13/cobra"
)

var milestone_getCmd = &cobra.Command{
	Use:   "get [<id>] [flags]",
	Short: "Get a milestone by ID in a project or group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(milestone_getCmd).Standalone()

	milestone_getCmd.Flags().String("group", "", "The ID or URL-encoded path of the group.")
	milestone_getCmd.Flags().String("jq", "", "Filter JSON output with a jq expression.")
	milestone_getCmd.Flags().StringP("output", "F", "", "Format output as: text, json.")
	milestone_getCmd.Flags().String("project", "", "The ID or URL-encoded path of the project.")
	milestoneCmd.AddCommand(milestone_getCmd)

	carapace.Gen(milestone_getCmd).FlagCompletion(carapace.ActionMap{
		"group":  action.ActionGroups(milestone_getCmd),
		"jq":     jq.ActionFilters(),
		"output": carapace.ActionValues("text", "json"),
	})
}
