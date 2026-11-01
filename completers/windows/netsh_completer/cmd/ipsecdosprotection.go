package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ipsecdosprotectionCmd = &cobra.Command{
	Use:   "ipsecdosprotection",
	Short: "IPsec DoS protection configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ipsecdosprotectionCmd).Standalone()
	rootCmd.AddCommand(ipsecdosprotectionCmd)
}
