package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bgpvpnCmd = &cobra.Command{
	Use:   "bgpvpn",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bgpvpnCmd).Standalone()

	rootCmd.AddCommand(bgpvpnCmd)
}
