package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/carapace-sh/carapace-jq/pkg/actions/tools/jq"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var runner_jobsCmd = &cobra.Command{
	Use:   "jobs <runner-id> [flags]",
	Short: "List jobs processed by a runner.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runner_jobsCmd).Standalone()

	runner_jobsCmd.Flags().String("jq", "", "Filter JSON output with a jq expression.")
	runner_jobsCmd.Flags().String("order-by", "", "Order jobs by: id.")
	runner_jobsCmd.Flags().StringP("output", "F", "", "Format output as: text, json.")
	runner_jobsCmd.Flags().StringP("page", "p", "", "Page number.")
	runner_jobsCmd.Flags().StringP("per-page", "P", "", "Number of items to list per page.")
	runner_jobsCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository. Can use either `OWNER/REPO` or `GROUP/NAMESPACE/REPO` format. Also accepts full URL or Git URL.")
	runner_jobsCmd.Flags().String("sort", "", "Sort order: asc or desc.")
	runner_jobsCmd.Flags().String("status", "", "Filter jobs by status: running, success, failed, canceled.")
	runnerCmd.AddCommand(runner_jobsCmd)

	carapace.Gen(runner_jobsCmd).FlagCompletion(carapace.ActionMap{
		"jq":       jq.ActionFilters(),
		"order-by": carapace.ActionValues("id"),
		"output":   carapace.ActionValues("text", "json"),
		"repo":     action.ActionRepo(runner_jobsCmd),
		"sort":     carapace.ActionValues("asc", "desc").StyleF(style.ForKeyword),
		"status":   carapace.ActionValues("running", "success", "failed", "canceled"),
	})
}
