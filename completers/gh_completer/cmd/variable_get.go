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
	variable_getCmd.Flags().StringP("org", "o", "", "Get a variable for an organization")
	variableCmd.AddCommand(variable_getCmd)

	carapace.Gen(variable_setCmd).FlagCompletion(carapace.ActionMap{
		"env": action.ActionEnvironments(variable_setCmd),
		"org": gh.ActionOrganizations(gh.HostOpts{}),
	})

	carapace.Gen(variable_getCmd).PositionalCompletion(
		action.ActionVariables(variable_getCmd),
	)
}
