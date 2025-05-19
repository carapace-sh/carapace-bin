package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var workflow_runCmd = &cobra.Command{
	Use:   "run [<workflow-id> | <workflow-name>]",
	Short: "Run a workflow by creating a workflow_dispatch event",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workflow_runCmd).Standalone()

	workflow_runCmd.Flags().StringSliceP("field", "F", nil, "Add a string parameter in `key=value` format, respecting @ syntax (see \"gh help api\").")
	workflow_runCmd.Flags().Bool("json", false, "Read workflow inputs as JSON via STDIN")
	workflow_runCmd.Flags().StringSliceP("raw-field", "f", nil, "Add a string parameter in `key=value` format")
	workflow_runCmd.Flags().StringP("ref", "r", "", "Branch or tag name which contains the version of the workflow file you'd like to run")
	workflowCmd.AddCommand(workflow_runCmd)

	carapace.Gen(workflow_runCmd).FlagCompletion(carapace.ActionMap{
		"ref": action.ActionBranches(workflow_runCmd),
	})

	carapace.Gen(workflow_runCmd).PositionalCompletion(
		action.ActionWorkflows(workflow_runCmd, action.WorkflowOpts{Enabled: true, Id: true, Name: true}),
	)
}
