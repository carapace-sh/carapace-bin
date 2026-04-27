package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deleteCacheCmd = &cobra.Command{
	Use:     "delete-cache",
	Short:   "Delete cache files",
	Aliases: []string{"dc"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deleteCacheCmd).Standalone()

	rootCmd.AddCommand(deleteCacheCmd)
}
