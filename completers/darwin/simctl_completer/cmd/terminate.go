package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/simctl"
	"github.com/spf13/cobra"
)

var terminateCmd = &cobra.Command{
	Use:   "terminate",
	Short: "Terminate an application by identifier on a device",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(terminateCmd).Standalone()
	rootCmd.AddCommand(terminateCmd)
	carapace.Gen(terminateCmd).PositionalCompletion(
		simctl.ActionDevices(),
	)
}
