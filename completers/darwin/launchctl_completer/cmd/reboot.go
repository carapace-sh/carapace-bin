package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rebootCmd = &cobra.Command{
	Use:   "reboot",
	Short: "Initiates a system reboot of the specified type",
	Run:   func(*cobra.Command, []string) {},
}

func init() {
	carapace.Gen(rebootCmd).Standalone()
	rootCmd.AddCommand(rebootCmd)
	carapace.Gen(rebootCmd).PositionalCompletion(
		carapace.ActionValues("system", "halt", "userspace", "logout", "obliterate", "panic"),
	)
}
