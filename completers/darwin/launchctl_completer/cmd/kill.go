package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/launchctl"
	"github.com/spf13/cobra"
)

var killCmd = &cobra.Command{
	Use:   "kill",
	Short: "Send a signal to the service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(killCmd).Standalone()
	killCmd.Flags().StringP("signal", "s", "", "Signal to send to the service")
	rootCmd.AddCommand(killCmd)
	carapace.Gen(killCmd).PositionalAnyCompletion(launchctl.ActionServices())
	carapace.Gen(killCmd).FlagCompletion(carapace.ActionMap{
		"signal": carapace.ActionValues("TERM", "KILL", "INT", "QUIT", "HUP", "USR1", "USR2"),
	})
}
