package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var address_groupCmd = &cobra.Command{
	Use:   "group",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(address_groupCmd).Standalone()

	addressCmd.AddCommand(address_groupCmd)
}
