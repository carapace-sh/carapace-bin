package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ipsecCmd = &cobra.Command{
	Use:   "ipsec",
	Short: "IPsec configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ipsecCmd).Standalone()
	rootCmd.AddCommand(ipsecCmd)
}
