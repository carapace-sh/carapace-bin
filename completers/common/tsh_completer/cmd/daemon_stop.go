package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var daemon_stopCmd = &cobra.Command{
	Use:    "stop",
	Short:  "Gracefully stops a process on Windows by sending Ctrl-Break to it.",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(daemon_stopCmd).Standalone()

	daemon_stopCmd.Flags().String("pid", "", "PID to be stopped.")
	daemonCmd.AddCommand(daemon_stopCmd)
}
