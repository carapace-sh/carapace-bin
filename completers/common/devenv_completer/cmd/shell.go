package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/devenv_completer/cmd/action"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var shellCmd = &cobra.Command{
	Use:   "shell",
	Short: "Activate the developer environment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shellCmd).Standalone()
	shellCmd.Flags().SetInterspersed(false)

	rootCmd.AddCommand(shellCmd)

	carapace.Gen(shellCmd).PositionalCompletion(
		carapace.Batch(
			action.ActionScripts(shellCmd),
			carapace.ActionExecutables(),
			carapace.ActionFiles(),
		).ToA(),
	)

	carapace.Gen(shellCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin(),
	)
}
