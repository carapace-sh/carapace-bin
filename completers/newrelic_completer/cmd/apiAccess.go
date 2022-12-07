package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var apiAccessCmd = &cobra.Command{
	Use:   "apiAccess",
	Short: "Manage New Relic API access keys",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apiAccessCmd).Standalone()
	rootCmd.AddCommand(apiAccessCmd)
}
