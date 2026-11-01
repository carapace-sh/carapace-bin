package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bgpvpn_networkCmd = &cobra.Command{
	Use:   "network",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bgpvpn_networkCmd).Standalone()

	bgpvpnCmd.AddCommand(bgpvpn_networkCmd)
}
