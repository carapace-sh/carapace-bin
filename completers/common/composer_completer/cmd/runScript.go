package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/composer_completer/cmd/action"
	"github.com/spf13/cobra"
)

var runScriptCmd = &cobra.Command{
	Use:     "run-script",
	Short:   "Runs the scripts defined in composer.json",
	Aliases: []string{"run"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runScriptCmd).Standalone()

	runScriptCmd.Flags().Bool("dev", false, "Sets the dev mode.")
	runScriptCmd.Flags().BoolP("list", "l", false, "List scripts.")
	runScriptCmd.Flags().Bool("no-dev", false, "Disables the dev mode.")
	runScriptCmd.Flags().String("timeout", "", "Sets script timeout in seconds, or 0 for never.")

	rootCmd.AddCommand(runScriptCmd)

	carapace.Gen(runScriptCmd).PositionalCompletion(
		action.ActionScripts(runScriptCmd),
	)
}
