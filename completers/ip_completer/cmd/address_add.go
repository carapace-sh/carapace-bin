package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var address_addCmd = &cobra.Command{
	Use:   "add",
	Short: "add new protocol address",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(address_addCmd).Standalone()

	addressCmd.AddCommand(address_addCmd)
}
