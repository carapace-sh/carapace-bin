package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var runner_assignCmd = &cobra.Command{
	Use:   "assign <runner-id> [flags]",
	Short: "Assign a runner to a project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runner_assignCmd).Standalone()

	runner_assignCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository. Can use either `OWNER/REPO` or `GROUP/NAMESPACE/REPO` format. Also accepts full URL or Git URL.")
	runnerCmd.AddCommand(runner_assignCmd)

	carapace.Gen(runner_assignCmd).FlagCompletion(carapace.ActionMap{
		"repo": action.ActionRepo(runner_assignCmd),
	})
}
