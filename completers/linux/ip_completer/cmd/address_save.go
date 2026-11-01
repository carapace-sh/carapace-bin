package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var address_saveCmd = &cobra.Command{
	Use:   "save",
	Short: "save address table to stdout",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(address_saveCmd).Standalone()

	addressCmd.AddCommand(address_saveCmd)
}
