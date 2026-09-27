package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var address_scope_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete address scope(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(address_scope_deleteCmd).Standalone()

	address_scopeCmd.AddCommand(address_scope_deleteCmd)
}
