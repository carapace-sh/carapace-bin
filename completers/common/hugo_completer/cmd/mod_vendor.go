package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mod_vendorCmd = &cobra.Command{
	Use:   "vendor",
	Short: "Vendor all module dependencies into the _vendor directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mod_vendorCmd).Standalone()
	modCmd.AddCommand(mod_vendorCmd)
}
