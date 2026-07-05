package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/simctl"
	"github.com/spf13/cobra"
)

var bootCmd = &cobra.Command{
	Use:   "boot",
	Short: "Boot a device",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootCmd).Standalone()
	rootCmd.AddCommand(bootCmd)
	carapace.Gen(bootCmd).PositionalCompletion(simctl.ActionDevices())
}
