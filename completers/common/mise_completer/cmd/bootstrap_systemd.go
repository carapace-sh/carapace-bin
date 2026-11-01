package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_systemdCmd = &cobra.Command{
	Use:    "systemd",
	Short:  "Manage systemd user units (use `mise bootstrap linux systemd-units`)",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_systemdCmd).Standalone()

	bootstrap_systemdCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrapCmd.AddCommand(bootstrap_systemdCmd)
}
