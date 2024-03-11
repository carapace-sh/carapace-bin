package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var daemonReexecCmd = &cobra.Command{
	Use:     "daemon-reexec",
	Short:   "Reexecute systemd manager",
	GroupID: "manager state",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(daemonReexecCmd).Standalone()

	rootCmd.AddCommand(daemonReexecCmd)
}
