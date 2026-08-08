package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var runnerController_token_rotateCmd = &cobra.Command{
	Use:   "rotate <controller-id> <token-id> [flags]",
	Short: "Rotate a token for a runner controller. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runnerController_token_rotateCmd).Standalone()

	runnerController_token_rotateCmd.Flags().BoolP("force", "f", false, "Skip confirmation prompt.")
	runnerController_token_rotateCmd.Flags().String("jq", "", "Filter JSON output with a jq expression.")
	runnerController_token_rotateCmd.Flags().StringP("output", "F", "", "Format output as: text, json.")
	runnerController_tokenCmd.AddCommand(runnerController_token_rotateCmd)

	carapace.Gen(runnerController_token_rotateCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionValues("text", "json"),
	})
}
