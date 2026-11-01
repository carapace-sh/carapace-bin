package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_linux_systemdUnitsCmd = &cobra.Command{
	Use:     "systemd-units",
	Short:   "Manage systemd user units from `[bootstrap.linux.systemd.units]`",
	Aliases: []string{"systemd"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_linux_systemdUnitsCmd).Standalone()

	bootstrap_linux_systemdUnitsCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrap_linuxCmd.AddCommand(bootstrap_linux_systemdUnitsCmd)
}
