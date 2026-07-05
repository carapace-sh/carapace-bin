package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/simctl"
	"github.com/spf13/cobra"
)

var openurlCmd = &cobra.Command{
	Use:   "openurl",
	Short: "Open a URL in a device",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(openurlCmd).Standalone()
	rootCmd.AddCommand(openurlCmd)
	carapace.Gen(openurlCmd).PositionalCompletion(simctl.ActionDevices())
}
