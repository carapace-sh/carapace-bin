package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/npm"
	"github.com/spf13/cobra"
)

var setScriptCmd = &cobra.Command{
	Use:   "set-script",
	Short: "Set tasks in the scripts section of package.json",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setScriptCmd).Standalone()
	addWorkspaceFlags(setScriptCmd)

	rootCmd.AddCommand(setScriptCmd)

	carapace.Gen(setScriptCmd).PositionalCompletion(
		npm.ActionScripts(),
	)
}
