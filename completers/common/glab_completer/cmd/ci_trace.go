package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var ci_traceCmd = &cobra.Command{
	Use:   "trace [<job-id>] [flags]",
	Short: "Trace a CI/CD job log in real time.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ci_traceCmd).Standalone()

	ci_traceCmd.Flags().StringP("branch", "b", "", "The branch to search for the job. (default current branch)")
	ci_traceCmd.Flags().StringP("pipeline-id", "p", "", "The pipeline ID to search for the job.")
	ciCmd.AddCommand(ci_traceCmd)

	carapace.Gen(ci_statusCmd).FlagCompletion(carapace.ActionMap{
		"branch": action.ActionBranches(ci_statusCmd),
	})

	// TODO complete job ids
}
