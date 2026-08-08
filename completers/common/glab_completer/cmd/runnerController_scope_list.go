package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var runnerController_scope_listCmd = &cobra.Command{
	Use:   "list <controller-id> [flags]",
	Short: "List scopes for a runner controller. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runnerController_scope_listCmd).Standalone()

	runnerController_scope_listCmd.Flags().String("jq", "", "Filter JSON output with a jq expression.")
	runnerController_scope_listCmd.Flags().StringP("output", "F", "", "Format output as: text, json.")
	runnerController_scopeCmd.AddCommand(runnerController_scope_listCmd)

	carapace.Gen(runnerController_scope_listCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionValues("text", "json"),
	})
}
