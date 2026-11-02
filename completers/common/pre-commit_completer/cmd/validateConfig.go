package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/pre-commit_completer/cmd/action"
	"github.com/spf13/cobra"
)

var validateConfigCmd = &cobra.Command{
	Use:   "validate-config [filenames...]",
	Short: "Validate .pre-commit-config.yaml files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(validateConfigCmd).Standalone()

	validateConfigCmd.Flags().String("color", "", "Whether to use color in output.")

	carapace.Gen(validateConfigCmd).FlagCompletion(carapace.ActionMap{
		"color": action.ActionColors(),
	})

	carapace.Gen(validateConfigCmd).PositionalAnyCompletion(
		carapace.ActionFiles(".yaml", ".yml"),
	)

	rootCmd.AddCommand(validateConfigCmd)
}
