package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var whoisCmd = &cobra.Command{
	Use:   "whois",
	Short: "Show the machine and user associated with a Tailscale IP (v4 or v6)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(whoisCmd).Standalone()

	whoisCmd.Flags().Bool("json", false, "output in JSON format")
	whoisCmd.Flags().String("proto", "", "protocol to use")
	rootCmd.AddCommand(whoisCmd)

	carapace.Gen(whoisCmd).FlagCompletion(carapace.ActionMap{
		"proto": carapace.ActionValues("tcp", "udp"),
	})
}
