package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var variable_updateCmd = &cobra.Command{
	Use:   "update <key> <value>",
	Short: "Update an existing variable for a project or group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(variable_updateCmd).Standalone()

	variable_updateCmd.Flags().StringP("description", "d", "", "Set description of a variable.")
	variable_updateCmd.Flags().StringP("group", "g", "", "Set variable for a group.")
	variable_updateCmd.Flags().BoolP("masked", "m", false, "Whether the variable is masked.")
	variable_updateCmd.Flags().BoolP("protected", "p", false, "Whether the variable is protected.")
	variable_updateCmd.Flags().BoolP("raw", "r", false, "Whether the variable is treated as a raw string.")
	variable_updateCmd.Flags().StringP("scope", "s", "", "The environment_scope of the variable. Values: all (*), or specific environments.")
	variable_updateCmd.Flags().StringP("type", "t", "", "The type of a variable: env_var, file.")
	variable_updateCmd.Flags().StringP("value", "v", "", "The value of a variable.")
	variableCmd.AddCommand(variable_updateCmd)

	carapace.Gen(variable_updateCmd).FlagCompletion(carapace.ActionMap{
		"group": action.ActionGroups(variable_updateCmd),
		"scope": carapace.Batch(
			carapace.ActionValues("*"),
			action.ActionEnvironments(variable_updateCmd),
		).ToA(),
		"type": carapace.ActionValues("env_var", "file"),
	})

	carapace.Gen(variable_updateCmd).PositionalCompletion(
		action.ActionVariables(variable_updateCmd),
	)
}
