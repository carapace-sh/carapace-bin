package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var setup_nixCmd = &cobra.Command{
	Use:   "nix",
	Short: "Install Nix",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setup_nixCmd).Standalone()

	setup_nixCmd.Flags().Bool("daemon", false, "Install Nix in multi-user mode. This flag is not supported if you are using DetSys installer")
	setupCmd.AddCommand(setup_nixCmd)
}
