package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/simctl"
	"github.com/spf13/cobra"
)

var pbsyncCmd = &cobra.Command{
	Use:   "pbsync",
	Short: "Sync pasteboard content from one device to another",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pbsyncCmd).Standalone()
	rootCmd.AddCommand(pbsyncCmd)
	carapace.Gen(pbsyncCmd).PositionalCompletion(
		simctl.ActionDevices(),
		simctl.ActionDevices(),
	)
}
