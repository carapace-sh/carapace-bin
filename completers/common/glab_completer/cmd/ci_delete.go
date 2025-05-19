package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var ci_deleteCmd = &cobra.Command{
	Use:   "delete <id> [flags]",
	Short: "Delete CI/CD pipelines.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ci_deleteCmd).Standalone()

	ci_deleteCmd.Flags().Bool("dry-run", false, "Simulate process, but do not delete anything.")
	ci_deleteCmd.Flags().String("older-than", "", "Filter pipelines older than the given duration. Valid units: h, m, s, ms, us, ns.")
	ci_deleteCmd.Flags().String("page", "", "Page number.")
	ci_deleteCmd.Flags().Bool("paginate", false, "Make additional HTTP requests to fetch all pages of projects before cloning. Respects '--per-page'.")
	ci_deleteCmd.Flags().String("per-page", "", "Number of items to list per page.")
	ci_deleteCmd.Flags().String("source", "", "Filter pipelines by source: api, chat, external, external_pull_request_event, merge_request_event, ondemand_dast_scan, ondemand_dast_validation, parent_pipeline, pipeline, push, schedule, security_orchestration_policy, trigger, web, webide.")
	ci_deleteCmd.Flags().StringP("status", "s", "", "Delete pipelines by status: running, pending, success, failed, canceled, skipped, created, manual.")
	ciCmd.AddCommand(ci_deleteCmd)

	carapace.Gen(ci_deleteCmd).FlagCompletion(carapace.ActionMap{
		"source": carapace.ActionValues(
			"api",
			"chat",
			"external",
			"external_pull_request_event",
			"merge_request_event",
			"ondemand_dast_scan",
			"ondemand_dast_validation",
			"parent_pipeline",
			"pipeline",
			"push",
			"schedule",
			"security_orchestration_policy",
			"trigger",
			"web",
			"webide",
		),
		"status": carapace.ActionValues(
			"running",
			"pending",
			"success",
			"failed",
			"canceled",
			"skipped",
			"created",
			"manual",
		).StyleF(style.ForKeyword),
	})

	carapace.Gen(ci_deleteCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionPipelines(ci_deleteCmd, ci_deleteCmd.Flag("status").Value.String()).UniqueList(",")
		}),
	)
}
