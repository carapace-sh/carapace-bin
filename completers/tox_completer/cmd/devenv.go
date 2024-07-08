package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tox"
	"github.com/spf13/cobra"
)

var devenvCmd = &cobra.Command{
	Use:     "devenv",
	Aliases: []string{"d"},
	Short:   "sets up a development environment at ENVDIR based on the tox configuration specified",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devenvCmd).Standalone()

	devenvCmd.Flags().StringS("e", "e", "", "environment to run (default: py)")
	devenvCmd.Flags().Bool("no-recreate-pkg", false, "if recreate is set do not recreate packaging tox environment(s) (default: False)")
	devenvCmd.Flags().String("skip-env", "", "exclude all environments selected that match this regular expression (default: '')")
	addCommonSubcommandFlags(devenvCmd)
	rootCmd.AddCommand(devenvCmd)

	carapace.Gen(devenvCmd).FlagCompletion(carapace.ActionMap{
		"e": tox.ActionEnvironments(),
	})

	carapace.Gen(devenvCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
