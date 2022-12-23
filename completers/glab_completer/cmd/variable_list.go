package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var variable_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List project or group variables",
	Aliases: []string{"new", "create"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(variable_listCmd).Standalone()
	variable_listCmd.PersistentFlags().StringP("group", "g", "", "Select a group/subgroup. This option is ignored if a repo argument is set.")
	variable_listCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository using the `OWNER/REPO` or `GROUP/NAMESPACE/REPO` format or full URL or git URL")
	variableCmd.AddCommand(variable_listCmd)

	carapace.Gen(variable_listCmd).FlagCompletion(carapace.ActionMap{
		"group": action.ActionGroups(variable_listCmd),
	})
}
