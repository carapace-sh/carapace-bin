package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var variable_setCmd = &cobra.Command{
	Use:   "set <variable-name>",
	Short: "Create or update variables",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(variable_setCmd).Standalone()

	variable_setCmd.Flags().StringP("body", "b", "", "The value for the variable (reads from standard input if not specified)")
	variable_setCmd.Flags().StringP("env", "e", "", "Set deployment `environment` variable")
	variable_setCmd.Flags().StringP("env-file", "f", "", "Load variable names and values from a dotenv-formatted `file`")
	variable_setCmd.Flags().StringP("org", "o", "", "Set `organization` variable")
	variable_setCmd.Flags().StringSliceP("repos", "r", nil, "List of `repositories` that can access an organization variable")
	variable_setCmd.Flags().StringP("visibility", "v", "", "Set visibility for an organization variable: {all|private|selected}")
	variableCmd.AddCommand(variable_setCmd)

	carapace.Gen(variable_setCmd).FlagCompletion(carapace.ActionMap{
		"env":        action.ActionEnvironments(variable_setCmd),
		"env-file":   carapace.ActionFiles(),
		"org":        gh.ActionOrganizations(gh.HostOpts{}),
		"repos":      gh.ActionOwnerRepositories(gh.HostOpts{}).UniqueList(","),
		"visibility": carapace.ActionValues("all", "private", "selected"),
	})

	carapace.Gen(variable_setCmd).PositionalCompletion(
		action.ActionVariables(variable_setCmd),
	)
}
