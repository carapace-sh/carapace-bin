package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var shim_addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a custom shim",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shim_addCmd).Standalone()
	shimCmd.AddCommand(shim_addCmd)
}
