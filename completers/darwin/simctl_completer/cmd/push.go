package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/simctl"
	"github.com/spf13/cobra"
)

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Send a simulated push notification",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pushCmd).Standalone()
	rootCmd.AddCommand(pushCmd)
	carapace.Gen(pushCmd).PositionalCompletion(
		simctl.ActionDevices(),
	)
	carapace.Gen(pushCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
