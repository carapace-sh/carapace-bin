package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/carapace-sh/carapace-jq/pkg/actions/tools/jq"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var ci_listCmd = &cobra.Command{
	Use:   "list [flags]",
	Short: "List CI/CD pipelines.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ci_listCmd).Standalone()

	ci_listCmd.Flags().String("jq", "", "Filter JSON output with a jq expression.")
	ci_listCmd.Flags().StringP("name", "n", "", "Return only pipelines with the given name.")
	ci_listCmd.Flags().StringP("order", "o", "id", "Order pipelines by this field. Options: id, status, ref, updated_at, user_id.")
	ci_listCmd.Flags().String("orderBy", "id", "Deprecated: use --order instead.")
	ci_listCmd.Flags().StringP("output", "F", "text", "Format output. Options: text, json.")
	ci_listCmd.Flags().StringP("page", "p", "1", "Page number.")
	ci_listCmd.Flags().StringP("per-page", "P", "", "Number of items to list per page. Defaults to the GitLab API default (20).")
	ci_listCmd.Flags().StringP("ref", "r", "", "Return only pipelines for the given ref.")
	ci_listCmd.Flags().String("scope", "", "Return only pipelines with the given scope. Options: running, pending, finished, branches, tags.")
	ci_listCmd.Flags().String("sha", "", "Return only pipelines with the given SHA.")
	ci_listCmd.Flags().String("sort", "desc", "Sort direction for '--order': asc or desc.")
	ci_listCmd.Flags().String("source", "", "Return only pipelines triggered by the given source. For the full list, see https://docs.gitlab.com/ci/jobs/job_rules/#ci_pipeline_source-predefined-variable. Commonly used options: merge_request_event, parent_pipeline, pipeline, push, trigger.")
	ci_listCmd.Flags().StringP("status", "s", "", "Filter pipelines by status. Options: running, pending, success, failed, canceled, skipped, created, manual, waiting_for_resource, preparing, scheduled.")
	ci_listCmd.Flags().StringP("updated-after", "a", "", "Return only pipelines updated after the specified date. Expected in ISO 8601 format (2019-03-15T08:00:00Z).")
	ci_listCmd.Flags().StringP("updated-before", "b", "", "Return only pipelines updated before the specified date. Expected in ISO 8601 format (2019-03-15T08:00:00Z).")
	ci_listCmd.Flags().StringP("username", "u", "", "Return only pipelines triggered by the given username.")
	ci_listCmd.Flags().BoolP("yaml-errors", "y", false, "Return only pipelines with invalid configurations.")
	ci_listCmd.Flag("orderBy").Hidden = true
	ciCmd.AddCommand(ci_listCmd)

	carapace.Gen(ci_listCmd).FlagCompletion(carapace.ActionMap{
		"jq":       jq.ActionFilters(),
		"order":    carapace.ActionValues("id", "status", "ref", "updated_at", "user_id"),
		"orderBy":  carapace.ActionValues("id", "status", "ref", "updated_at", "user_id"),
		"output":   carapace.ActionValues("text", "json"),
		"ref":      carapace.Batch(action.ActionBranches(ci_listCmd), action.ActionTags(ci_listCmd)).ToA(),
		"scope":    carapace.ActionValues("running", "pending", "finished", "branches", "tags"),
		"sort":     carapace.ActionValues("asc", "desc").StyleF(style.ForKeyword),
		"source":   carapace.ActionValues("merge_request_event", "parent_pipeline", "pipeline", "push", "trigger"),
		"status":   carapace.ActionValues("running", "pending", "success", "failed", "canceled", "skipped", "created", "manual", "waiting_for_resource", "preparing", "scheduled"),
		"username": action.ActionUsers(ci_listCmd),
	})
}
