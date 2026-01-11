package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mod_vendorCmd = &cobra.Command{
	Use:   "vendor",
	Short: "make vendored copy of dependencies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mod_vendorCmd).Standalone()
	mod_vendorCmd.Flags().SetInterspersed(false)

	mod_vendorCmd.Flags().BoolS("e", "e", false, "attempt to proceed despite errors encountered")
	mod_vendorCmd.Flags().StringS("o", "o", "", "create the vendor directory at the given path")
	mod_vendorCmd.Flags().BoolS("v", "v", false, "print the names of vendored modules and packages")
	modCmd.AddCommand(mod_vendorCmd)

	carapace.Gen(mod_vendorCmd).FlagCompletion(carapace.ActionMap{
		"o": carapace.ActionDirectories(),
	})
}
