package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/carapace-sh/carapace-jq/pkg/actions/tools/jq"
	"github.com/spf13/cobra"
)

var runner_managersCmd = &cobra.Command{
	Use:   "managers <runner-id> [flags]",
	Short: "List runner managers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runner_managersCmd).Standalone()

	runner_managersCmd.Flags().String("jq", "", "Filter JSON output with a jq expression.")
	runner_managersCmd.Flags().StringP("output", "F", "", "Format output as: text, json.")
	runner_managersCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository. Can use either `OWNER/REPO` or `GROUP/NAMESPACE/REPO` format. Also accepts full URL or Git URL.")
	runnerCmd.AddCommand(runner_managersCmd)

	carapace.Gen(runner_managersCmd).FlagCompletion(carapace.ActionMap{
		"jq":     jq.ActionFilters(),
		"output": carapace.ActionValues("text", "json"),
		"repo":   action.ActionRepo(runner_managersCmd),
	})
}
