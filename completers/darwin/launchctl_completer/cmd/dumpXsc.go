package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dumpXscCmd = &cobra.Command{
	Use:   "dump-xsc",
	Short: "Dumps launchd XPC string caches to stdout",
	Run:   func(*cobra.Command, []string) {},
}

func init() {
	carapace.Gen(dumpXscCmd).Standalone()
	rootCmd.AddCommand(dumpXscCmd)
}
