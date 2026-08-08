package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-jq/pkg/actions/tools/jq"
	"github.com/spf13/cobra"
)

var runnerController_token_listCmd = &cobra.Command{
	Use:   "list <controller-id> [flags]",
	Short: "List tokens for a runner controller. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runnerController_token_listCmd).Standalone()

	runnerController_token_listCmd.Flags().String("jq", "", "Filter JSON output with a jq expression.")
	runnerController_token_listCmd.Flags().StringP("output", "F", "", "Format output as: text, json.")
	runnerController_token_listCmd.Flags().StringP("page", "p", "", "Page number.")
	runnerController_token_listCmd.Flags().StringP("per-page", "P", "", "Number of items per page.")
	runnerController_tokenCmd.AddCommand(runnerController_token_listCmd)

	carapace.Gen(runnerController_token_listCmd).FlagCompletion(carapace.ActionMap{
		"jq":     jq.ActionFilters(),
		"output": carapace.ActionValues("text", "json"),
	})
}
