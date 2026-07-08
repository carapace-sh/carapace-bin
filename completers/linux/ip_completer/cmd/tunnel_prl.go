package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tunnel_prlCmd = &cobra.Command{
	Use:   "prl",
	Short: "potential router list (ISATAP only)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tunnel_prlCmd).Standalone()

	tunnelCmd.AddCommand(tunnel_prlCmd)
}
