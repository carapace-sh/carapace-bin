package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var shim_infoCmd = &cobra.Command{
	Use:   "info",
	Short: "show a shim's information",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shim_infoCmd).Standalone()
	shimCmd.AddCommand(shim_infoCmd)
}
