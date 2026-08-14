package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var listAliasesCmd = &cobra.Command{
	Use:     "list-aliases",
	Short:   "List existing plugin aliases",
	Aliases: []string{"la"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listAliasesCmd).Standalone()

	listAliasesCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(listAliasesCmd)
}
