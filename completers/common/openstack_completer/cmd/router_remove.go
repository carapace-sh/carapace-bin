package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var router_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(router_removeCmd).Standalone()

	routerCmd.AddCommand(router_removeCmd)
}
