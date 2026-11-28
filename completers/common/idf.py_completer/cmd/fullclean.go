package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fullcleanCmd = &cobra.Command{
	Use:   "fullclean",
	Short: "Delete the entire build directory contents",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fullcleanCmd).Standalone()
	rootCmd.AddCommand(fullcleanCmd)
}
