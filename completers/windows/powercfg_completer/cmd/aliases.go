package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var aliasesCmd = &cobra.Command{
	Use:   "aliases",
	Short: "display aliases and corresponding identifiers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(aliasesCmd).Standalone()
	rootCmd.AddCommand(aliasesCmd)
}
