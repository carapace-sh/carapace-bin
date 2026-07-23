package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tabCmd = &cobra.Command{
	Use:   "tab",
	Short: "Manage tabs over the socket API",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tabCmd).Standalone()

	rootCmd.AddCommand(tabCmd)
}
