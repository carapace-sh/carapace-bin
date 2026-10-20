package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_linux_systemdUnits_statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show the state of the configured systemd user units",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_linux_systemdUnits_statusCmd).Standalone()

	bootstrap_linux_systemdUnits_statusCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrap_linux_systemdUnits_statusCmd.Flags().BoolP("json", "J", false, "Output in JSON format")
	bootstrap_linux_systemdUnits_statusCmd.Flags().Bool("missing", false, "Exit with status 1 if any unit differs from the config (all are still listed)")
	bootstrap_linux_systemdUnitsCmd.AddCommand(bootstrap_linux_systemdUnits_statusCmd)
}
