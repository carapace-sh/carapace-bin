package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-jq/pkg/actions/tools/jq"
	"github.com/spf13/cobra"
)

var runnerController_token_createCmd = &cobra.Command{
	Use:   "create <controller-id> [flags]",
	Short: "Create a token for a runner controller. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runnerController_token_createCmd).Standalone()

	runnerController_token_createCmd.Flags().StringP("description", "d", "", "Description of the token.")
	runnerController_token_createCmd.Flags().String("jq", "", "Filter JSON output with a jq expression.")
	runnerController_token_createCmd.Flags().StringP("output", "F", "", "Format output as: text, json.")
	runnerController_tokenCmd.AddCommand(runnerController_token_createCmd)

	carapace.Gen(runnerController_token_createCmd).FlagCompletion(carapace.ActionMap{
		"jq":     jq.ActionFilters(),
		"output": carapace.ActionValues("text", "json"),
	})
}
