package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/pre-commit_completer/cmd/action"
	"github.com/spf13/cobra"
)

var validateManifestCmd = &cobra.Command{
	Use:   "validate-manifest [filenames...]",
	Short: "Validate .pre-commit-hooks.yaml files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(validateManifestCmd).Standalone()

	validateManifestCmd.Flags().String("color", "", "Whether to use color in output.")

	carapace.Gen(validateManifestCmd).FlagCompletion(carapace.ActionMap{
		"color": action.ActionColors(),
	})

	carapace.Gen(validateManifestCmd).PositionalAnyCompletion(
		carapace.ActionFiles(".yaml", ".yml"),
	)

	rootCmd.AddCommand(validateManifestCmd)
}
