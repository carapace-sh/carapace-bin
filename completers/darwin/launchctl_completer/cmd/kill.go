package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var killCmd = &cobra.Command{
	Use:   "kill",
	Short: "Sends a signal to the service instance",
	Run:   func(*cobra.Command, []string) {},
}

func init() {
	carapace.Gen(killCmd).Standalone()
	rootCmd.AddCommand(killCmd)
	carapace.Gen(killCmd).PositionalCompletion(
		carapace.ActionValues("SIGHUP", "SIGINT", "SIGKILL", "SIGTERM", "SIGSTOP", "SIGCONT", "SIGUSR1", "SIGUSR2"),
		actionDomainTargets(),
	)
}
