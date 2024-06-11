package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
	"github.com/spf13/cobra"
)

var flake_initCmd = &cobra.Command{
	Use:   "init",
	Short: "create a flake in current directory from a template",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(flake_initCmd).Standalone()

	flake_initCmd.Flags().StringP("template", "t", "default", "The template to use")

	addEvaluationFlags(flake_initCmd)
	addLoggingFlags(flake_initCmd)

	carapace.Gen(flake_initCmd).FlagCompletion(carapace.ActionMap{
		"template": nix.ActionTemplates(),
	})

	flakeCmd.AddCommand(flake_initCmd)
}
