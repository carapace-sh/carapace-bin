package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_viaCmd = &cobra.Command{
	Use:   "via",
	Short: "Convert between site-specific IPv4 CIDRs and IPv6 'via' routes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_viaCmd).Standalone()

	debugCmd.AddCommand(debug_viaCmd)
}
