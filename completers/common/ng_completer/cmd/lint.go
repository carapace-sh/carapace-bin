package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/ng_completer/cmd/action"
	"github.com/spf13/cobra"
)

var lintCmd = &cobra.Command{
	Use:   "lint",
	Short: "Runs linting tools on Angular app code in a given project folder",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lintCmd).Standalone()

	lintCmd.Flags().StringP("configuration", "c", "", "One or more named builder configurations")

	rootCmd.AddCommand(lintCmd)

	carapace.Gen(lintCmd).FlagCompletion(carapace.ActionMap{
		"configuration": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) < 1 {
				return carapace.ActionValues()
			}
			return action.ActionBuilderConfigurations(c.Args[0], "lint").UniqueList(",")
		}),
	})

	carapace.Gen(lintCmd).PositionalCompletion(
		action.ActionProjects(),
	)
}
