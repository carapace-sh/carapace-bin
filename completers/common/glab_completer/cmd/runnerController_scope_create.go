package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-jq/pkg/actions/tools/jq"
	"github.com/spf13/cobra"
)

var runnerController_scope_createCmd = &cobra.Command{
	Use:   "create <controller-id> [flags]",
	Short: "Create a scope for a runner controller. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runnerController_scope_createCmd).Standalone()

	runnerController_scope_createCmd.Flags().Bool("instance", false, "Add an instance-level scope.")
	runnerController_scope_createCmd.Flags().String("jq", "", "Filter JSON output with a jq expression.")
	runnerController_scope_createCmd.Flags().StringP("output", "F", "text", "Format output as: text, json.")
	runnerController_scope_createCmd.Flags().StringSlice("runner", nil, "Add a runner-level scope for the specified runner ID. Multiple IDs can be comma-separated or specified by repeating the flag.")
	runnerController_scopeCmd.AddCommand(runnerController_scope_createCmd)

	carapace.Gen(runnerController_scope_createCmd).FlagCompletion(carapace.ActionMap{
		"jq":     jq.ActionFilters(),
		"output": carapace.ActionValues("text", "json"),
	})
}
