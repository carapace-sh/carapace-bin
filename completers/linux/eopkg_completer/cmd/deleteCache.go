package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deleteCacheCmd = &cobra.Command{
	Use:     "delete-cache",
	Aliases: []string{"dc"},
	Short:   "clear out any temporary caches still held by eopkg for downloads and package files",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deleteCacheCmd).Standalone()

	rootCmd.AddCommand(deleteCacheCmd)
}
