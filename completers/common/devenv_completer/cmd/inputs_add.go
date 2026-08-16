package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/devenv"
	"github.com/spf13/cobra"
)

var inputs_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add an input to devenv.yaml",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inputs_addCmd).Standalone()

	inputs_addCmd.Flags().StringP("follows", "f", "", "What inputs should follow your inputs?")

	inputsCmd.AddCommand(inputs_addCmd)

	carapace.Gen(inputs_addCmd).FlagCompletion(carapace.ActionMap{
		"follows": devenv.ActionInputs(),
	})

	carapace.Gen(inputs_addCmd).PositionalCompletion(
		carapace.ActionValues(),
		actionSources(),
	)
}
