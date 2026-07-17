package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_resolveCmd = &cobra.Command{
	Use:   "resolve",
	Short: "Does a DNS lookup",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_resolveCmd).Standalone()

	debug_resolveCmd.Flags().String("net", "", "network type (ip, ip4, or ip6)")
	debugCmd.AddCommand(debug_resolveCmd)

	carapace.Gen(debug_resolveCmd).FlagCompletion(carapace.ActionMap{
		"net": carapace.ActionValues("ip", "ip4", "ip6"),
	})
}
