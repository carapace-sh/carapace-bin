package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/carapace-sh/carapace-jq/pkg/actions/tools/jq"
	"github.com/spf13/cobra"
)

var runner_listCmd = &cobra.Command{
	Use:     "list [flags]",
	Short:   "List runners.",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runner_listCmd).Standalone()

	runner_listCmd.Flags().StringP("group", "g", "", "List runners for a group. Ignored if -R/--repo is set.")
	runner_listCmd.Flags().BoolP("instance", "i", false, "List all runners available to the user (instance scope).")
	runner_listCmd.Flags().String("jq", "", "Filter JSON output with a jq expression.")
	runner_listCmd.Flags().StringP("output", "F", "text", "Format output as: text, json.")
	runner_listCmd.Flags().StringP("page", "p", "1", "Page number.")
	runner_listCmd.Flags().StringP("per-page", "P", "30", "Number of items to list per page.")
	runner_listCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository. You can use either OWNER/REPO or GROUP/NAMESPACE/REPO. The full URL or Git URL is also accepted.")
	runnerCmd.AddCommand(runner_listCmd)

	carapace.Gen(runner_listCmd).FlagCompletion(carapace.ActionMap{
		"group":  action.ActionGroups(runner_listCmd),
		"jq":     jq.ActionFilters(),
		"output": carapace.ActionValues("text", "json"),
		"repo":   action.ActionRepo(runner_listCmd),
	})
}
