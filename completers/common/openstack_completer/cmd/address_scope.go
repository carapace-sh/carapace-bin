package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var address_scopeCmd = &cobra.Command{
	Use:   "scope",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(address_scopeCmd).Standalone()

	addressCmd.AddCommand(address_scopeCmd)
}
