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

	ci_listCmd.Flags().StringP("orderBy", "o", "", "Order pipelines by this field. Options: id, status, ref, updated_at, user_id.")
	ci_listCmd.Flags().StringP("output", "F", "", "Format output. Options: text, json.")
	ci_listCmd.Flags().StringP("page", "p", "", "Page number.")
	ci_listCmd.Flags().StringP("per-page", "P", "", "Number of items to list per page.")
	ci_listCmd.Flags().String("sort", "", "Sort pipelines. Options: asc, desc.")
	ci_listCmd.Flags().StringP("status", "s", "", "Get pipeline with this status. Options: running, pending, success, failed, canceled, skipped, created, manual, waiting_for_resource, preparing, scheduled}")
	ciCmd.AddCommand(ci_listCmd)

	carapace.Gen(ci_listCmd).FlagCompletion(carapace.ActionMap{
		"orderBy": carapace.ActionValues("id", "status", "ref", "updated_at", "user_id"),
		"output":  carapace.ActionValues("text", "json"),
		"sort":    carapace.ActionValues("asc", "desc"),
		"status":  carapace.ActionValues("running", "pending", "success", "failed", "canceled", "skipped", "created", "manual", "waiting_for_resource", "preparing", "scheduled"),
	})
}
