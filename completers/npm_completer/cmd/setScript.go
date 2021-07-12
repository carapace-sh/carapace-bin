package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/npm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var setScriptCmd = &cobra.Command{
	Use:   "set-script",
	Short: "Set tasks in the scripts section of package.json",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setScriptCmd).Standalone()
	setScriptCmd.Flags().StringP("workspace", "w", "", "Enable running a command in the context of the given workspace")
	setScriptCmd.Flags().Bool("workspaces", false, "Enable running a command in the context fo all workspaces")

	rootCmd.AddCommand(setScriptCmd)

	carapace.Gen(setScriptCmd).PositionalCompletion(
		action.ActionScripts(setScriptCmd),
	)
}
