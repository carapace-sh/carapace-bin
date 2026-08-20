package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_listAliasesCmd = &cobra.Command{
	Use:   "list-aliases",
	Short: "List existing plugin aliases",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_listAliasesCmd).Standalone()

	helpCmd.AddCommand(help_listAliasesCmd)
}
