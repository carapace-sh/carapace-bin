package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tuntap_showCmd = &cobra.Command{
	Use:     "show",
	Aliases: []string{"list"},
	Short:   "list TUN/TAP devices",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tuntap_showCmd).Standalone()

	tuntapCmd.AddCommand(tuntap_showCmd)
}
