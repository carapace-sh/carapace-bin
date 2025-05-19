package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var workflow_enableCmd = &cobra.Command{
	Use:   "enable [<workflow-id> | <workflow-name>]",
	Short: "Enable a workflow",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workflow_enableCmd).Standalone()

	workflowCmd.AddCommand(workflow_enableCmd)

	carapace.Gen(workflow_enableCmd).PositionalCompletion(
		action.ActionWorkflows(workflow_enableCmd, action.WorkflowOpts{Disabled: true, Id: true, Name: true}),
	)
}
