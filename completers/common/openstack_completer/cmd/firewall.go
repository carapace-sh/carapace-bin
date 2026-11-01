package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var firewallCmd = &cobra.Command{
	Use:   "firewall",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(firewallCmd).Standalone()

	rootCmd.AddCommand(firewallCmd)
}
