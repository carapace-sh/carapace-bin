package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/simctl"
	"github.com/spf13/cobra"
)

var ioCmd = &cobra.Command{
	Use:   "io",
	Short: "Set up a device IO operation (screenshot, recordVideo)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ioCmd).Standalone()
	rootCmd.AddCommand(ioCmd)
	carapace.Gen(ioCmd).PositionalCompletion(
		simctl.ActionDevices(),
		carapace.ActionValues("screenshot", "recordVideo"),
	)
}
