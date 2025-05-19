package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var variable_deleteCmd = &cobra.Command{
	Use:     "delete <variable-name>",
	Short:   "Delete variables",
	Aliases: []string{"remove"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(variable_deleteCmd).Standalone()

	variable_deleteCmd.Flags().StringP("env", "e", "", "Delete a variable for an environment")
	variable_deleteCmd.Flags().StringP("org", "o", "", "Delete a variable for an organization")
	variableCmd.AddCommand(variable_deleteCmd)

	carapace.Gen(variable_deleteCmd).FlagCompletion(carapace.ActionMap{
		"env": action.ActionEnvironments(variable_deleteCmd),
		"org": gh.ActionOrganizations(gh.HostOpts{}),
	})

	carapace.Gen(variable_deleteCmd).PositionalCompletion(
		action.ActionVariables(variable_deleteCmd),
	)
}
