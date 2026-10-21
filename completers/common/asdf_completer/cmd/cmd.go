package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/asdf_completer/cmd/action"
	"github.com/spf13/cobra"
)

var cmdCmd = &cobra.Command{
	Use:   "cmd",
	Short: "Runs extension commands for a plugin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cmdCmd).Standalone()

	rootCmd.AddCommand(cmdCmd)

	carapace.Gen(cmdCmd).PositionalCompletion(
		action.ActionPlugins(),
		action.ActionExtensionCommands(""),
	)
}
