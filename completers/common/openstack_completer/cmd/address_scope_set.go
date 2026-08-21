package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var address_scope_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set address scope properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(address_scope_setCmd).Standalone()

	address_scope_setCmd.Flags().String("extra-property", "", "Additional parameters can be passed using this property.")
	address_scope_setCmd.Flags().String("name", "", "Set address scope name")
	address_scope_setCmd.Flags().Bool("no-share", false, "Do not share the address scope between projects")
	address_scope_setCmd.Flags().Bool("share", false, "Share the address scope between projects")
	address_scopeCmd.AddCommand(address_scope_setCmd)
}
