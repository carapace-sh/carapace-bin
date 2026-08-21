package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var volume_host_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set volume host properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_host_setCmd).Standalone()

	volume_host_setCmd.Flags().Bool("disable", false, "Freeze and disable the specified volume host")
	volume_host_setCmd.Flags().Bool("enable", false, "Thaw and enable the specified volume host")
	volume_hostCmd.AddCommand(volume_host_setCmd)
}
