package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/launchctl"
	"github.com/spf13/cobra"
)

var kickstartCmd = &cobra.Command{
	Use:   "kickstart",
	Short: "Force a service to start in the specified domain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kickstartCmd).Standalone()
	kickstartCmd.Flags().BoolP("kill", "k", false, "Kill the running service before starting")
	kickstartCmd.Flags().BoolP("terminate", "p", false, "Terminate the service by sending SIGTERM")
	rootCmd.AddCommand(kickstartCmd)
	carapace.Gen(kickstartCmd).PositionalAnyCompletion(launchctl.ActionServices())
}
