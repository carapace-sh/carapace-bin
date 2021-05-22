package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var run_viewCmd = &cobra.Command{
	Use:   "view [<run-id>]",
	Short: "View a summary of a workflow run",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	run_viewCmd.Flags().Bool("exit-status", false, "Exit with non-zero status if run failed")
	run_viewCmd.Flags().StringP("job", "j", "", "View a specific job ID from a run")
	run_viewCmd.Flags().Bool("log", false, "View full log for either a run or specific job")
	run_viewCmd.Flags().Bool("log-failed", false, "View the log for any failed steps in a run or specific job")
	run_viewCmd.Flags().BoolP("verbose", "v", false, "Show job steps")
	run_viewCmd.Flags().BoolP("web", "w", false, "Open run in the browser")
	runCmd.AddCommand(run_viewCmd)

	carapace.Gen(run_viewCmd).FlagCompletion(carapace.ActionMap{
		"job": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 0 {
				return action.ActionWorkflowJobs(run_viewCmd, c.Args[0], action.RunOpts{All: true})
			}
			return carapace.ActionValues()
		}),
	})

	carapace.Gen(run_viewCmd).PositionalCompletion(
		action.ActionWorkflowRuns(run_viewCmd, action.RunOpts{All: true}),
	)
}
