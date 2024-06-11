package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
	"github.com/spf13/cobra"
)

var flake_newCmd = &cobra.Command{
	Use:   "new [flags] DIR",
	Short: "create a flake in target directory from a template",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(flake_newCmd).Standalone()

	flake_newCmd.Flags().StringP("template", "t", "default", "The template to use")

	addEvaluationFlags(flake_newCmd)
	addLoggingFlags(flake_newCmd)

	carapace.Gen(flake_newCmd).FlagCompletion(carapace.ActionMap{
		"template": nix.ActionTemplates(),
	})
	carapace.Gen(flake_newCmd).PositionalCompletion(carapace.ActionDirectories())

	flakeCmd.AddCommand(flake_newCmd)
}
