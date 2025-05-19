package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var variable_getCmd = &cobra.Command{
	Use:   "get <key>",
	Short: "Get a variable for a project or group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(variable_getCmd).Standalone()

	variable_getCmd.Flags().StringP("group", "g", "", "Get variable for a group.")
	variable_getCmd.Flags().StringP("output", "F", "", "Format output as: text, json.")
	variable_getCmd.Flags().StringP("scope", "s", "", "The environment_scope of the variable. Values: all (*), or specific environments.")
	variableCmd.AddCommand(variable_getCmd)

	carapace.Gen(variable_getCmd).FlagCompletion(carapace.ActionMap{
		"group": action.ActionGroups(variable_getCmd),
		"scope": carapace.Batch(
			carapace.ActionValues("*"),
			action.ActionEnvironments(variable_getCmd),
		).ToA(),
	})

	carapace.Gen(variable_getCmd).PositionalCompletion(
		action.ActionVariables(variable_getCmd),
	)
}
