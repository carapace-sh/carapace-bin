package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var workflow_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List workflows",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workflow_listCmd).Standalone()

	workflow_listCmd.Flags().BoolP("all", "a", false, "Include disabled workflows")
	workflow_listCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	workflow_listCmd.Flags().StringSlice("json", nil, "Output JSON with the specified `fields`")
	workflow_listCmd.Flags().StringP("limit", "L", "", "Maximum number of workflows to fetch")
	workflow_listCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template; see \"gh help formatting\"")
	workflowCmd.AddCommand(workflow_listCmd)

	carapace.Gen(workflow_listCmd).FlagCompletion(carapace.ActionMap{
		"json": gh.ActionWorkflowFields().UniqueList(","),
	})
}
