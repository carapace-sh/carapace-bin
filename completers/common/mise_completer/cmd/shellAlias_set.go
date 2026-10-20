package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var shellAlias_setCmd = &cobra.Command{
	Use:     "set",
	Short:   "Set a shell alias in the global config",
	Aliases: []string{"add", "create"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shellAlias_setCmd).Standalone()

	shellAlias_setCmd.Flags().BoolP("help", "h", false, "Print help")
	shellAliasCmd.AddCommand(shellAlias_setCmd)

	carapace.Gen(shellAlias_setCmd).PositionalCompletion(
		carapace.ActionValues(),
		bridge.ActionCarapaceBin().Split(),
	)
}
