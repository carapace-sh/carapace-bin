package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sharedmemoryCmd = &cobra.Command{
	Use:   "sharedmemory",
	Short: "Show shared memory information",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sharedmemoryCmd).Standalone()
	rootCmd.AddCommand(sharedmemoryCmd)
}
