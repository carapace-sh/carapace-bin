package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/carapace-sh/carapace-jq/pkg/actions/tools/jq"
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

	variable_exportCmd.Flags().String("format", "json", "Format of output: json, export, env.")
	variable_exportCmd.PersistentFlags().StringP("group", "g", "", "Select a group or subgroup. Ignored if a repository argument is set.")
	variable_exportCmd.Flags().String("jq", "", "Filter JSON output with a jq expression.")
	variable_exportCmd.Flags().StringP("output", "F", "json", "Format output as: json, export, env.")
	variable_exportCmd.Flags().StringP("page", "p", "1", "Page number.")
	variable_exportCmd.Flags().StringP("per-page", "P", "100", "Number of items to list per page.")
	variable_exportCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository. You can use either OWNER/REPO or GROUP/NAMESPACE/REPO. The full URL or Git URL is also accepted.")
	variable_exportCmd.Flags().StringP("scope", "s", "*", "The environment_scope of the variables. Values: '*' (default), or specific environments.")
	variable_exportCmd.Flag("format").Hidden = true
	variableCmd.AddCommand(variable_exportCmd)

	carapace.Gen(variable_exportCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "export", "env"),
		"group":  action.ActionGroups(variable_exportCmd),
		"jq":     jq.ActionFilters(),
		"output": carapace.ActionValues("json", "export", "env"),
		"repo":   action.ActionRepo(variable_exportCmd),
	})
}
