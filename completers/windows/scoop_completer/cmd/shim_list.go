package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var shim_listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all shims or matching shims",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shim_listCmd).Standalone()
	shimCmd.AddCommand(shim_listCmd)
}
