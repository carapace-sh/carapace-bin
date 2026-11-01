package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dhcpclientCmd = &cobra.Command{
	Use:   "dhcpclient",
	Short: "DHCP client configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dhcpclientCmd).Standalone()
	rootCmd.AddCommand(dhcpclientCmd)
}
