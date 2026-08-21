package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var router_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete router(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(router_deleteCmd).Standalone()

	routerCmd.AddCommand(router_deleteCmd)
}
