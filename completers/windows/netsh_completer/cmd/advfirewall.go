package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var advfirewallCmd = &cobra.Command{
	Use:   "advfirewall",
	Short: "advanced firewall configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(advfirewallCmd).Standalone()
	rootCmd.AddCommand(advfirewallCmd)
}
