package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var listSessionsCmd = &cobra.Command{
	Use:     "list-sessions",
	Aliases: []string{"ls"},
	Short:   "list sessions managed by server",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listSessionsCmd).Standalone()

	listSessionsCmd.Flags().StringS("F", "F", "", "specify output format")
	listSessionsCmd.Flags().StringS("O", "O", "", "initial sort order")
	listSessionsCmd.Flags().StringS("f", "f", "", "filter items")
	listSessionsCmd.Flags().BoolS("r", "r", false, "reverse sort order")
	rootCmd.AddCommand(listSessionsCmd)
}
