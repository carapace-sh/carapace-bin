package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
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
	runner_listCmd.Flags().StringP("output", "F", "", "Format output as: text, json.")
	runner_listCmd.Flags().StringP("page", "p", "", "Page number.")
	runner_listCmd.Flags().StringP("per-page", "P", "", "Number of items to list per page.")
	runner_listCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository. Can use either `OWNER/REPO` or `GROUP/NAMESPACE/REPO` format. Also accepts full URL or Git URL.")
	runnerCmd.AddCommand(runner_listCmd)

	carapace.Gen(runner_listCmd).FlagCompletion(carapace.ActionMap{
		"group":  action.ActionGroups(runner_listCmd),
		"output": carapace.ActionValues("text", "json"),
		"repo":   action.ActionRepo(runner_listCmd),
	})
}
