package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var variable_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List variables for a project or group.",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(variable_listCmd).Standalone()

	variable_listCmd.PersistentFlags().StringP("group", "g", "", "Select a group or subgroup. Ignored if a repository argument is set.")
	variable_listCmd.Flags().StringP("output", "F", "", "Format output as: text, json.")
	variable_listCmd.Flags().StringP("page", "p", "", "Page number.")
	variable_listCmd.Flags().StringP("per-page", "P", "", "Number of items to list per page.")
	variable_listCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository. Can use either `OWNER/REPO` or `GROUP/NAMESPACE/REPO` format. Also accepts full URL or Git URL.")
	variableCmd.AddCommand(variable_listCmd)

	carapace.Gen(variable_listCmd).FlagCompletion(carapace.ActionMap{
		"group": action.ActionGroups(variable_listCmd),
		"repo":  action.ActionRepo(variable_listCmd),
	})
}
