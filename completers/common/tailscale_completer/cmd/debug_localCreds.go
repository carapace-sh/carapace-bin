package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_localCredsCmd = &cobra.Command{
	Use:   "local-creds",
	Short: "Print how to access Tailscale LocalAPI",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_localCredsCmd).Standalone()

	debugCmd.AddCommand(debug_localCredsCmd)
}
