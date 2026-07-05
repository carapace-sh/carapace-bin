package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/simctl"
	"github.com/spf13/cobra"
)

var addmediaCmd = &cobra.Command{
	Use:   "addmedia",
	Short: "Add photos, videos, or contacts to a device's library",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(addmediaCmd).Standalone()
	rootCmd.AddCommand(addmediaCmd)
	carapace.Gen(addmediaCmd).PositionalCompletion(
		simctl.ActionDevices(),
		carapace.ActionFiles(),
	)
}
