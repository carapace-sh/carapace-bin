package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debugfsinfoCmd = &cobra.Command{
	Use:    "debugfsinfo",
	Short:  "show information detected about current filesystem",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugfsinfoCmd).Standalone()

	rootCmd.AddCommand(debugfsinfoCmd)
}
