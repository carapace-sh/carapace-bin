package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var accountsCmd = &cobra.Command{
	Use:   "accounts",
	Short: "Listing of your collaborations",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(accountsCmd).Standalone()

	rootCmd.AddCommand(accountsCmd)
}
