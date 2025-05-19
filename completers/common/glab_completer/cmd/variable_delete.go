package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var variable_deleteCmd = &cobra.Command{
	Use:     "delete <key>",
	Short:   "Delete a variable for a project or group.",
	Aliases: []string{"remove"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(variable_deleteCmd).Standalone()

	variable_deleteCmd.Flags().StringP("group", "g", "", "Delete variable from a group.")
	variable_deleteCmd.Flags().StringP("scope", "s", "", "The 'environment_scope' of the variable. Options: all (*), or specific environments.")
	variableCmd.AddCommand(variable_deleteCmd)

	carapace.Gen(variable_deleteCmd).FlagCompletion(carapace.ActionMap{
		"group": action.ActionGroups(variable_deleteCmd),
		"scope": carapace.Batch(
			carapace.ActionValues("*"),
			action.ActionEnvironments(variable_updateCmd),
		).ToA(),
	})

	carapace.Gen(variable_deleteCmd).PositionalCompletion(
		action.ActionVariables(variable_deleteCmd),
	)
}
