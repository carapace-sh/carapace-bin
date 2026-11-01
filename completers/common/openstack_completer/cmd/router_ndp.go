package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var router_ndpCmd = &cobra.Command{
	Use:   "ndp",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(router_ndpCmd).Standalone()

	routerCmd.AddCommand(router_ndpCmd)
}
