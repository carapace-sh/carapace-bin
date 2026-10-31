package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var aliases_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove an alias",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(aliases_removeCmd).Standalone()

	aliasesCmd.AddCommand(aliases_removeCmd)
}
