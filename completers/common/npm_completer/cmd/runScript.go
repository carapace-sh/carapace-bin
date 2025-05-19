package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/npm"
	"github.com/spf13/cobra"
)

var runScriptCmd = &cobra.Command{
	Use:     "run-script",
	Short:   "Run arbitrary package scripts",
	Aliases: []string{"run"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runScriptCmd).Standalone()
	runScriptCmd.Flags().Bool("if-present", false, "")
	runScriptCmd.Flags().Bool("ignore-scripts", false, "do not run scripts specified in package.json")
	runScriptCmd.Flags().String("script-shell", "", "shell to use for scripts")
	addWorkspaceFlags(runScriptCmd)

	rootCmd.AddCommand(runScriptCmd)

	carapace.Gen(runScriptCmd).FlagCompletion(carapace.ActionMap{
		// TODO flag completion
		"script-shell": os.ActionShells(),
	})

	carapace.Gen(runScriptCmd).PositionalCompletion(
		npm.ActionScripts(),
	)
}
