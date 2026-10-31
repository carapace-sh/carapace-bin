package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var aliasesCmd = &cobra.Command{
	Use:   "aliases",
	Short: "Manage model aliases",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(aliasesCmd).Standalone()

	rootCmd.AddCommand(aliasesCmd)
}
