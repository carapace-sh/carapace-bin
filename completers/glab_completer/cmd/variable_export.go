package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var variable_exportCmd = &cobra.Command{
	Use:     "export",
	Short:   "Export variables from a project or group.",
	Aliases: []string{"ex"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(variable_exportCmd).Standalone()

	variable_exportCmd.Flags().StringP("format", "F", "", "Format of output: json, export, env.")
	variable_exportCmd.PersistentFlags().StringP("group", "g", "", "Select a group or subgroup. Ignored if a repository argument is set.")
	variable_exportCmd.Flags().StringP("page", "p", "", "Page number.")
	variable_exportCmd.Flags().StringP("per-page", "P", "", "Number of items to list per page.")
	variable_exportCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository. Can use either `OWNER/REPO` or `GROUP/NAMESPACE/REPO` format. Also accepts full URL or Git URL.")
	variable_exportCmd.Flags().StringP("scope", "s", "", "The environment_scope of the variables. Values: '*' (default), or specific environments.")
	variableCmd.AddCommand(variable_exportCmd)

	// TODO complete scope
	carapace.Gen(variable_exportCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "export", "env"),
		"group":  action.ActionGroups(variable_exportCmd),
		"repo":   action.ActionRepo(variable_exportCmd),
	})
}
