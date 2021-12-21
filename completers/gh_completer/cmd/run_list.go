package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var run_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List recent workflow runs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(run_listCmd).Standalone()
	run_listCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	run_listCmd.Flags().StringSlice("json", []string{}, "Output JSON with the specified `fields`")
	run_listCmd.Flags().IntP("limit", "L", 20, "Maximum number of runs to fetch")
	run_listCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template")
	run_listCmd.Flags().StringP("workflow", "w", "", "Filter runs by workflow")
	runCmd.AddCommand(run_listCmd)

	carapace.Gen(run_listCmd).FlagCompletion(carapace.ActionMap{
		"json":     action.ActionRunFields(),
		"workflow": action.ActionWorkflows(run_listCmd, action.WorkflowOpts{Enabled: true, Id: true, Name: true}),
	})
}
