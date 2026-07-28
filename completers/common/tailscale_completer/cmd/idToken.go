package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var idTokenCmd = &cobra.Command{
	Use:   "id-token",
	Short: "Fetch an OIDC id-token for the Tailscale machine",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(idTokenCmd).Standalone()

	rootCmd.AddCommand(idTokenCmd)
}
