package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tuntap_deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"del"},
	Short:   "remove a TUN or TAP device",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tuntap_deleteCmd).Standalone()

	tuntapCmd.AddCommand(tuntap_deleteCmd)
}
