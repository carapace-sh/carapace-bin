package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var variable_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List variables",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(variable_listCmd).Standalone()

	variable_listCmd.Flags().StringP("env", "e", "", "List variables for an environment")
	variable_listCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	variable_listCmd.Flags().StringSlice("json", nil, "Output JSON with the specified `fields`")
	variable_listCmd.Flags().StringP("org", "o", "", "List variables for an organization")
	variable_listCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template; see \"gh help formatting\"")
	variableCmd.AddCommand(variable_listCmd)

	carapace.Gen(variable_listCmd).FlagCompletion(carapace.ActionMap{
		"env":  action.ActionEnvironments(variable_listCmd),
		"json": gh.ActionVariableFields().UniqueList(","),
		"org":  gh.ActionOrganizations(gh.HostOpts{}),
	})
}
