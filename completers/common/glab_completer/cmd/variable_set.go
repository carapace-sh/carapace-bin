package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var variable_setCmd = &cobra.Command{
	Use:     "set <key> <value>",
	Short:   "Create a new variable for a project or group.",
	Aliases: []string{"new", "create"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(variable_setCmd).Standalone()

	variable_setCmd.Flags().StringP("description", "d", "", "Set description of a variable.")
	variable_setCmd.Flags().StringP("group", "g", "", "Set variable for a group.")
	variable_setCmd.Flags().BoolP("masked", "m", false, "Whether the variable is masked.")
	variable_setCmd.Flags().BoolP("protected", "p", false, "Whether the variable is protected.")
	variable_setCmd.Flags().BoolP("raw", "r", false, "Whether the variable is treated as a raw string.")
	variable_setCmd.Flags().StringP("scope", "s", "", "The environment_scope of the variable. Values: all (*), or specific environments.")
	variable_setCmd.Flags().StringP("type", "t", "", "The type of a variable: env_var, file.")
	variable_setCmd.Flags().StringP("value", "v", "", "The value of a variable.")
	variableCmd.AddCommand(variable_setCmd)

	carapace.Gen(variable_setCmd).FlagCompletion(carapace.ActionMap{
		"group": action.ActionGroups(variableCmd),
		"scope": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return carapace.Batch(
				action.ActionEnvironments(variable_setCmd),
				carapace.ActionValues("*"),
			).Invoke(c).Merge().ToA()
		}),
		"type": carapace.ActionValues("env_var", "file"),
	})

	carapace.Gen(variable_setCmd).PositionalCompletion(
		action.ActionVariables(variable_setCmd),
	)
}
