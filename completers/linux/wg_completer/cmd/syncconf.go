package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/wg_completer/cmd/action"
	"github.com/spf13/cobra"
)

var syncconfCmd = &cobra.Command{
	Use:   "syncconf",
	Short: "Synchronizes a configuration file to a WireGuard interface",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(syncconfCmd).Standalone()

	rootCmd.AddCommand(syncconfCmd)

	carapace.Gen(syncconfCmd).PositionalCompletion(
		action.ActionInterfaces(),
		carapace.ActionFiles(),
	)
}
