package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var variable_getCmd = &cobra.Command{
	Use:   "get <variable-name>",
	Short: "Get variables",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(variable_getCmd).Standalone()

	variable_getCmd.Flags().StringP("env", "e", "", "Get a variable for an environment")
	variable_getCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	variable_getCmd.Flags().StringSlice("json", nil, "Output JSON with the specified `fields`")
	variable_getCmd.Flags().StringP("org", "o", "", "Get a variable for an organization")
	variable_getCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template; see \"gh help formatting\"")
	variableCmd.AddCommand(variable_getCmd)

	carapace.Gen(variable_getCmd).FlagCompletion(carapace.ActionMap{
		"env":  action.ActionEnvironments(variable_setCmd),
		"json": gh.ActionVariableFields().UniqueList(","),
		"org":  gh.ActionOrganizations(gh.HostOpts{}),
	})

	carapace.Gen(variable_getCmd).PositionalCompletion(
		action.ActionVariables(variable_getCmd),
	)
}
