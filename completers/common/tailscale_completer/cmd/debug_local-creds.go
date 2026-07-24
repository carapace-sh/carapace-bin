package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_local_credsCmd = &cobra.Command{
	Use:   "local-creds",
	Short: "Print how to access Tailscale LocalAPI",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_local_credsCmd).Standalone()

	debugCmd.AddCommand(debug_local_credsCmd)
}
