package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ci_listCmd = &cobra.Command{
	Use:   "list [flags]",
	Short: "Get the list of CI/CD pipelines",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ci_listCmd).Standalone()

	ci_listCmd.Flags().StringP("orderBy", "o", "", "Order pipeline by {id|status|ref|updated_at|user_id}")
	ci_listCmd.Flags().StringP("page", "p", "", "Page number")
	ci_listCmd.Flags().StringP("per-page", "P", "", "Number of items to list per page")
	ci_listCmd.Flags().String("sort", "", "Sort pipeline by {asc|desc}")
	ci_listCmd.Flags().StringP("status", "s", "", "Get pipeline with status: {running|pending|success|failed|canceled|skipped|created|manual|waiting_for_resource|preparing|scheduled}")
	ciCmd.AddCommand(ci_listCmd)

	carapace.Gen(ci_listCmd).FlagCompletion(carapace.ActionMap{
		"orderBy": carapace.ActionValues("id", "status", "ref", "updated_at", "user_id"),
		"sort":    carapace.ActionValues("asc", "desc"),
		"status":  carapace.ActionValues("running", "pending", "success", "failed", "canceled", "skipped", "created", "manual", "waiting_for_resource", "preparing", "scheduled"),
	})
}
