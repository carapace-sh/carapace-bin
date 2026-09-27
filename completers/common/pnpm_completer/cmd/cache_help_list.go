package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cache_help_listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists the available packages metadata cache. Supports filtering by glob",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cache_help_listCmd).Standalone()

	cache_helpCmd.AddCommand(cache_help_listCmd)
}
