package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var work_vendorCmd = &cobra.Command{
	Use:   "vendor",
	Short: "make vendored copy of dependencies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(work_vendorCmd).Standalone()

	work_vendorCmd.Flags().BoolS("e", "e", false, "attempt to proceed despite errors")
	work_vendorCmd.Flags().StringS("o", "o", "", "create the vendor directory at the given path")
	work_vendorCmd.Flags().BoolS("v", "v", false, "print the names of vendored modules and packages")
	workCmd.AddCommand(work_vendorCmd)

	carapace.Gen(work_vendorCmd).FlagCompletion(carapace.ActionMap{
		"o": carapace.ActionDirectories(),
	})
}
