package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var accountCmd = &cobra.Command{
	Use:   "account",
	Short: "Display commands that retrieve account details",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(accountCmd).Standalone()
	rootCmd.AddCommand(accountCmd)
}
