package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_firewallCmd = &cobra.Command{
	Use:   "firewall",
	Short: "Manage the Linux host firewall from `[bootstrap.linux.firewall]`",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_firewallCmd).Standalone()

	bootstrap_firewallCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrapCmd.AddCommand(bootstrap_firewallCmd)
}
