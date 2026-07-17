package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dumpXscCmd = &cobra.Command{
	Use:   "dump-xsc",
	Short: "Dump launchd XPC string caches to stdout",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dumpXscCmd).Standalone()
	rootCmd.AddCommand(dumpXscCmd)
}
