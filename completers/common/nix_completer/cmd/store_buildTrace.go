package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var store_buildTraceCmd = &cobra.Command{
	Use:     "build-trace",
	Short:   "manipulate a Nix build trace",
	Aliases: []string{"realisation"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(store_buildTraceCmd).Standalone()

	storeCmd.AddCommand(store_buildTraceCmd)
}
