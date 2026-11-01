package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debugupdatecachesCmd = &cobra.Command{
	Use:    "debugupdatecaches",
	Short:  "warm all known caches in the repository",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugupdatecachesCmd).Standalone()

	rootCmd.AddCommand(debugupdatecachesCmd)
}
