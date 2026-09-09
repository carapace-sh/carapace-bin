package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var host_setCmd = &cobra.Command{
	Use:   "set",
	Short: "DEPRECATED: Set host properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(host_setCmd).Standalone()

	host_setCmd.Flags().Bool("disable", false, "Disable the host")
	host_setCmd.Flags().Bool("disable-maintenance", false, "Disable maintenance mode for the host")
	host_setCmd.Flags().Bool("enable", false, "Enable the host")
	host_setCmd.Flags().Bool("enable-maintenance", false, "Enable maintenance mode for the host")
	hostCmd.AddCommand(host_setCmd)
}
