package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cache_help_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes metadata cache for the specified package(s). Supports patterns",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cache_help_deleteCmd).Standalone()

	cache_helpCmd.AddCommand(cache_help_deleteCmd)
}
