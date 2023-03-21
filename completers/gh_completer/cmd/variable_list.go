package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/gh"
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
	variable_listCmd.Flags().StringP("org", "o", "", "List variables for an organization")
	variableCmd.AddCommand(variable_listCmd)

	carapace.Gen(variable_listCmd).FlagCompletion(carapace.ActionMap{
		"env": action.ActionEnvironments(variable_listCmd),
		"org": gh.ActionOrganizations(gh.HostOpts{}),
	})
}
