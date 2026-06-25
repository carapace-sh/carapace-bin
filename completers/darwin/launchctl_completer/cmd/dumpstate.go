package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dumpstateCmd = &cobra.Command{
	Use:   "dumpstate",
	Short: "Dumps launchd state to stdout",
	Run:   func(*cobra.Command, []string) {},
}

func init() {
	carapace.Gen(dumpstateCmd).Standalone()
	rootCmd.AddCommand(dumpstateCmd)
}
