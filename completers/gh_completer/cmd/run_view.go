package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var run_viewCmd = &cobra.Command{
	Use:   "view [<run-id>]",
	Short: "View a summary of a workflow run",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(run_viewCmd).Standalone()

	run_viewCmd.Flags().StringP("attempt", "a", "", "The attempt number of the workflow run")
	run_viewCmd.Flags().Bool("exit-status", false, "Exit with non-zero status if run failed")
	run_viewCmd.Flags().StringP("job", "j", "", "View a specific job ID from a run")
	run_viewCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	run_viewCmd.Flags().StringSlice("json", nil, "Output JSON with the specified `fields`")
	run_viewCmd.Flags().Bool("log", false, "View full log for either a run or specific job")
	run_viewCmd.Flags().Bool("log-failed", false, "View the log for any failed steps in a run or specific job")
	run_viewCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template; see \"gh help formatting\"")
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
		"json": carapace.Batch(
			action.ActionRunFields(),
			carapace.ActionValues("jobs"),
		).ToA().UniqueList(","),
	})

	carapace.Gen(run_viewCmd).PositionalCompletion(
		action.ActionWorkflowRuns(run_viewCmd, action.RunOpts{All: true}),
	)
}
