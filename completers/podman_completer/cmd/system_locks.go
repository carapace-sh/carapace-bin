package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var system_locksCmd = &cobra.Command{
	Use:    "locks",
	Short:  "Debug Libpod's use of locks, identifying any potential conflicts",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(system_locksCmd).Standalone()

	systemCmd.AddCommand(system_locksCmd)
}
