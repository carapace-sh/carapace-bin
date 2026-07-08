package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tunnel_changeCmd = &cobra.Command{
	Use:   "change",
	Short: "change an existing tunnel",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tunnel_changeCmd).Standalone()

	tunnelCmd.AddCommand(tunnel_changeCmd)
}
