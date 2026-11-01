package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cache_help_pathCmd = &cobra.Command{
	Use:   "path",
	Short: "Prints the path to the cache directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cache_help_pathCmd).Standalone()

	cache_helpCmd.AddCommand(cache_help_pathCmd)
}
