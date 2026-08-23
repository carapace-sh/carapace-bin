package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(evictCmd)
}

var evictCmd = &cobra.Command{
	Use:   "evict",
	Short: "evict the local copy of the document",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(evictCmd).Standalone()
}