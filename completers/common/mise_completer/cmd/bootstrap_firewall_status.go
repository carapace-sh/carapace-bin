// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_firewall_statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show the state of the firewall from `[bootstrap.linux.firewall]`",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_firewall_statusCmd).Standalone()

	bootstrap_firewall_statusCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrap_firewall_statusCmd.Flags().BoolP("json", "J", false, "Output in JSON format")
	bootstrap_firewall_statusCmd.Flags().Bool("missing", false, "Exit with status 1 if the firewall differs from the config")
	bootstrap_firewallCmd.AddCommand(bootstrap_firewall_statusCmd)
}
