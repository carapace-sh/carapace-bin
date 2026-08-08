package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var runner_unassignCmd = &cobra.Command{
	Use:   "unassign <runner-id> [flags]",
	Short: "Unassign a runner from a project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runner_unassignCmd).Standalone()

	runner_unassignCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository. Can use either `OWNER/REPO` or `GROUP/NAMESPACE/REPO` format. Also accepts full URL or Git URL.")
	runnerCmd.AddCommand(runner_unassignCmd)

	carapace.Gen(runner_unassignCmd).FlagCompletion(carapace.ActionMap{
		"repo": action.ActionRepo(runner_unassignCmd),
	})
}
