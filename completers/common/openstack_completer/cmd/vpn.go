package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpnCmd = &cobra.Command{
	Use:   "vpn",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpnCmd).Standalone()

	rootCmd.AddCommand(vpnCmd)
}
