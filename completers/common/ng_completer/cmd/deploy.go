package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/ng_completer/cmd/action"
	"github.com/spf13/cobra"
)

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Invokes the deploy builder for a specified project or for the default project in the workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deployCmd).Standalone()
	deployCmd.Flags().StringP("configuration", "c", "", "One or more named builder configurations")

	rootCmd.AddCommand(deployCmd)

	carapace.Gen(deployCmd).FlagCompletion(carapace.ActionMap{
		"configuration": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) < 1 {
				return carapace.ActionValues()
			}
			return action.ActionBuilderConfigurations(c.Args[0], "").UniqueList(",")
		}),
	})

	carapace.Gen(deployCmd).PositionalCompletion(
		action.ActionProjects(),
	)
}
