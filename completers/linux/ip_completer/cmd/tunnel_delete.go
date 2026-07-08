package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tunnel_deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"del"},
	Short:   "destroy a tunnel",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tunnel_deleteCmd).Standalone()

	tunnelCmd.AddCommand(tunnel_deleteCmd)
}
