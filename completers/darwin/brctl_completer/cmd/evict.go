package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var evictCmd = &cobra.Command{
	Use:   "evict",
	Short: "evict the local copy of the document",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(evictCmd).Standalone()
	rootCmd.AddCommand(evictCmd)
}
