package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var address_showdumpCmd = &cobra.Command{
	Use:   "showdump",
	Short: "dump saved address data",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(address_showdumpCmd).Standalone()

	addressCmd.AddCommand(address_showdumpCmd)
}
