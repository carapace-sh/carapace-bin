package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cache_help_viewCmd = &cobra.Command{
	Use:   "view",
	Short: "Views information from the specified package's cache",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cache_help_viewCmd).Standalone()

	cache_helpCmd.AddCommand(cache_help_viewCmd)
}
