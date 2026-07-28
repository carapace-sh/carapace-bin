package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tunnel_showCmd = &cobra.Command{
	Use:     "show",
	Aliases: []string{"list"},
	Short:   "list tunnels",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tunnel_showCmd).Standalone()

	tunnelCmd.AddCommand(tunnel_showCmd)
}
