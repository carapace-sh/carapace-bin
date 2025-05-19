package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var extension_execCmd = &cobra.Command{
	Use:   "exec <name> [args]",
	Short: "Execute an installed extension",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(extension_execCmd).Standalone()

	extensionCmd.AddCommand(extension_execCmd)

	carapace.Gen(extension_execCmd).PositionalCompletion(
		action.ActionExtensions(),
	)

	carapace.Gen(extension_execCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("gh"),
	)
}
