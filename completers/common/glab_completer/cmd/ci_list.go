package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ci_listCmd = &cobra.Command{
	Use:   "list [flags]",
	Short: "Get the list of CI/CD pipelines.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ci_listCmd).Standalone()

	ci_listCmd.Flags().StringP("name", "n", "", "Return only pipelines with the given name.")
	ci_listCmd.Flags().StringP("orderBy", "o", "", "Order pipelines by this field. Options: id, status, ref, updated_at, user_id.")
	ci_listCmd.Flags().StringP("output", "F", "", "Format output. Options: text, json.")
	ci_listCmd.Flags().StringP("page", "p", "", "Page number.")
	ci_listCmd.Flags().StringP("per-page", "P", "", "Number of items to list per page.")
	ci_listCmd.Flags().StringP("ref", "r", "", "Return only pipelines for given ref.")
	ci_listCmd.Flags().String("scope", "", "Return only pipelines with the given scope: {running|pending|finished|branches|tags}")
	ci_listCmd.Flags().String("sha", "", "Return only pipelines with the given SHA.")
	ci_listCmd.Flags().String("sort", "", "Sort pipelines. Options: asc, desc.")
	ci_listCmd.Flags().String("source", "", "Return only pipelines triggered via the given source. See https://docs.gitlab.com/ci/jobs/job_rules/#ci_pipeline_source-predefined-variable for full list. Commonly used options: {merge_request_event|parent_pipeline|pipeline|push|trigger}")
	ci_listCmd.Flags().StringP("status", "s", "", "Get pipeline with this status. Options: running, pending, success, failed, canceled, skipped, created, manual, waiting_for_resource, preparing, scheduled")
	ci_listCmd.Flags().StringP("updated-after", "a", "", "Return only pipelines updated after the specified date. Expected in ISO 8601 format (2019-03-15T08:00:00Z).")
	ci_listCmd.Flags().StringP("updated-before", "b", "", "Return only pipelines updated before the specified date. Expected in ISO 8601 format (2019-03-15T08:00:00Z).")
	ci_listCmd.Flags().StringP("username", "u", "", "Return only pipelines triggered by the given username.")
	ci_listCmd.Flags().BoolP("yaml-errors", "y", false, "Return only pipelines with invalid configurations.")
	ciCmd.AddCommand(ci_listCmd)

	// TODO complete new flags
	carapace.Gen(ci_listCmd).FlagCompletion(carapace.ActionMap{
		"orderBy": carapace.ActionValues("id", "status", "ref", "updated_at", "user_id"),
		"output":  carapace.ActionValues("text", "json"),
		"sort":    carapace.ActionValues("asc", "desc"),
		"status":  carapace.ActionValues("running", "pending", "success", "failed", "canceled", "skipped", "created", "manual", "waiting_for_resource", "preparing", "scheduled"),
	})
}
