package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/simctl"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install an application on a device",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installCmd).Standalone()
	rootCmd.AddCommand(installCmd)
	carapace.Gen(installCmd).PositionalCompletion(
		simctl.ActionDevices(),
		carapace.ActionDirectories(),
	)
}
