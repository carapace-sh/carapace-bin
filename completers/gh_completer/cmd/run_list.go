package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/time"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var run_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List recent workflow runs",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(run_listCmd).Standalone()

	run_listCmd.Flags().BoolP("all", "a", false, "Include disabled workflows")
	run_listCmd.Flags().StringP("branch", "b", "", "Filter runs by branch")
	run_listCmd.Flags().StringP("commit", "c", "", "Filter runs by the `SHA` of the commit")
	run_listCmd.Flags().String("created", "", "Filter runs by the `date` it was created")
	run_listCmd.Flags().StringP("event", "e", "", "Filter runs by which `event` triggered the run")
	run_listCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	run_listCmd.Flags().StringSlice("json", []string{}, "Output JSON with the specified `fields`")
	run_listCmd.Flags().StringP("limit", "L", "", "Maximum number of runs to fetch")
	run_listCmd.Flags().StringP("status", "s", "", "Filter runs by status: {queued|completed|in_progress|requested|waiting|pending|action_required|cancelled|failure|neutral|skipped|stale|startup_failure|success|timed_out}")
	run_listCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template; see \"gh help formatting\"")
	run_listCmd.Flags().StringP("user", "u", "", "Filter runs by user who triggered the run")
	run_listCmd.Flags().StringP("workflow", "w", "", "Filter runs by workflow")
	runCmd.AddCommand(run_listCmd)

	carapace.Gen(run_listCmd).FlagCompletion(carapace.ActionMap{
		"branch": action.ActionBranches(run_listCmd),
		"commit": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionBranchCommits(run_listCmd, run_listCmd.Flag("branch").Value.String())
		}),
		"created":  time.ActionDate(),
		"event":    gh.ActionWorkflowEvents(),
		"json":     action.ActionRunFields(),
		"status":   carapace.ActionValues("queued", "completed", "in_progress", "requested", "waiting", "pending", "action_required", "cancelled", "failure", "neutral", "skipped", "stale", "startup_failure", "success", "timed_out"),
		"user":     gh.ActionUsers(gh.HostOpts{}),
		"workflow": action.ActionWorkflows(run_listCmd, action.WorkflowOpts{Enabled: true, Id: true, Name: true}),
	})
}
