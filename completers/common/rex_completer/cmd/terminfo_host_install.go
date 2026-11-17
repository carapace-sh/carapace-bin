package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var terminfoHostInstallCmd = &cobra.Command{
	Use:   "install DESTINATION [-- SSH_OPTIONS...]",
	Short: "Install Rex's terminfo entry on a remote host",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(terminfoHostInstallCmd).Standalone()

	terminfoHostInstallCmd.Flags().SetInterspersed(false)

	terminfoHostCmd.AddCommand(terminfoHostInstallCmd)

	carapace.Gen(terminfoHostInstallCmd).DashAnyCompletion(
		bridge.ActionCarapaceBin("ssh"),
	)
}
